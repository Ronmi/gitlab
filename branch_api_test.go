package gitlab

import (
	"fmt"
	"reflect"
	"testing"
)

// find a branch in branch array and return pointer to it, nil if not found
func findBranch(brs []Branch, name string) *Branch {
	for i, b := range brs {
		if b.Name == name {
			return &(brs[i])
		}
	}
	return nil
}

func TestBranch(t *testing.T) {
	c := makeClient()
	if err := BranchTestList(c); err != "" {
		t.Fatal(err)
	}

	if err := BranchTestCreate(c, "orz"); err != "" {
		t.Fatal(err)
	}
	if err := BranchTestDelete(c, "orz"); err != "" {
		t.Fatal(err)
	}
}

func BranchTestList(c *GitLab) string {
	brs, page, err := c.ListBranches(RepoID, nil)
	if err != nil {
		return fmt.Sprintf("Unexpected error when calling GET /projects/:id/branches: %s", err)
	}

	expectPage := Pagination{}
	if p := *page; !reflect.DeepEqual(expectPage, p) {
		return fmt.Sprintf("Pagination info should be zero value, got %#v", p)
	}

	// There should be only 2 branches, one is master and another is test_branch
	if l := len(brs); l != 2 {
		return fmt.Sprintf("Expected to have 2 branches, got %d", l)
	}
	if b := findBranch(brs, "master"); b == nil {
		return fmt.Sprintf("Branch master not found")
	}
	if b := findBranch(brs, "test_branch"); b == nil {
		return fmt.Sprintf("Branch test_branch not found")
	}

	return ""
}

func BranchTestCreate(c *GitLab, name string) string {
	br, err := c.CreateBranch(RepoID, name, "") // based on master
	if err != nil {
		return fmt.Sprintf("Unexpected error occured when creating branch: %s", err)
	}
	if br.Name != name {
		return fmt.Sprintf("Expected to create branch %s, but %s created instead", name, br.Name)
	}

	brs, _, err := c.ListBranches(RepoID, nil)
	if err != nil {
		return fmt.Sprintf("Unexpected error when fetching branches: %s", err)
	}

	if l := len(brs); l != 3 {
		return fmt.Sprintf("Expected to have 3 branches, got %d", l)
	}
	if b := findBranch(brs, name); b == nil {
		return fmt.Sprintf("Expected to have %s branch, but not found", name)
	}

	return ""
}

func BranchTestDelete(c *GitLab, name string) string {
	err := c.DeleteBranch(RepoID, name)
	if err != nil {
		return fmt.Sprintf("Unexpected error occured when deleteing branch: %s", err)
	}

	brs, _, err := c.ListBranches(RepoID, nil)
	if err != nil {
		return fmt.Sprintf("Unexpected error when fetching branches: %s", err)
	}

	if l := len(brs); l != 2 {
		return fmt.Sprintf("Expected to have 2 branches, got %d", l)
	}
	if b := findBranch(brs, name); b != nil {
		return fmt.Sprintf("Expected not to have %s branch, but found", name)
	}

	return ""
}
