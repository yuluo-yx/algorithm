package TwoSum1

// {2, 5, 5, 10} 10 : error
func twoSum(nums []int, target int) []int {

	set := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		set[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		t := target - nums[i]
		// 值存在
		if _, exist := set[t]; exist {
			// 判断值是否重复
			if t == nums[i] {
				// 判断下标是否相同，避免 {3, 3} 和 {3, 2, 4}
				if i == set[t] {
					continue
				}
			}
		}

		return []int{i, set[t]}
	}

	return nil
}

func twoNumPlus(nums []int, target int) []int {

	set := make(map[int]int)

	for i, num := range nums {
		t := target - num
		if idx, exist := set[t]; exist {
			return []int{idx, i}
		} else {
			set[num] = i
		}
	}

	return nil
}
