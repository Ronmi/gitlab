package gitlab

import "net/url"

// ListBranches retrieves all branches in a project
//
// See http://docs.gitlab.com/ce/api/branches.html#list-repository-branches
func (g *GitLab) ListBranches(pid string, opts *ListOption) (ret []Branch, page *Pagination, err error) {
	uri := g.uri("/projects/:id/repository/branches", map[string]string{":id": pid})
	resp, page, err := g.get(uri, opts)
	err = forgeRet(resp, &ret, err)
	return
}

// CreateBranch creates new branch in a project
//
// See http://docs.gitlab.com/ce/api/branches.html#create-repository-branch
func (g *GitLab) CreateBranch(pid string, name, ref string) (ret Branch, err error) {
	if ref == "" {
		ref = "master"
	}

	uri := g.uri("/projects/:id/repository/branches", map[string]string{":id": pid})

	opts := url.Values{}
	opts.Set("branch_name", name)
	opts.Set("ref", ref)

	resp, _, err := g.postForm(uri, opts)
	err = forgeRet(resp, &ret, err)
	return
}

// DeleteBranch deletes a branch from project
//
// See http://docs.gitlab.com/ce/api/branches.html#delete-repository-branch
func (g *GitLab) DeleteBranch(pid string, name string) (err error) {
	uri := g.uri("/projects/:id/repository/branches/:br", map[string]string{":id": pid, ":br": name})

	resp, _, err := g.delete(uri, nil)
	err = forgeRet(resp, nil, err)
	return
}
