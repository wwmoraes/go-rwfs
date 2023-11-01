package rwfs

import (
	"io"
	"io/fs"
	"syscall"
)

// We do like the standard os package and re-export the open syscall Flags used
// by OpenFile. Not all flags may be implemented on a given system.
const (
	// O_RDONLY open the file read-only.
	//
	// O_RDONLY, O_WRONLY and O_RDWR are mutually exclusive.
	O_RDONLY int = syscall.O_RDONLY

	// O_RDONLY open the file write-only.
	//
	// O_RDONLY, O_WRONLY and O_RDWR are mutually exclusive.
	O_WRONLY int = syscall.O_WRONLY

	// O_RDONLY open the file read-write.
	//
	// O_RDONLY, O_WRONLY and O_RDWR are mutually exclusive.
	O_RDWR int = syscall.O_RDWR

	// O_APPEND append data to the file when writing.
	O_APPEND int = syscall.O_APPEND

	// O_CREATE create a new file if none exists.
	O_CREATE int = syscall.O_CREAT

	// O_EXCL used with O_CREATE, file must not exist.
	O_EXCL int = syscall.O_EXCL

	// O_SYNC open for synchronous I/O.
	O_SYNC int = syscall.O_SYNC

	// O_TRUNC truncate regular writable file when opened.
	O_TRUNC int = syscall.O_TRUNC
)

type File interface {
	fs.File
	io.Seeker
	io.StringWriter
	io.Writer
	io.WriterAt
}
