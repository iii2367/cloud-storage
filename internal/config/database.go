package config

type Database struct {
	Host		string
	Port		string
	User		string
	Password	string
	DBName		string
}

func MustLoadDatabase(prefix string) *Database {
	return &Database {
		Host: mustEnv(prefix + "_DATABASE_HOST"),
		Port: mustEnv(prefix + "_DATABASE_PORT"),
		User: mustEnv(prefix + "_DATABASE_USER"),
		Password: mustEnv(prefix + "_DATABASE_PASSWORD"),
		DBName: mustEnv(prefix + "_DATABASE_DBNAME"),
	}
}
