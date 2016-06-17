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
const datapath = "data/"

func main() {

	db := initDb()
	defer db.Close()
	router := app.Setup(db, datapath)
	router.Run(":8080")
}

func initDb() *gorm.DB {

	dbfilepath := datapath + "db.sqlite"

	err := os.MkdirAll(datapath, 0777)
	checkErr(err, "Failed to create datapath.")
	// for now we will delete the db.sqlite file
	if removeOldDB {

		err = os.Remove(dbfilepath)
		checkErr(err, "Removing previous database file failed.")

	}

	db, err := gorm.Open("sqlite3", dbfilepath)
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
