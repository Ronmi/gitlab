package webhook

// CommentEvent defines payload format for comment on commit
type CommentEvent struct {
	ObjectKind       string       `json:"object_kind"`
	User             User         `json:"user"`
	Project          Project      `json:"project"`
	Commit           Commit       `json:"commit,omitempty"`
	MergeRequest     MergeRequest `json:"merge_request,omitempty"`
	Issue            Issue        `json:"issue,omitempty"`
	Snippet          Snippet      `json:"snippet,omitempty"`
	ObjectAttributes struct {
		CommonAttribute
		ObjectAttribute
		Note         string `json:"note"`
		NoteableType string `json:"noteable_type"`
		Attachment   string `json:"attachment,omitempty"`
		LineCode     string `json:"line_code,omitempty"`
		CommitID     string `json:"commit_id,omitempty"`
		NoteableID   int    `json:"noteable_id"`
		System       bool   `json:"system,omitempty"`
		STDiff       Diff   `json:"st_diff,omitempty"`
	} `json:"object_attributes"`
}
