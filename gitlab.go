package gitlab

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// helper to forge query string
func forgeURL(url string, opts APIOption) string {
	if opts != nil {
		queryString := opts.Encode(nil).Encode()
		prefix := "?"
		if strings.Contains(url, "?") {
			prefix = "&"
		}
		url += prefix + queryString
	}
	return url
}

// GitLab maps all gitlab apis to method calls
type GitLab struct {
	d requestDecorator
	c *http.Client
}

func (g *GitLab) do(req *http.Request) (*http.Response, error) {
	if g.d == nil { // raw client
		return g.c.Do(req)
	}

	if err := g.d.decorate(req); err != nil {
		return nil, err
	}
	return g.c.Do(req)
}
func (g *GitLab) get(url string, opts APIOption) (resp *http.Response, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) put(url string, opts APIOption) (resp *http.Response, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("PUT", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) delete(url string, opts APIOption) (resp *http.Response, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("DELETE", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err == nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	return g.do(req)
}
func (g *GitLab) postForm(url string, data url.Values) (resp *http.Response, err error) {
	return g.post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func newGitLab(d requestDecorator, c *http.Client) *GitLab {
	client := c
	if c == nil {
		client = http.DefaultClient
	}
	return &GitLab{
		d: d,
		c: client,
	}
}

// FromPAT creates GitLab API client which authorizes by personal access token
func FromPAT(token string, client *http.Client) *GitLab {
	return newGitLab(patDecorator(token), client)
}

// FromOAuth creates GitLab API client which authorized by oauth access token
func FromOAuth(source oauth2.TokenSource, client *http.Client) *GitLab {
	return newGitLab(&oauthDecorator{
		source: oauth2.ReuseTokenSource(nil, source),
	}, client)
}

// RawClient creates non-wrapped GitLab API client.
// Since most apis need authorization, use this with caution.
//
// Note: passing client from oauth2.Config.Client to this will make api calls
// authorized by oauth token.
func RawClient(client *http.Client) *GitLab {
	return newGitLab(nil, client)
}
