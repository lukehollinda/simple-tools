package main

// torso reads the 'middle' of a file - the bytes around a given offset.
// it's not the head of the file, and it's not the tail - it's the torso.
// usage:
//
//	torso -offset n -before [b=128] -after [a=128] -from file [-newline]
//
// if no file is given, reads from standard input.

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

func usage() {
	fmt.Println("torso -offset n -before [b=128] -after [a=128] -file file [-newline]")
}

func main() {
	var offset, before, after int
	var filename string
	var newline bool
	flag.StringVar(&filename, "file", "", "file to read from; if empty read from stdin")
	flag.IntVar(&offset, "offset", -1, "offset to read from, must be specified")
	flag.IntVar(&before, "before", 128, "bytes to read before offset")
	flag.IntVar(&after, "after", 128, "bytes to read after offset")
	flag.BoolVar(&newline, "newline", false, "append newline to output")
	flag.Parse()

	// Validate args
	if offset < 0 {
		fmt.Printf("invalid offset: %v\n", offset)
		usage()
		os.Exit(1)
	}
	if filename == "" {
		fmt.Printf("please provide a filename\n")
		usage()
		os.Exit(1)
	}

	before = max(before, 0)
	before = min(before, offset) // Can't go beyond start of file
	after = max(after, 0)

	// Torso
	err := torso(filename, offset, before, after, newline)
	if err != nil {
		fmt.Printf("torso: %v\n", err)
		os.Exit(1)
	}
}

func torso(filename string, offset, before, after int, newline bool) error {
	start := offset - before
	size := before + after
	if size == 0 {
		// Nothing to do
		return nil
	}

	// File
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}
	defer func(){
		file.Close()
	}()

	buf := make([]byte, size)
	_, err = file.ReadAt(buf, int64(start))
	if err != nil && !errors.Is(err, io.EOF){
		return fmt.Errorf("read file: %w", err)
	}

	_, err = os.Stdout.Write(buf)
	if err != nil {
		return fmt.Errorf("write to stdout: %w", err)
	}

	if newline {
		fmt.Println()
	}

	return nil
}
