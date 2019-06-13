package oscommon

import (
	"fmt"
	"os"

	"strings"

	"github.com/go-errors/errors"
)

type PathError struct {
	Name     string
	Err      error
	Messages []string
}

func (this *PathError) AddMessage(msg ...string) *PathError {
	this.Messages = append(this.Messages, msg...)
	return this
}

var (
	errNotIsDir  = errors.New("not is dir")
	errNotIsFile = errors.New("not is file")
)

func (ar *PathError) Error() string {
	return fmt.Sprintf("Asset %q: %v", ar.Name, strings.Join(append([]string{ar.Err.Error()}, ar.Messages...), " -- "))
}

func ErrNotFound(name string) error {
	return &PathError{Name: name, Err: os.ErrNotExist}
}

func ErrNotDir(name string) error {
	return &PathError{Name: name, Err: errNotIsDir}
}

func ErrNotFile(name string) error {
	return &PathError{Name: name, Err: errNotIsFile}
}

func IsErr(this error, err ...error) bool {
	if this == nil {
		return false
	}
	for _, err := range err {
		if err == nil {
			continue
		}
		if this == err {
			return true
		}
		if ae, ok := this.(*PathError); ok {
			if ae.Err == err {
				return true
			}
		}
		if comp, ok := err.(interface{ Is(error) bool }); ok && comp.Is(err) {
			return true
		}
	}
	return false
}

func IsNotFound(err error) (ok bool) {
	return IsErr(err, errNotIsFile, os.ErrNotExist)
}

func IsNotDir(err error) (ok bool) {
	return IsErr(err, errNotIsDir)
}

func IsNotFile(err error) (ok bool) {
	return IsErr(err, errNotIsFile)
}
