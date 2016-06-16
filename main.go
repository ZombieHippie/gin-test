package main

import (
	"database/sql"
	"log"
	"time"

	"gopkg.in/gorp.v1"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // used by gorp ?
	"net/http"
	"os"
	"strconv"
)

var dbmap = initDb()

func main() {

	db, err := gorm.Open("")

	defer dbmap.Db.Close()

	router := gin.Default()

	router.GET("/articles", ArticlesList)
	router.POST("/articles", ArticlePost)
	router.GET("/articles/:id", ArticlesDetail)

	// Summary get by ID
	router.GET("/summary/:id", func(c *gin.Context) {
		summaryID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusOK, SummaryResult{
				Result: fmt.Sprintf("Summary '%s' is not an integer.", c.Param("id")),
			})
		}
		status, result := SummaryGet(Summary{
			SummaryID: int64(summaryID),
		})
		c.JSON(status, result)
	})

	// Summary get by ID
	router.DELETE("/summary/:id", func(c *gin.Context) {
		summaryID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusOK, SummaryResult{
				Result: fmt.Sprintf("Summary '%s' is not an integer.", c.Param("id")),
			})
		}
		status, result := SummaryGet(Summary{
			SummaryID: int64(summaryID),
		})
		c.JSON(status, result)
	})

	// Create summary
	router.POST("/summary", func(c *gin.Context) {
		var json Summary

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.
		status, result := SummaryPost(json)
		c.JSON(status, result)
	})

	// Summary find by repo/commit/pr
	router.POST("/find/summary", func(c *gin.Context) {
		var json Summary

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.
		status, result := SummaryGet(json)
		c.JSON(status, result)
	})

	router.Static("/public", "./public")
	router.Run(":8080")
}

func initDb() *gorp.DbMap {

	// for now we will delete the db.sqlite file
	rmerr := os.Remove("db.sqlite3")

	checkErr(rmerr, "Removing previous database file failed.")

	db, err := gorm.Open("sqlite3", "db.sqlite3")
	checkErr(err, "gorm.Open failed")
	/*
		dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

		dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "ID")
		dbmap.AddTableWithName(Summary{}, "summaries").SetKeys(true, "SummaryID")

		err = dbmap.CreateTablesIfNotExists()
		checkErr(err, "Create tables failed")
	*/
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
