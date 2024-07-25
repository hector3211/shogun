package shogun

type Driver string

const (
	POSTGRES Driver = "postgres"
	SQLITE   Driver = "sqlite3"
)

var DefaultDriver = SQLITE
