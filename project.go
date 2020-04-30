package jira

// Project info
type Project struct {
	Key        string            `json:"key,omitempty"`
	Name       string            `json:"name,omitempty"`
	AvatarURLs map[string]string `json:"avatarUrls,omitempty"`
	IssueTypes []IssueType       `json:"issueTypes,omitempty"`
}

// ProjectList array of projects
type ProjectList []Project
