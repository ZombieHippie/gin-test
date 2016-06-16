package app

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func postWebhook(c *gin.Context, db *gorm.DB) {
	var json summary.Summary

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.

	hasRepoName := !shared.IsZero(json.Repository.ID)

	if !hasRepoName {
		c.JSON(http.StatusBadRequest, webhookResp{
			Message: "Error: No Repository.Name specified.",
		})
		return
	}

	json.Repository = repo.EnsureRepository(db, json.Repository)

	var err string
	hasArtifacts := len(json.Artifacts) > 0
	if !hasArtifacts {
		err += "No Artifacts present. "
	}

	hasPullRequest := !shared.IsZero(json.PullRequestID)
	if !hasPullRequest {
		err += "No PullRequestID specified. "
	}

	hasCommit := !shared.IsZero(json.Commit)
	if !hasCommit {
		err += "No Commit SHA specified. "
	}

	hasCreated := !shared.IsZero(json.Created)
	if !hasCreated {
		err += "No Created provided. "
	}

	if err != "" {
		c.JSON(http.StatusBadRequest, webhookResp{
			Message: "Error: " + string(err),
		})
		return
	}

	created, result := summary.CreateSummary(db, json)

	var status int
	if created {
		status = http.StatusCreated
	} else {
		status = http.StatusInternalServerError
	}

	c.JSON(status, webhookResp{
		Message: "Successfully created summary.",
		Summary: result,
	})
}
