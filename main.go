package main

import (
	"fmt"

	"github.com/bamdadam/zipripgo/internal/compression/huffman"
)

func main() {
	he := huffman.HuffmanEncoder{}
	data := []byte{}
	for _, v := range "ABRACADABRA" {
		data = append(data, byte(v))
	}
	res, err := he.Encode(data)
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
	fmt.Println([]byte("ABRACADABRA"))
	fmt.Println(res)
}
