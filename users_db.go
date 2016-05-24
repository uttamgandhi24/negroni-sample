package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func getUserIDByNameAndPassword(username string, password string) string {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	stmt, err := db.Prepare(`SELECT UserID FROM  users WHERE username = ?
        	AND password= ?`)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	var userID string
	err = stmt.QueryRow(username, password).Scan(&userID)

	if err != nil {
		panic(err.Error())
	}
	return userID
}
