package gitlab

// Me retrieve current user info
//
// See http://docs.gitlab.com/ce/api/users.html#current-user
func (g *GitLab) Me() (ret User, err error) {
	uri := g.uri("/user")
	resp, _, err := g.get(uri, nil)
	err = forgeRet(resp, &ret, err)
	return
}
