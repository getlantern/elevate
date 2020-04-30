// this is a simple demo program that runs whatever command you give it with elevated privileges,
// like
//   ./demo ls /tmp
//
package main

import (
	"fmt"
	"os"

	"github.com/getlantern/elevate"
)

func main() {
	out, err := elevate.Command(os.Args[1], os.Args[2:]...).CombinedOutput()
	fmt.Println(string(out))
	if err != nil {
		fmt.Printf("Unexpected error '%v'\n", err)
		os.Exit(1)
	}
}
