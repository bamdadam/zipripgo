package huffman

type frequencyCount map[byte]int

type huffmanTree struct {
	byteTostring map[byte]string
	stringToByte map[string]byte
}

type HuffmanEncoder struct {
	pq priorityQueue
}

func (h *HuffmanEncoder) Encode(data []byte) ([]byte, error) {
	frequencyMap := h.frequency_count(data)
	hft := createHuffmanTree(frequencyMap, true)
	res := ""
	for _, v := range data {
		res += hft.byteTostring[v]
	}
	return bitsToBytes(res)
}

func (h *HuffmanEncoder) frequency_count(data []byte) frequencyCount {
	res := map[byte]int{}
	for _, v := range data {
		res[v] += 1
	}
	return res
}
