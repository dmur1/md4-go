package main

import (
	"fmt"
	"encoding/binary"
)

func md4PadMessage(message []byte) ([]byte) {
	messageLenghtBytes := len(message)
	messageLengthBits := messageLenghtBytes * 8

	numZeroBits := (448 - (messageLengthBits + 1)) % 512
	if numZeroBits < 0 {
		numZeroBits += 512
	}

	numBitsPadding := 64 + 1 + numZeroBits
	numPaddingBytes := numBitsPadding / 8

	paddedMessage := make([]byte, messageLenghtBytes + numPaddingBytes)

	copy(paddedMessage, message)

	paddedMessage[messageLenghtBytes] = 0x80

	binary.LittleEndian.PutUint64(
		paddedMessage[len(paddedMessage) - 8:], uint64(messageLengthBits))

	return paddedMessage
}

func md4(message []byte) ([]byte) {
	paddedMessage := md4PadMessage(message)

	numBlocks := len(paddedMessage) / 64

	A := uint32(0x67452301)
	B := uint32(0xefcdab89)
	C := uint32(0x98badcfe)
	D := uint32(0x10325476)

	var X [16]uint32

	for i := 0; i < numBlocks; i++ {
		for j := 0; j < 16; j++ {
			X[j] = binary.LittleEndian.Uint32(
				paddedMessage[(i * 64) + (j * 4):])
		}

		AA := A
		BB := B
		CC := C
		DD := D

		for j := 0; j < 4; j++ {
			A += ((B & C) | ((^B) & D)) + X[(j * 4) + 0]
			A = (A << 3) | (A >> (32 - 3))
			D += ((A & B) | ((^A) & C)) + X[(j * 4) + 1]
			D = (D << 7) | (D >> (32 - 7))
			C += ((D & A) | ((^D) & B)) + X[(j * 4) + 2]
			C = (C << 11) | (C >> (32 - 11))
			B += ((C & D) | ((^C) & A)) + X[(j * 4) + 3]
			B = (B << 19) | (B >> (32 - 19))
		}

		for j := 0; j < 4; j++ {
			A += ((B & C) | (B & D) | (C & D)) + X[0 + j] + 0x5A827999
			A = (A << 3) | (A >> (32 - 3))
			D += ((A & B) | (A & C) | (B & C)) + X[4 + j] + 0x5A827999
			D = (D << 5) | (D >> (32 - 5))
			C += ((D & A) | (D & B) | (A & B)) + X[8 + j] + 0x5A827999
			C = (C << 9) | (C >> (32 - 9))
			B += ((C & D) | (C & A) | (D & A)) + X[12 + j] + 0x5A827999
			B = (B << 13) | (B >> (32 - 13))
		}

		r3x := func(l uint32) uint32 {
			return ((l & 0x1) << 3) | ((l & 0x2) << 1) | ((l & 0x4) >> 1) | ((l & 0x8) >> 3)
		}

		for j := 0; j < 4; j++ {
			A += (B ^ C ^ D) + X[r3x(uint32(j * 4))] + 0x6ED9EBA1
			A = (A << 3) | (A >> (32 - 3))
			D += (A ^ B ^ C) + X[r3x(uint32(j * 4) + 1)] + 0x6ED9EBA1
			D = (D << 9) | (D >> (32 - 9))
			C += (D ^ A ^ B) + X[r3x(uint32(j * 4) + 2)] + 0x6ED9EBA1
			C = (C << 11) | (C >> (32 - 11))
			B += (C ^ D ^ A) + X[r3x(uint32(j * 4) + 3)] + 0x6ED9EBA1
			B = (B << 15) | (B >> (32 - 15))
		}

		A = AA + A
		B = BB + B
		C = CC + C
		D = DD + D
	}

	hash := make([]byte, 16)
	binary.LittleEndian.PutUint32(hash[0:], A)
	binary.LittleEndian.PutUint32(hash[4:], B)
	binary.LittleEndian.PutUint32(hash[8:], C)
	binary.LittleEndian.PutUint32(hash[12:], D)

	return hash
}

func main() {
	fmt.Printf("%x ~ 31d6cfe0d16ae931b73c59d7e0c089c0\n", md4([]byte("")))
	fmt.Printf("%x ~ bde52cb31de33e46245e05fbdbd6fb24\n", md4([]byte("a")))
	fmt.Printf("%x ~ a448017aaf21d8525fc10ae87aa6729d\n", md4([]byte("abc")))
	fmt.Printf("%x ~ d9130a8164549fe818874806e1c7014b\n", md4([]byte("message digest")))
	fmt.Printf("%x ~ d79e1c308aa5bbcdeea8ed63df412da9\n", md4([]byte("abcdefghijklmnopqrstuvwxyz")))
	fmt.Printf("%x ~ 043f8582f241db351ce627e153e7f0e4\n", md4([]byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")))
	fmt.Printf("%x ~ e33b4ddc9c38f2199c3e7b164fcc0536\n", md4([]byte("12345678901234567890123456789012345678901234567890123456789012345678901234567890")))
}
