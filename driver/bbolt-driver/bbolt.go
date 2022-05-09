// Package bbolt-driver implements a bbolt driver for jidoco
package bboltdriver

import "strings"
import "go.etcd.io/bbolt"

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

type Path string

const PathSeparator = "/"

func (p Path) Parts() []string {
	return strings.Split(string(p), PathSeparator)
}

func (p Path) Name() string {
	return string(p)
}

type Driver struct {
}

type Bucket struct {
	*bbolt.Tx
	*bbolt.Bucket
}

type Storage struct {
	*bbolt.DB
}

type Cursor struct {
}

type Tx struct {
	*bbolt.Tx
}

// Bucket returns a bucket or sub bucket. Returns nil if the bucket doesn't
// exist. The path / should always exist and return a usable bucket.
func (s *Storage) Bucket(name Path) Bucket {

}

// CreateBucket creates a bucket or sub bucket. It may return an error if the
// bucket exists.
func (s *Storage) CreateBucket(name Path) (Bucket, error) {

}

// CreateBucketIfNotExists creates a bucket or sub bucket. If the bucket
// exists, it is returned without error.
func (s *Storage) CreateBucketIfNotExists(name Path) (*Bucket, error) {

}

// ForEachBucket executes a function for each bucket in the root.
// If the provided function returns an error then the iteration is stopped
// and the error is returned to the caller.
func (s *Storage) ForEachBucket(fn func(p Path, b *Bucket) error) error {

}

// Close closes the storage
func (s *Storage) Close() error {
	if s.DB != nil {
		res := s.DB.Close()
		s.DB = nil
		return res
	}
	return nil
}

func (Driver) Open(storageName string) (*Storage, error) {
	var err error
	res := &Storage{}
	res.DB, err = bbolt.Open(storageName+".bbolt", 0o600, nil)
	return res, err
}

func New() Driver {
	return Driver{}
}
