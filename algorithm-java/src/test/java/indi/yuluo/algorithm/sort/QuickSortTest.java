package indi.yuluo.algorithm.sort;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class QuickSortTest {

	@Test
	void quickSortTest() {

		int[] input = {2, 1, 3, 4, 5};
		int[] expected = {1, 2, 3, 4, 5};

		int[] sorted = QuickSort.quickSort(input, 0, input.length - 1);
		assertArrayEquals(sorted, expected);

	}


}
