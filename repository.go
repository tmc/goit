package goit

//#cgo LDFLAGS: -lgit2
//#include <git2.h>
import "C"

import (
	"fmt"
	"unsafe"
)

// A git repository
type Repository struct {
	repo *C.git_repository
}

// Open an existing git repository located at the given filesystem path.
func OpenRepository(path string) (*Repository, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	r := new(Repository)
	if err := gitError(C.git_repository_open(&r.repo, p)); err != nil {
		return nil, err
	}
	return r, nil
}

// Initalize a new git repository at the given filesystem path.
// The bare flag controls if the repository should be created as bare or not.
func InitRepository(path string, bare bool) (*Repository, error) {
	p := C.CString(path)
	defer C.free(unsafe.Pointer(p))

	r := new(Repository)

	var bareFlag C.uint
	if bare {
		bareFlag = 1
	}
	if err := gitError(C.git_repository_init(&r.repo, p, bareFlag)); err != nil {
		return nil, err
	}
	return r, nil
}

// Frees memory associated with the repository. The repository handle should
// be considered invalid after this operation.
func (r Repository) Free() {
	C.git_repository_free(r.repo)
}

// Returns the Commit that is the current HEAD of the repository.
// If orphaned is true the Commit is
func (r Repository) Head() (head *Reference, orphaned bool, err error) {
	ref := new(Reference)
	if err := gitError(C.git_repository_head(&ref.ref, r.repo)); err != nil {
		if err.(*GitError).klass == err_EORPHANEDHEAD {
			orphaned = true
		} else {
			return nil, false, err
		}
	}
	// sucessfully got head ref
	return ref, orphaned, nil
}

// Returns whether or not a repository is bare.
func (r Repository) IsBare() bool {
	bare := int(C.git_repository_is_bare(r.repo))
	return bare == 1
}

// Returns whether or not a repository is empty.
func (r Repository) IsEmpty() bool {
	empty := int(C.git_repository_is_empty(r.repo))
	return empty == 1
}

// Returns the filesystem path the repository resides at.
func (r Repository) Path() string {
	return C.GoString(C.git_repository_path(r.repo))
}

// Returns the working directory associated with the repository
// Will be an empty string if the repository is bare.
func (r Repository) Workdir() string {
	return C.GoString(C.git_repository_workdir(r.repo))
}

func (r Repository) String() string {
	return fmt.Sprintf("<goit.Repository %s>", r.Path())
}

func (r Repository) LookupObject(id string) (*Object, error) {
	o := new(Object)
	oid, err := NewOidFromString(id)
	if err != nil {
		return nil, err
	}
	if err := gitError(C.git_object_lookup_prefix(&o.obj, r.repo, oid.oid, C.uint(len(id)), _GIT_OBJ_ANY)); err != nil {
		return nil, err
	}
	return o, nil
}
