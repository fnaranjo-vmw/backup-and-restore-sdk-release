package azure

import (
	"fmt"
	"net/url"

	"context"

	"encoding/base64"

	"time"

	"github.com/Azure/azure-storage-blob-go/2017-07-29/azblob"
)

//go:generate counterfeiter -o fakes/fake_container.go . Container
type Container interface {
	Name() string
	SoftDeleteEnabled() (bool, error)
	ListBlobs() ([]BlobId, error)
	CopyBlobsFrom(containerName string, blobIds []BlobId) error
}

type SDKContainer struct {
	name    string
	service azblob.ServiceURL
}

func NewContainer(name, storageAccount, storageKey string) (container SDKContainer, err error) {
	credential, err := buildCredential(storageAccount, storageKey)
	if err != nil {
		return SDKContainer{}, err
	}

	pipeline := azblob.NewPipeline(credential, azblob.PipelineOptions{})
	azureURL, err := url.Parse(fmt.Sprintf("https://%s.blob.core.windows.net", storageAccount))
	if err != nil {
		return SDKContainer{}, fmt.Errorf("invalid account name: '%s'", storageAccount)
	}

	service := azblob.NewServiceURL(*azureURL, pipeline)
	return SDKContainer{
		name:    name,
		service: service,
	}, nil
}

func (c SDKContainer) Name() string {
	return c.name
}

func (c SDKContainer) CopyBlobsFrom(sourceContainerName string, blobIds []BlobId) error {
	sourceContainerURL := c.service.NewContainerURL(sourceContainerName)

	blobs, err := fetchBlobs(sourceContainerURL)
	if err != nil {
		return err
	}

	for _, blobId := range blobIds {
		sourceBlob, ok := blobs[blobId]
		if !ok {
			return fmt.Errorf("no \"%s\" blob with \"%s\" ETag found in container \"%s\"", blobId.Name, blobId.ETag, sourceContainerName)
		}

		err = c.copyBlob(sourceContainerURL, sourceBlob)
		if err != nil {
			return err
		}
	}

	return nil
}

func fetchBlobs(sourceContainerURL azblob.ContainerURL) (map[BlobId]azblob.Blob, error) {
	var blobs = map[BlobId]azblob.Blob{}

	for marker := (azblob.Marker{}); marker.NotDone(); {
		page, err := sourceContainerURL.ListBlobs(
			context.Background(),
			marker,
			azblob.ListBlobsOptions{
				Details: azblob.BlobListingDetails{
					Snapshots: true,
					Deleted:   true,
				},
			},
		)
		if err != nil {
			return nil, err
		}

		marker = page.NextMarker

		for _, blob := range page.Blobs.Blob {
			blobId := BlobId{Name: blob.Name, ETag: string(blob.Properties.Etag)}
			blobs[blobId] = blob
		}
	}

	return blobs, nil
}

func (c SDKContainer) copyBlob(sourceContainerURL azblob.ContainerURL, blob azblob.Blob) error {
	ctx := context.Background()

	sourceBlobURL := sourceContainerURL.NewBlobURL(blob.Name)
	destinationContainerURL := c.service.NewContainerURL(c.name)
	destinationBlobURL := destinationContainerURL.NewBlobURL(blob.Name)

	_, err := sourceBlobURL.Undelete(ctx)
	if err != nil {
		return err
	}

	response, err := destinationBlobURL.StartCopy(
		ctx,
		sourceBlobURL.WithSnapshot(blob.Snapshot).URL(),
		azblob.Metadata{},
		azblob.BlobAccessConditions{},
		azblob.BlobAccessConditions{},
	)
	if err != nil {
		return err
	}

	copyStatus := response.CopyStatus()
	for copyStatus == azblob.CopyStatusPending {
		time.Sleep(time.Second * 2)
		getMetadata, err := destinationBlobURL.GetPropertiesAndMetadata(ctx, azblob.BlobAccessConditions{})
		if err != nil {
			return err
		}

		copyStatus = getMetadata.CopyStatus()
	}

	return nil
}

func (c SDKContainer) SoftDeleteEnabled() (bool, error) {
	properties, err := c.service.GetProperties(context.Background())
	if err != nil {
		return false, fmt.Errorf("failed fetching properties for storage account: '%s'", err)
	}

	if properties.DeleteRetentionPolicy == nil {
		return true, nil
	}

	return properties.DeleteRetentionPolicy.Enabled, nil
}

func (c SDKContainer) ListBlobs() ([]BlobId, error) {
	var blobs []BlobId
	client := c.service.NewContainerURL(c.name)

	for marker := (azblob.Marker{}); marker.NotDone(); {
		page, err := client.ListBlobs(context.Background(), marker, azblob.ListBlobsOptions{})
		if err != nil {
			return nil, fmt.Errorf("failed listing blobs in container '%s': %s", c.name, err)
		}

		marker = page.NextMarker

		for _, blobInfo := range page.Blobs.Blob {
			blobs = append(blobs, BlobId{Name: blobInfo.Name, ETag: string(blobInfo.Properties.Etag)})
		}
	}

	return blobs, nil
}

func buildCredential(storageAccount, storageKey string) (*azblob.SharedKeyCredential, error) {
	_, err := base64.StdEncoding.DecodeString(storageKey)
	if err != nil {
		return nil, fmt.Errorf("invalid storage key: '%s'", err)
	}

	return azblob.NewSharedKeyCredential(storageAccount, storageKey), nil
}