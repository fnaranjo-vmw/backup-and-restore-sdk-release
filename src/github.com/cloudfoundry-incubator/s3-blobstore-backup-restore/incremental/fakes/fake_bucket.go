// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/incremental"
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
	ListBlobsStub        func() ([]incremental.Blob, error)
	listBlobsMutex       sync.RWMutex
	listBlobsArgsForCall []struct{}
	listBlobsReturns     struct {
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
	CopyBlobToBucketStub        func(bucket incremental.Bucket, src, dst string) error
	copyBlobToBucketMutex       sync.RWMutex
	copyBlobToBucketArgsForCall []struct {
		bucket incremental.Bucket
		src    string
		dst    string
	}
	copyBlobToBucketReturns struct {
		result1 error
	}
	copyBlobToBucketReturnsOnCall map[int]struct {
		result1 error
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

func (fake *FakeBucket) ListBlobs() ([]incremental.Blob, error) {
	fake.listBlobsMutex.Lock()
	ret, specificReturn := fake.listBlobsReturnsOnCall[len(fake.listBlobsArgsForCall)]
	fake.listBlobsArgsForCall = append(fake.listBlobsArgsForCall, struct{}{})
	fake.recordInvocation("ListBlobs", []interface{}{})
	fake.listBlobsMutex.Unlock()
	if fake.ListBlobsStub != nil {
		return fake.ListBlobsStub()
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

func (fake *FakeBucket) CopyBlobToBucket(bucket incremental.Bucket, src string, dst string) error {
	fake.copyBlobToBucketMutex.Lock()
	ret, specificReturn := fake.copyBlobToBucketReturnsOnCall[len(fake.copyBlobToBucketArgsForCall)]
	fake.copyBlobToBucketArgsForCall = append(fake.copyBlobToBucketArgsForCall, struct {
		bucket incremental.Bucket
		src    string
		dst    string
	}{bucket, src, dst})
	fake.recordInvocation("CopyBlobToBucket", []interface{}{bucket, src, dst})
	fake.copyBlobToBucketMutex.Unlock()
	if fake.CopyBlobToBucketStub != nil {
		return fake.CopyBlobToBucketStub(bucket, src, dst)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.copyBlobToBucketReturns.result1
}

func (fake *FakeBucket) CopyBlobToBucketCallCount() int {
	fake.copyBlobToBucketMutex.RLock()
	defer fake.copyBlobToBucketMutex.RUnlock()
	return len(fake.copyBlobToBucketArgsForCall)
}

func (fake *FakeBucket) CopyBlobToBucketArgsForCall(i int) (incremental.Bucket, string, string) {
	fake.copyBlobToBucketMutex.RLock()
	defer fake.copyBlobToBucketMutex.RUnlock()
	return fake.copyBlobToBucketArgsForCall[i].bucket, fake.copyBlobToBucketArgsForCall[i].src, fake.copyBlobToBucketArgsForCall[i].dst
}

func (fake *FakeBucket) CopyBlobToBucketReturns(result1 error) {
	fake.CopyBlobToBucketStub = nil
	fake.copyBlobToBucketReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyBlobToBucketReturnsOnCall(i int, result1 error) {
	fake.CopyBlobToBucketStub = nil
	if fake.copyBlobToBucketReturnsOnCall == nil {
		fake.copyBlobToBucketReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyBlobToBucketReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.listBlobsMutex.RLock()
	defer fake.listBlobsMutex.RUnlock()
	fake.listDirectoriesMutex.RLock()
	defer fake.listDirectoriesMutex.RUnlock()
	fake.copyBlobWithinBucketMutex.RLock()
	defer fake.copyBlobWithinBucketMutex.RUnlock()
	fake.copyBlobToBucketMutex.RLock()
	defer fake.copyBlobToBucketMutex.RUnlock()
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