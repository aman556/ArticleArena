package config

import "os"

type Config struct {
	DbHostPort    string `yaml:"DbHostPort"`
	DbName        string `yaml:"DbName"`
	DbPass        string `yaml:"DbPass"`
	DbUser        string `yaml:"DbUser"`
	DbServiceName string `yaml:"DbServiceName"`
}

func NewConfig() Config {
	var config Config
	config.GetEnv()

	return config
}

func (c *Config) GetEnv() {
	c.DbUser = os.Getenv("CONFIG_DBUSER")
	c.DbName = os.Getenv("CONFIG_DBNAME")
	c.DbPass = os.Getenv("MYSQL_ROOT_PASSWORD")
	c.DbServiceName = os.Getenv("SERVICE_NAME")
	c.DbHostPort = os.Getenv("CONFIG_DBHOST_PORT")
}
