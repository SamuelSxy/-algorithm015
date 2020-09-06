import "container/heap"

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

//hashmap nlong(n) 不满足条件2
//@todo：快排方法
//@todo:heap的基础实现(大小顶）

// @lc code=start
func topKFrequent(nums []int, k int) []int {

	res := make([]int, k)
	timesList := make(map[int]int)
	for _, num := range nums {
		timesList[num]++
	}

	h := &Iheap{}
	heap.Init(h)

	for num, times := range timesList {

		heap.Push(h, [2]int{num, times})

		if h.Len() > k {
			heap.Pop(h)
		}

	}

	for i := 0; i < k; i++ {
		res[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return res

}

type Iheap [][2]int

func (h Iheap) Len() int {
	return len(h)
}

func (h Iheap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

func (h Iheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i] //语法糖
}

func (h *Iheap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *Iheap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end

