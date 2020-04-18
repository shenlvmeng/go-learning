package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func seachIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	res, err := http.Get(Issueurl + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", res.Status)
	}
	var issueRes IssuesSearchResult
	if err := json.NewDecoder(res.Body).Decode(&issueRes); err != nil {
		res.Body.Close()
		return nil, err
	}
	res.Body.Close()
	return &issueRes, nil
}
