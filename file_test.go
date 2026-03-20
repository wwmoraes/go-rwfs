package rwfs_test

import (
	"fmt"
	"os"

	"github.com/wwmoraes/go-rwfs"
)

func ExampleFile_WriteString() {
	testContent := "lorem ipsum"

	tmpDir, err := os.MkdirTemp("", "*")
	if err != nil {
		fmt.Println(err)

		return
	}

	defer os.RemoveAll(tmpDir)

	fsys := rwfs.OSDirFS(tmpDir)

	rwfd, err := fsys.OpenFile("bar.txt", os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		fmt.Println(err)

		return
	}

	defer rwfd.Close()

	writtenBytes, err := rwfd.WriteString(testContent)
	if err != nil {
		fmt.Println(err)

		return
	}

	err = rwfd.Close()
	if err != nil {
		fmt.Println(err)

		return
	}

	if writtenBytes != len(testContent) {
		fmt.Printf("wrote %d bytes, expected %d\n", writtenBytes, len(testContent))

		return
	}

	// Output:
}
