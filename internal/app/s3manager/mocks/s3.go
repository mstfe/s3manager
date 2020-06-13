// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mocks

import (
	"io"
	"sync"

	"github.com/minio/minio-go"
)

var (
	lockS3MockGetObject     sync.RWMutex
	lockS3MockListBuckets   sync.RWMutex
	lockS3MockListObjectsV2 sync.RWMutex
	lockS3MockMakeBucket    sync.RWMutex
	lockS3MockPutObject     sync.RWMutex
	lockS3MockRemoveBucket  sync.RWMutex
	lockS3MockRemoveObject  sync.RWMutex
)

// S3Mock is a mock implementation of S3.
//
//     func TestSomethingThatUsesS3(t *testing.T) {
//
//         // make and configure a mocked S3
//         mockedS3 := &S3Mock{
//             GetObjectFunc: func(bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
// 	               panic("TODO: mock out the GetObject method")
//             },
//             ListBucketsFunc: func() ([]minio.BucketInfo, error) {
// 	               panic("TODO: mock out the ListBuckets method")
//             },
//             ListObjectsV2Func: func(bucketName string, objectPrefix string, recursive bool, doneCh <-chan struct{}) <-chan minio.ObjectInfo {
// 	               panic("TODO: mock out the ListObjectsV2 method")
//             },
//             MakeBucketFunc: func(bucketName string, location string) error {
// 	               panic("TODO: mock out the MakeBucket method")
//             },
//             PutObjectFunc: func(bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (int64, error) {
// 	               panic("TODO: mock out the PutObject method")
//             },
//             RemoveBucketFunc: func(bucketName string) error {
// 	               panic("TODO: mock out the RemoveBucket method")
//             },
//             RemoveObjectFunc: func(bucketName string, objectName string) error {
// 	               panic("TODO: mock out the RemoveObject method")
//             },
//         }
//
//         // TODO: use mockedS3 in code that requires S3
//         //       and then make assertions.
//
//     }
type S3Mock struct {
	// GetObjectFunc mocks the GetObject method.
	GetObjectFunc func(bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error)

	// ListBucketsFunc mocks the ListBuckets method.
	ListBucketsFunc func() ([]minio.BucketInfo, error)

	// ListObjectsV2Func mocks the ListObjectsV2 method.
	ListObjectsV2Func func(bucketName string, objectPrefix string, recursive bool, doneCh <-chan struct{}) <-chan minio.ObjectInfo

	// MakeBucketFunc mocks the MakeBucket method.
	MakeBucketFunc func(bucketName string, location string) error

	// PutObjectFunc mocks the PutObject method.
	PutObjectFunc func(bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (int64, error)

	// RemoveBucketFunc mocks the RemoveBucket method.
	RemoveBucketFunc func(bucketName string) error

	// RemoveObjectFunc mocks the RemoveObject method.
	RemoveObjectFunc func(bucketName string, objectName string) error

	// calls tracks calls to the methods.
	calls struct {
		// GetObject holds details about calls to the GetObject method.
		GetObject []struct {
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Opts is the opts argument value.
			Opts minio.GetObjectOptions
		}
		// ListBuckets holds details about calls to the ListBuckets method.
		ListBuckets []struct {
		}
		// ListObjectsV2 holds details about calls to the ListObjectsV2 method.
		ListObjectsV2 []struct {
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectPrefix is the objectPrefix argument value.
			ObjectPrefix string
			// Recursive is the recursive argument value.
			Recursive bool
			// DoneCh is the doneCh argument value.
			DoneCh <-chan struct{}
		}
		// MakeBucket holds details about calls to the MakeBucket method.
		MakeBucket []struct {
			// BucketName is the bucketName argument value.
			BucketName string
			// Location is the location argument value.
			Location string
		}
		// PutObject holds details about calls to the PutObject method.
		PutObject []struct {
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
			// Reader is the reader argument value.
			Reader io.Reader
			// ObjectSize is the objectSize argument value.
			ObjectSize int64
			// Opts is the opts argument value.
			Opts minio.PutObjectOptions
		}
		// RemoveBucket holds details about calls to the RemoveBucket method.
		RemoveBucket []struct {
			// BucketName is the bucketName argument value.
			BucketName string
		}
		// RemoveObject holds details about calls to the RemoveObject method.
		RemoveObject []struct {
			// BucketName is the bucketName argument value.
			BucketName string
			// ObjectName is the objectName argument value.
			ObjectName string
		}
	}
}

// GetObject calls GetObjectFunc.
func (mock *S3Mock) GetObject(bucketName string, objectName string, opts minio.GetObjectOptions) (*minio.Object, error) {
	if mock.GetObjectFunc == nil {
		panic("S3Mock.GetObjectFunc: method is nil but S3.GetObject was just called")
	}
	callInfo := struct {
		BucketName string
		ObjectName string
		Opts       minio.GetObjectOptions
	}{
		BucketName: bucketName,
		ObjectName: objectName,
		Opts:       opts,
	}
	lockS3MockGetObject.Lock()
	mock.calls.GetObject = append(mock.calls.GetObject, callInfo)
	lockS3MockGetObject.Unlock()
	return mock.GetObjectFunc(bucketName, objectName, opts)
}

// GetObjectCalls gets all the calls that were made to GetObject.
// Check the length with:
//     len(mockedS3.GetObjectCalls())
func (mock *S3Mock) GetObjectCalls() []struct {
	BucketName string
	ObjectName string
	Opts       minio.GetObjectOptions
} {
	var calls []struct {
		BucketName string
		ObjectName string
		Opts       minio.GetObjectOptions
	}
	lockS3MockGetObject.RLock()
	calls = mock.calls.GetObject
	lockS3MockGetObject.RUnlock()
	return calls
}

// ListBuckets calls ListBucketsFunc.
func (mock *S3Mock) ListBuckets() ([]minio.BucketInfo, error) {
	if mock.ListBucketsFunc == nil {
		panic("S3Mock.ListBucketsFunc: method is nil but S3.ListBuckets was just called")
	}
	callInfo := struct {
	}{}
	lockS3MockListBuckets.Lock()
	mock.calls.ListBuckets = append(mock.calls.ListBuckets, callInfo)
	lockS3MockListBuckets.Unlock()
	return mock.ListBucketsFunc()
}

// ListBucketsCalls gets all the calls that were made to ListBuckets.
// Check the length with:
//     len(mockedS3.ListBucketsCalls())
func (mock *S3Mock) ListBucketsCalls() []struct {
} {
	var calls []struct {
	}
	lockS3MockListBuckets.RLock()
	calls = mock.calls.ListBuckets
	lockS3MockListBuckets.RUnlock()
	return calls
}

// ListObjectsV2 calls ListObjectsV2Func.
func (mock *S3Mock) ListObjectsV2(bucketName string, objectPrefix string, recursive bool, doneCh <-chan struct{}) <-chan minio.ObjectInfo {
	if mock.ListObjectsV2Func == nil {
		panic("S3Mock.ListObjectsV2Func: method is nil but S3.ListObjectsV2 was just called")
	}
	callInfo := struct {
		BucketName   string
		ObjectPrefix string
		Recursive    bool
		DoneCh       <-chan struct{}
	}{
		BucketName:   bucketName,
		ObjectPrefix: objectPrefix,
		Recursive:    recursive,
		DoneCh:       doneCh,
	}
	lockS3MockListObjectsV2.Lock()
	mock.calls.ListObjectsV2 = append(mock.calls.ListObjectsV2, callInfo)
	lockS3MockListObjectsV2.Unlock()
	return mock.ListObjectsV2Func(bucketName, objectPrefix, recursive, doneCh)
}

// ListObjectsV2Calls gets all the calls that were made to ListObjectsV2.
// Check the length with:
//     len(mockedS3.ListObjectsV2Calls())
func (mock *S3Mock) ListObjectsV2Calls() []struct {
	BucketName   string
	ObjectPrefix string
	Recursive    bool
	DoneCh       <-chan struct{}
} {
	var calls []struct {
		BucketName   string
		ObjectPrefix string
		Recursive    bool
		DoneCh       <-chan struct{}
	}
	lockS3MockListObjectsV2.RLock()
	calls = mock.calls.ListObjectsV2
	lockS3MockListObjectsV2.RUnlock()
	return calls
}

// MakeBucket calls MakeBucketFunc.
func (mock *S3Mock) MakeBucket(bucketName string, location string) error {
	if mock.MakeBucketFunc == nil {
		panic("S3Mock.MakeBucketFunc: method is nil but S3.MakeBucket was just called")
	}
	callInfo := struct {
		BucketName string
		Location   string
	}{
		BucketName: bucketName,
		Location:   location,
	}
	lockS3MockMakeBucket.Lock()
	mock.calls.MakeBucket = append(mock.calls.MakeBucket, callInfo)
	lockS3MockMakeBucket.Unlock()
	return mock.MakeBucketFunc(bucketName, location)
}

// MakeBucketCalls gets all the calls that were made to MakeBucket.
// Check the length with:
//     len(mockedS3.MakeBucketCalls())
func (mock *S3Mock) MakeBucketCalls() []struct {
	BucketName string
	Location   string
} {
	var calls []struct {
		BucketName string
		Location   string
	}
	lockS3MockMakeBucket.RLock()
	calls = mock.calls.MakeBucket
	lockS3MockMakeBucket.RUnlock()
	return calls
}

// PutObject calls PutObjectFunc.
func (mock *S3Mock) PutObject(bucketName string, objectName string, reader io.Reader, objectSize int64, opts minio.PutObjectOptions) (int64, error) {
	if mock.PutObjectFunc == nil {
		panic("S3Mock.PutObjectFunc: method is nil but S3.PutObject was just called")
	}
	callInfo := struct {
		BucketName string
		ObjectName string
		Reader     io.Reader
		ObjectSize int64
		Opts       minio.PutObjectOptions
	}{
		BucketName: bucketName,
		ObjectName: objectName,
		Reader:     reader,
		ObjectSize: objectSize,
		Opts:       opts,
	}
	lockS3MockPutObject.Lock()
	mock.calls.PutObject = append(mock.calls.PutObject, callInfo)
	lockS3MockPutObject.Unlock()
	return mock.PutObjectFunc(bucketName, objectName, reader, objectSize, opts)
}

// PutObjectCalls gets all the calls that were made to PutObject.
// Check the length with:
//     len(mockedS3.PutObjectCalls())
func (mock *S3Mock) PutObjectCalls() []struct {
	BucketName string
	ObjectName string
	Reader     io.Reader
	ObjectSize int64
	Opts       minio.PutObjectOptions
} {
	var calls []struct {
		BucketName string
		ObjectName string
		Reader     io.Reader
		ObjectSize int64
		Opts       minio.PutObjectOptions
	}
	lockS3MockPutObject.RLock()
	calls = mock.calls.PutObject
	lockS3MockPutObject.RUnlock()
	return calls
}

// RemoveBucket calls RemoveBucketFunc.
func (mock *S3Mock) RemoveBucket(bucketName string) error {
	if mock.RemoveBucketFunc == nil {
		panic("S3Mock.RemoveBucketFunc: method is nil but S3.RemoveBucket was just called")
	}
	callInfo := struct {
		BucketName string
	}{
		BucketName: bucketName,
	}
	lockS3MockRemoveBucket.Lock()
	mock.calls.RemoveBucket = append(mock.calls.RemoveBucket, callInfo)
	lockS3MockRemoveBucket.Unlock()
	return mock.RemoveBucketFunc(bucketName)
}

// RemoveBucketCalls gets all the calls that were made to RemoveBucket.
// Check the length with:
//     len(mockedS3.RemoveBucketCalls())
func (mock *S3Mock) RemoveBucketCalls() []struct {
	BucketName string
} {
	var calls []struct {
		BucketName string
	}
	lockS3MockRemoveBucket.RLock()
	calls = mock.calls.RemoveBucket
	lockS3MockRemoveBucket.RUnlock()
	return calls
}

// RemoveObject calls RemoveObjectFunc.
func (mock *S3Mock) RemoveObject(bucketName string, objectName string) error {
	if mock.RemoveObjectFunc == nil {
		panic("S3Mock.RemoveObjectFunc: method is nil but S3.RemoveObject was just called")
	}
	callInfo := struct {
		BucketName string
		ObjectName string
	}{
		BucketName: bucketName,
		ObjectName: objectName,
	}
	lockS3MockRemoveObject.Lock()
	mock.calls.RemoveObject = append(mock.calls.RemoveObject, callInfo)
	lockS3MockRemoveObject.Unlock()
	return mock.RemoveObjectFunc(bucketName, objectName)
}

// RemoveObjectCalls gets all the calls that were made to RemoveObject.
// Check the length with:
//     len(mockedS3.RemoveObjectCalls())
func (mock *S3Mock) RemoveObjectCalls() []struct {
	BucketName string
	ObjectName string
} {
	var calls []struct {
		BucketName string
		ObjectName string
	}
	lockS3MockRemoveObject.RLock()
	calls = mock.calls.RemoveObject
	lockS3MockRemoveObject.RUnlock()
	return calls
}
