package gitlab

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
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

// helper to normalize GitLab return value
func normResp(c *http.Client, req *http.Request) (resp *http.Response, page *Pagination, err error) {
	if resp, err = c.Do(req); err != nil {
		return
	}
	if resp.StatusCode >= 400 {
		err = APIError(resp.StatusCode)
		return
	}

	page = ParsePagination(resp.Header)
	return
}

type APIError int

func (e APIError) Error() string {
	return "GitLab server returns status code " + strconv.Itoa(int(e))
}

const (
	ErrMissingAttr   APIError = 400 // at least one required attribute is missing
	ErrUnauthorized  APIError = 401 // need authorization
	ErrForbidden     APIError = 403 // action is not allowed for current user
	ErrNotFound      APIError = 404 // resource cannot be accessed
	ErrNotSupported  APIError = 405 // action is not supported
	ErrConflict      APIError = 409 // resource confliction
	ErrUnprocessable APIError = 422 // entity cannot be processed
	ErrServerError   APIError = 500 // internal server error
)

// GitLab maps all gitlab apis to method calls
type GitLab struct {
	d    requestDecorator
	c    *http.Client
	base string // base uri
	path string // api path
}

func (g *GitLab) uri(entry string) string {
	return g.base + "/" + g.path + entry
}
func (g *GitLab) do(req *http.Request) (*http.Response, *Pagination, error) {
	if g.d == nil { // raw client
		return normResp(g.c, req)
	}

	if err := g.d.decorate(req); err != nil {
		return nil, nil, err
	}
	return normResp(g.c, req)
}
func (g *GitLab) get(url string, opts APIOption) (resp *http.Response, page *Pagination, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) put(url string, opts APIOption) (resp *http.Response, page *Pagination, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("PUT", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) delete(url string, opts APIOption) (resp *http.Response, page *Pagination, err error) {
	url = forgeURL(url, opts)
	req, err := http.NewRequest("DELETE", url, nil)
	if err == nil {
		return
	}
	return g.do(req)
}
func (g *GitLab) post(url string, bodyType string, body io.Reader) (resp *http.Response, page *Pagination, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err == nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	return g.do(req)
}
func (g *GitLab) postForm(url string, data url.Values) (resp *http.Response, page *Pagination, err error) {
	return g.post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}

func newGitLab(base, path string, d requestDecorator, c *http.Client) *GitLab {
	client := c
	if c == nil {
		client = http.DefaultClient
	}

	return &GitLab{
		// strip trailing / from base
		base: strings.TrimRight(base, "/"),
		// strip heading and trailing / from path
		path: strings.TrimRight(strings.TrimLeft(path, "/"), "/"),
		d:    d,
		c:    client,
	}
}

// FromPAT creates GitLab API client which authorizes by personal access token
func FromPAT(base, path, token string, client *http.Client) *GitLab {
	return newGitLab(base, path, patDecorator(token), client)
}

// FromOAuth creates GitLab API client which authorized by oauth access token
func FromOAuth(base, path string, source oauth2.TokenSource, client *http.Client) *GitLab {
	return newGitLab(base, path, &oauthDecorator{
		source: oauth2.ReuseTokenSource(nil, source),
	}, client)
}

// RawClient creates non-wrapped GitLab API client.
// Since most apis need authorization, use this with caution.
//
// Note: passing client from oauth2.Config.Client to this will make api calls
// authorized by oauth token.
func RawClient(base, path string, client *http.Client) *GitLab {
	return newGitLab(base, path, nil, client)
}
