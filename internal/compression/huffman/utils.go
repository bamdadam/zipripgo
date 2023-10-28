package huffman

import (
	"container/heap"
	"fmt"
	"strconv"
	"strings"
)

func createHuffmanTree(fq frequencyCount, encode bool) *huffmanTree {
	tree := new(huffmanTree)
	elems := priorityQueue{}
	for b, c := range fq {
		elems = append(elems, pqElem{
			elemCount: c,
			elem:      []byte{b},
		})
	}
	heap.Init(&elems)
	for elems.Len() >= 2 {
		item1 := heap.Pop(&elems)
		item2 := heap.Pop(&elems)
		eb := []byte{}
		eb = append(append(eb, item1.(pqElem).elem...), item2.(pqElem).elem...)
		newElem := pqElem{
			elemCount: item1.(pqElem).elemCount + item2.(pqElem).elemCount,
			elem:      eb,
		}
		heap.Push(&elems, newElem)
	}
	root := heap.Pop(&elems).(pqElem)
	bts := map[byte]string{}
	stb := map[string]byte{}
	fmt.Println(root)
	for index, v := range root.elem {
		path := strings.Repeat("1", index)
		if index < len(root.elem)-1 {
			path += "0"
		}
		if encode {
			bts[v] = path
		} else {
			stb[path] = v
		}
	}
	tree.byteTostring = bts
	tree.stringToByte = stb
	return tree
}

func bitsToBytes(bits string) ([]byte, error) {
	var bytes []byte

	// Pad the string with leading zeros to make its length a multiple of 8
	for len(bits)%8 != 0 {
		bits = "0" + bits
	}

	// Iterate through the string, 8 bits at a time
	for i := 0; i < len(bits); i += 8 {
		chunk := bits[i : i+8]

		// Convert the 8-bit chunk to a byte
		val, err := strconv.ParseUint(chunk, 2, 8)
		if err != nil {
			return nil, err
		}

		bytes = append(bytes, byte(val))
	}

	return bytes, nil
}
