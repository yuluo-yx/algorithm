package romanToInt13

func romanToInt(s string) int {

	var set = make(map[string]int, 1)

	set["I"] = 1
	set["V"] = 5
	set["X"] = 10
	set["L"] = 50
	set["C"] = 100
	set["D"] = 500
	set["M"] = 1000

	var sum int
	preNum := set[string(s[0])]
	for i := 1; i < len(s); i++ {
		num := set[string(s[i])]
		if preNum < num {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = num
	}
	sum += preNum

	return sum
}
