package app

import (
	"github.com/ZombieHippie/test-gin/server/src/repo"
	"github.com/ZombieHippie/test-gin/server/src/shared"
	"github.com/ZombieHippie/test-gin/server/src/summary"
	"github.com/ZombieHippie/test-gin/server/src/upload"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

type uploadResp struct {
	Message string
	Summary summary.Summary
}

func postUpload(c *gin.Context, db *gorm.DB, savedir string) {
	var json upload.SummaryUpload

	c.Bind(&json) // This will infer what binder to use depending on the content-type header.

	hasRepoName := !shared.IsZero(json.Repository.ID)

	if !hasRepoName {
		c.JSON(http.StatusBadRequest, uploadResp{
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

	hasBranch := !shared.IsZero(json.BranchID)
	if !hasBranch {
		err += "No BranchID specified. "
	}

	hasBuild := !shared.IsZero(json.BuildID)
	if !hasBuild {
		err += "No BuildID specified. "
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
		c.JSON(http.StatusBadRequest, uploadResp{
			Message: "Error: " + string(err),
		})
		return
	}

	// Save files if length longer than 255 chars
	for _, art := range json.Artifacts {
		contents, err := art.ReadFile()

		if err != nil {
			log.Fatalln(err)
			continue
		}

		if len(contents) >= 255 {
			// Somehow make sure that this is not the same file
			filepath := path.Join(savedir, json.Repository.ID, strconv.Itoa(json.BuildID), art.FileName)

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

	c.JSON(status, uploadResp{
		Message: "Successfully created summary.",
		Summary: result,
	})
}
