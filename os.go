package rwfs

import (
	"io/fs"
	"os"
	"path/filepath"
)

type osFS string

func OSDirFS(dir string) FS {
	return osFS(dir)
}

func (fsys osFS) Open(name string) (fs.File, error) {
	return os.Open(filepath.Join(string(fsys), name))
}

func (fsys osFS) OpenFile(name string, flag int, perm fs.FileMode) (File, error) {
	return os.OpenFile(filepath.Join(string(fsys), name), flag, perm)
}
