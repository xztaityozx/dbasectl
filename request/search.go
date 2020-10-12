package request

import "github.com/xztaityozx/dbasectl/output"

type (
	SearchRequest struct {
		query           string
		selectedColumns []string
		outputFormat    output.Format
	}

	PostSearchRequest struct {
		SearchRequest
	}
)

const (
	PostSearch  EndPoint = "posts"
	UserSearch  EndPoint = "users"
	GroupSearch EndPoint = "groups"
)
