package indi.yuluo.algorithm.leetcode;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.leetcode.IsPalindrome9.isPalindrome;
import static org.junit.jupiter.api.Assertions.assertEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class IsPalindrome9Test {

	@Test
	void testIsPalindrome9() {

		List<Data> data = init();

		for (Data datum : data) {
			boolean actual = isPalindrome(datum.input);
			assertEquals(actual, datum.expected);
		}

	}

	private static List<Data> init() {

		ArrayList<Data> data = new ArrayList<>();
		data.add(new Data(121, true));
		data.add(new Data(0, true));
		data.add(new Data(1, true));
		data.add(new Data(10, true));
		data.add(new Data(-121, true));

		return data;
	}

	static class Data {

		private int input;

		private boolean expected;

		Data(int input, boolean expected) {
			this.input = input;
			this.expected = expected;
		}
	}

}
