package main

import (
	"fmt"
	"os"
)

func usage() {
	fmt.Println("Usage: findoffset <filename> <string>")
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	filename := os.Args[1]
	find := os.Args[2]

	buf, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("open file: %v", err)
		os.Exit(1)
	}

	// Check all bytes
	for i := range(len(buf)-len(find)) {
		for j := range(len(find)) {
			if buf[i+j] != find[j] {
				break
			}

			if j == len(find) - 1 {
				// Found string
				fmt.Printf("%v", i)
				os.Exit(0)
			}

		}
	}

	os.Exit(1)
}


