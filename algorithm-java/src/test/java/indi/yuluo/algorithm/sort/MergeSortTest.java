package indi.yuluo.algorithm.sort;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.sort.MergeSort.mergeSort;
import static org.junit.jupiter.api.Assertions.assertArrayEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class MergeSortTest {

	int n = 5;
	int[] input = new int[] {3, 1, 2, 4, 5};
	int[] tmp = new int[n];
	int[] expected = {1, 2, 3, 4, 5};

	@Test
	void mergeSortTest() {

		int[] res = mergeSort(input, tmp, 0, n - 1);
		assertArrayEquals(res, expected);
	}

}
