package lib

import (
	"fmt"
	"time"	
	"errors"	

	"database/sql"	
	"github.com/go-sql-driver/mysql"
)

type DB struct {
	db *sql.DB
}

func NewConnection() *DB {
	db, err := sql.Open("mysql", "chaos:1234@/www")
	if err != nil {
		panic(err)
	}

	fmt.Println("DB:Connection Established")
	// See "Important settings" section.
	// db.SetConnMaxLifetime(time.Second * 60)
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &DB{db} 
}

func (self *DB) Close() {
	if self != nil && self.db != nil {
		self.db.Close()
	}
}

func (self *DB) IsOpened() error {
	if self == nil || self.db == nil {
		return errors.New("[ERROR]Connection is aleady closed")
	} else {
		if err := self.db.Ping(); err != nil {
			return err
		}

		return nil
	}
}

func (self *DB) GetDB() (*DB, error) {
	if err :=  self.IsOpened(); err != nil {
		fmt.Println("DB:Existed Connection Returned")
		return self, err
	} 
	
	return NewConnection(), nil
}

func (self *DB) GetRow(query string, args ...interface{}) error {
	if err := self.db.QueryRow(query).Scan(args...); err != nil && err != sql.ErrNoRows {
		return err
	} 
	return nil
}

func (self *DB) Exec(query string, args ...interface{}) (int64, error) {
	if result, err := self.db.Exec(query); err != nil {
		if mysqlError, ok := err.(*mysql.MySQLError); ok {
			if mysqlError.Number == 1062 { return 0, nil } //ER_DUP_ENTRY 
			return -1, errors.New(fmt.Sprintln("[%s]%s", mysqlError.Number, mysqlError.Message))
		} 
		return -1, err
	} else {
		return result.RowsAffected()
	}
}
	// rows, err := self.db.Query(query, args...)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return err
	// }
	// defer rows.Close()

	// fmt.Println("Rows:::", rows)
	// return nil

	// for rows.Next() {
	// 	var (
	// 		id   int64
	// 		name string
	// 	)
	// 	if err := rows.Scan(&id, &name); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("id %d name is %s\n", id, name)
	// }
	// if !rows.NextResultSet() {
	// 	log.Fatalf("expected more result sets: %v", rows.Err())
	// }
	// var roleMap = map[int64]string{
	// 	1: "user",
	// 	2: "admin",
	// 	3: "gopher",
	// }
	// for rows.Next() {
	// 	var (
	// 		id   int64
	// 		role int64
	// 	)
	// 	if err := rows.Scan(&id, &role); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Printf("id %d has role %s\n", id, roleMap[role])
	// }
	// if err := rows.Err(); err != nil {
	// 	log.Fatal(err)
	// }


// func (self *DB) GetRow(query string, args ...interface{}) {
// 	rows, err := self.db.Query(query, args); err != nil {
//     if err != nil {
// 			log.Fatal(err)
// 	}
// 	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

// 	for rows.Next() {
// 			err := rows.Scan(args)
// 			if err != nil {
// 					log.Fatal(err)
// 			}
// 			fmt.Println(args)
// 	}
// }

// func (self *DB) GetAFiled(query as string, fld as interface{}) as (interface{}, error) {
// 	if err := db.QueryRow(query).Scan(&fld); err != nil {
// 		return nil, err;
// 	}

// 	fmt.Println("value:::::::::::::", fld)
// 	return fld, nil
// }