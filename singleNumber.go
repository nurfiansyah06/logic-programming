package logicprogramming

func singleNumber(nums []int) int {
    counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	for num, count := range counts {
		if count == 1 {
			return num
		}
	}

	return -1 
}