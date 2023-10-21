package indi.yuluo.algorithm.doublePoint;

import java.util.Scanner;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * acwing 799
 */

public class DoublePoint {

	static final int N = (int) (1e5 + 10);
	static int n;
	static int[] a = new int[N];
	static int[] s = new int[N];
	public static void main(String[] args) {
		Scanner scanner = new Scanner(System.in);
		scanner.nextLine();
		String data = scanner.nextLine();
		String[] s1 = data.split(" ");
		for (int i = 0; i < s1.length; i++) {
			a[i] = Integer.parseInt(s1[i]);
		}
		// a = new int[] {1, 2, 2, 3, 5};

		int res = 0;
		for (int i = 0, j = 0; i < a.length; i ++) {
			s[a[i]] ++;
			while (j < i && s[a[i]] > 1) {
				s[a[j]] --;
				j ++;
			}
			res = Math.max(res, i - j + 1);
		}
		System.out.println(res);
	}

}
