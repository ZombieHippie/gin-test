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

// SummaryPostHandler accepts a post request to create an article
func SummaryPostHandler(c *gin.Context) {
	var json Summary

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	status, result := SummaryPost(json)
	c.JSON(status, result)
}

func createArticle(title, body string) Article {
	article := Article{
		Created: time.Now().UnixNano(),
		Title:   title,
		Content: body,
	}

	err := dbmap.Insert(&article)
	checkErr(err, "Insert failed")
	return article
}

func getArticle(articleID int) Article {
	article := Article{}
	err := dbmap.SelectOne(&article, "select * from articles where article_id=?", articleID)
	checkErr(err, "SelectOne failed")
	return article
}

func initDb() *gorp.DbMap {

	// for now we will delete the db.sqlite file
	rmerr := os.Remove("db.sqlite3")

	if rmerr != nil {
		checkErr(rmerr, "Removing previous database file failed.")
		panic(rmerr)
	}

	db, err := sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "ID")
	dbmap.AddTableWithName(Summary{}, "summaries").SetKeys(true, "SummaryID")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
