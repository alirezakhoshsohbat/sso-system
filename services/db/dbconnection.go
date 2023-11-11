package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func DatabaseConection() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:9268@/users")
	if err != nil {
		fmt.Println( "connection failed ",err)
		return nil, err
	}
	return db, nil
}
