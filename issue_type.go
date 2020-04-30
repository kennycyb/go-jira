package jira

// IssueType bug, etc.
type IssueType struct {
	Name    string `json:"name"`
	IconURL string `json:"iconUrl,omitempty"`
	SubTask bool   `json:"subtask"`
}

// IssueTypeList list of IssueType
type IssueTypeList []IssueType
