package models

type psqlConfig struct {
	DriverName string `yaml:"DriverName" default:"postgres"`
	Host       string `yaml:"Host" default:"localhost"`
	Port       string `yaml:"Port" default:"5432"`
	User       string `yaml:"User" default:"postgres"`
	Password   string `yaml:"Password" default:"postgres"`
	DbName     string `yaml:"DbName" default:"postgres"`
	TableName  string `yaml:"TableName" default:"users"`
}

type dbConfig struct {
	ConnStr string `yaml:"ConnStr"`
}

type serverConfig struct {
	Host string `yaml:"Host" default:"localhost"`
	Port string `yaml:"Port" default:"8080"`
}

type AppConfig struct {
	Psql     psqlConfig
	DataBase dbConfig
	Server   serverConfig
}
