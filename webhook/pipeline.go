package webhook

// PipelineEvent defines payload format for pipeline event
type PipelineEvent struct {
	ObjectKind       string  `json:"object_kind"`
	User             User    `json:"user"`
	Project          Project `json:"project"`
	Commit           Commit  `json:"commit"`
	Builds           []Build `json:"builds"`
	ObjectAttributes struct {
		ID         int      `json:"id"`
		Ref        string   `json:"ref"`
		Tag        bool     `json:"tag"`
		SHA        string   `json:"sha"`
		BeforeSHA  string   `json:"before_sha"`
		Status     string   `json:"status"`
		Stages     []string `json:"stages"`
		CreatedAt  string   `json:"created_at"`
		FinishedAt string   `json:"finished_at"`
		Duration   int      `json:"duration"`
	} `json:"object_attributes"`
}
