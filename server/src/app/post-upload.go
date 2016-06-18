package app

import (
	"encoding/json"
	"github.com/ZombieHippie/test-gin/server/src/artifact"
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
	"regexp"
	"strconv"
)

type uploadResp struct {
	Message   string
	Summary   summary.Summary
	Artifacts []artifact.Artifact
}

func postUpload(c *gin.Context, db *gorm.DB, savedir string) {
	log.SetPrefix("postUpload > ")

	sSumUp := c.Request.FormValue("SummaryUpload")
	var sumUp upload.SummaryUpload

	json.Unmarshal([]byte(sSumUp), &sumUp)

	hasRepoName := !shared.IsZero(sumUp.Repository.ID)
	if !hasRepoName {
		log.Println(sumUp)
		c.JSON(http.StatusBadRequest, uploadResp{
			Message: "Error: No Repository.ID specified.",
		})
		return
	}
	sumUp.Repository = repo.EnsureRepository(db, sumUp.Repository)

	hasReqdFields := hasRequiredFields(sumUp)
	if hasReqdFields != "" {
		c.JSON(http.StatusBadRequest, uploadResp{
			Message: "Error: " + string(hasReqdFields),
		})
		return
	}

	uploadedArtifacts := make([]artifact.Artifact, 4)

	// Save files if length longer than 255 chars
	for _, art := range sumUp.Artifacts {

		file, header, err := c.Request.FormFile(art.Path)
		defer file.Close()
		filename := header.Filename
		log.Println("Found file:", art.Path, filename)
		if err != nil {
			log.Fatalln("Fatal form file", err)
		}

		// Make sure that this is not the same file
		r := regexp.MustCompile(`[^\w\-]+`)
		safeFilename := r.ReplaceAllString(art.Label, "-")
		filepath := path.Join(savedir, sumUp.Repository.ID, strconv.Itoa(sumUp.BuildID), safeFilename)

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
		upArt, err := art.SaveUpload(filepath, file, false)
		if err != nil {
			log.Fatalln("Error on save upload", err)
		} else {
			uploadedArtifacts = append(uploadedArtifacts, upArt)
		}
	}

	// create summary based on SummaryUpload
	sum := summary.Summary{
		Repository: sumUp.Repository,
		BranchID:   sumUp.BranchID,
		BuildID:    sumUp.BuildID,
		Created:    sumUp.Created,
		Commit:     sumUp.Commit,
		Author:     sumUp.Author,
		Message:    sumUp.Message,
		Success:    sumUp.Success,
	}

	db.Create(&sum)

	for _, art := range uploadedArtifacts {
		art.Summary = sum
		db.Create(&art)
	}

	status := 200

	c.JSON(status, uploadResp{
		Message:   "Successfully created summary.",
		Summary:   sum,
		Artifacts: uploadedArtifacts,
	})
}

func hasRequiredFields(json upload.SummaryUpload) string {
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

	return err
}
