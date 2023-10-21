package dfs

// dfs 模板
// int dfs(int u)
// {
//     st[u] = true; // st[u] 表示点u已经被遍历过

//     for (int i = h[u]; i != -1; i = ne[i])
//     {
//         int j = e[i];
//         if (!st[j]) dfs(j);
//     }
// }

const N int = 7 + 5

//var (
//	n int
//	// 恢复现场
//	path []int
//	// 状态数组，记录已经使用过的数字
//	status []bool
//)

//	func init() {
//		path = make([]int, N)
//		status = make([]bool, N)
//	}
//func main() {
//
//	_, err := fmt.Scanf("%d", &n)
//	if err != nil {
//		return
//	}
//
//	// 以任意位置开始遍历
//	res := dfs(0, n)
//	for i := range res {
//		fmt.Println(i)
//	}
//
//}

//
//func dfs(u int) {
//
//	// 递归出口，当 u == n 的时候，就是最后一个叶子节点
//	if u == n {
//		for i := 0; i < n; i++ {
//			fmt.Printf("%d ", path[i])
//		}
//		fmt.Println()
//	}
//
//	// 找到一个没有用过的点
//	for i := 1; i <= n; i++ {
//		if status[i] != true {
//			path[u] = i
//			status[i] = true
//			// 递归下一层
//			dfs(u + 1)
//			// 恢复现场
//			status[i] = false
//		}
//	}
//
//}

//func dfs(u, n int) (res [][]int) {
//
//	path := make([]int, N)
//	status := make([]bool, N)
//
//	// 递归出口，当 u == n 的时候，就是最后一个叶子节点
//	if u == n {
//		return res
//	}
//
//	// 找到一个没有用过的点
//	for i := 1; i <= n; i++ {
//		if status[i] != true {
//			path[u] = i
//			status[i] = true
//			// 递归下一层
//			dfs(u+1, n)
//			// 恢复现场
//			status[i] = false
//		}
//	}
//
//	return res
//}
