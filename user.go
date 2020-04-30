package jira

// User of JIRA
type User struct {
	Self        string            `json:"self"`
	Key         string            `json:"key"`
	Name        string            `json:"name"`
	Email       string            `json:"emailAddress"`
	AvatarUrls  map[string]string `json:"avatarUrls"`
	DisplayName string            `json:"displayName"`
	Active      bool              `json:"active"`
	TimeZone    string            `json:"timeZone"`
	Locale      string            `json:"locale"`
}
