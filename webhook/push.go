package webhook

// PushEvent defines payload format of push event
type PushEvent struct {
	ObjectKind       string   `json:"object_kind"`
	Before           string   `json:"before"`
	After            string   `json:"after"`
	Ref              string   `json:"ref"`
	CheckoutSHA      string   `json:"checkout_sha"`
	UserID           int      `json:"user_id"`
	UserName         string   `json:"user_name"`
	UserEmail        string   `json:"user_email"`
	UserAvatar       string   `json:"user_avatar"`
	ProjectID        int      `json:"project_id"`
	Project          Project  `json:"project"`
	Commits          []Commit `json:"commits"`
	TotalCommitCount int      `json:"total_commit_count"`
}
