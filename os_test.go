package rwfs_test

import (
	"fmt"
	"os"

	"github.com/wwmoraes/go-rwfs"
)

func ExampleOSDirFS_error() {
	_, err := rwfs.OSDirFS("").OpenFile("test", os.O_CREATE|os.O_WRONLY, 0o640)
	if err != nil {
		fmt.Println(err)
	}

	// Output:
	// open : no such file or directory
}
