package prefixand

import "fmt"

const N = 1e6 + 10

//var (
//	// n 数组长度，m 询问个数
//	n, m int
//	// a 数据数组，s 前缀和数组
//	a, s [N]int
//)
//
//func main() {
//
//	_, err := fmt.Scanf("%d%d", &n, &m)
//	if err != nil {
//		return
//	}
//
//	// 前缀和的下标一定要从 1 开始，避免下标的转换 s[0] = 0
//	for i := 1; i <= n; i++ {
//		// 读入数组数据，不 - 1
//		_, err = fmt.Scanf("%d", &a[i])
//		if err != nil {
//			return
//		}
//		// 初始化前缀和数组
//		s[i] = s[i-1] + a[i]
//	}
//
//	// 处理询问
//	for i := 0; i < m; i++ {
//		var l, r int
//		_, err = fmt.Scanf("%d%d", &l, &r)
//		if err != nil {
//			return
//		}
//
//		fmt.Printf("%d\n", s[r]-s[l-1])
//	}
//
//}

// prefixAnd() 快速求出元素组中某段区间的和  这里是传入数组，为了避免 a 下标越界，故 - 1
// s[i] = s[i - 1] + a[i -  1]
// fmt.Println(s[r] - s[l - 1])
// a 数据数组，qu 询问数组, m 询问数组大小
func prefixAnd(a []int, n int, qu [][]int, m int) []int {

	var s = make([]int, m+10)
	var res = make([]int, m)

	for i := 1; i <= n; i++ {
		// 初始话前缀和数组
		s[i] = s[i-1] + a[i-1]
	}

	fmt.Print(s)

	// 处理询问
	for i := 0; i < m; i++ {
		l, r := qu[i][0], qu[i][1]
		res[i] = s[r] - s[l-1]
	}

	return res

}
