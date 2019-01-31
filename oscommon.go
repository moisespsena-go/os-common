package oscommon

import (
	"os"
	"time"
)

type VirtualDirFileInfo struct {
	name string
}

func NewVirtualDirFileInfo(name string) os.FileInfo {
	return &VirtualDirFileInfo{name}
}

func (d VirtualDirFileInfo) Name() string {
	return d.name
}
func (d VirtualDirFileInfo) Size() int64 {
	return -1
}
func (d VirtualDirFileInfo) Mode() os.FileMode {
	return os.ModeDir | 0500
}
func (d VirtualDirFileInfo) ModTime() (t time.Time) {
	return t
}
func (d VirtualDirFileInfo) IsDir() bool {
	return true
}
func (d VirtualDirFileInfo) Sys() interface{} {
	return nil
}
