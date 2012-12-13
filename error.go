package goit

//#include <git2.h>
import "C"

const (
	err_OK = -iota
	err_ERROR
	err_ENOTFOUND
	err_EEXISTS
	err_EAMBIGUOUS
	err_EBUFS
	err_EUSER
	err_EBAREREPO
	err_EORPHANEDHEAD
	err_EUNMERGED
	err_ENONFASTFORWARD
	err_EINVALIDSPEC
)

// Represents an error generated by libgit2
type GitError struct {
	message string
	klass   int
}

func (ge GitError) Error() string {
	return ge.message
}

// Returns a GitError or nil
func gitError(errorCode C.int) error {
	if errorCode != err_OK {
		ge := C.giterr_last()
		return &GitError{C.GoString(ge.message), int(ge.klass)}
	}
	return nil
}
