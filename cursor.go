package jidoco

import "encoding/binary"
import "sort"
import "github.com/beoran/jidoco/interface/cursor"

type Cursor = cursor.Cursor

// ArrayCursor implements the cursor interface using a backing array.
// The keys are little-endian binary encoded uint64 values of the array index.
// This type is probably most useful fot testing and simulating a key/value
// store.
type ArrayCursor struct {
	values [][]byte
	index  int
}

func NewArrayCursor(values [][]byte) *ArrayCursor {
	return &ArrayCursor{values, 0}
}

func (c *ArrayCursor) First() (key []byte, value []byte) {
	c.index = 0
	return c.This()
}

func (c ArrayCursor) This() (key []byte, value []byte) {
	if c.index >= len(c.values) {
		return nil, nil
	}
	key = make([]byte, 4)
	binary.LittleEndian.PutUint64(key, uint64(c.index))
	return key, c.values[c.index]
}

func (c *ArrayCursor) Next() (key []byte, value []byte) {
	if c.index < len(c.values) {
		c.index++
	}
	return c.This()
}

func (c *ArrayCursor) Prev() (key []byte, value []byte) {
	if c.index > 0 {
		c.index--
	}
	return c.This()
}

func (c *ArrayCursor) Seek(seek []byte) (key []byte, value []byte) {
	i := int(binary.LittleEndian.Uint64(seek))
	if i < len(c.values) {
		c.index = i
	}
	return c.This()
}

// MapCursor implements the cursor infterface using a backing map.
// The keys are strings, the values binary.
// This type is probably most useful for testing and simulating a key/value
// store.
type MapCursor struct {
	values map[string][]byte
	keys   []string
	index  int
}

func keysOfStringMap(values map[string][]byte) []string {
	keys := []string{}
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

func NewMapCursor(values map[string][]byte) *MapCursor {
	return &MapCursor{values, keysOfStringMap(values), 0}
}

func (c MapCursor) Get(key []byte) []byte {
	value, ok := c.values[string(key)]
	if ok {
		return value
	} else {
		return nil
	}
}

func (c *MapCursor) Put(key, value []byte) {
	c.values[string(key)] = value
	c.keys = keysOfStringMap(c.values)
}

func (c MapCursor) This() ([]byte, []byte) {
	if c.index >= len(c.values) {
		return nil, nil
	}
	key := c.keys[c.index]
	value := c.Get([]byte(key))
	return []byte(key), value
}

func (c *MapCursor) First() (key []byte, value []byte) {
	c.index = 0
	return c.This()
}

func (c *MapCursor) Next() (key []byte, value []byte) {
	if c.index < len(c.values) {
		c.index++
	}
	return c.This()
}

func (c *MapCursor) Prev() (key []byte, value []byte) {
	if c.index > 0 {
		c.index--
	}
	return c.This()
}

func (c *MapCursor) Seek(seek []byte) (key []byte, value []byte) {
	i := sort.SearchStrings(c.keys, string(seek))
	if i < len(c.values) {
		c.index = i
	}
	return c.This()
}

// MatchFunc returns true if the key and value "match" and should be kept,
// or false if they are not needed.
type MatchFunc func(key, value []byte) bool

// MatchCursor applies a boolean match function on an underlying cursor
// and only iterates over or seeks values that match.
type MatchCursor struct {
	underlying Cursor
	MatchFunc
}

func NewMatchCursor(underlying Cursor, f MatchFunc) *MatchCursor {
	return &MatchCursor{underlying, f}
}

func (c *MatchCursor) Next() ([]byte, []byte) {
	var k, v []byte
	// skip non-matching key/value pairs
	for k, v = c.underlying.Next(); k != nil && !c.MatchFunc(k, v); k, v = c.underlying.Next() {
	}
	return k, v
}

func (c *MatchCursor) Prev() (key []byte, value []byte) {
	var k, v []byte
	// skip non-matching key/value pairs
	for k, v = c.underlying.Prev(); k != nil && !c.MatchFunc(k, v); k, v = c.underlying.Prev() {
	}
	return k, v
}

func (c *MatchCursor) Seek(seek []byte) (key []byte, value []byte) {
	var k, v []byte
	for k, v := c.underlying.Seek(seek); k != nil && !c.MatchFunc(k, v); k, v = c.underlying.Next() {
	}
	return k, v

}

// SubstituteFunc returns the value that should be used in stead of the
// original value. This may also be nil. The key is not subsituted.
// This may be useful, for instamcem for decrypting encrypted documents.
type SubstituteFunc func(key, value []byte) []byte

// SubstituteCursor applies a substitution filter on an underlying cursor
// and iterates over or seeks values with that substitution applied.
// This could be useful for instance for only returning a part of the document.
type SubstituteCursor struct {
	underlying Cursor
	SubstituteFunc
}

func (c *SubstituteCursor) Next() (key []byte, value []byte) {
	k, v := c.underlying.Next()
	return k, c.SubstituteFunc(k, v)
}

func (c *SubstituteCursor) Prev() (key []byte, value []byte) {
	k, v := c.underlying.Prev()
	return k, c.SubstituteFunc(k, v)
}

func (c *SubstituteCursor) Seek(seek []byte) (key []byte, value []byte) {
	k, v := c.underlying.Seek(seek)
	return k, c.SubstituteFunc(k, v)
}

var _ Cursor = &ArrayCursor{}
var _ Cursor = &MapCursor{}
var _ Cursor = &MatchCursor{}
var _ Cursor = &SubstituteCursor{}
var _ Cursor = &Iterator{}
