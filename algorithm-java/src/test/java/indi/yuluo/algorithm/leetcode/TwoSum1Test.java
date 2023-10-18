package indi.yuluo.algorithm.leetcode;

import java.util.ArrayList;
import java.util.List;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.leetcode.TwoSum1.twoSum;
import static org.junit.jupiter.api.Assertions.assertArrayEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class TwoSum1Test {

	@Test
	void testTwoSum() {

		List<Data> data = init();

		data.forEach(v -> {
			int[] ints = twoSum(v.getArray(), v.getTarget());
			assertArrayEquals(ints, v.getExpected());
		});


	}

	private static List<Data> init() {
		List<Data> data = new ArrayList<>();
		data.add(new Data(new int[]{3, 3}, 6, new int[]{0, 1}));
		data.add(new Data(new int[]{3, 2, 4}, 6, new int[]{1, 2}));
		data.add(new Data(new int[]{2, 7, 11, 5}, 9, new int[]{0, 1}));

		return data;
	}

	static class Data {

		private int[] array;

		private int target;

		private int[] expected;

		Data(int[] array, int target, int[] expected) {
			this.array = array;
			this.target = target;
			this.expected = expected;
		}

		int[] getArray() {
			return array;
		}

		void setArray(int[] array) {
			this.array = array;
		}

		int getTarget() {
			return target;
		}

		void setTarget(int target) {
			this.target = target;
		}

		int[] getExpected() {
			return expected;
		}

		void setExpected(int[] expected) {
			this.expected = expected;
		}
	}

}
