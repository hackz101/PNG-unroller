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

func ColorTypeBitDepthCheck(color uint8, bit uint8) bool {
	if color == 0 {
		if bit != 1 && bit != 2 && bit != 4 && bit != 8 && bit != 16 {
			return false
		}
	} else if color == 2 {
		if bit != 8 && bit != 16 {
			return false
		}
	} else if color == 3 {
		if bit != 1 && bit != 2 && bit != 4 && bit != 8 {
			return false
		}
	} else if color == 4 {
		if bit != 8 && bit != 16 {
			return false
		}
	} else {
		if bit != 8 && bit != 16 {
			return false
		}
	}
	return true
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
	} else if typecode == "IDAT" {

	} else {
		fmt.Println("Sorry haven't implemented this chunk processing yet!")
	}
}

func ProcessIHDR(chunk Chunk) IHDR {
	data := read.OpenBitstream(&chunk.data)
	var ihdr IHDR
	ihdr.width = read.ReadUint32Bitstream(&data)
	ihdr.height = read.ReadUint32Bitstream(&data)
	//size checks
	if (ihdr.width == 0 || ihdr.height == 0) || (ihdr.width > uint32(math.Pow(2, 32)-1) || ihdr.height > uint32(math.Pow(2, 32)-1)) {
		fmt.Println("Image size format error")
		os.Exit(0)
	}
	ihdr.bitdepth = read.ReadUint8Bitstream(&data)
	//value check
	if ihdr.bitdepth != 1 && ihdr.bitdepth != 2 && ihdr.bitdepth != 4 && ihdr.bitdepth != 8 && ihdr.bitdepth != 16 {
		fmt.Println("Bit depth formatting error")
		os.Exit(0)
	}
	ihdr.colortype = read.ReadUint8Bitstream(&data)
	//value check
	if ihdr.colortype != 0 && ihdr.colortype != 2 && ihdr.colortype != 3 && ihdr.colortype != 4 && ihdr.colortype != 6 {
		fmt.Println("Color type formatting error")
		os.Exit(0)
	}
	//is bit depth allowed based on color type?
	if !(ColorTypeBitDepthCheck(ihdr.colortype, ihdr.bitdepth)) {
		fmt.Println("Invalid bit depth for color type")
		os.Exit(0)
	}
	ihdr.compressionmethod = read.ReadUint8Bitstream(&data)
	//make sure method is 0 infalte/deflate
	if ihdr.compressionmethod != 0 {
		fmt.Println("Compression method not supported")
		os.Exit(0)
	}
	ihdr.filtermethod = read.ReadUint8Bitstream(&data)
	//make sure method is 0 adaptive filtering
	if ihdr.filtermethod != 0 {
		fmt.Println("Filter method not supported")
		os.Exit(0)
	}
	ihdr.interlacemethod = read.ReadUint8Bitstream(&data)
	//make sure method is 0 or 1 no interlace or Adam7 interlace
	if ihdr.interlacemethod != 0 && ihdr.interlacemethod != 1 {
		fmt.Println("Interlace method not supported")
		os.Exit(0)
	}
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
