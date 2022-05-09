package cursor

// Cursor is the interface that the underlying data store has to support
// for use with Jidoco. This not coincidentally matches the bbolt Cursor
// struct. It is an interface which should allow iterating or seeking
// documents in the underlying document store.
type Cursor interface {
	// Next moves the cursor to the next item in the collection and returns its
	// key andvalue. If the cursor is at the end of the bucket then a nil key
	// and valueshould be returned.
	Next() (key []byte, value []byte)
	// Prev moves the cursor to the previous item in the collection and returns
	// its keyand value. If the cursor is at the beginning of the collection
	// then a nil keyand value should be returned.
	Prev() (key []byte, value []byte)
	// Seek moves the cursor to a given key and returns it. If the key does not
	// exist then the next key is used. If no keys follow, a nil key is returned.
	Seek(seek []byte) (key []byte, value []byte)
}
