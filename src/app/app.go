package app

import (
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

type webhookResp struct {
	Message string
	Summary summary.Summary
}

type summaryListResp struct {
	Message   string
	Summaries []summary.Summary
}

// Setup creates our router and returns it
func Setup(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	// Create summary
	router.POST("/webhook", func(c *gin.Context) {
		postWebhook(c, db)
	})

	router.GET("/summary/list", func(c *gin.Context) {
		sums := summary.GetAllSummaries(db)

		c.JSON(http.StatusOK, summaryListResp{
			Summaries: sums,
			Message:   "Successfully retrieved summaries.",
		})
	})

	router.Static("/public", "./public")

	return router
}
