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
	c.DbUser = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	c.DbName = "article_arena_database"
	c.DbPass = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	c.DbServiceName = os.Getenv("SERVICE_NAME")
	c.DbHostPort = "27017"
}
