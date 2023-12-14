package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// no of args
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("File name missing")
	} else if len(args) > 1 {
		fmt.Println("Too many arguments")
	} else {
		// get filename from command line arguments
		filePath := args[0]
		// read content of file
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Print("Unable to read file:", err)
			return
		}
		// print content of standard output
		fmt.Print(string(content))
	}
}
