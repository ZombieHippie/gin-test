package main

import (
	"net/http"
)

// SummaryResult is result from different routes
type SummaryResult struct {
	Result  string
	Summary Summary
}

// SummaryListResult is result from routes returning lists of SummaryList
type SummaryListResult struct {
	Result      string
	SummaryList []Summary
}

// SummaryPost adds Summary to database
func SummaryPost(input Summary) (int, SummaryResult) {
	hasSufficientParameters := input.Commit != "" && input.RepoID != "" && input.PullRequestID != 0
	if hasSufficientParameters {
		// All properties are available
		summary := createSummary(input)
		if summary.RepoID == input.RepoID {
			return http.StatusCreated, SummaryResult{
				Result:  "Created summary",
				Summary: summary,
			}
		}
		return http.StatusInternalServerError, SummaryResult{Result: "An error occurred"}
	}
	return http.StatusBadRequest, SummaryResult{
		Result: "Insufficient parameters.",
	}
}

// SummaryGet retrieves the summary report information based on given information about the Summary from
// its SummaryID, or its RepoID and either PullRequestID or Commit. If PullRequestID is provided, but
// not Commit, then the most recently created Summary of that pull request is retrieved.
func SummaryGet(input Summary) (int, SummaryResult) {
	var summary Summary
	var err error
	if input.SummaryID != 0 {
		// by SummaryID
		summary, err = getSummaryByID(input.SummaryID)

	} else if input.RepoID != "" && input.Commit != "" {
		// by Repo and Commit
		summary, err = getSummaryByCommit(input.RepoID, input.Commit)

	} else if input.RepoID != "" && input.PullRequestID != 0 {
		// by Repo and Pull request
		var summaries []Summary
		summaries, err = getSummaryListByPullRequest(input.RepoID, input.PullRequestID)
		summary = summaries[0]
	} else {
		// Insufficient parameters
		return http.StatusBadRequest, SummaryResult{
			Result: "Insufficient parameters.",
		}
	}
	if err != nil {
		return http.StatusInternalServerError, SummaryResult{
			Result: "Database error retrieving summary",
		}
	}
	if summary.SummaryID != 0 {
		return http.StatusOK, SummaryResult{
			Summary: summary,
		}
	}
	return http.StatusNotFound, SummaryResult{
		Summary: summary,
	}
}

// SummaryList retrieves the summary reports for a given RepoID and optional PullRequestID.
func SummaryList(input Summary) (int, SummaryListResult) {
	hasSufficientParameters := input.RepoID != ""
	if hasSufficientParameters {
		if input.PullRequestID != 0 {
			summary, err := getSummaryListByPullRequest(input.RepoID, input.PullRequestID)
			if err != nil {
				return http.StatusOK, SummaryListResult{
					SummaryList: summary,
				}
			}
			return http.StatusNotFound, SummaryListResult{
				SummaryList: summary,
			}
		}

	}
	return http.StatusBadRequest, SummaryListResult{
		Result: "Insufficient parameters.",
	}
}
