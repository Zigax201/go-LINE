package config

type Configuration struct {
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

type ServerConfiguration struct {
	Host string
	Port int
}

type DatabaseConfiguration struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBIPAddr   string
	DBDialect  string
}
