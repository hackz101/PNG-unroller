package read

import (
	"os"
)

func ReadByte(file *os.File) byte {
	temp := make([]byte, 1)
	file.Read(temp)
	return temp[0]
}

func ReadSignature(file *os.File) []byte {
	temp := make([]byte, 8)
	file.Read(temp)
	return temp
}
