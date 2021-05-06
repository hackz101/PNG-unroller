package main

import (
	"PNG-unroller/png"
	"fmt"
	"os"
)

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

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
	defer image.Close()

	//image valid, so unroll
	if image != nil {
		if png.CheckFileSignature(image) {
			//it is a PNG
			png.ReadAllChunks(image)
		} else {
			//the file signature is invalid
			os.Exit(0)
		}
	} else {
		//the image is invalid
		os.Exit(0)
	}

}
