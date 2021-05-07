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

type IHDR struct {
	width             uint32
	height            uint32
	bitdepth          uint8
	colortype         uint8
	compressionmethod uint8
	filtermethod      uint8
	interlacemethod   uint8
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
	//originally string type is not used because png format specifies not to
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

func ProccessChunk(chunk Chunk) {
	//check chunk type
	typecode := StringifyType(chunk.typecode)
	if typecode == "IHDR" {
		ihdr := ProcessIHDR(chunk)
		fmt.Println("\tWidth: " + fmt.Sprint(ihdr.width))
		fmt.Println("\tHeight: " + fmt.Sprint(ihdr.height))
		fmt.Println("\tBit depth: " + fmt.Sprint(ihdr.bitdepth))
		fmt.Println("\tColor type: " + fmt.Sprint(ihdr.colortype))
		fmt.Println("\tCompression method: " + fmt.Sprint(ihdr.compressionmethod))
		fmt.Println("\tFilter method: " + fmt.Sprint(ihdr.filtermethod))
		fmt.Println("\tInterlace method: " + fmt.Sprint(ihdr.interlacemethod))
	} else {
		fmt.Println("Sorry haven't implemented this chunk processing yet!")
	}
}

func ProcessIHDR(chunk Chunk) IHDR {
	data := read.OpenBitstream(&chunk.data)
	var ihdr IHDR
	ihdr.width = read.ReadUint32Bitstream(&data)
	ihdr.height = read.ReadUint32Bitstream(&data)
	ihdr.bitdepth = read.ReadUint8Bitstream(&data)
	ihdr.colortype = read.ReadUint8Bitstream(&data)
	ihdr.compressionmethod = read.ReadUint8Bitstream(&data)
	ihdr.filtermethod = read.ReadUint8Bitstream(&data)
	ihdr.interlacemethod = read.ReadUint8Bitstream(&data)
	return ihdr
}

func ReadAllChunks(file *os.File) {
	chunknum := 0 //used to check first chunk

	//go through all chunks
	chunktype := ""
	for chunktype != "IEND" {
		chunknum++
		//read next chunk
		chunk := ReadChunk(file)
		chunktype = StringifyType(chunk.typecode)
		//make sure IHDR is first chunk
		if chunknum == 1 && chunktype != "IHDR" {
			fmt.Println("Misformatted png file.")
			os.Exit(0)
		}
		fmt.Println("Reading " + chunktype + "...")

		//do chunk data manipulation here
		ProccessChunk(chunk)
	}
}
