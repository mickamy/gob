package database

type Database interface {
	Name() string
	Create() error
	Exists() (bool, error)
}
