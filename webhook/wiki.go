package webhook

// WikiEvent defines payload format of wiki event
type WikiEvent struct {
	ObjectKind       string  `json:"object_kind"`
	User             User    `json:"user"`
	Project          Project `json:"project"`
	Wiki             Wiki    `json:"wiki"`
	ObjectAttributes struct {
		ObjectAttribute

		Title   string `json:"title"`
		Content string `json:"content"`
		Format  string `json:"format"`
		Message string `json:"message"`
		Slug    string `json:"slug"`
	} `json:"object_attributes"`
}
