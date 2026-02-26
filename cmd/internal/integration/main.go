// Binary integration uses go-rwfs in an example usage of it. It also doubles as
// an integration test coverage, hence its name.
package main

import (
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"runtime"
	"sync/atomic"

	"github.com/wwmoraes/go-rwfs"
)

const (
	testFileName  = "test.txt"
	testContent   = "lorem ipsum"
	testDirName   = "testdir"
	testNestedDir = "nested/dir"
)

var (
	ErrAssertion       = errors.New("assertion failed")
	ErrContentMismatch = errors.New("content mismatch")
	exitCode           atomic.Int32
)

//nolint:funlen // TODO refactor
func main() {
	defer func() {
		os.Exit(int(exitCode.Load()))
	}()

	_, err := rwfs.OSDirFS("").OpenFile("test", os.O_CREATE|os.O_WRONLY, 0o640)
	if err == nil {
		AssertWith(ErrAssertion, "opening empty root name")
	}

	log.Println("creating temp directory")

	tmpDir, err := os.MkdirTemp("", "*")
	AssertWith(err, "making temporary directory tree")

	log.Println("temp directory:", tmpDir)
	defer os.RemoveAll(tmpDir)

	fsys := rwfs.OSDirFS(tmpDir)

	log.Println("opening RW file")

	rwfd, err := fsys.OpenFile(testFileName, os.O_CREATE|os.O_WRONLY, 0o640)
	AssertWith(err, "opening file for write")

	defer rwfd.Close()

	log.Println("writing test string")

	writtenBytes, err := rwfd.WriteString(testContent)
	AssertWith(err, "writing content")

	log.Println("closing RW file")
	AssertWith(rwfd.Close(), "closing file")

	log.Println("checking written bytes")

	if writtenBytes != len(testContent) {
		AssertWith(io.ErrShortWrite, fmt.Sprintf("wrote %d bytes, expected %d", writtenBytes, len(testContent)))
	}

	log.Println("creating directory")

	err = fsys.MkdirAll(testDirName, 0o755)
	AssertWith(err, "creating directory")

	log.Println("creating nested directory")

	err = fsys.MkdirAll(testNestedDir, 0o755)
	AssertWith(err, "creating nested directory")

	log.Println("reading file contents")

	content, err := fs.ReadFile(fsys, testFileName)
	AssertWith(err, "reading file")

	log.Println("checking content")

	if string(content) != testContent {
		AssertWith(ErrContentMismatch, fmt.Sprintf("test file contains '%s', expected '%s'", string(content), testContent))
	}

	log.Println("done!")
}

func AssertWith(err error, message string) {
	if err == nil {
		return
	}

	log.SetOutput(os.Stderr)
	log.Printf("%s: %s\n", message, err)
	exitCode.CompareAndSwap(0, 1)
	runtime.Goexit()
}
