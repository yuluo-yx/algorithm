package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 快速幂
 */

public class MyPow50 {

	public static double myPow(double x, int n) {

		if (x == 0.0f) {
			return 0.0d;
		}

		long b = n;
		double res = 1.0;

		if (b < 0) {
			x = 1 / x;
			b = -b;
		}
		while (b > 0) {
			// 取余数 n%2 等价于判断二进制最右位 n&1
			if ((b & 1) == 1) {
				res *= x;
			}
			x *= x;
			// 向下整除 n//2n 等价于 右移一位 n>>1
			b >>= 1;
		}

		return res;
	}

}
