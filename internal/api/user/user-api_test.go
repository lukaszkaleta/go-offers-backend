package user

import "testing"

func Test(t *testing.T) {
	user := User{Archived: false}
	error := user.Archive()
	if error != nil {
		t.Error("Error should be nil")
	}
	if user.Archived == false {
		t.Fail()
	}
}
