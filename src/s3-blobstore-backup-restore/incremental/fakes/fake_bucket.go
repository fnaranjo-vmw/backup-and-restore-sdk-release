// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"s3-blobstore-backup-restore/incremental"
)

type FakeBucket struct {
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct{}
	nameReturns     struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	RegionStub        func() string
	regionMutex       sync.RWMutex
	regionArgsForCall []struct{}
	regionReturns     struct {
		result1 string
	}
	regionReturnsOnCall map[int]struct {
		result1 string
	}
	ListBlobsStub        func(path string) ([]incremental.Blob, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct {
		path string
	}
	listBlobsReturns struct {
		result1 []incremental.Blob
		result2 error
	}
	listBlobsReturnsOnCall map[int]struct {
		result1 []incremental.Blob
		result2 error
	}
	ListDirectoriesStub        func() ([]string, error)
	listDirectoriesMutex       sync.RWMutex
	listDirectoriesArgsForCall []struct{}
	listDirectoriesReturns     struct {
		result1 []string
		result2 error
	}
	listDirectoriesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	CopyBlobWithinBucketStub        func(src, dst string) error
	copyBlobWithinBucketMutex       sync.RWMutex
	copyBlobWithinBucketArgsForCall []struct {
		src string
		dst string
	}
	copyBlobWithinBucketReturns struct {
		result1 error
	}
	copyBlobWithinBucketReturnsOnCall map[int]struct {
		result1 error
	}
	CopyBlobFromBucketStub        func(bucket incremental.Bucket, src, dst string) error
	copyBlobFromBucketMutex       sync.RWMutex
	copyBlobFromBucketArgsForCall []struct {
		bucket incremental.Bucket
		src    string
		dst    string
	}
	copyBlobFromBucketReturns struct {
		result1 error
	}
	copyBlobFromBucketReturnsOnCall map[int]struct {
		result1 error
	}
	UploadBlobStub        func(path, contents string) error
	uploadBlobMutex       sync.RWMutex
	uploadBlobArgsForCall []struct {
		path     string
		contents string
	}
	uploadBlobReturns struct {
		result1 error
	}
	uploadBlobReturnsOnCall map[int]struct {
		result1 error
	}
	HasBlobStub        func(path string) (bool, error)
	hasBlobMutex       sync.RWMutex
	hasBlobArgsForCall []struct {
		path string
	}
	hasBlobReturns struct {
		result1 bool
		result2 error
	}
	hasBlobReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct{}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.nameReturns.result1
}

func (fake *FakeBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBucket) NameReturns(result1 string) {
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) NameReturnsOnCall(i int, result1 string) {
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) Region() string {
	fake.regionMutex.Lock()
	ret, specificReturn := fake.regionReturnsOnCall[len(fake.regionArgsForCall)]
	fake.regionArgsForCall = append(fake.regionArgsForCall, struct{}{})
	fake.recordInvocation("Region", []interface{}{})
	fake.regionMutex.Unlock()
	if fake.RegionStub != nil {
		return fake.RegionStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.regionReturns.result1
}

func (fake *FakeBucket) RegionCallCount() int {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	return len(fake.regionArgsForCall)
}

func (fake *FakeBucket) RegionReturns(result1 string) {
	fake.RegionStub = nil
	fake.regionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) RegionReturnsOnCall(i int, result1 string) {
	fake.RegionStub = nil
	if fake.regionReturnsOnCall == nil {
		fake.regionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.regionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) ListBlobs(path string) ([]incremental.Blob, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("ListBlobs", []interface{}{path})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listBlobsReturns.result1, fake.listBlobsReturns.result2
}

func (fake *FakeBucket) ListBlobsCallCount() int {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return len(fake.listBlobsArgsForCall)
}

func (fake *FakeBucket) ListBlobsArgsForCall(i int) string {
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	return fake.listBlobsArgsForCall[i].path
}

func (fake *FakeBucket) ListBlobsReturns(result1 []incremental.Blob, result2 error) {
	fake.ListBlobsStub = nil
	fake.listBlobsReturns = struct {
		result1 []incremental.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListBlobsReturnsOnCall(i int, result1 []incremental.Blob, result2 error) {
	fake.ListBlobsStub = nil
	if fake.listBlobsReturnsOnCall == nil {
		fake.listBlobsReturnsOnCall = make(map[int]struct {
			result1 []incremental.Blob
			result2 error
		})
	}
	fake.listBlobsReturnsOnCall[i] = struct {
		result1 []incremental.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListDirectories() ([]string, error) {
	fake.listDirectoriesMutex.Lock()
	ret, specificReturn := fake.listDirectoriesReturnsOnCall[len(fake.listDirectoriesArgsForCall)]
	fake.listDirectoriesArgsForCall = append(fake.listDirectoriesArgsForCall, struct{}{})
	fake.recordInvocation("ListDirectories", []interface{}{})
	fake.listDirectoriesMutex.Unlock()
	if fake.ListDirectoriesStub != nil {
		return fake.ListDirectoriesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.listDirectoriesReturns.result1, fake.listDirectoriesReturns.result2
}

func (fake *FakeBucket) ListDirectoriesCallCount() int {
	fake.listDirectoriesMutex.RLock()
	defer fake.listDirectoriesMutex.RUnlock()
	return len(fake.listDirectoriesArgsForCall)
}

func (fake *FakeBucket) ListDirectoriesReturns(result1 []string, result2 error) {
	fake.ListDirectoriesStub = nil
	fake.listDirectoriesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListDirectoriesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.ListDirectoriesStub = nil
	if fake.listDirectoriesReturnsOnCall == nil {
		fake.listDirectoriesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.listDirectoriesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) CopyBlobWithinBucket(src string, dst string) error {
	fake.copyBlobWithinBucketMutex.Lock()
	ret, specificReturn := fake.copyBlobWithinBucketReturnsOnCall[len(fake.copyBlobWithinBucketArgsForCall)]
	fake.copyBlobWithinBucketArgsForCall = append(fake.copyBlobWithinBucketArgsForCall, struct {
		src string
		dst string
	}{src, dst})
	fake.recordInvocation("CopyBlobWithinBucket", []interface{}{src, dst})
	fake.copyBlobWithinBucketMutex.Unlock()
	if fake.CopyBlobWithinBucketStub != nil {
		return fake.CopyBlobWithinBucketStub(src, dst)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobWithinBucketReturns.result1
}

func (fake *FakeBucket) CopyBlobWithinBucketCallCount() int {
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	return len(fake.copyBlobWithinBucketArgsForCall)
}

func (fake *FakeBucket) CopyBlobWithinBucketArgsForCall(i int) (string, string) {
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	return fake.copyBlobWithinBucketArgsForCall[i].src, fake.copyBlobWithinBucketArgsForCall[i].dst
}

func (fake *FakeBucket) CopyBlobWithinBucketReturns(result1 error) {
	fake.CopyBlobWithinBucketStub = nil
	fake.copyBlobWithinBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobWithinBucketReturnsOnCall(i int, result1 error) {
	fake.CopyBlobWithinBucketStub = nil
	if fake.copyBlobWithinBucketReturnsOnCall == nil {
		fake.copyBlobWithinBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobWithinBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobFromBucket(bucket incremental.Bucket, src string, dst string) error {
	fake.copyBlobFromBucketMutex.Lock()
	ret, specificReturn := fake.copyBlobFromBucketReturnsOnCall[len(fake.copyBlobFromBucketArgsForCall)]
	fake.copyBlobFromBucketArgsForCall = append(fake.copyBlobFromBucketArgsForCall, struct {
		bucket incremental.Bucket
		src    string
		dst    string
	}{bucket, src, dst})
	fake.recordInvocation("CopyBlobFromBucket", []interface{}{bucket, src, dst})
	fake.copyBlobFromBucketMutex.Unlock()
	if fake.CopyBlobFromBucketStub != nil {
		return fake.CopyBlobFromBucketStub(bucket, src, dst)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobFromBucketReturns.result1
}

func (fake *FakeBucket) CopyBlobFromBucketCallCount() int {
	fake.copyBlobFromBucketMutex.RLock()
	defer fake.copyBlobFromBucketMutex.RUnlock()
	return len(fake.copyBlobFromBucketArgsForCall)
}

func (fake *FakeBucket) CopyBlobFromBucketArgsForCall(i int) (incremental.Bucket, string, string) {
	fake.copyBlobFromBucketMutex.RLock()
	defer fake.copyBlobFromBucketMutex.RUnlock()
	return fake.copyBlobFromBucketArgsForCall[i].bucket, fake.copyBlobFromBucketArgsForCall[i].src, fake.copyBlobFromBucketArgsForCall[i].dst
}

func (fake *FakeBucket) CopyBlobFromBucketReturns(result1 error) {
	fake.CopyBlobFromBucketStub = nil
	fake.copyBlobFromBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobFromBucketReturnsOnCall(i int, result1 error) {
	fake.CopyBlobFromBucketStub = nil
	if fake.copyBlobFromBucketReturnsOnCall == nil {
		fake.copyBlobFromBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobFromBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) UploadBlob(path string, contents string) error {
	fake.uploadBlobMutex.Lock()
	ret, specificReturn := fake.uploadBlobReturnsOnCall[len(fake.uploadBlobArgsForCall)]
	fake.uploadBlobArgsForCall = append(fake.uploadBlobArgsForCall, struct {
		path     string
		contents string
	}{path, contents})
	fake.recordInvocation("UploadBlob", []interface{}{path, contents})
	fake.uploadBlobMutex.Unlock()
	if fake.UploadBlobStub != nil {
		return fake.UploadBlobStub(path, contents)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.uploadBlobReturns.result1
}

func (fake *FakeBucket) UploadBlobCallCount() int {
	fake.uploadBlobMutex.RLock()
	defer fake.uploadBlobMutex.RUnlock()
	return len(fake.uploadBlobArgsForCall)
}

func (fake *FakeBucket) UploadBlobArgsForCall(i int) (string, string) {
	fake.uploadBlobMutex.RLock()
	defer fake.uploadBlobMutex.RUnlock()
	return fake.uploadBlobArgsForCall[i].path, fake.uploadBlobArgsForCall[i].contents
}

func (fake *FakeBucket) UploadBlobReturns(result1 error) {
	fake.UploadBlobStub = nil
	fake.uploadBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) UploadBlobReturnsOnCall(i int, result1 error) {
	fake.UploadBlobStub = nil
	if fake.uploadBlobReturnsOnCall == nil {
		fake.uploadBlobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.uploadBlobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) HasBlob(path string) (bool, error) {
	fake.hasBlobMutex.Lock()
	ret, specificReturn := fake.hasBlobReturnsOnCall[len(fake.hasBlobArgsForCall)]
	fake.hasBlobArgsForCall = append(fake.hasBlobArgsForCall, struct {
		path string
	}{path})
	fake.recordInvocation("HasBlob", []interface{}{path})
	fake.hasBlobMutex.Unlock()
	if fake.HasBlobStub != nil {
		return fake.HasBlobStub(path)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.hasBlobReturns.result1, fake.hasBlobReturns.result2
}

func (fake *FakeBucket) HasBlobCallCount() int {
	fake.hasBlobMutex.RLock()
	defer fake.hasBlobMutex.RUnlock()
	return len(fake.hasBlobArgsForCall)
}

func (fake *FakeBucket) HasBlobArgsForCall(i int) string {
	fake.hasBlobMutex.RLock()
	defer fake.hasBlobMutex.RUnlock()
	return fake.hasBlobArgsForCall[i].path
}

func (fake *FakeBucket) HasBlobReturns(result1 bool, result2 error) {
	fake.HasBlobStub = nil
	fake.hasBlobReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) HasBlobReturnsOnCall(i int, result1 bool, result2 error) {
	fake.HasBlobStub = nil
	if fake.hasBlobReturnsOnCall == nil {
		fake.hasBlobReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.hasBlobReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	fake.listDirectoriesMutex.RLock()
	defer fake.listDirectoriesMutex.RUnlock()
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	fake.copyBlobFromBucketMutex.RLock()
	defer fake.copyBlobFromBucketMutex.RUnlock()
	fake.uploadBlobMutex.RLock()
	defer fake.uploadBlobMutex.RUnlock()
	fake.hasBlobMutex.RLock()
	defer fake.hasBlobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucket) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ incremental.Bucket = new(FakeBucket)