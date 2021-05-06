package png

import (
	"PNG-unroller/read"
	"bytes"
	"fmt"
	"math"
	"os"
)

type Chunk struct {
	length   uint32
	typecode []byte
	data     []byte
	crc      []byte
}

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

func StringifyType(typecode []byte) string {
	//returns chunk type as string
	//originally string type is not used because png formats specifies not to
	return string(typecode)
}

func ReadChunk(file *os.File) Chunk {
	var chunk Chunk
	length := read.ReadUint32(file)

	//check if length meets png requirements
	if length > uint32(math.Pow(2, 32)-1) {
		fmt.Println("Chunk length exceeds maximum. Chunk reading failed")
		os.Exit(0)
	} else {
		//continue reading chunk
		chunk.length = length

		typecode := read.ReadLengthOfBytes(file, 4)
		chunk.typecode = typecode

		data := read.ReadLengthOfBytes(file, length)
		chunk.data = data

		crc := read.ReadLengthOfBytes(file, 4)
		chunk.crc = crc
	}

	return chunk
}

func ReadAllChunks(file *os.File) {
	//go through all chunks
	chunktype := ""
	for chunktype != "IEND" {
		//read next chunk
		chunk := ReadChunk(file)
		chunktype = StringifyType(chunk.typecode)
		fmt.Println(chunktype)

	}
}
