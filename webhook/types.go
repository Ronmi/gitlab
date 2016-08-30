package webhook

// Wiki represents wiki info in webhook payload
type Wiki struct {
	WebURL            string `json:"web_url"`
	GitSshURL         string `json:"git_ssh_url,omitempty"`
	GitHttpURL        string `json:"git_http_url,omitempty"`
	PathWithNamespace string `json:"name_with_namespace"`
	DefaultBranch     string `json:"default_branch"`
}

// Project represents project info in webhook payload
type Project struct {
	Wiki
	Name            string `json:"name"`
	Description     string `json:"description"`
	AvatarURL       string `json:"avatar_url,omitempty"`
	Namespace       string `json:"namespace"`
	VisibilityLevel int    `json:"visibility_level"`
	Homepage        string `json:"homepage,omitempty"`
	URL             string `json:"url,omitempty"`
}

// Commit represents commit info in webhook payload
type Commit struct {
	ID        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	URL       string `json:"url"`
	Authos    struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	} `json:"author"`
	Added    []string `json:"added,omitempty"`
	Modified []string `json:"modified,omitempty"`
	Removed  []string `json:"removed,omitempty"`
}

// Diff represents diff info in webhook payload
type Diff struct {
	Diff        string `json:"diff"`
	NewPath     string `json:"new_path"`
	OldPath     string `json:"old_path"`
	AMode       string `json:"a_mode"`
	BMode       string `json:"b_mode"`
	NewFile     bool   `json:"new_file,omitempty"`
	RenamedFile bool   `json:"renamed_file,omitempty"`
	DeletedFile bool   `json:"deleted_file,omitempty"`
}

// User represents user info in webhook payload
type User struct {
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
}

// CommonAttribute represents common attributes
type CommonAttribute struct {
	ID        int    `json:"id"`
	AuthorID  int    `json:"author_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// ObjectAttribute represents common object attributes.
//
// To prevent name collision, CommonAttribute is not embbeded.
type ObjectAttribute struct {
	ProjectID int    `json:"project_id,omitempty"`
	URL       string `json:"url"`
	Action    string `json:"action,omitempty"`
}

// Issue represents issue info in webhook payload
type Issue struct {
	CommonAttribute
	Title       string `json:"title"`
	AssigneeID  int    `json:"assignee_id"`
	Position    int    `json:"position,omitempty"`
	BranchName  string `json:"branch_name,omitempty"`
	Description string `json:"description"`
	MilestoneID int    `json:"milestone_id"`
	State       string `json:"state"`
	IID         int    `json:"iid"`
}

// MergeRequest represents merge request in webhook payload
type MergeRequest struct {
	Issue
	TargetBranch    string  `json:"target_branch"`
	SourceBranch    string  `json:"source_branch"`
	TargetProjectID int     `json:"target_project_id"`
	SourceProjectID int     `json:"source_project_id"`
	MergeStatus     string  `json:"merge_status"`
	LockedAt        string  `json:"locaked_at,omitempty"`
	Target          Project `json:"target"`
	Source          Project `json:"source"`
	LastCommit      Commit  `json:"last_commit"`
	WorkInProgress  bool    `json:"work_in_progress"`
	Assignee        User    `json:"assignee"`
}

// Snippet represents snippet info in webhook payload
type Snippet struct {
	CommonAttribute

	Title           string `json:"title"`
	Content         string `json:"content"`
	ProjectID       int    `json:"project_id"`
	FileName        string `json:"file_name"`
	ExpiresAt       string `json:"expires_at"`
	Type            string `json:"type"`
	VisibilityLevel int    `json:"visibility_level"`
}

// Build represents build info in webhook payload
type Build struct {
	ID            int    `json:"id"`
	Stage         string `json:"stage"`
	Name          string `json:"name"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
	StartedAt     string `json:"started_at,omitempty"`
	FinishedAt    string `json:"finished_at,omitempty"`
	When          string `json:"when"`
	Manual        bool   `json:"manual"`
	User          User   `json:"user"`
	Runner        string `json:"runner"`
	ArtifactsFile struct {
		Filename string `json:"filename,omitempty"`
		Size     int    `json:"size,omitempty"`
	} `json:"artifacts_file"`
}
