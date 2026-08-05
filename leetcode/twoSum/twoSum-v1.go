package twosum

func twoSum(nums []int, target int) []int {
    for i := range nums {
        for j, val := range nums[i + 1:] {
            if nums[i] + val == target{
                return []int{i, i + j + 1}
            }
        }
    }
    return nil
}

// nums = [2, 7, 11, 15], target = 22
// [1, 3]

// i = 0, val1 = 2
// nums[i +1:] [7, 11, 15]
// j = 0, val2 = 7 no, i+j+1 = 1
// j = 1, val2 = 11 no, i+j+1 = 2
// j = 2, val2 = 15 no, i+j+1 = 3

// i = 1, val1 = 7
// nums[i +1:] [11, 15]
// j = 0, val2 = 11 no, i+j+1 = 2
// j = 1, val2 = 15 yes i+j+1 = 3


