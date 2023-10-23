package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * x 的平方根 牛顿迭代法
 * 首先随便猜一个近似值 x，然后不断令 x 等于 x 和 a/x 的平均数，迭代个六七次后 x 的值就已经相当精确了。
 */

public class MySqrt69 {

	public static int mySqrt(int x) {

		if (x == 0) {

			return x;
		}

		double C = x, x0 = x;
		while (true) {
			double xi = 0.5 * (x0 + C / x0);
			if (Math.abs(x0 - xi) < 1e-7) {
				break;
			}
			x0 = xi;
		}
		return (int) x0;
	}

}
