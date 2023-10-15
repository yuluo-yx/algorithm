package binary

import (
	"fmt"
)

const N int = 1e5 + 10

var (
	// n 数组长度，m 询问数字个数
	n, m, qu int
	// 数据数组
	q []int
)

func init() {
	q = make([]int, N)
}

func main() {

	_, err := fmt.Scanf("%d%d", &n, &m)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d", &q[i])
		if err != nil {
			return
		}
	}

	for ; m > 0; m-- {
		_, err = fmt.Scanf("%d", &qu)
		if err != nil {
			return
		}
		res := binarySearch(q, n, qu)

		fmt.Printf("%d %d\n", res[0], res[1])
	}

}

// binarySearch() 当询问的数据不存在时，返回 -1 -1 存在返回元素的 起始位置 和 终止位置
func binarySearch(q []int, n, qu int) []int {

	res := make([]int, 2)

	// 左右边界
	l, r := 0, n-1

	for l < r {
		mid := (l + r) >> 1
		if q[mid] >= qu {
			r = mid
		} else {
			l = mid + 1
		}
	}

	if q[l] != qu {
		return []int{-1, -1}
	} else {
		res[0] = l
		l, r := 0, n-1
		for l < r {
			mid := (l + r + 1) >> 1
			if q[mid] <= qu {
				l = mid
			} else {
				r = mid - 1
			}
		}
		res[1] = l
	}

	return res

}
