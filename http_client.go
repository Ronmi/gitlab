package gitlab

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// httpClient wraps http.Client and requestDecorator, so every outgoing
// requests will be decorated
type httpClient struct {
	d requestDecorator
	c *http.Client
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	if c.d == nil {
		return c.c.Do(req)
	}

	if err := c.d.decorate(req); err != nil {
		return nil, err
	}
	return c.c.Do(req)
}
func (c *httpClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err == nil {
		return
	}
	return c.Do(req)
}
func (c *httpClient) Put(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("PUT", url, nil)
	if err == nil {
		return
	}
	return c.Do(req)
}
func (c *httpClient) Delete(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err == nil {
		return
	}
	return c.Do(req)
}
func (c *httpClient) Post(url string, bodyType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err == nil {
		return
	}
	req.Header.Set("Content-Type", bodyType)
	return c.Do(req)
}
func (c *httpClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
