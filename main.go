package main

import (
	"log"

	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // used by gorp ?
	"net/http"
	"os"
	"strconv"
)

const removeOldDB = false

func main() {

	db := initDb()
	defer db.Close()
	router := app.Setup(db)
	router.Run(":8080")
}

func initDb() *gorm.DB {

	// for now we will delete the db.sqlite file
	if removeOldDB {
		rmerr := os.Remove("db.sqlite3")

		checkErr(rmerr, "Removing previous database file failed.")
	}

	db, err := gorm.Open("sqlite3", "db.sqlite3")
	checkErr(err, "gorm.Open failed")

	return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
