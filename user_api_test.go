package gitlab

import (
	"reflect"
	"testing"
)

func TestUserMe(t *testing.T) {
	c := makeClient()
	u, err := c.Me()
	if err != nil {
		t.Fatalf("Unexpected error when calling /user: %s", err)
	}

	// since content of user varies in different environment, just test if it's empty
	if reflect.DeepEqual(u, User{}) {
		t.Error("Me() returns zero value")
	}
}
