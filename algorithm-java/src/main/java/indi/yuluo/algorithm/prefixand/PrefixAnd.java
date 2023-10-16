package indi.yuluo.algorithm.prefixand;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class PrefixAnd {

	static int n = 5;
	static int m = 3;
	static int[] a = {2, 1, 3, 6, 4};
	static int[] s = new int[n + 10];
	static int[][] qu = {{1, 2}, {1, 3}, {2, 4}};

	public static void main(String[] args) {

		for (int i = 1; i <= n; i++) {
			s[i] = s[i - 1] + a[i - 1];
		}

		for (int i = 0; i < m; i++) {
			int l = qu[i][0];
			int r = qu[i][1];
			System.out.println(s[r] - s[l - 1]);
		}

	}

}
