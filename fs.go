package rwfs

import "io/fs"

type FS interface {
	fs.FS

	OpenFile(name string, flag int, perm fs.FileMode) (File, error)
}
