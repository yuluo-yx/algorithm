package indi.yuluo.algorithm.leetcode;

import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.Objects;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class TwoSum1 {

	public static void main(String[] args) {

		/*
		测试序列：[3, 2, 4 6], [3, 3, 6]
		 */
		System.out.println(Arrays.toString(twoSum(new int[] {3, 3}, 6)));

	}

	public static int[] twoSum(int[] nums, int target) {

		Map<Integer, Integer> map = new HashMap<>();
		for (int i = 0; i < nums.length; i++) {
			map.put(nums[i], i);
		}

		for (int i = 0; i < nums.length; i++) {
			int t = target - nums[i];
			if (map.containsKey(t)) {
				// 数字重复
				if (Objects.equals(t, nums[i])) {
					// 下标重复
					if (Objects.equals(i, map.get(t)))
						continue;
				}
				return new int[]{i, map.get(t)};
			}
		}

		return null;
	}

}
