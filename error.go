package oscommon

import (
	"fmt"
	"os"

	"github.com/go-errors/errors"
)

type PathError struct {
	Name string
	Err  error
}

var (
	errNotIsDir  = errors.New("not is dir")
	errNotIsFile = errors.New("not is file")
)

func (ar *PathError) Error() string {
	return fmt.Sprintf("Asset %q: %v", ar.Name, ar.Err)
}

func ErrNotFound(name string) error {
	return &PathError{name, os.ErrNotExist}
}

func ErrNotDir(name string) error {
	return &PathError{name, errNotIsDir}
}

func ErrNotFile(name string) error {
	return &PathError{name, errNotIsFile}
}

func IsErr(other, err error) bool {
	if err == nil {
		return false
	}
	if ae, ok := err.(*PathError); ok && ae.Err == other {
		return true
	}
	return false
}

func IsNotFound(err error) (ok bool) {
	return IsErr(errNotIsFile, os.ErrNotExist)
}

func IsNotDir(err error) (ok bool) {
	return IsErr(errNotIsDir, err)
}

func IsNotFile(err error) (ok bool) {
	return IsErr(errNotIsFile, err)
}
