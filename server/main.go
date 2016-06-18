package main

import (
	"log"

	"github.com/ZombieHippie/test-gin/server/src/app"
	"github.com/ZombieHippie/test-gin/server/src/artifact"
	"github.com/ZombieHippie/test-gin/server/src/repo"
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" // used by gorp ?
	"os"
	"strings"
)

const removeOldDB = true
const datapath = "data/"

func main() {
	port := ":8080"
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		if pair[0] == "PORT" {
			port = ":" + pair[1]
		}
	}

	db := initDb()
	defer db.Close()
	router := app.Setup(db, datapath)
	router.Run(port)
}

func initDb() *gorm.DB {

	dbfilepath := datapath + "db.sqlite3"
	var err error

	// for now we will delete the db.sqlite file
	if removeOldDB {
		err = os.RemoveAll(datapath)
		checkErr(err, "Removing previous database file failed.")

	}
	err = os.MkdirAll(datapath, 0777)
	checkErr(err, "Failed to create datapath.")

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
