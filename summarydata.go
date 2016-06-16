package main

import (
	"time"
)

func createSummary(input Summary) Summary {
	if input.Created == 0 {
		input.Created = time.Now().UnixNano()
	}

	err := dbmap.Insert(&input)
	checkErr(err, "Insert Summary failed")
	return input
}

func getSummaryList(repoID string) ([]Summary, error) {
	var err error
	var summaries []Summary
	_, err = dbmap.Select(&summaries, "select * from summaries where repo_id=? order by created", repoID)

	if err != nil {
		return nil, err
	}

	return summaries, nil
}

func getSummaryByID(summaryID int64) (Summary, error) {
	var err error
	var summary Summary
	query := "select * from summaries where summary_id=?"
	err = dbmap.SelectOne(&summary, query, summaryID)

	return summary, err
}

func getSummaryByCommit(repoID, commit string) (Summary, error) {
	var err error
	var summary Summary
	query := "select * from summaries where repo_id=? and where commit=?"
	err = dbmap.SelectOne(&summary, query, repoID, commit)

	return summary, err
}

func getSummaryListByPullRequest(repoID string, pullRequestID int64) ([]Summary, error) {
	var err error
	var summaries []Summary
	query := "select * from summaries where repo_id=? and where pull_request_id=? order by created"
	_, err = dbmap.Select(&summaries, query, repoID, pullRequestID)

	if err != nil {
		return nil, err
	}

	return summaries, nil
}
