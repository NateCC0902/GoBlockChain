package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand" 
)

func main() {

    
	s_key0, s_key1, pub_key0, pub_key1 := GenerateKey()
	sign := Sign(s_key0, s_key1, "hello world")

	fmt.Println(len(s_key0), len(s_key1), len(pub_key0), len(pub_key1), len(sign))
	// fmt.Printf("%x", sign[0])
    // fmt.Printf("%x", s_key0[0])
    fmt.Println(Verify(pub_key0, pub_key1,"hello world", sign))
    fmt.Println(signedMessage1[0])
}

func Block(length int) []byte {
	//a := rand.Intn(2)
	block := make([]byte, length)
    for i := 0; i < length; i++ {
		block[i] = byte(rand.Intn(2))
	}
	return block
}

func HashBlock(block []byte) [32]byte {
	str := fmt.Sprintf("%v", block)
	h := sha256.Sum256([]byte(str))
	return h
}

func GenerateKey() ([][]byte, [][]byte, [][32]byte, [][32]byte) {

	pair0 := make([][]byte, 256)
	pair1 := make([][]byte, 256)

	pub_pair0 := make([][32]byte, 256)
	pub_pair1 := make([][32]byte, 256)

	for i := 0; i < 256; i++ {
		pair0[i] = Block(256)
		pair1[i] = Block(256)

		pub_pair0[i] = HashBlock(pair0[i])
		pub_pair1[i] = HashBlock(pair1[i])
	}

	return pair0, pair1, pub_pair0, pub_pair1
}

func Sign(secret_key0 [][]byte, secret_key1 [][]byte, stringMessage string) [][]byte {
	message := []byte(stringMessage)
	h := sha256.Sum256([]byte(message))
	binaryStr := ""

	for i := 0; i < len(h); i++ {
		binaryStr += toBinaryBlock(h[i])
	}

	sign := make([][]byte, 256)
	for i := 0; i < len(binaryStr); i++ {
		if binaryStr[i] == '0' {
			sign[i] = secret_key0[i]
		} else if binaryStr[i] == '1' {
			sign[i] = secret_key1[i]
		}
	}

	return sign
}

func toBinaryBlock(input byte) string {
	str := fmt.Sprintf("%b", input)
	if len(str) < 8 {
		gap := 8 - len(str)
		for i := 0; i < gap; i++ {
			str = "0" + str
		}
	}

	return str
}

func Verify(pub_key0 [][32]byte,pub_key1 [][32]byte, message string, sign [][]byte ) bool {
    hashSign := make([][32]byte, 256)

   for i := 0; i < len(sign); i++ {
        hashSign[i] = HashBlock(sign[i])
   } 

    for i := 0; i < len(hashSign); i++ {
       for j := 0; j < len(hashSign[i]); j++ { 
           if hashSign[i][j] != pub_key0[i][j] && hashSign[i][j] != pub_key1[i][j] {
               return false
           }
       }
    }
    return true
}

