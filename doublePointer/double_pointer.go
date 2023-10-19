//package doublePointer
//
//import "fmt"
//
//const N int = 1e5 + 10
//
//var (
//	n int
//	// 数据序列
//	a [N]int
//	// j - i 区间中每个数出现的次数
//	s [N]int
//)
//
//// acwing 799 最长连续不重复子序列
//// 5，{1 2 2 3 5}， 3
//func main() {
//
//	// 读取输入数据
//	_, err := fmt.Scanf("%d", &n)
//	if err != nil {
//		return
//	}
//
//	for i := 0; i < n; i++ {
//		_, err2 := fmt.Scanf("%d", &a[i])
//		if err2 != nil {
//			return
//		}
//	}
//
//	// res 答案
//	var res int
//	// 使用双指针算法
//	for i, j := 0, 0; i < n; i++ {
//		// 将 a[i] 加入 s 数组中
//		s[a[i]]++
//		// 当 a[i] 出现过 1 次的时候
//		for s[a[i]] > 1 {
//			// 弹出 s[a[j]]
//			s[a[j]]--
//			// j 指针后移
//			j++
//		}
//
//		res = max(res, i-j+1)
//	}
//
//	fmt.Println(res)
//
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
