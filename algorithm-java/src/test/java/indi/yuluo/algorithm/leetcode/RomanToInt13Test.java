package indi.yuluo.algorithm.leetcode;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.leetcode.RomanToInt13.romanToInt;
import static org.junit.jupiter.api.Assertions.assertEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */
public class RomanToInt13Test {

	@Test
	void testRomanToInt() {

		List<Data> init = init();
		for (Data data : init) {
			int actual = romanToInt(data.input);
			assertEquals(actual, data.expected);
		}

	}

	private static List<Data> init() {

		ArrayList<Data> data = new ArrayList<>();
		data.add(new Data("III", 3));
		data.add(new Data("LVIII", 58));
		data.add(new Data("MCMXCIV", 1994));

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
