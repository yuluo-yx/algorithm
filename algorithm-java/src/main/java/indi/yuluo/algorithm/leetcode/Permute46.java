package indi.yuluo.algorithm.leetcode;

import java.util.ArrayDeque;
import java.util.ArrayList;
import java.util.Deque;
import java.util.List;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 全排列回溯算法
 * 回溯（暴搜） -> dfs
 */

public class Permute46 {

	public List<List<Integer>> permute(int[] nums) {

		int len = nums.length;

		List<List<Integer>> res = new ArrayList<>();
		if (len == 0) {
			return res;
		}

		boolean[] used = new boolean[len];
		Deque<Integer> path = new ArrayDeque<>(len);

		dfs(nums, len, 0, path, used, res);

		return res;

	}

	private void dfs(int[] nums, int len, int depth, Deque<Integer> path, boolean[] used, List<List<Integer>> res) {

		if (depth == len) {
			res.add(new ArrayList<>(path));
			return;
		}

		for (int i = 0; i < len; i ++) {
			if (!used[i]) {
				path.addLast(nums[i]);
				used[i] = true;

				dfs(nums, len, depth + 1, path, used, res);

				used[i] = false;
				path.removeLast();
			}
		}

	}

}
