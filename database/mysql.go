package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/xerrors"
)

type SetupData interface {
	DSN() string
}

type Mysql struct {
	dsn string
	*sql.DB
}

func NewMySQL(d SetupData) (*Mysql, error) {
	dsn := d.DSN()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, xerrors.Errorf("database connect open error: %w", err)
	}
	if err = db.Ping(); err != nil {
		return nil, xerrors.Errorf("database ping error: %w", err)
	}

	return &Mysql{
		dsn: dsn,
		DB:  db,
	}, nil
}

func (d *Mysql) Close() {
	err := d.DB.Close()
	if err != nil {
		fmt.Printf("database close error: %+v", err)
	}
}
