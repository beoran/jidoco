package storage

import "github.com/beoran/jidoco/interface/cursor"

type Cursor struct {
	cursor.Cursor
}

// Bucket is an interface for low level groups of documents.
// Jidoco assumes that the underlying storage supports grouping
// documents together in buckets and sub-buckets, much like
// folders with files in them. If the underlying storage does not support
// this directly this can be implemented, for example, by using prefixed keys
// or by physically splitting the files used by the data store.
type Bucketer interface {
	// The Cursor function must return a Cursor for iterating over the bucket.
	Cursor() Cursor

	// Delete removes a key from the bucket. If the key does not exist then
	// nothingis done and a nil error is returned. May return an error if the
	// underlying datastore did not allow writing.
	Delete(key []byte) error

	// ForEach executes a function for each key/value pair in a bucket. If the
	// provided function returns an error then the iteration is stopped and the
	// error is returned to the caller. The provided function must not modify
	// the bucket. If v is nil it is a sub-bucket.
	ForEach(fn func(k, v []byte) error) error

	// Get retrieves the value for a key in the bucket. Returns a nil value
	// if thekey does not exist or if the key is a nested bucket.
	// The returned value must be copied out.
	Get(key []byte) []byte

	// Put sets the value for a key in the bucket. If the key exist then its
	// previous value will be overwritten. May return an error if the
	// underlying datastore did not allow writing.
	Put(key []byte, value []byte) error
}

type Bucket struct {
	Bucketer
}

type Pather interface {
	Parts() []string
	Name() string
}

type Path struct {
	Pather
}

// Tx is an inteface to abstract the low level transactions used by the low
// level storage the JSON documents are stored in.
// Transactions themselves are not wholly abstracted because of the separation
// between reading and writing.
type Txer interface {
	// Bucket returns a bucket or sub bucket. Returns nil if the bucket doesn't
	// exist. The path / should always exist and return a usable bucket.
	Bucket(name Path) Bucket
	// CreateBucket creates a bucket or sub bucket. It may return an error if the
	// bucket exists. Only for write transactions.
	CreateBucket(name Path) (Bucket, error)
	// CreateBucketIfNotExists creates a bucket or sub bucket. If the bucket
	// exists, it is returned without error. Only for write transactions.
	CreateBucketIfNotExists(name Path) (Bucket, error)
	// ForEachBucket executes a function for each bucket in the root.
	// If the provided function returns an error then the iteration is stopped
	// and the error is returned to the caller.
	ForEachBucket(fn func(p Path, b Bucket) error) error
}

type Tx struct {
	Txer
}

// Storage in an interface to abstract the low level data storage used
// for storing JSON documents in.
type Storager interface {
	// View executes a transaction in read only mode
	View(func(t Tx) error) error
	// Update executes a transaction in read write mode
	Update(func(t Tx) error) error
	// Batch executes a tranaction in batch read write mode
	Batch(func(t Tx) error) error
	// Close closes the storage
	Close() error
}

type Storage struct {
	Storager
}
