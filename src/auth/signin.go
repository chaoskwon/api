package auth

import (
	"fmt"

	"lib"		
)

func SignIn (db *lib.DB, email string, pwd string) (bool, error) {
	query := fmt.Sprintf("SELECT id FROM users WHERE email = '%s' and pwd = md5('%s')", email, pwd)
	fmt.Println("== SignIn query :::::::: \n", query)

	var id int
	if err := db.GetRow(query, &id); err != nil || id < 0 { return false, err }
	if id == 0 { return false, nil } 

	return true, nil
}

func Register (db *lib.DB, email string, pwd string) (int64, error) {
	query := fmt.Sprintf("INSERT INTO users (email, pwd) VALUES('%s', md5('%s'))", email, pwd)
	fmt.Println("== Register query :::::::: \n", query)

	return db.Exec(query)
	// bln, err := db.Exec(query)
	// fmt.Println("bln, error :: ", bln, err)
	// return bln, err
}