package indi.yuluo.algorithm.leetcode;

/**
 * @author yuluo
 * @author 1481556636@qq.com
 */
public class HIndex274 {

	public static void main(String[] args) {
//		System.out.println(hIndex(new int[]{3,0,6,1,5}));
		System.out.println(hIndex(new int[]{1,3,1}));
	}

	public static int hIndex(int[] citations) {

		int total = 0, n = citations.length;

		int[] cnt = new int[n + 5];

		for (int citation : citations) {
			if (citation >= n) {
				cnt[n]++;
			}
			else {
				cnt[citation]++;
			}
		}

		for (int i = n; i >= 0; i --) {
			total += cnt[i];
			if (total >= i) {
				return i;
			}
		}

		return 0;

	}

}
