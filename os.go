package rwfs

import (
	"io/fs"
	"os"
)

type osFS string

// OSDirFS returns an OS file system (an fs.FS) for the tree of files rooted at
// the directory dir.
//
// The result implements [fs.FS] + OpenFile, which allows opening files for
// writing.
func OSDirFS(dir string) FS {
	return osFS(dir)
}

// Open opens the named file for reading. If successful, methods on
// the returned file can be used for reading; the associated file
// descriptor has mode [O_RDONLY].
// If there is an error, it will be of type [*PathError].
func (fsys osFS) Open(name string) (fs.File, error) {
	return os.OpenInRoot(string(fsys), name)
}

// OpenFile is the generalized open call; most users will use Open
// or Create instead. It opens the named file with specified flag
// ([O_RDONLY] etc.). If the file does not exist, and the [O_CREATE] flag
// is passed, it is created with mode perm (before umask);
// the containing directory must exist. If successful,
// methods on the returned File can be used for I/O.
// If there is an error, it will be of type [*PathError].
func (fsys osFS) OpenFile(name string, flag int, perm fs.FileMode) (File, error) {
	root, err := os.OpenRoot(string(fsys))
	if err != nil {
		return nil, err
	}

	return root.OpenFile(name, flag, perm)
}

// MkdirAll creates a new directory in the root, along with any necessary parents.
// See [MkdirAll] for more details.
//
// If perm contains bits other than the nine least-significant bits (0o777),
// MkdirAll returns an error.
func (fsys osFS) MkdirAll(path string, perm fs.FileMode) error {
	root, err := os.OpenRoot(string(fsys))
	if err != nil {
		return err
	}

	return root.MkdirAll(path, perm)
}
