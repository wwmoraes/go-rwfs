package rwfs

import "io/fs"

// An FS provides access to a hierarchical file system.
//
// The FS interface is the minimum implementation required of the file system.
// A file system may implement additional interfaces,
// such as [ReadFileFS], to provide additional or optimized functionality.
//
// [testing/fstest.TestFS] may be used to test implementations of an FS for
// correctness.
type FS interface {
	fs.FS

	OpenFile(name string, flag int, perm fs.FileMode) (File, error)
	MkdirAll(path string, perm fs.FileMode) error
}
