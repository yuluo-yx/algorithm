package sort

import "fmt"

const N = 1e6 + 10

func main() {

	q := make([]int, N)
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}

	for i := 0; i < n; i++ {
		_, err = fmt.Scanf("%d", &q[i])
		if err != nil {
			return
		}
	}

	quickSort(q, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}

}

func quickSort(q []int, l, r int) []int {

	// 递归出口
	if l >= r {
		return q
	}

	x := q[(l+r)>>1]
	i, j := l-1, r+1

	for i < j {
		for {
			i++
			if q[i] >= x {
				break
			}
		}

		for {
			j--
			if q[j] <= x {
				break
			}
		}

		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	// 递归左右
	quickSort(q, l, j)
	quickSort(q, j+1, r)

	return q
}
