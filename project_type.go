package gitlab

// Webhook represents a project hook
type Webhook struct {
	// These fields are listed at
	// http://docs.gitlab.com/ce/api/projects.html#list-project-hooks
	ID                    int    `json:"id"`
	URL                   string `json:"url"`
	ProjectID             int    `json:"project_id"`
	PushEvents            bool   `json:"push_events"`
	IssuesEvents          bool   `json:"issues_events"`
	MergeRequestsEvents   bool   `json:"merge_requests_events"`
	TagPushEvents         bool   `json:"tag_push_events"`
	NoteEvents            bool   `json:"note_events"`
	BuildEvents           bool   `json:"build_events"`
	PipelineEvents        bool   `json:"pipeline_events"`
	WikiPageEvents        bool   `json:"wiki_page_events"`
	EnableSSLVerification bool   `json:"enable_ssl_verification"`
	CreatedAt             string `json:"created_at"`
}
