package indi.yuluo.algorithm.binary;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class BinarySearch {

	public static void main(String[] args) {

		int n = 6;
		int m = 3;
		int[] input = {1, 3, 3, 4, 4, 6};
		int[] qu = {2, 8, 3};

		for (int i = 0; i < m; i ++) {
			int[] res = binarySearch(input, n, qu[i]);
			System.out.println(res[0] + " " + res[1]);
		}

	}

	public static int[] binarySearch(int[] q, int n, int qu) {

		int[] res = {-1, -1};

		int l = 0;
		int r = n - 1;

		while (l < r) {
			int mid = (l + r) >> 1;
			if (q[mid] >= qu) {
				r = mid;
			} else {
				l = mid + 1;
			}
		}

		if (q[l] != qu) {
			return res;
		} else {
			res[0] = l;
			l = 0;
			r = n - 1;
			while (l < r) {
				int mid = (l + r + 1) >> 1;
				if (q[mid] <= qu) {
					l = mid;
 				} else {
					r = mid - 1;
				}
			}

			res[1] = l;
		}

		return res;
	}

}
