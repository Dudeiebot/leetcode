package heap

import (
	"testing"
)

func TestKthLargest(t *testing.T) {
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)

	if obj.Add(3) != 4 {
		t.Errorf("Expected 4, got %d", obj.Add(3))
	}

	if obj.Add(5) != 5 {
		t.Errorf("Expected 5, got %d", obj.Add(5))
	}

	if obj.Add(10) != 5 {
		t.Errorf("Expected 5, got %d", obj.Add(10))
	}

	if obj.Add(9) != 8 {
		t.Errorf("Expected 8, got %d", obj.Add(9))
	}

	if obj.Add(4) != 8 {
		t.Errorf("Expected 8, got %d", obj.Add(4))
	}
}

// for finding the KthLargest number, we are using minHeap and it is available in the standard library(the min heap kind of sort the array for us easily with the nums)
