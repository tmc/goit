package goit

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func mktmpDir() string {
	r, err := ioutil.TempDir("/tmp", "goit")
	if err != nil {
		panic(err)
	}
	return r
}

func TestOpenRepository(t *testing.T) {
	_, err := OpenRepository("nonexistent.git")
	expected := "Failed to resolve path 'nonexistent.git': No such file or directory"
	if expected != err.Error() {
		t.Errorf("error unexpected: '%s' != '%s'", err, expected)
	}
}

func TestInitRepository(t *testing.T) {
	tmpDir := mktmpDir()
	defer os.RemoveAll(tmpDir)

	r, err := InitRepository(tmpDir, false)
	if err != nil {
		t.Error(err)
	}
	if r.IsBare() == true {
		t.Error("repository should not be bare")
	}
	if cleanupPath(r.Workdir()) != cleanupPath(tmpDir) {
		t.Error("workdir mismatch, expected", tmpDir, "got", r.Workdir())
	}
	if r.Path() == "" {
		t.Error("path should not be empty")
	}
}

func TestInitRepositoryBare(t *testing.T) {
	tmpDir := mktmpDir()
	defer os.RemoveAll(tmpDir)
	r, err := InitRepository(tmpDir, true)
	if err != nil {
		t.Error(err)
	}
	if r.IsBare() == false {
		t.Error("repository should be bare")
	}
	if r.Workdir() != "" {
		t.Error("Workdir should be empty for a bare repository:", r.Workdir())
	}
	if r.IsEmpty() != true {
		t.Error("New bare repos should be empty")
	}
}

func TestExistingRepo(t *testing.T) {
	path := "/usr/local/src/git"
	r, err := OpenRepository(path)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(r)
	fmt.Println(r.Workdir())
	h, _, e := r.Head()
	if e != nil {
		t.Fatal(e)
	}
	fmt.Println("head", h)
	fmt.Println("target", h.Target())
	fmt.Println("oid", h.Oid())
	//s,e := r.LookupCommit(h.Oid())
	//fmt.Println("commit", s, e)
}
