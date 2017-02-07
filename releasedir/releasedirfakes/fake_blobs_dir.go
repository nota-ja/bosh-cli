// This file was generated by counterfeiter
package releasedirfakes

import (
	"io"
	"sync"

	"github.com/cloudfoundry/bosh-cli/releasedir"
)

type FakeBlobsDir struct {
	InitStub        func() error
	initMutex       sync.RWMutex
	initArgsForCall []struct{}
	initReturns     struct {
		result1 error
	}
	BlobsStub        func() ([]releasedir.Blob, error)
	blobsMutex       sync.RWMutex
	blobsArgsForCall []struct{}
	blobsReturns     struct {
		result1 []releasedir.Blob
		result2 error
	}
	SyncBlobsStub        func(numOfParallelWorkers int) error
	syncBlobsMutex       sync.RWMutex
	syncBlobsArgsForCall []struct {
		numOfParallelWorkers int
	}
	syncBlobsReturns struct {
		result1 error
	}
	UploadBlobsStub        func() error
	uploadBlobsMutex       sync.RWMutex
	uploadBlobsArgsForCall []struct{}
	uploadBlobsReturns     struct {
		result1 error
	}
	TrackBlobStub        func(string, io.ReadCloser) (releasedir.Blob, error)
	trackBlobMutex       sync.RWMutex
	trackBlobArgsForCall []struct {
		arg1 string
		arg2 io.ReadCloser
	}
	trackBlobReturns struct {
		result1 releasedir.Blob
		result2 error
	}
	UntrackBlobStub        func(string) error
	untrackBlobMutex       sync.RWMutex
	untrackBlobArgsForCall []struct {
		arg1 string
	}
	untrackBlobReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlobsDir) Init() error {
	fake.initMutex.Lock()
	fake.initArgsForCall = append(fake.initArgsForCall, struct{}{})
	fake.recordInvocation("Init", []interface{}{})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub()
	}
	return fake.initReturns.result1
}

func (fake *FakeBlobsDir) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeBlobsDir) InitReturns(result1 error) {
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) Blobs() ([]releasedir.Blob, error) {
	fake.blobsMutex.Lock()
	fake.blobsArgsForCall = append(fake.blobsArgsForCall, struct{}{})
	fake.recordInvocation("Blobs", []interface{}{})
	fake.blobsMutex.Unlock()
	if fake.BlobsStub != nil {
		return fake.BlobsStub()
	}
	return fake.blobsReturns.result1, fake.blobsReturns.result2
}

func (fake *FakeBlobsDir) BlobsCallCount() int {
	fake.blobsMutex.RLock()
	defer fake.blobsMutex.RUnlock()
	return len(fake.blobsArgsForCall)
}

