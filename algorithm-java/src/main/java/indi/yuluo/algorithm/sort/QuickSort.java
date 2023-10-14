package indi.yuluo.algorithm.sort;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class QuickSort {

	public static void main(String[] args) {

		int[] input = {2, 1, 3, 4 ,5};
		int[] sorted = quickSort(input, 0, input.length - 1);

		for (int i : sorted) {
			System.out.print(i + " ");
		}

	}

	public static int[] quickSort(int[]q, int l, int r) {

		if (l >= r) {

			return q;
		}

		int x = q[(l + r) >> 1];
		int i = l - 1;
		int j = r + 1;

		while (i < j) {
			do {
				i++;
			}
			while (q[i] < x);
			do {
				j --;
			}
			while(q[j] > x);

			if (i < j) {
				int temp = q[i];
				q[i] = q[j];
				q[j] = temp;
			}
		}

		quickSort(q, l, j);
		quickSort(q, j + 1, r);

		return q;
	}

}
