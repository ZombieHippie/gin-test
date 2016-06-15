package main

import (
	"database/sql"
	"log"
	"time"

	"gopkg.in/gorp.v1"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3" // used by gorp ?
)

var dbmap = initDb()

func main() {

	defer dbmap.Db.Close()

	router := gin.Default()
	router.GET("/articles", ArticlesList)
	router.POST("/articles", ArticlePost)
	router.GET("/articles/:id", ArticlesDetail)
	router.Static("/public", "./public")
	router.Run(":8080")
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

func createSummary(repo, commit string) Summary {
	summary := Summary{
		Created:  time.Now().UnixNano(),
		RepoID:   repo,
		Commit:   commit,
	}

	err := dbmap.Insert(&summary)
	checkErr(err, "Insert Summary failed")
	return summary
}

func getArticle(articleID int) Article {
	article := Article{}
	err := dbmap.SelectOne(&article, "select * from articles where article_id=?", articleID)
	checkErr(err, "SelectOne failed")
	return article
}

func initDb() *gorp.DbMap {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	checkErr(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Article{}, "articles").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
