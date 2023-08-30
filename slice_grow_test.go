package ds

import "testing"

func TestSliceGrow(t *testing.T) {
	list := []int{1,2}
	t.Log(len(list), cap(list))
	list = append(list, 3, 4, 5)
	t.Log(len(list), cap(list))
	list = append(list, 6, 7, 8)
	t.Log(len(list), cap(list))
	list = append(list, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25)
	t.Log(len(list), cap(list))
}