package app

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
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
	Message string
	Summary summary.Summary[]
}

// Setup creates our router and returns it
func Setup(db *gorm.DB) {
	router := gin.Default()

	// Create summary
	router.POST("/webhook", func(c *gin.Context) {
		var json summary.Summary

		c.Bind(&json) // This will infer what binder to use depending on the content-type header.

		hasRepoName := !shared.IsZero(json.Repository.Name)

		if !hasRepoName {
			c.JSON(http.StatusBadRequest, webhookResp{
				Message: "Error: No Repository.Name specified."
			})
			return
		}

		json.Repository = repo.EnsureRepository(db, json.Repository)

		hasCommit := !shared.IsZero(json.Commit)
		hasCreated := !shared.IsZero(json.Created)
		hasPullRequest := !shared.IsZero(json.PullRequestID)
		hasArtifacts := len(json.Artifacts) > 0

		var err error
		if !hasArtifacts {
			err = "No Artifacts present."
		} else if !hasPullRequest {
			err = "No PullRequestID specified."
		} else if !hasCommit {
			err = "No Commit SHA specified."
		} else if !hasArtifacts {
			err = "No Artifacts provided."
		}

		if err != nil {
			c.JSON(http.StatusBadRequest, webhookResp{
				Message: "Error: " + string(err)
			})
			return
		}

		created, result := summary.CreateSummary(json)
		
		status := created ? http.StatusCreated : http.StatusInternalServerError
		c.JSON(status, webhookResp{
			Message: "Successfully created summary.",
			Summary: result,
		})
	})

	router.GET("/summary/list", func(c *gin.Context) {
		sums := summary.GetAllSummaries(db)
		
		c.JSON(http.StatusOK, summaryListResp{
			Summaries: sums,
			Message: "Successfully retrieved summaries."
		})
	})

	router.Static("/public", "./public")

	return router
}
