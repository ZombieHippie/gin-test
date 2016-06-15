package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

const TestToken = "HELLO"

func TestSummaryPostGetsId(t *testing.T) {
	router := gin.New()
	router.POST("/summary", SummaryPost)

	repoID := "ZombieHippie/testing"
	commit := "577cf826d6a622fc62d9cec456b14bdb2a3664bf"
	var pr int64 = 4

	data, _ := json.Marshal(gin.H{
		"RepoID":        repoID,
		"Commit":        commit,
		"PullRequestID": pr,
	})

	req, _ := http.NewRequest("POST", "/summary", bytes.NewBuffer(data))
	req.Header.Add("Authorization", fmt.Sprintf("auth_token=\"%s\"", TestToken))
	req.Header.Add("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	fmt.Println(resp.Body.String())
	respData := SummaryPostResult{}
	err := json.Unmarshal([]byte(resp.Body.String()), &respData)

	fmt.Println(resp.Body.Bytes(), respData)

	respSummary := respData.Summary
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(respSummary)
		assert.Equal(t, commit, respSummary.Commit)
		assert.Equal(t, pr, respSummary.PullRequestID)
		assert.Equal(t, repoID, respSummary.RepoID)
		assert.NotZero(t, respSummary.SummaryID, "SummaryID is assigned")
		assert.NotZero(t, respSummary.Created, "Created is assigned")
	}

}
