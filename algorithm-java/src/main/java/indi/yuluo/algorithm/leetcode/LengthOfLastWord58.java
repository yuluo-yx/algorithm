package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 * 最后一个单词的长度
 */

public class LengthOfLastWord58 {

	public static void main(String[] args) {
		System.out.println(lengthOfLastWord("   fly me   to   the moon  "));
		System.out.println(lengthOfLastWord("Hello World"));
	}

	public static int lengthOfLastWord(String s) {

		return s.split(" ")[s.split(" ").length - 1].length();
	}

}
