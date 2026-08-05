package twosum

func twoSum(nums []int, target int) []int {
	// key num value pos
	m := make(map[int]int)

	// for every number
    for i, val := range nums {
		// get how much they need to sum target
		diff := target - val

		// if the amount is in m
		if j, ok := m[diff]; ok {
			// access its pos and return it
			return []int{i, j}
		}

		// if not add it to look later
		m[val] = i
    }

    return nil
}
