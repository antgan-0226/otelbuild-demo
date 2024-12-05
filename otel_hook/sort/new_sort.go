package sort

import (
	"fmt"
	"github.com/antgan-0226/kotelbuild/pkg/api"
)

func newSortImpl(call api.CallContext, x []int) {
	fmt.Println(fmt.Sprintf("[new sort impl]采用冒泡排序算法"))
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	fmt.Println("Sorted x:", x)

	//跳过原方法执行
	call.SetSkipCall(true)
}
