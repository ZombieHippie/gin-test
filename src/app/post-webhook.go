package app

import (
	"github.com/ZombieHippie/test-gin/src/repo"
	"github.com/ZombieHippie/test-gin/src/shared"
	"github.com/ZombieHippie/test-gin/src/summary"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

type webhookResp struct {
	Message string
	Summary summary.Summary
}

func postWebhook(c *gin.Context, db *gorm.DB, savedir string) {
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

	// Save files if length longer than 255 chars
	for _, art := range json.Artifacts {
		contents, err := art.ReadFile()
		log.Println(contents, err)
		if len(contents) >= 255 {
			// Somehow make sure that this is not the same file
			filepath := path.Join(savedir, json.Repository.ID, json.Commit, art.Label)

			// test different paths, file.txt, file.1.txt, file.2.txt etc
			for existsSuffix := 0; existsSuffix < 100; existsSuffix++ {
				testpath := filepath
				ext := path.Ext(testpath)
				if existsSuffix > 0 {
					testpath = testpath[0:len(testpath)-len(ext)] + "." + strconv.Itoa(existsSuffix) + ext
				}
				if _, err := os.Stat(testpath); os.IsNotExist(err) {
					filepath = testpath
					break
				}
			}

			art.SaveIntoFile(filepath)
		}
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
