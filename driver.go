package jidoco

type Driver interface {
	Open(storageName string) (Storage, error)
}
