//package intervalMerging
//
//import (
//	"bufio"
//	"fmt"
//	"math"
//	"os"
//	"sort"
//	"strconv"
//	"strings"
//)
//
//const N int = 1e5 + 10
//
//var (
//	a       []int
//	n       int
//	scanner = *bufio.NewScanner(os.Stdin)
//)
//
//// 模拟 pair 对，存储区间左右端点
//type pair struct {
//	first  int
//	second int
//}
//
//func main() {
//
//	bs := make([]byte, 20000*1024)
//	scanner.Buffer(bs, len(bs))
//
//	n = readLine()[0]
//	segs := make([]pair, 0)
//
//	for i := 0; i < n; i++ {
//		x := readLine()
//		segs = append(segs, pair{x[0], x[1]})
//	}
//
//	res := merge(segs)
//
//	fmt.Println(len(res))
//}
//
//func merge(segs []pair) []pair {
//
//	// 定义区间合并后的结果
//	var res []pair
//	sort.Slice(segs, func(i, j int) bool {
//		if segs[i].first == segs[j].first {
//			return segs[i].second < segs[j].second
//		}
//		return segs[i].first < segs[j].first
//	})
//
//	start, end := math.MinInt32, math.MinInt32
//	for _, seg := range segs {
//		if end < seg.first {
//			if start != math.MinInt32 {
//				res = append(res, pair{start, end})
//			}
//			start, end = seg.first, seg.second
//		} else {
//			if seg.second > end {
//				end = seg.second
//			}
//		}
//	}
//
//	if start != math.MinInt32 {
//		res = append(res, pair{start, end})
//	}
//
//	return res
//
//}
//
//func readLine() []int {
//	scanner.Scan()
//	l := strings.Split(scanner.Text(), " ")
//	res := make([]int, len(l))
//	for i, s := range l {
//		x, _ := strconv.Atoi(s)
//		res[i] = x
//	}
//
//	return res
//}
