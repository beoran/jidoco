package jidoco

// Database is an abstracted database in which there are several collections
// each which may have several documents in them, stored in an underlying
// storage system.

type Database struct {
	Storage     `json:"-"`
	Collections map[string]Collection `json:"collections"`
}
