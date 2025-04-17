package db

type Database interface {
	Name() string
	Create() error
	Exists() (bool, error)
}
