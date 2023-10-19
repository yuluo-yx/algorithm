package indi.yuluo.algorithm.difference;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class Difference {

	static int N = 100010;

	static int n = 6;
	// 数据数组
	static int[] data = {1, 2, 2, 1, 2, 1};
	// 差分数组
	static int[] b = new int[N];
	static int m = 3;
	static int[][] qu = {{1, 3, 1}, {3, 5, 1}, {1, 6, 1}};

	public static void main(String[] args) {

		int[] a = new int[N];

		// 构造差分数组 下标从 1 开始，避免前缀和数组中的下标转换
		for (int i = 1; i <= n; i++) {
			a[i] = data[i - 1];
			b[i] = a[i] - a[i - 1];
		}

		for (int i = 0; i < m; i ++) {
			int l = qu[i][0];
			int r = qu[i][1];
			int c = qu[i][2];
			insert(l, r, c);
		}

		// 前缀和运算
		for (int i = 1; i <= n; i ++) {
			a[i] = b[i] + a[i - 1];
			System.out.print(a[i] + " ");
		}

	}

	public static void insert(int l, int r, int c) {

		b[l] += c;
		b[r + 1] -= c;
	}

}
