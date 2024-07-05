package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args[0])
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