func (fake *FakeBlobsDir) BlobsReturns(result1 []releasedir.Blob, result2 error) {
	fake.BlobsStub = nil
	fake.blobsReturns = struct {
		result1 []releasedir.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobsDir) SyncBlobs(numOfParallelWorkers int) error {
	fake.syncBlobsMutex.Lock()
	fake.syncBlobsArgsForCall = append(fake.syncBlobsArgsForCall, struct {
		numOfParallelWorkers int
	}{numOfParallelWorkers})
	fake.recordInvocation("SyncBlobs", []interface{}{numOfParallelWorkers})
	fake.syncBlobsMutex.Unlock()
	if fake.SyncBlobsStub != nil {
		return fake.SyncBlobsStub(numOfParallelWorkers)
	}
	return fake.syncBlobsReturns.result1
}

func (fake *FakeBlobsDir) SyncBlobsCallCount() int {
	fake.syncBlobsMutex.RLock()
	defer fake.syncBlobsMutex.RUnlock()
	return len(fake.syncBlobsArgsForCall)
}

func (fake *FakeBlobsDir) SyncBlobsArgsForCall(i int) int {
	fake.syncBlobsMutex.RLock()
	defer fake.syncBlobsMutex.RUnlock()
	return fake.syncBlobsArgsForCall[i].numOfParallelWorkers
}

func (fake *FakeBlobsDir) SyncBlobsReturns(result1 error) {
	fake.SyncBlobsStub = nil
	fake.syncBlobsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) UploadBlobs() error {
	fake.uploadBlobsMutex.Lock()
	fake.uploadBlobsArgsForCall = append(fake.uploadBlobsArgsForCall, struct{}{})
	fake.recordInvocation("UploadBlobs", []interface{}{})
	fake.uploadBlobsMutex.Unlock()
	if fake.UploadBlobsStub != nil {
		return fake.UploadBlobsStub()
	}
	return fake.uploadBlobsReturns.result1
}

func (fake *FakeBlobsDir) UploadBlobsCallCount() int {
	fake.uploadBlobsMutex.RLock()
	defer fake.uploadBlobsMutex.RUnlock()
	return len(fake.uploadBlobsArgsForCall)
}

func (fake *FakeBlobsDir) UploadBlobsReturns(result1 error) {
	fake.UploadBlobsStub = nil
	fake.uploadBlobsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) TrackBlob(arg1 string, arg2 io.ReadCloser) (releasedir.Blob, error) {
	fake.trackBlobMutex.Lock()
	fake.trackBlobArgsForCall = append(fake.trackBlobArgsForCall, struct {
		arg1 string
		arg2 io.ReadCloser
	}{arg1, arg2})
	fake.recordInvocation("TrackBlob", []interface{}{arg1, arg2})
	fake.trackBlobMutex.Unlock()
	if fake.TrackBlobStub != nil {
		return fake.TrackBlobStub(arg1, arg2)
	}
	return fake.trackBlobReturns.result1, fake.trackBlobReturns.result2
}

func (fake *FakeBlobsDir) TrackBlobCallCount() int {
	fake.trackBlobMutex.RLock()
	defer fake.trackBlobMutex.RUnlock()
	return len(fake.trackBlobArgsForCall)
}

func (fake *FakeBlobsDir) TrackBlobArgsForCall(i int) (string, io.ReadCloser) {
	fake.trackBlobMutex.RLock()
	defer fake.trackBlobMutex.RUnlock()
	return fake.trackBlobArgsForCall[i].arg1, fake.trackBlobArgsForCall[i].arg2
}

func (fake *FakeBlobsDir) TrackBlobReturns(result1 releasedir.Blob, result2 error) {
	fake.TrackBlobStub = nil
	fake.trackBlobReturns = struct {
		result1 releasedir.Blob
		result2 error
	}{result1, result2}
}

func (fake *FakeBlobsDir) UntrackBlob(arg1 string) error {
	fake.untrackBlobMutex.Lock()
	fake.untrackBlobArgsForCall = append(fake.untrackBlobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("UntrackBlob", []interface{}{arg1})
	fake.untrackBlobMutex.Unlock()
	if fake.UntrackBlobStub != nil {
		return fake.UntrackBlobStub(arg1)
	}
	return fake.untrackBlobReturns.result1
}

func (fake *FakeBlobsDir) UntrackBlobCallCount() int {
	fake.untrackBlobMutex.RLock()
	defer fake.untrackBlobMutex.RUnlock()
	return len(fake.untrackBlobArgsForCall)
}

func (fake *FakeBlobsDir) UntrackBlobArgsForCall(i int) string {
	fake.untrackBlobMutex.RLock()
	defer fake.untrackBlobMutex.RUnlock()
	return fake.untrackBlobArgsForCall[i].arg1
}

func (fake *FakeBlobsDir) UntrackBlobReturns(result1 error) {
	fake.UntrackBlobStub = nil
	fake.untrackBlobReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBlobsDir) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.blobsMutex.RLock()
	defer fake.blobsMutex.RUnlock()
	fake.syncBlobsMutex.RLock()
	defer fake.syncBlobsMutex.RUnlock()
	fake.uploadBlobsMutex.RLock()
	defer fake.uploadBlobsMutex.RUnlock()
	fake.trackBlobMutex.RLock()
	defer fake.trackBlobMutex.RUnlock()
	fake.untrackBlobMutex.RLock()
	defer fake.untrackBlobMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBlobsDir) recordInvocation(key string, args []interface{}) {
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

var _ releasedir.BlobsDir = new(FakeBlobsDir)
