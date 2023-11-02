package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 买卖股票的最佳时机2
 */

public class MaxProfit122 {

	public static void main(String[] args) {

		System.out.println(maxProfit(new int[]{7,1,5,3,6,4}));
		System.out.println(maxProfitDp(new int[]{7,1,5,3,6,4}));
	}

	/**
	 * 贪心算法，做当前情况下最好的选择
	 *
	 * @param prices
	 * @return
	 */
	public static int maxProfit(int[] prices) {
		int res = 0;
		int profit;

		for (int i = 1; i < prices.length; i ++) {
			// 总是做当前最好的选择，只累加是正数的值
			profit = (prices[i] - prices[i - 1]);
			if (profit > 0) {
				res += profit;
			}
		}

		return res;
	}

	/**
	 * 动规解法
	 *
	 * @param prices
	 * @return
	 */
	public static int maxProfitDp(int[] prices) {

		int len = prices.length;
		if (len < 2) {
			return 0;
		}

		// cash 持有现金
		int[] cash = new int[len];
		// 持有股票
		int[] hold = new int[len];

		// 状态数组
		// 状态转移：cash -> hold -> case -> hold -> case ....
		cash[0] = 0;
		// 持有股票，当前拥有的现金数是当天股价的相反数，花钱买入 -7
		hold[0] = -prices[0];

		for (int i = 1; i < len; i ++) {
			cash[i] = Math.max(cash[i - 1], hold[i - 1] + prices[i]);
			hold[i] = Math.max(hold[i - 1], cash[i - 1] - prices[i]);
		}

		return cash[len - 1];
	}

}
