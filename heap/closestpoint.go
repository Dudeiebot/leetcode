package heap

import (
	"container/heap"
	"math"
)

type Point struct {
	distance float64
	x        int
	y        int
}
type PointHeap []Point

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return h[i].distance > h[j].distance
}

func (h PointHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PointHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}

func (h *PointHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func KCloestPoint(points [][]int, k int) [][]int {
	h := &PointHeap{}
	heap.Init(h)

	for _, point := range points {
		distance := math.Sqrt(float64(point[0]*point[0] + point[1]*point[1]))
		heap.Push(h, Point{distance: distance, x: point[0], y: point[1]})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	res := make([][]int, k)
	for i := 0; i < k; i++ {
		temp := heap.Pop(h).(Point)
		res[i] = []int{temp.x, temp.y}
	}
	return res
}
