package gitlab

import "strconv"

// ListBranches retrieves all branches in a project
//
// See http://docs.gitlab.com/ce/api/branches.html#list-repository-branches
func (g *GitLab) ListBranches(pid int, opts *ListOption) (ret []Branch, page *Pagination, err error) {
	uri := g.uri("/projects/:id/repository/branches", map[string]string{":id": strconv.Itoa(pid)})
	resp, page, err := g.get(uri, opts)
	err = forgeRet(resp, &ret, err)
	return
}
