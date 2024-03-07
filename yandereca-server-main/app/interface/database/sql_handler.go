package database

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
}

type Result interface {
}

type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
