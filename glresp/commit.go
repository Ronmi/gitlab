package glresp

// CommitStats represents states of a commit
type CommitStats struct {
	Additions int `json:"additions,omitempty"`
	Deletions int `json:"deletions,omitempty"`
	Total     int `json:"total,omitempty"`
}

// Commit represents a commit in a repository
type Commit struct {
	// These fields are listed at
	// http://docs.gitlab.com/ce/api/commits.html#list-repository-commits
	ID           string `json:"id"`
	ShortID      string `json:"short_id,omitempty"`
	Title        string `json:"title,omitempty"`
	AuthorName   string `json:"author_name,omitempty"`
	AuthorEmail  string `json:"author_email,omitempty"`
	CreatedAt    string `json:"created_at,omitempty"`
	Message      string `json:"message,omitempty"`
	AllowFailure string `json:"allow_failure,omitempty"`

	// These fields are listed at
	// http://docs.gitlab.com/ce/api/commits.html#get-a-single-commit
	CommittedDate string      `json:"committed_date,omitempty"`
	AuthoredDate  string      `json:"authored_date,omitempty"`
	ParentIDs     []string    `json:"parent_ids,omitempty"`
	Stats         CommitStats `json:"stats,omitempty"`
	Status        string      `json:"status,omitempty"`
}
