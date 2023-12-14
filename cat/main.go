package main

import (
	"bufio"
	"fmt"
	"os"
)

// func to read and print file contents
func PrintFile(filename string) {
	file, err := os.Open(filename)
	// handle file opening error
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // print each line in file
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		// handle scanning error
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}
}

// func to read and print from stdinput
func PrintStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// print each line from stdinput
		fmt.Println(scanner.Text())
	}
	// scanner error
	if err := scanner.Err(); err != nil {
		// handle scanning error
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
	}
}

// main function
func main() {
	// check for cmd line args
	args := os.Args[1:]
	if len(args) > 1 {
		// if args are present , read and print content of each file
		for _, arg := range args {
			PrintFile(arg)
		}
	} else {
		// if no args are present, read and print from stdin
		PrintStdin()
	}
}
