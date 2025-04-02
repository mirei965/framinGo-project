package framingo

import "database/sql"

type initPaths struct {
	rootPath    string
	folderNames []string
}

// cookieConfig holds cookie config values
type cookieConfig struct {
	name     string
	lifetime string
	persist  string
	secure   string
	domain   string
}

type databaseConfig struct {
	dsn string
	database string
}

type Database struct {
	DataType string
	Pool *sql.DB
}
type redisConfig struct {
	Host string
	Password string
	prefix string
}
