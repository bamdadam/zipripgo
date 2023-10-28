package huffman

type priorityQueue []pqElem

type pqElem struct {
	elemCount int
	elem      []byte
}

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].elemCount < pq[j].elemCount
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x any) {
	item := x.(pqElem)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	// old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}
