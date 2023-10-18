package indi.yuluo.algorithm.leetcode;

import java.util.HashMap;
import java.util.Map;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class lengthOfLongestSubstring3 {

	public static void main(String[] args) {

		int res = lengthOfLongestSubstring("abcabcbb");
		System.out.println(res);

	}

	// 滑动窗口
	public static int lengthOfLongestSubstring(String s) {

		if (s.length() <= 1) {
			return s.length();
		}

		// hash 表存储数据，记录每个字符是否出现过
		Map<Character, Boolean> map = new HashMap<>();

		int l = 0;
		int r = 0;
		int max = 0;

		// 当 r < s长度时，一直执行
		while (r < s.length()) {
			// 如果 map 中包含当前字符，（这意味着在当前的滑动窗口中已经有了这个字符），那么我们需要更新窗口的左边界 left。
			if (map.containsKey(s.charAt(r))) {
				// 左指针向右移动一格，移除一个字符
				map.remove(s.charAt(l));
				l ++;
			} else {
				// 不包含，放进 map 中，标记为 true，
				map.put(s.charAt(r), true);
				// 计算长度最大值（第 l 到 r 个字符）
				max = Math.max(max, r - l + 1);
				// 移动右指针
				r ++;
			}
		}

		return max;
	}

}
