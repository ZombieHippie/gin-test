package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func postJSONToHandler(handler gin.HandlerFunc, data interface{}, responseData interface{}) error {
	route := "/arbitrary-route"
	// set up new gin
	router := gin.New()
	router.POST(route, handler)

	jsonData, _ := json.Marshal(data)

	req, reqerr := http.NewRequest("POST", route, bytes.NewBuffer(jsonData))

	if reqerr != nil {
		return reqerr
	}

	req.Header.Add("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	return json.Unmarshal([]byte(resp.Body.String()), &responseData)
}

func TestSummaryPostHandlerGetsId(t *testing.T) {
	repoID := "ZombieHippie/testing"
	commit := "577cf826d6a622fc62d9cec456b14bdb2a3664bf"
	pr := int64(4)

	respData := SummaryResult{}

	err := postJSONToHandler(SummaryPostHandler, gin.H{
		"RepoID":        repoID,
		"Commit":        commit,
		"PullRequestID": pr,
	}, &respData)

	respSummary := respData.Summary

	if err != nil {
		t.Error(err)
	} else {
		assert.Equal(t, commit, respSummary.Commit)
		assert.Equal(t, pr, respSummary.PullRequestID)
		assert.Equal(t, repoID, respSummary.RepoID)
		assert.NotZero(t, respSummary.SummaryID, "SummaryID is assigned")
		assert.NotZero(t, respSummary.Created, "Created is assigned")
	}

}

func createData(repoID string, pullRequestID int, count int) []Summary {
	commits := []string{
		"c9627d332b58cbdb0a1afa29e3999ffcaabd49c",
		"0df91757c9ca4f058f87850175dec440f56a698",
		"06e1b3566a250cb23fdca84fc31a1b9b8e9b5c1",
		"c0ed8514524f32412c45ead9ba98d030167e499",
	}
	messages := []string{
		"Test post",
		"Created summary testing",
		"Some REpo code",
		"Initial repo",
	}
	authors := []string{
		"adam@gmail.com",
		"beck@gmail.com",
		"cole@gmail.com",
		"dale@gmail.com",
	}
	success := []bool{
		false,
		false,
		true,
		true,
	}

	summaries := make([]Summary, count)

	for i := 0; i < count; i++ {
		summaries[i] = Summary{
			RepoID:        repoID,
			Commit:        string(pullRequestID) + commits[i], // use the same commit hashes with different prs
			Message:       messages[i],
			Author:        authors[i],
			Success:       success[i],
			PullRequestID: int64(pullRequestID),
			Created:       time.Now().UnixNano() - int64(i*100), // youngest first
		}
	}
	return summaries
}

func TestSummaryCreate(t *testing.T) {

	repoID := "ZombieHippie/testing"
	pr := 4
	count := 3

	summaryList := createData(repoID, pr, count)

	for i := 0; i < len(summaryList); i++ {
		status, res := SummaryPost(summaryList[i])
		if status > 299 || status < 200 {
			t.Error("Didn't expect SummaryPost() to fail.", res)
		}
	}

	pr = 5
	count = 2
	summaryList = createData(repoID, pr, count)

	for i := 0; i < len(summaryList); i++ {
		status, res := SummaryPost(summaryList[i])
		if status > 299 || status < 200 {
			t.Error("Didn't expect SummaryPost() to fail.", res)
		}
	}
	var err error
	var sum []Summary

	sum, err = getSummaryList(repoID)
	checkErrT(t, err, "getSummaryList error")

	assert.Equal(t, 5, len(sum), "Expect that 5 summaries have been created and are listed.")
}

func checkErrT(t *testing.T, err error, message string) {
	if err != nil {
		t.Error(message)
	}
}
