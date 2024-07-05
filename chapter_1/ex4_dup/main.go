package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileLines := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileLines)
	} else {
		for _, arg := range os.Args[1:] {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Println(err)
			}
			countLines(f, counts, fileLines)
			f.Close()
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Println(fileLines[line])
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileLines map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		fileLines[input.Text()] = f.Name()
	}
}
