package png

import (
	"PNG-unroller/read"
	"bytes"
	"fmt"
	"os"
)

func CheckFileSignature(file *os.File) bool {
	fmt.Println("Checking File Signature")
	pngsignature := []byte{byte(137), byte(80), byte(78), byte(71), byte(13), byte(10), byte(26), byte(10)}
	if bytes.Equal(read.ReadSignature(file), pngsignature) {
		fmt.Println("Verified PNG")
		return true
	} else {
		fmt.Println("This is not in PNG format")
		return false
	}
}
