package app

import (
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type summaryListResp struct {
	Message   string
	Summaries []summary.Summary
	Count     int
}

// Setup creates our router and returns it
func Setup(db *gorm.DB, savedir string) *gin.Engine {
	router := gin.Default()

	// Create summary
	router.POST("/summary/webhook", func(c *gin.Context) {
		postWebhook(c, db, savedir)
	})

	router.GET("/summary/list", func(c *gin.Context) {
		sums, count := summary.GetAllSummaries(db)

		c.JSON(http.StatusOK, summaryListResp{
			Message:   "Successfully retrieved summaries.",
			Summaries: sums,
			Count:     count,
		})
	})

	router.Static("/public", "./public")

	// indexed files
	router.StaticFS("/data", gin.Dir("./data", true))

	return router
}
