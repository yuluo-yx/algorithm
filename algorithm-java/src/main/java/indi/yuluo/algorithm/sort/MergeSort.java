package indi.yuluo.algorithm.sort;

import java.util.Arrays;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class MergeSort {

	private static final int N = (int) 1e6 + 10;

	static int n;
	static int[] input = new int[N];
	static int[] tmp = new int[N];

	public static void main(String[] args) {

		n = 5;
		input = new int[] {3, 1, 2, 4, 5};

		int[] res = mergeSort(input, tmp, 0, n - 1);

		System.out.println(Arrays.toString(res));
	}


	public static int[] mergeSort(int[] q, int[] tmp, int l, int r) {

		if (l >= r) {
			return q;
		}

		int mid = l + r >> 1;

		mergeSort(q, tmp, l, mid);
		mergeSort(q, tmp,mid + 1, r);

		int k = 0, i = l, j = mid + 1;
		while(i <= mid && j <= r) {
			if (q[i] <= q[j]) {
				tmp[k ++] = q[i ++];
			} else {
				tmp[k ++] = q[j ++];
			}
		}

		while (i <= mid) {
			tmp[k ++] = q[i ++];
		}
		while(j <= r) {
			tmp[k ++] = q[j ++];
		}

		for (i = l, j = 0; i <= r; i ++, j++) {
			q[i] = tmp[j];
		}

		return q;

	}

}
