package qsort

import "testing"

func TestQuickSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	QuickSort(values)

	for i := 0; i < len(values); i++ {
		if values[i] != i+1 {
			t.Error("QuickSort() failed. GOT", values, "Expected 1 2 3 4 5")
		}
	}
}

func TestQuickSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	QuickSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("QuickSort() failed. GOT", values, "Expected 1 2 3 4 5")
	}
}

func TestQuickSort3(t *testing.T) {
	values := []int{5}
	QuickSort(values)

	if values[0] != 5 {
		t.Error("QuickSort() failed. GOT", values, "Expected 5 5 5 5 5")
	}
}
