package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name          string
		mapping, nums []int
		expect        []int
	}{
		{"TestCase1", []int{8, 9, 4, 0, 2, 1, 3, 5, 7, 6}, []int{991, 338, 38}, []int{338, 38, 991}},
		{"TestCase2", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{789, 456, 123}, []int{123, 456, 789}},
		{"TestCase3", []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.mapping, c.nums)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v",
					c.expect, got, c.mapping, c.nums)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}
