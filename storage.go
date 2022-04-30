package jidoco

import "strings"

/*
func (b *Bucket) Cursor() *Cursor
func (b *Bucket) Delete(key []byte) error
func (b *Bucket) DeleteBucket(key []byte) error
func (b *Bucket) ForEach(fn func(k, v []byte) error) error
func (b *Bucket) Get(key []byte) []byte
func (b *Bucket) NextSequence() (uint64, error)
func (b *Bucket) Put(key []byte, value []byte) error
func (b *Bucket) Root() pgidfunc
func (b *Bucket) Sequence() uint64
func (b *Bucket) SetSequence(v uint64) errorf
func (b *Bucket) Stats() BucketStats
func (b *Bucket) Tx() *Tx
func (b *Bucket) Writable() bool
*/

// Bucket is an interface for low level groups of documents.
// Jidoco assumes that the underlying storage supports grouping
// documents together in buckets and sub-buckets, much like
// folders with files in them. If the underlying storage does not support
// this directly this can be implemented, for example, by using prefixed keys
// or by physically splitting the files used by the data store.
type Bucket interface {
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

const PathSeparator = "/"

// Path is the path where a bucket can be found. A / separated path is used
// for this. The top level path / signifies the root bucket, which may be
// the top level data store itself.
type Path string

func (p Path) Parts() []string {
	return strings.Split(string(p), PathSeparator)
}

// Storage in an interface to abstract the low level data storage used
// for storing JSON documents in.
type Storage interface {
	// Bucket returns a bucket or sub bucket. Returns nil if the bucket doesn't
	// exist. The path / should always exist and return a usable bucket.
	Bucket(name Path) Bucket
	// CreateBucket creates a bucket or sub bucket. It may return an error if the
	// bucket exists.
	CreateBucket(name Path) (Bucket, error)
	// CreateBucketIfNotExists creates a bucket or sub bucket. If the bucket
	// exists, it is returned without error.
	CreateBucketIfNotExists(name Path) (Bucket, error)

	// ForEachBucket executes a function for each bucket in the root.
	// If the provided function returns an error then the iteration is stopped
	// and the error is returned to the caller.
	ForEachBucket(fn func(p Path, b Bucket) error) error
}
