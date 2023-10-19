package indi.yuluo.algorithm.leetcode;

import java.util.Objects;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class IsPalindrome9 {

	public static void main(String[] args) {

		System.out.println(isPalindrome(11));

	}

	public static boolean isPalindrome(int x) {

		if (x < 0) {
			return false;
		}

		String str = String.valueOf(x);
		String prefix = new StringBuffer(str.substring(0, str.length() / 2)).reverse().toString();
		String suffix = null;
		if (!isEven(str.length())) {
			suffix = str.substring(str.length() / 2 + 1);
		} else {
			suffix = str.substring(str.length() / 2);
		}

		System.out.println(prefix + " " + suffix);

		return Objects.equals(prefix, suffix);
	}

	private static boolean isEven(int a) {

		return a % 2 == 0;
	}

}
