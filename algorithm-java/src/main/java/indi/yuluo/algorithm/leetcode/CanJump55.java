package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */
public class CanJump55 {

	public static void main(String[] args) {
		System.out.println(canJump(new int[]{2,3,1,1,4}));
	}

	public static boolean canJump(int[] nums) {

		if (nums == null) {

			return false;
		}

		// 前 n - 1 个元素能到达的最远长度
		int len = 0;
		// 依次遍历，找到每个点可以到达的最远长度。
		// 如果最远长度大于等于 数组长度，则证明可以。反则不行
		for (int i = 0; i <= len; i ++) {
			int step = i + nums[i];
			// 更新最远距离
			len = Math.max(len, step);
			if (len >= nums.length - 1) {
				// 大于数组长度，则证明能到达数组尽头
				return true;
			}
		}

		return false;
	}

}
