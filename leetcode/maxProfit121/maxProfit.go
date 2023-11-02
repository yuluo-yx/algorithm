package maxProfit121

import "math"

func maxProfit(prices []int) int {
	var res int
	minNum := 1 << 30

	for _, price := range prices {
		res = int(math.Max(float64(res), float64(price-minNum)))
		minNum = int(math.Min(float64(minNum), float64(price)))
	}

	return res
}
