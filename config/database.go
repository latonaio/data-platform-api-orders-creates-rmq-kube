package config

import (
	"fmt"
	"os"
)

type Database struct {
	user     string
	password string
	dbName   string
	address  string
	port     string
}

func newDatabase() *Database {
	return &Database{
		user:     os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		dbName:   os.Getenv("DB_NAME"),
		address:  os.Getenv("DATA_PLATFORM_MASTERS_AND_TRANSACTIONS_MYSQL_KUBE"),
		port:     os.Getenv("MYSQL_PORT"),
	}
}
func (c Database) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		c.user, c.password, c.address, c.port, c.dbName,
	)
}
