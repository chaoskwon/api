package auth

import (
	"fmt"

	"lib"		
)

func SignIn (db *lib.DB, email string, pwd string) (bool, error) {
	var err error

	db, err = db.GetDB() 
	if err != nil {
		return false, err
	} 

	query := fmt.Sprintf("SELECT id FROM users WHERE email = '%s' and pwd = md5('%s')", email, pwd)
	fmt.Println("query :::::::: \n", query)

	var id int
	if err = db.GetRow(query, &id); err != nil || id <= 0 {
		return false, err
	}
	
	return true, nil
}

func Register (db *lib.DB, email string, pwd string) (bool, error) {
	var err error

	db, err = db.GetDB() 
	if err != nil {
		return false, err
	} 

	// INSERT 문 실행
	query := fmt.Sprintf("INSERT INTO users (email, pwd) VALUES('%s', md5('%s'))", email, pwd)
	fmt.Println("query :::::::: \n", query)

	if err := db.Exec(query); err != nil {
		return false, err		
	} else {
		return true, nil
	}
}