package webhook

// MergeRequestEvent defines payload format of merge request event
type MergeRequestEvent struct {
	ObjectKind       string `json:"object_kind"`
	User             User   `json:"user,omitempty"`
	ObjectAttributes struct {
		MergeRequest
		ObjectAttribute
		STCommits []Commit `json:"st_commits,omitempty"`
		STDiffs   []Diff   `json:"st_diffs,omitempty"`
	} `json:"object_attributes"`
}
