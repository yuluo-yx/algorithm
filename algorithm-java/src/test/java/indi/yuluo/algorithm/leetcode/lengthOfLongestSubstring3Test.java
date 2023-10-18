package indi.yuluo.algorithm.leetcode;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.leetcode.lengthOfLongestSubstring3.lengthOfLongestSubstring;
import static org.junit.jupiter.api.Assertions.assertEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class lengthOfLongestSubstring3Test {

	@Test
	void testLengthOfLongestSubstring() {

		List<Data> init = init();

		init.forEach(v -> {
			int res = lengthOfLongestSubstring(v.input);
			assertEquals(v.expected, res);
		});



	}

	private static List<Data> init() {

		ArrayList<Data> data = new ArrayList<>();
		data.add(new Data("abcabcbb", 3));
		data.add(new Data("bbbbb", 1));
		data.add(new Data("pwwkew", 3));

		return data;
	}

	static class Data {

		private String input;

		private int expected;

		Data(String input, int expected) {
			this.input = input;
			this.expected = expected;
		}
	}

}


