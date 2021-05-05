package main

import (
	"fmt"
	"os"
)

func checkForArgs(args []string) *os.File {
	fmt.Println("==PNG UNROLLER==")
	//handle argument input
	if len(args) == 1 {
		//no external arguments added, get input
		var filename string
		fmt.Print("Image to unroll: ")
		fmt.Scanln(&filename)

		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("There was a problem opening the png.")
		} else {
			fmt.Println("Successfully opened " + filename)
		}
		return file

	} else {
		//external arguments added, no input needed
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Println("There was a problem opening the png.")
		} else {
			fmt.Println("Successfully opened " + args[1])
		}
		return file
	}
}

func main() {
	image := checkForArgs(os.Args)
	image.Close()
}
