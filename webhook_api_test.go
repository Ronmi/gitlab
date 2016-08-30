package gitlab

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWebhook(t *testing.T) {
	c := makeClient()
	var err string
	if err = WebhookTestList(c); err != "" {
		t.Fatal(err)
	}
	hookid, err := WebhookTestAdd(c, "http://localhost:1234")
	if err != "" {
		t.Fatal(err)
	}
	if err = WebhookTestEdit(c, hookid, "http://google.com"); err != "" {
		t.Fatal(err)
	}
	if err = WebhookTestDelete(c, hookid); err != "" {
		t.Fatal(err)
	}
}

func WebhookTestList(c *GitLab) string {
	hooks, page, err := c.ProjectHooks(RepoID, nil)
	if err != nil {
		return fmt.Sprintf("Unexpected error when calling GET /projects/:id/hooks: %s", err)
	}

	if page.Total != 0 {
		return fmt.Sprintf("There should be no webhook, but got %#v in pagination info", page.Total)
	}

	// There should be no webhook
	if l := len(hooks); l != 0 {
		return fmt.Sprintf("Expected to have no hook, got %d", l)
	}
	return ""
}

func WebhookTestAdd(c *GitLab, url string) (int, string) {
	hook, err := c.AddProjectHook(RepoPath, url, &ProjectHookOption{PushEvents: true})
	if err != nil {
		return 0, fmt.Sprintf("Unexpected error when calling POST /projects/hooks: %s", err)
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
		return 0, fmt.Sprintf("Returned webhook info differences with expection: %#v", hook)
	}

	return hook.ID, ""
}

func WebhookTestDelete(c *GitLab, id int) string {
	err := c.DeleteProjectHook(RepoPath, id)
	if err != nil {
		return fmt.Sprintf("Unexpected error when calling DELETE /porjects/:pid/hooks/:hid: %s", err)
	}

	// check if we really deleted the webhook
	hooks, _, err := c.ProjectHooks(RepoID, nil)
	if err != nil {
		return fmt.Sprintf("Unexpected error when checking result of DELETE /porjects/:pid/hooks/:hid: %s", err)
	}
	if l := len(hooks); l > 0 {
		return fmt.Sprintf("Expected no webhook left after deletion, got %d", l)
	}
	return ""
}

func WebhookTestEdit(c *GitLab, id int, url string) string {
	hook, err := c.EditProjectHook(RepoPath, id, url, &ProjectHookOption{TagPushEvents: true})
	if err != nil {
		return fmt.Sprintf("Unexpected error when calling POST /projects/hooks: %s", err)
	}

	expect := Webhook{
		ProjectID:     RepoID,
		URL:           url,
		PushEvents:    true,
		TagPushEvents: true,
		ID:            id,

		// these fields are dynamically generated, so just copy
		CreatedAt: hook.CreatedAt,
	}

	if !reflect.DeepEqual(expect, hook) {
		return fmt.Sprintf("Returned webhook info differences with expection: %#v", hook)
	}

	return ""
}
