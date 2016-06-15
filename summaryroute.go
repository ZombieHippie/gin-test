package main


// SummaryPost accepts a post request to create an article
func SummaryPost(c *gin.Context) {
	var json Summary

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	article := createArticle(json.RepoID, json.Content)
	if article.Title == json.Title {
		content := gin.H{
			"result":  "Success",
			"repo":   article.Title,
			"content": article.Content,
		}
		c.JSON(201, content)
	} else {
		c.JSON(500, gin.H{"result": "An error occured"})
	}
}
