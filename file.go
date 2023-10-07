package rwfs

import (
	"io"
	"io/fs"
)

type File interface {
	fs.File
	io.Seeker
	io.StringWriter
	io.Writer
	io.WriterAt
}
