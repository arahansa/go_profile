// flagstudy
package main

import (
	"flag"
	"fmt"
)

func main() {
	wordPtr := flag.String("p", "test", "a string")
	flag.Parse()
	fmt.Println("your flag is :", *wordPtr)
}
