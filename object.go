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

	n := len(from)
	if n > _GIT_OID_HEXSZ {
		n = _GIT_OID_HEXSZ
	}
	err := gitError(C.git_oid_fromstrn(oid.oid, cfrom, C.size_t(n)))
	if err != nil {
		return nil, err
	}
	return oid, nil
}

func (o Oid) String() string {
	if o.oid == nil {
		return "(invalid)"
	}
	ptr := C.git_oid_allocfmt(o.oid)
	defer C.free(unsafe.Pointer(ptr))
	return C.GoString(ptr)
}

type Object struct {
	obj *C.struct_git_object
}

func (o Object) Oid() Oid {
	oid := Oid{}
	oid.oid = C.git_object_id(o.obj)
	return oid
}

func (o Object) Type() string {
	return C.GoString(C.git_object_type2string(C.git_object_type(o.obj)))
}

func (o Object) String() string {
	return o.Oid().String()
}

}
