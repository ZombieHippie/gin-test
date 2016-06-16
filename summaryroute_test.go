package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
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

func TestSummaryPostGetsId(t *testing.T) {
	repoID := "ZombieHippie/testing"
	commit := "577cf826d6a622fc62d9cec456b14bdb2a3664bf"
	pr := int64(4)

	respData := SummaryPostResult{}

	err := postJSONToHandler(SummaryPost, gin.H{
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
