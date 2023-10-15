package indi.yuluo.algorithm.binary;

import org.junit.jupiter.api.Test;

import static indi.yuluo.algorithm.binary.BinarySearch.binarySearch;
import static org.junit.jupiter.api.Assertions.assertArrayEquals;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */

public class BinarySearchTest {

	@Test
	void binarySearchTest() {

		int n = 6;
		int m = 3;
		int[] input = {1, 3, 3, 4, 4, 6};
		int[] qu = {2, 8, 3};
		int[][] expected = {{-1, -1}, {-1, -1}, {1, 2}};

		for (int i = 0; i < m; i ++) {
			int[] res = binarySearch(input, n, qu[i]);

			assertArrayEquals(res, expected[i]);
		}

	}

}
