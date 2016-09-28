package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)

	for i := 0; i < len(values); i++ {
		if values[i] != i+1 {
			t.Error("BubbleSort() failed. GOT", values, "Expected 1 2 3 4 5")
		}
	}
}

func TestBubbleSort2(t *testing.T) {
	values := []int{5, 5, 3, 2, 1}
	BubbleSort(values)

	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 5 || values[4] != 5 {
		t.Error("BubbleSort() failed. GOT", values, "Expected 1 2 3 4 5")
	}
}

func TestBubbleSort3(t *testing.T) {
	values := []int{5}
	BubbleSort(values)

	if values[0] != 5 {
		t.Error("BubbleSort() failed. GOT", values, "Expected 5 5 5 5 5")
	}
}
