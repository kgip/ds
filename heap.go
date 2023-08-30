package ds

type Heap struct {
	Cap             int //容量
	Count           int //元素个数
	Buffer          []int
	IsSmallRootHeap bool //是否小根堆
}

func NewHeap(cap int) *Heap {
	return &Heap{Cap: cap, IsSmallRootHeap: true, Buffer: make([]int, cap)}
}

func NewBigRootHeap(cap int) *Heap {
	return &Heap{Cap: cap}
}

// Add 往堆中添加一个元素
func (h *Heap) Add(item int) {
	if h.Count >= h.Cap {
		return
	}
	//将元素放到堆尾
	h.Buffer[h.Count] = item
	h.Count++
	//调整堆
	var rootIndex, currIndex = 0, h.Count - 1
	for {
		if currIndex%2 == 0 { //下标为偶数
			rootIndex = (currIndex - 2) / 2
		} else {
			rootIndex = (currIndex - 1) / 2
		}
		if rootIndex < 0 {
			break
		}
		if h.IsSmallRootHeap && h.Buffer[rootIndex] > item || !h.IsSmallRootHeap && h.Buffer[rootIndex] < item {
			h.Buffer[rootIndex], h.Buffer[currIndex] = h.Buffer[currIndex], h.Buffer[rootIndex]
			currIndex = rootIndex
		} else {
			break
		}
	}
}

// Pop 弹出堆顶元素
func (h *Heap) Pop() int {
	if h.Count <= 0 {
		return 0
	}
	//获取堆顶元素
	top := h.Buffer[0]
	//将堆尾元素放到堆顶，调整堆
	h.Buffer[0] = h.Buffer[h.Count-1]
	h.Count--
	var currIndex int
	for currIndex <= h.Count-1 {
		leftIndex, rightIndex := currIndex*2+1, currIndex*2+2
		if h.IsSmallRootHeap {
			if rightIndex <= h.Count-1 {
				//选出一个较小者
				if h.Buffer[leftIndex] < h.Buffer[rightIndex] {
					if h.Buffer[leftIndex] < h.Buffer[currIndex] {
						h.Buffer[currIndex], h.Buffer[leftIndex] = h.Buffer[leftIndex], h.Buffer[currIndex]
						currIndex = leftIndex
					} else {
						break
					}
				} else {
					if h.Buffer[rightIndex] < h.Buffer[currIndex] {
						h.Buffer[currIndex], h.Buffer[rightIndex] = h.Buffer[rightIndex], h.Buffer[currIndex]
						currIndex = rightIndex
					} else {
						break
					}
				}
			} else if leftIndex <= h.Count-1 {
				if h.Buffer[leftIndex] < h.Buffer[currIndex] {
					h.Buffer[leftIndex], h.Buffer[currIndex] = h.Buffer[currIndex], h.Buffer[leftIndex]
					currIndex = leftIndex
				} else {
					break
				}
			} else {
				break
			}
		} else {
			if rightIndex <= h.Count-1 {
				//选出一个较小者
				if h.Buffer[leftIndex] > h.Buffer[rightIndex] {
					if h.Buffer[leftIndex] > h.Buffer[currIndex] {
						h.Buffer[currIndex], h.Buffer[leftIndex] = h.Buffer[leftIndex], h.Buffer[currIndex]
						currIndex = leftIndex
					} else {
						break
					}
				} else {
					if h.Buffer[rightIndex] > h.Buffer[currIndex] {
						h.Buffer[currIndex], h.Buffer[rightIndex] = h.Buffer[rightIndex], h.Buffer[currIndex]
						currIndex = rightIndex
					} else {
						break
					}
				}
			} else if leftIndex <= h.Count-1 {
				if h.Buffer[leftIndex] > h.Buffer[currIndex] {
					h.Buffer[leftIndex], h.Buffer[currIndex] = h.Buffer[currIndex], h.Buffer[leftIndex]
					currIndex = leftIndex
				} else {
					break
				}
			} else {
				break
			}
		}
	}
	return top
}
