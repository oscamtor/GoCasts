package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	//fmt.Println(os.Args)

	// Arg[0] is executable => Arg[1] is myfile.txt
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
