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

func ReadUint32(file *os.File) uint32 {
	temp := make([]byte, 4)
	file.Read(temp)
	var retval uint32
	retval |= uint32(temp[0]) << 24
	retval |= uint32(temp[1]) << 16
	retval |= uint32(temp[2]) << 8
	retval |= uint32(temp[3])
	return retval
}

func ReadLengthOfBytes(file *os.File, length uint32) []byte {
	temp := make([]byte, length)
	file.Read(temp)
	return temp
}
