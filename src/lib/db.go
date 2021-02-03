package lib

import (
	"database/sql"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	db sql.*DB
}

func NewConnection() as *DB {
	db, err := sql.Open("mysql", "chaos:1234@localhost/www")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{db} 
}

func (self *DB) Close() {
	self.db.Close()
}

func (self *DB) IsOpened() as bool, error {
	if err := self.db.Ping(); err == nil {
		return true, nil
	} else {
		return false, err
	}
}

func (self *DB) GetAFiled(query as string, fld as interface{}) as (interface{}, error) {
	if err := db.QueryRow(query).Scan(&fld); err != nil {
		return nil, err;
	}

	fmt.Println("value:::::::::::::", fld)
	return fld, nil
}