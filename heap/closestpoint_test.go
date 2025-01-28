package heap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKClosestPoint(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{-2, 4}, {3, 3}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %v", tt.points, tt.k), func(t *testing.T) {
			got := KCloestPoint(tt.points, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("Got %v, Want %v", got, tt.want)
			}
		})
	}
}
