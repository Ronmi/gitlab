package gitlab

import (
	"net/http"

	"golang.org/x/oauth2"
)

type requestDecorator interface {
	decorate(req *http.Request) error
}

type patDecorator string

func (d patDecorator) decorate(req *http.Request) error {
	req.Header.Set("PRIVATE-TOKEN", string(d))
	return nil
}

type oauthDecorator struct {
	source oauth2.TokenSource
}

func (d *oauthDecorator) decorate(req *http.Request) error {
	t, err := d.source.Token()
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", t.Type()+" "+t.AccessToken)
	return nil
}
