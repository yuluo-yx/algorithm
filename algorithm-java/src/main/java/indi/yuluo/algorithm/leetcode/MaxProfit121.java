package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 买卖股票的最佳时机
 */

public class MaxProfit121 {

	public static void main(String[] args) {

		System.out.println(maxProfit(new int[] {7, 1, 5, 3, 6, 4}));
	}

	public static int maxProfit(int[] prices) {
		int res = 0;
		int min = 1 << 30;

		for (int price : prices) {
			res = Math.max(res, price - min);
			min = Math.min(min, price);
		}

		return res;
	}

}
