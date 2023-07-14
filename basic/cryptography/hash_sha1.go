package cryptography

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func Gorunhash() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("We shall overcome!")
	n, err := hasher.Write(data)
	if n != len(data) || err != nil {
		log.Printf("Hash write error: %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("Result: %x\n", checksum)

	hasher1 := md5.New()
	b1 := []byte{}
	io.WriteString(hasher1, "test")
	fmt.Printf("Result: %x\n", hasher1.Sum(b1))
	fmt.Printf("Result: %d\n", hasher1.Sum(b1))
}
