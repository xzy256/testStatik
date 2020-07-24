//go:generate statik -src=../assets
//go:generate go fmt statik/statik.go
package pkg

import (
	"fmt"
	"github.com/rakyll/statik/fs"
	"io/ioutil"
	"os"

	_ "testStatik/pkg/statik"
)

func Test() {
	statikFS, err := fs.New()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	file, err := statikFS.Open("/1.sh")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("content: %s\n", content)
}
