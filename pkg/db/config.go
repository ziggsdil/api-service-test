package db

type Config struct {
	Database string `config:"POSTGRES_DB"`
	User     string `config:"POSTGRES_USER"`
	Password string `config:"POSTGRES_PASSWORD"`
	Host     string `config:"POSTGRES_HOST"`
	Port     int    `config:"POSTGRES_PORT"`
}
