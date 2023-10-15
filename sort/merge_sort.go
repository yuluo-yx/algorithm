package sort

var n int

// 用于排序合并时的临时数组
var t []int

func init() {
	t = make([]int, N)
}

//func main() {
//
//	_, err := fmt.Scanf("%d", &n)
//	if err != nil {
//		return
//	}
//
//	// 数字序列数组
//	q := make([]int, N)
//
//	for i := 0; i < n; i++ {
//		_, err2 := fmt.Scanf("%d", &q[i])
//		if err2 != nil {
//			return
//		}
//	}
//
//	sorted := mergeSort(q, t, 0, n-1)
//
//	for i := 0; i < n; i++ {
//		fmt.Printf("%d ", sorted[i])
//	}
//
//}

func mergeSort(q, t []int, l, r int) []int {

	// 递归出口
	if l >= r {
		return q
	}

	mid := l + (r-l)>>1

	// 递归排序左右
	mergeSort(q, t, l, mid)
	mergeSort(q, t, mid+1, r)

	k, i, j := 0, l, mid+1

	for i <= mid && j <= r {
		if q[i] <= q[j] {
			t[k] = q[i]
			k, i = k+1, i+1
		} else {
			t[k] = q[j]
			k, j = k+1, j+1
		}
	}

	// 判断左右两边是否已经合并完成
	for i <= mid {
		t[k] = q[i]
		k, i = k+1, i+1
	}

	for j <= r {
		t[k] = q[j]
		k, j = k+1, j+1
	}

	// 还原 t 数组到之前的 q 数组中去
	for i, j := l, 0; i <= r; i, j = i+1, j+1 {
		q[i] = t[j]
	}

	return q

}
