package webhook

// IssueEvent defines payload format of issue event
type IssueEvent struct {
	ObjectKind           string `json:"object_kind"`
	User                 User   `json:"user"`
	Assignee             User   `json:"assignee"`
	ObjectKindAttributes struct {
		Issue
		ObjectAttribute
	} `json:"object_attributes"`
}
