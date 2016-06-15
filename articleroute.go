package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ArticlesList responds with JSON of list of articles
func ArticlesList(c *gin.Context) {
	var articles []Article
	_, err := dbmap.Select(&articles, "select * from articles order by article_id")
	checkErr(err, "Select failed")
	content := gin.H{}
	for k, v := range articles {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)
}

// ArticlesDetail responds with specific Article
func ArticlesDetail(c *gin.Context) {
	articleID := c.Params.ByName("id")
	aID, _ := strconv.Atoi(articleID)
	article := getArticle(aID)
	content := gin.H{"title": article.Title, "content": article.Content}
	c.JSON(200, content)
}

// ArticlePost accepts a post request to create an article
func ArticlePost(c *gin.Context) {
	var json Article

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	article := createArticle(json.Title, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"title":   article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}
