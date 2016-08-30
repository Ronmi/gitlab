package gitlab

import (
	"reflect"
	"testing"
)

func TestWebhook(t *testing.T) {
	c := makeClient()
	WebhookTestList(t, c)
	_ = WebhookTestAdd(t, c, "http://localhost:1234")
}

func WebhookTestList(t *testing.T, c *GitLab) {
	hooks, page, err := c.ProjectHooks(RepoID, nil)
	if err != nil {
		t.Fatalf("Unexpected error when calling GET /projects/:id/hooks: %s", err)
	}

	if page.Total != 0 {
		t.Errorf("There should be no webhook, but got %#v in pagination info", page.Total)
	}

	// There should be no webhook
	if l := len(hooks); l != 0 {
		t.Errorf("Expected to have no hook, got %d", l)
	}
}

func WebhookTestAdd(t *testing.T, c *GitLab, url string) int {
	hook, err := c.AddProjectHook(RepoPath, url, &AddProjectHookOption{PushEvents: true})
	if err != nil {
		t.Fatalf("Unexpected error when calling POST /projects/hooks: %s", err)
	}

	expect := Webhook{
		ProjectID:  RepoID,
		URL:        url,
		PushEvents: true,

		// these fields are dynamically generated, so just copy
		CreatedAt: hook.CreatedAt,
		ID:        hook.ID,
	}

	if !reflect.DeepEqual(expect, hook) {
		t.Fatalf("Returned webhook info differences with expection: %#v", hook)
	}

	return hook.ID
}
