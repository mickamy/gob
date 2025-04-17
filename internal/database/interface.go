package database

type Database interface {
	Name() string
	Exists() (bool, error)
	Create() error
	Drop() error
}
