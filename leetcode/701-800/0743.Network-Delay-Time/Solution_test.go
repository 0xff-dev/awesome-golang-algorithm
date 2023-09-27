package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		n, k   int
		expect int
	}{
		{"TestCase1", [][]int{
			{2, 1, 1}, {2, 3, 1}, {3, 4, 1},
		}, 4, 2, 2},
		{"TestCase2", [][]int{
			{1, 2, 1},
		}, 2, 1, 1},
		{"TestCase3", [][]int{
			{1, 2, 1},
		}, 2, 2, -1},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs, c.n, c.k)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v %v %v",
					c.expect, got, c.inputs, c.n, c.k)
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
