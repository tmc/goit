package goit

//#include <git2.h>
import "C"
import "unsafe"

const _GIT_OID_HEXSZ = C.GIT_OID_HEXSZ

// Unique git object identifier (sha)
type Oid struct {
	oid *C.git_oid
}

func NewOid() *Oid {
	return &Oid{new(C.git_oid)}
}

func NewOidFromString(from string) (*Oid, error) {
	oid := NewOid()
	cfrom := C.CString(from)
	defer C.free(unsafe.Pointer(cfrom))

	err := gitError(C.git_oid_fromstr(oid.oid, cfrom))
	if err != nil {
		return nil, err
	}
	return oid, nil
}

func (o Oid) String() string {
	ptr := C.git_oid_allocfmt(o.oid)
	defer C.free(unsafe.Pointer(ptr))
	return C.GoString(ptr)
}

type ObjectType int

type GitObject interface {
	Oid() Oid
	Type() ObjectType
}

type Object struct {
	obj *C.struct_git_object
	oid Oid
}

func (o Object) Oid() Oid {
	return o.oid
}

func (o Object) Type() ObjectType {
	return _GIT_OBJ_ANY
}

type Commit struct {
	Object
}
