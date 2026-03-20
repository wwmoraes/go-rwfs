package rwfs_test

import (
	"fmt"
	"os"

	"github.com/wwmoraes/go-rwfs"
)

func ExampleFS_Open_error() {
	tmpDir, err := os.MkdirTemp("", "*")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer os.RemoveAll(tmpDir)

	fsys := rwfs.OSDirFS(tmpDir)

	fd, err := fsys.Open("test.txt")
	if err != nil {
		fmt.Println(err)

		return
	}

	fd.Close()

	// Output:
	// openat test.txt: no such file or directory
}

func ExampleFS_MkdirAll() {
	tmpDir, err := os.MkdirTemp("", "*")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer os.RemoveAll(tmpDir)

	fsys := rwfs.OSDirFS(tmpDir)

	err = fsys.MkdirAll("foo/bar", 0o750)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
}

func ExampleFS_MkdirAll_root_error() {
	fsys := rwfs.OSDirFS("")

	err := fsys.MkdirAll("foo/bar", 0o750)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// open : no such file or directory
}
