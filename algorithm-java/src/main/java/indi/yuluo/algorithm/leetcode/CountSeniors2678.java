package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class CountSeniors2678 {

	public static int countSeniors(String[] details) {

		int sum = 0;
		for (String str : details) {
			int age = Integer.parseInt(str.substring(11, 13));
			System.out.println(age);
			if (age > 60) {
				sum ++;
			}
		}

		return sum;

	}

}
