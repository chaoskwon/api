package auth

import (
	"fmt"

	"lib"		
)

func SignIn (db *lib.DB, email string, pwd string) (bool, error) {
	var id int
	var err error

	db, err = db.GetDB() 
	if db != nil {
		query := fmt.Sprintf("SELECT id FROM users WHERE email = '%s' and pwd = md5('%s')", email, pwd)
		fmt.Println("query :::::::: \n", query)
    id = db.GetRow(query, id)
    fmt.Println("** id::::::::::::", id)

		return true, err
	} else {
		fmt.Println("disconnected have to stop")
		return false, err
	}
}