package database

import "github.com/shrikar007/01-mongo-example/types"

type Database interface {
	Open() error
	Close() error
	Save(req types.Request) error
	Retrieve() ([]types.Request, error)
	//Delete(Id string) error
}

