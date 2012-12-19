package goit

import (
	"os"
	"reflect"
	"testing"
)

func TestObject(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Error(err)
	}
	repo, err := OpenRepository(path)
	if err != nil {
		t.Error(err)
		return
	}
	h, _, err := repo.Head()
	if err != nil {
		t.Error(err)
		return
	}
	oid := h.Oid().String()
	o1, err := repo.LookupObject(oid)
	if err != nil {
		t.Fatal(err)
	}
	repo2 := o1.Repository()
	if !reflect.DeepEqual(repo, repo2) {
		t.Error(repo, "!=", repo2)
	}
}
