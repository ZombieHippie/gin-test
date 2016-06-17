package main

import (
	"log"

	"github.com/ZombieHippie/test-gin/src/app"
	"github.com/ZombieHippie/test-gin/src/artifact"
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // used by gorp ?
	"os"
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
	db.LogMode(true)
	checkErr(err, "gorm.Open failed")

	db.CreateTable(
		new(repo.Repository),
		new(summary.Summary),
		new(artifact.Artifact),
	)

	return db
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
