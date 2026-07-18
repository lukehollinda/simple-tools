package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)


func usage() {
	fmt.Println("usage: binpatch <file_name> <offset> <replacement_string>")

}

func main() {
	if len(os.Args) != 4 {
		usage()
		os.Exit(1)
	}

	filename := os.Args[1]
	offset, err := strconv.ParseInt(os.Args[2], 0, 64)
	if err != nil {
		log.Fatal("parsing offset argument:", err)
	}
	patch := os.Args[3]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("opening file:", err)
	}
	defer func()  {
		_ = file.Close()
	}()

	// Copy pre-offset
	_, err = io.CopyN(os.Stdout, file, offset)
	if err != nil {
		log.Fatal("writing file:", err)
	}

	// Write patch
	os.Stdout.Write([]byte(patch))

	// Discard len(patch) bytes
	_, err = io.CopyN(io.Discard, file, int64(len(patch)))
	if err != nil {
		log.Fatal("writing file:", err)
	}

	// Copy remaining file
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		log.Fatal("writing file:", err)
	}

	os.Exit(0)
}
