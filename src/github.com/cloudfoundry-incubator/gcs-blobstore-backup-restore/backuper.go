package gcs

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Backuper struct {
	buckets map[string]BucketPair
}

const liveBucketBackupArtifactName = "temporary-backup-artifact"
const commonBlobsName = "common_blobs.json"

func NewBackuper(buckets map[string]BucketPair) Backuper {
	return Backuper{
		buckets: buckets,
	}
}

func (b *Backuper) Unlock() (map[string]BackupBucketAddress, error) {

	backupBuckets, err := b.TransferBlobsToBackupBucket()

	err = b.CopyBlobsWithinBackupBucket(backupBuckets)

	return backupBuckets, err
}

func (b *Backuper) CreateLiveBucketSnapshot() error {
	for _, bucketPair := range b.buckets {
		var commonBlobs []Blob
		bucket := bucketPair.Bucket

		blobs, err := bucket.ListBlobs()
		if err != nil {
			return err
		}

		lastBackupBlobs, err := bucketPair.BackupBucket.ListLastBackupBlobs()
		if err != nil {
			return err
		}

		inLastBackup := make(map[string]Blob)
		for _, blob := range lastBackupBlobs {
			nameParts := strings.Split(blob.Name, "/")
			inLastBackup[nameParts[len(nameParts)-1]] = blob
		}

		for _, blob := range blobs {
			if blobFromBackup, ok := inLastBackup[blob.Name]; ok {
				commonBlobs = append(commonBlobs, blobFromBackup)
			} else {
				err := bucket.CopyBlobWithinBucket(blob.Name, fmt.Sprintf("%s/%s", liveBucketBackupArtifactName, blob.Name))
				if err != nil {
					return err
				}
			}
		}

		j, err := json.Marshal(commonBlobs)
		if err != nil {
			return err
		}

		err = bucket.CreateFile(liveBucketBackupArtifactName+"/"+commonBlobsName, j)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *Backuper) CleanupLiveBuckets() error {
	for _, bucketPair := range b.buckets {
		blobs, err := bucketPair.Bucket.ListBlobs()
		if err != nil {
			return err
		}

		for _, blob := range blobs {
			if strings.HasPrefix(blob.Name, fmt.Sprintf("%s/", liveBucketBackupArtifactName)) {
				err = bucketPair.Bucket.DeleteBlob(blob.Name)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (b *Backuper) TransferBlobsToBackupBucket() (map[string]BackupBucketAddress, error) {
	timestamp := time.Now().Format("2006_01_02_15_04_05")

	backupBuckets := make(map[string]BackupBucketAddress)

	for id, bucketPair := range b.buckets {
		bucket := bucketPair.Bucket
		backupBucket := bucketPair.BackupBucket

		backupBlobs, err := bucket.ListBlobs()

		if err != nil {
			return nil, err
		}

		backupBuckets[id] = BackupBucketAddress{
			BucketName: backupBucket.Name(),
			Path:       fmt.Sprintf("%s/%s", timestamp, id),
		}
		for _, blob := range backupBlobs {
			if strings.HasPrefix(blob.Name, fmt.Sprintf("%s/", liveBucketBackupArtifactName)) {

				blobName := strings.Replace(blob.Name, liveBucketBackupArtifactName, fmt.Sprintf("%s/%s", timestamp, id), 1)

				err := bucket.CopyBlobBetweenBuckets(backupBucket, blob.Name, blobName)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return backupBuckets, nil
}

func (b *Backuper) CopyBlobsWithinBackupBucket(backupBucketAddresses map[string]BackupBucketAddress) error {
	for bucketId, backupBucketAddress := range backupBucketAddresses {
		commonBlobBytes, err := b.buckets[bucketId].BackupBucket.GetBlob(backupBucketAddress.Path + "/common_blobs.json")
		if err != nil {
			return fmt.Errorf("failed to get %s/common_blobs.json: %v", backupBucketAddress.Path, err)
		}

		var commonBlobs []Blob
		err = json.Unmarshal(commonBlobBytes, &commonBlobs)
		if err != nil {
			return err
		}

		err = b.buckets[bucketId].BackupBucket.DeleteBlob(backupBucketAddress.Path + "/common_blobs.json")

		if err != nil {
			return err
		}

		for _, blob := range commonBlobs {
			nameParts := strings.Split(blob.Name, "/")
			destinationBlobName := fmt.Sprintf("%s/%s", backupBucketAddress.Path, nameParts[len(nameParts)-1])
			err = b.buckets[bucketId].BackupBucket.CopyBlobWithinBucket(blob.Name, destinationBlobName)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
