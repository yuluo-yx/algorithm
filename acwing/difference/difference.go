package difference

//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//const N int = 1e5 + 10
//
//var (
//	n, m int
//	a    = make([]int, N)
//	b    = make([]int, N)
//)
//
//// 差分算法：在时间复杂度为 O(n) 下，给区间中的每个数组加上常熟 C
//// 和 前缀和 互为逆运算
//func main() {
//
//	in := bufio.NewReader(os.Stdin)
//	out := bufio.NewWriter(os.Stdout)
//	// 刷新流
//	defer func(out *bufio.Writer) {
//		err := out.Flush()
//		if err != nil {
//			panic("error")
//		}
//	}(out)
//
//	_, err := fmt.Scan(&n, &m)
//	if err != nil {
//		return
//	}
//
//	for i := 1; i <= n; i++ {
//
//		_, err2 := fmt.Fscan(in, &a[i])
//		// 构建差分数组
//		b[i] = a[i] - a[i-1]
//
//		if err2 != nil {
//			return
//		}
//	}
//
//	for m > 0 {
//		var l, r, c int
//		_, err2 := fmt.Fscan(in, &l, &r, &c)
//		if err2 != nil {
//			return
//		}
//		// 加入常数 C
//		Insert(l, r, c)
//		m--
//	}
//
//	//前缀和运算
//	for i := 1; i <= n; i++ {
//		a[i] = b[i] + a[i-1]
//	}
//
//	for i := 1; i <= n; i++ {
//		_, err2 := fmt.Fprint(out, a[i], " ")
//		if err2 != nil {
//			return
//		}
//	}
//}
//
//// Insert 在 l - r 区间内加入常数 C
//func Insert(l, r, c int) {
//	b[l] += c
//	b[r+1] -= c
//}
