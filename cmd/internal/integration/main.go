package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/wwmoraes/go-rwfs"
)

const (
	testFileName  = "test.txt"
	testContent   = "lorem ipsum"
	testDirName   = "testdir"
	testNestedDir = "nested/dir"
)

func main() {
	defer handleExit()

	log.Println("creating temp directory")
	tmpDir, err := os.MkdirTemp("", "*")
	assert(err)

	log.Println("temp directory:", tmpDir)
	defer os.RemoveAll(tmpDir)

	fsys := rwfs.OSDirFS(tmpDir)

	log.Println("opening RW file")
	rwfd, err := fsys.OpenFile(testFileName, os.O_CREATE|os.O_WRONLY, 0640)
	assert(err)
	defer rwfd.Close()

	log.Println("writing test string")
	n, err := rwfd.WriteString(testContent)
	assert(err)

	log.Println("closing RW file")
	assert(rwfd.Close())

	log.Println("checking written bytes")
	if n != len(testContent) {
		assert(fmt.Errorf("wrote %d bytes, expected %d", n, len(testContent)))
	}

	log.Println("creating directory")
	err = fsys.MkdirAll(testDirName, 0755)
	assert(err)

	log.Println("creating nested directory")
	err = fsys.MkdirAll(testNestedDir, 0755)
	assert(err)

	log.Println("reading file contents")
	content, err := fs.ReadFile(fsys, testFileName)
	assert(err)

	log.Println("checking content")
	if string(content) != testContent {
		assert(fmt.Errorf("test file contains '%s', expected '%s'", string(content), testContent))
	}

	log.Println("done!")
}

func assert(err error) {
	if err == nil {
		return
	}

	log.SetOutput(os.Stderr)
	log.Println(err)
	panic(1)
}

func handleExit() {
	if e := recover(); e != nil {
		if exit, ok := e.(int); ok {
			os.Exit(exit)
		}

		panic(e)
	}
}
