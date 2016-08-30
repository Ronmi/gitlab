package gitlab

import (
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

func TestListBranches(t *testing.T) {
	c := makeClient()
	brs, page, err := c.ListBranches(RepoID, nil)
	if err != nil {
		t.Fatalf("Unexpected error when calling /user: %s", err)
	}

	expectPage := Pagination{}
	if p := *page; !reflect.DeepEqual(expectPage, p) {
		t.Errorf("Pagination info should be zero value, got %#v", p)
	}

	// There should be only 2 branches, one is master and another is test_branch
	if l := len(brs); l != 2 {
		t.Errorf("Expected to have 2 branches, got %d", l)
	}
	if b := findBranch(brs, "master"); b == nil {
		t.Errorf("Branch master not found")
	}
	if b := findBranch(brs, "test_branch"); b == nil {
		t.Errorf("Branch test_branch not found")
	}
}
