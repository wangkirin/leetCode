package queue_stack

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// tcsnumSquares is testcase slice
//https://github.com/aQuaYi/LeetCode-in-Go/blob/master/Algorithms/0752.open-the-lock/open-the-lock_test.go
var tcs1 = []struct {
	deadends []string
	target   string
	ans      int
}{
	{
		[]string{"0201", "0101", "0102", "1212", "2002"},
		"0202",
		6,
	},

	{
		[]string{"8888"},
		"0000",
		0,
	},

	{
		[]string{"8888"},
		"0009",
		1,
	},

	{
		[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"},
		"8888",
		-1,
	},

	{
		[]string{"0000"},
		"8888",
		-1,
	},

	// 可以有多个 testcase
}

func Test_fn(t *testing.T) {
	ast := assert.New(t)

	for _, tc := range tcs1 {
		fmt.Printf("~~%v~~\n", tc)
		ast.Equal(tc.ans, openLock(tc.deadends, tc.target), "输入:%v", tc)
	}
}

func Benchmark_fn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range tcs1 {
			openLock(tc.deadends, tc.target)
		}
	}
}
