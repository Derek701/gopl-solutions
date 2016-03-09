// Package github provides a Go API for the GitHub issue tracker.
// See https://developer.github.com/v3/search/#search-issues.
package github

import "time"

// IssuesURL is the issue tracker web-service interface URL.
const IssuesURL = "https://api.github.com/search/issues"

// IssuesSearchResult is for issue tracker response data.
type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

// Issue variables for each result.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

// User variables for each issue.
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
