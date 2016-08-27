package glresp

// Branch represents info of a branch
type Branch struct {
	Name               string `json:"name"`
	Protected          bool   `json:"protected,omitempty"`
	DevelopersCanPush  bool   `json:"developers_can_push,omitempty"`
	DevelopersCanMerge bool   `json:"developers_can_merge,omitempty"`
	Commit             Commit `json:"commit"`
}
