package read

import (
	"os"
)

//used for reading functions that read from byte arrays
//i = index, v = value
type Bitstream struct {
	currentbyte byte
	byteI       uint32
	bitI        uint8
	bitV        byte
	data        *[]byte
}

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

func readBit(b byte, pos uint8) byte {
	return ((b >> pos) & byte(1))
}

func OpenBitstream(data *[]byte) Bitstream {
	var temp Bitstream
	temp.data = data
	temp.currentbyte = (*temp.data)[0]
	temp.byteI = 0
	temp.bitI = 0
	temp.bitV = readBit(temp.currentbyte, temp.bitI)
	return temp
}

func ReadByteBitstream(stream *Bitstream) byte {
	var temp byte
	if stream.byteI < uint32(len(*stream.data)-1) {
		temp = stream.currentbyte
		stream.byteI++
		stream.currentbyte = (*stream.data)[stream.byteI]
	} else if stream.byteI == uint32(len(*stream.data)-1) {
		temp = stream.currentbyte
	}
	return temp
}

func ReadUint8Bitstream(stream *Bitstream) byte {
	var temp byte
	if stream.byteI < uint32(len(*stream.data)-1) {
		temp = stream.currentbyte
		stream.byteI++
		stream.currentbyte = (*stream.data)[stream.byteI]
	} else if stream.byteI == uint32(len(*stream.data)-1) {
		temp = stream.currentbyte
	}
	return uint8(temp)
}

func ReadUint32Bitstream(stream *Bitstream) uint32 {
	var temp uint32
	if stream.byteI <= uint32(len(*stream.data)-4) {
		temp |= uint32((*stream.data)[stream.byteI]) << 24
		temp |= uint32((*stream.data)[stream.byteI+1]) << 16
		temp |= uint32((*stream.data)[stream.byteI+2]) << 8
		temp |= uint32((*stream.data)[stream.byteI+3])
		stream.byteI += 4
		stream.currentbyte = (*stream.data)[stream.byteI]
	} else if stream.byteI == uint32(len(*stream.data)-4) {
		temp |= uint32((*stream.data)[stream.byteI]) << 24
		temp |= uint32((*stream.data)[stream.byteI+1]) << 16
		temp |= uint32((*stream.data)[stream.byteI+2]) << 8
		temp |= uint32((*stream.data)[stream.byteI+3])
	}
	return temp
}
