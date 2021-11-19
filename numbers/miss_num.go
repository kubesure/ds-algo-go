package numbers

func findMissingNum(input []int) int {
	var sumInput int
	for _, num := range input {
		sumInput += num
	}
	length := len(input) + 1
	sumActual := (length * (length + 1)) / 2
	return sumActual - sumInput
}
