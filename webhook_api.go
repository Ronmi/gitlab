package gitlab

import (
	"net/url"
	"strconv"
)

// ProjectHooks retrieves all webhooks of a project
func (g *GitLab) ProjectHooks(pid string, opts *ListOption) (ret []Webhook, page *Pagination, err error) {
	uri := g.uri("/projects/:id/hooks", map[string]string{":id": pid})
	resp, page, err := g.get(uri, opts)
	err = forgeRet(resp, &ret, err)
	return
}

// ProjectHookOption represents optional parameters needed for AddProjectHook
type ProjectHookOption struct {
	url                   string
	PushEvents            bool
	IssuesEvents          bool
	MergeRequestsEvents   bool
	TagPushEvents         bool
	NoteEvents            bool
	BuildEvents           bool
	PipelineEvents        bool
	WikiPageEvents        bool
	EnableSSLVerification bool
}

func (p *ProjectHookOption) Encode(v url.Values) url.Values {
	if p == nil {
		return url.Values{}
	}
	ret := v
	if v == nil {
		ret = url.Values{}
	}

	// id and url are required
	ret.Set("url", p.url)

	optBool(p.PushEvents, ret, "push_events")
	optBool(p.IssuesEvents, ret, "issues_events")
	optBool(p.MergeRequestsEvents, ret, "merge_requests_events")
	optBool(p.TagPushEvents, ret, "tag_push_events")
	optBool(p.NoteEvents, ret, "note_events")
	optBool(p.BuildEvents, ret, "build_events")
	optBool(p.PipelineEvents, ret, "pipeline_events")
	optBool(p.WikiPageEvents, ret, "wiki_page_events")

	// ssl verify defaults to true, so we must pass this param to set it false
	ret.Set("enable_ssl_verification", "false")
	optBool(p.EnableSSLVerification, ret, "enable_ssl_verification")
	return ret
}

// AddProjectHook add a webhook to project
func (g *GitLab) AddProjectHook(pid, url string, opts *ProjectHookOption) (ret Webhook, err error) {
	opt := opts
	if opts == nil {
		opt = &ProjectHookOption{}
	}
	opt.url = url

	uri := g.uri("/projects/:id/hooks", map[string]string{":id": pid})
	resp, _, err := g.postForm(uri, opt.Encode(nil))
	err = forgeRet(resp, &ret, err)
	return
}

// DeleteProjectHook deletes specified project webhook
func (g *GitLab) DeleteProjectHook(pid string, hid int) (err error) {
	uri := g.uri("/projects/:pid/hooks/:hid", map[string]string{":pid": pid, ":hid": strconv.Itoa(hid)})
	resp, _, err := g.delete(uri, nil)
	err = forgeRet(resp, nil, err)
	return
}

// EditProjectHook updates project hook info
func (g *GitLab) EditProjectHook(pid string, hid int, url string, opts *ProjectHookOption) (ret Webhook, err error) {
	uri := g.uri("/projects/:pid/hooks/:hid", map[string]string{":pid": pid, ":hid": strconv.Itoa(hid)})
	opt := opts
	if opts == nil {
		opt = &ProjectHookOption{}
	}
	opt.url = url
	resp, _, err := g.putForm(uri, opt.Encode(nil))
	err = forgeRet(resp, &ret, err)
	return
}
