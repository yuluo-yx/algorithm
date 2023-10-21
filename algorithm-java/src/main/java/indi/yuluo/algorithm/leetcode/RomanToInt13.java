package indi.yuluo.algorithm.leetcode;

import java.util.HashMap;
import java.util.Map;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class RomanToInt13 {

	public static int romanToInt(String s) {

		Map<String, Integer> map = new HashMap<>();
		map.put("I", 1);
		map.put("V", 5);
		map.put("X", 10);
		map.put("L", 50);
		map.put("C", 100);
		map.put("D", 500);
		map.put("M", 1000);

		int preNum = map.get(String.valueOf(s.charAt(0)));
		int sum = 0;
		for (int i = 1; i < s.length(); i ++) {
			int num = map.get(String.valueOf(s.charAt(i)));
			if (preNum < num) {
				sum -= preNum;
			} else {
				sum += preNum;
			}
			preNum = num;
		}

		sum += preNum;

		return sum;
	}

}
