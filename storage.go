package jidoco

import "strings"
import driver "go.etcd.io/bbolt"

// Path is the path of a bucket.
type Path string

// PathSeparator is used to separate parts of paths to buckets.
const PathSeparator = "/"

func (p Path) Parts() []string {
	return strings.Split(string(p), PathSeparator)
}

// Bucket is a wrapper around the driver's bucket.
type Bucket struct {
	*driver.Bucket
}

// Tx is a wrapper around the driver's transactions.
type Tx struct {
	*driver.Tx
}

// Storage is a wrapper around the driver's database or storage engine
type Storage struct {
	*driver.DB
}

// Iterator is a wrapper around the driver's cursor
type Iterator struct {
	*driver.Cursor
}
