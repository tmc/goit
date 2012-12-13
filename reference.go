package goit

//#cgo LDFLAGS: -lgit2
//#include <git2.h>
import "C"

import (
	"fmt"
)

// A git ref
type Reference struct {
	ref *C.git_reference
}

func (r Reference) String() string {
	return fmt.Sprintf("ref:%s", r.Name())
}

// Frees memory associated with the Reference. The repository should be
// considered invalid after this operation.
func (r Reference) Free() {
	C.git_repository_free(r.ref)
}

/* future api (not in stable release yet)
// Returns whether or not the reference is a branch.
func (r Reference) IsBranch() bool {
	branch := int(C.git_reference_is_branch(r.ref))
	return branch == 1
}

// Returns whether or not the reference is a remote ref.
func (r Reference) IsRemote() bool {
	branch := int(C.git_reference_is_remote(r.ref))
	return branch == 1
}
*/

// Returns the name of the reference
func (r Reference) Name() string {
	return C.GoString(C.git_reference_name(r.ref))
}

// Returns the Oid of the reference
//func (r Reference) Oid() OID {
//	return C.git_reference_oid(r.ref)
//}

// Returns the name of the Target of the reference. Only
// valid if the reference is symbolic.
func (r Reference) Target() string {
	if cs := C.git_reference_target(r.ref); cs != nil {
		return C.GoString(cs)
	}
	return ""
}

func (r Reference) Oid() *Oid {
	o := NewOid()
	o.oid = C.git_reference_oid(r.ref)
	//fmt.Println(o.oid, o)
	//if C.int(o.oid) == 0 {
	//		return nil
	//	}
	return o
}
