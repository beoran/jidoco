package jidoco

// An index is a sub-collection of a collection where indices are stored
// to speed up looking up data in a collection.
type Index struct {
	Collection `json:"collection"`
	Unique     bool `json:unique`
	Fulltext   bool `json:fulltext`
}
