package leethcode

func twoSum(nums []int, target int) []int {
	indices := make([]int, 2)
	dataMap := make(map[int]int)
	for i, n := range nums {
		if val, ok := dataMap[target-n]; ok {
			indices[0] = val
			indices[1] = i
			break
		} else {
			dataMap[nums[i]] = i
		}
	}
	return indices
}
