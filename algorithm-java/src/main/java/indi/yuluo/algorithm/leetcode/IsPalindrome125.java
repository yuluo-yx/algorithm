package indi.yuluo.algorithm.leetcode;

import java.util.Locale;
import java.util.Objects;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 如果在将所有大写字符转换为小写字符、并移除所有非字母数字字符之后，
 * 短语正着读和反着读都一样。则可以认为该短语是一个 回文串 。
 */
public class IsPalindrome125 {

	public static void main(String[] args) {
		System.out.println(isPalindrome("A man, a plan, a canal: Panama"));
		System.out.println(isPalindrome("*0P"));

		System.out.println(isPalindromePlus("A man, a plan, a canal: Panama"));
		System.out.println(isPalindromePlus("*0P"));
	}

	/**
	 * 双指针
	 * @return
	 */
	public static boolean isPalindromePlus(String s) {

		s = s.toLowerCase(Locale.forLanguageTag(s))
				.replaceAll("[^0-9a-z]", "");

		int len = s.length();

		boolean flag = true;
		for (int i = 0, j = len - 1; i < len / 2; i ++, j --) {
			if (!Objects.equals(s.charAt(i), s.charAt(j))) {
				flag = false;
			}
		}

		return flag;

	}

	/**
	 * 朴素做法
	 */
	public static boolean isPalindrome(String s) {

		boolean flag = true;

		s = s.toLowerCase(Locale.forLanguageTag(s))
				.replaceAll("[^0-9a-z]", "");

		if (s.length() == 1) {
			return flag;
		}

		String start;
		String end;

		if (isPrim(s.length())) {
			start = s.substring(0, s.length() / 2);
		} else {
			start = s.substring(0, s.length() / 2 + 1);
		}
		end = s.substring(s.length() / 2);

		StringBuilder sb = new StringBuilder(end);
		end = sb.reverse().toString();

		if (!Objects.equals(end, start)) {
			flag = false;
		}

		return flag;

	}

	private static boolean isPrim(int s) {

		return s % 2 == 0;
	}

}
