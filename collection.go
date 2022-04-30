package jidoco

// Codec is used to encode and decode data stored in and loaded from
// a collection
type Codec interface {
	Encode(decoded interface{}, options ...interface{}) ([]string, error)
	Decode(decoded interface{}, encoded []byte) error
}

// Collections is a group of similar documents stored together.
// Normally these are stored as JSON documents, but the data may also be
// encrypted, or binary for attachments.
//
// All normal collections use a ULID as their key and a JSON document
// as their value. ULID makes creation timestamps redundant.
//
// Attachments use a ULID for their key and the binary data prefixed with
// name and mime string as their value.
//
// Indexes use the indexed value of one or more documents as their key,
// and the binary ulids of the documents that match that index value
// as their value. Since binary ULID is fixed length this simplifies
// index lookup.
type Collection struct {
	// Bucket for underlying storage.
	Bucket `json:"-"`
	// Path to the bucket this collection is stored in.
	Path string `json:"path"`
	// Codec to use to store and retrieve the data.
	Codec
	// Name of codec to use to store and retrieve the data.
	CodecName string `json:"codec_name"`
	// Indexes on this collection.
	Indexes map[string]Index `json:"indexes"`
	// Attachment sub-collection for this collection.
	*Attachment `json:"attachment,omitempty"`
}
