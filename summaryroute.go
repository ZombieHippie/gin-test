package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SummaryPostResult is result from the SummaryPost route
type SummaryPostResult struct {
	Result  string
	Summary Summary
}

// SummaryPost accepts a post request to create an article
func SummaryPost(c *gin.Context) {
	var json Summary

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.
	if json.Commit != "" && json.RepoID != "" && json.PullRequestID != 0 {
		// All properties are available
		summary := createSummary(json.RepoID, json.Commit, json.PullRequestID)
		if summary.RepoID == json.RepoID {
			c.JSON(http.StatusCreated, SummaryPostResult{
				Result:  "success",
				Summary: summary,
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"Result": "An error occurred"})
		}
	} else {
		c.JSON(http.StatusBadRequest, SummaryPostResult{
			Result: "Insufficient parameters.",
		})
	}
}
