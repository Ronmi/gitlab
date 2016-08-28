package gitlab

import (
	"encoding/json"
	"io/ioutil"
	"strconv"
)

// ListBranches retrieves all branches in a project
//
// See http://docs.gitlab.com/ce/api/branches.html#list-repository-branches
func (g *GitLab) ListBranches(pid int, opts *ListOption) (ret []Branch, page *Pagination, err error) {
	uri := g.uri("/projects/" + strconv.Itoa(pid) + "/repository/branches")
	resp, page, err := g.get(uri, opts)
	if err != nil {
		if _, ok := err.(APIError); ok {
			resp.Body.Close()
		}
		return
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(buf, &ret)
	return
}
