package config

// Config contains basic config of any server
type Config struct {
	ServiceName    string
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
	LoggerConfig   LoggerConfig
}

// ServerConfig has only server specific configuration
type ServerConfig struct {
	Port string
}

// DatabaseConfig has database related configuration.
type DatabaseConfig struct {
	DbType   string
	Host     string
	DbName   string
	UserName string
	Password string
}

// LoggerConfig has logger related configuration.
type LoggerConfig struct {
	LogFilePath string
}

// SaveDefaultConfigParams func save default params if not set already.
func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
	if o.DatabaseConfig.Host == "" {
		o.DatabaseConfig.Host = "localhost"
	}
	if o.DatabaseConfig.DbName == "" {
		o.DatabaseConfig.DbName = "service"
	}
}
