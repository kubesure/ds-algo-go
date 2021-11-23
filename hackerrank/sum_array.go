package hackerrank

/*
 * Complete the 'simpleArraySum' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY ar as parameter.
 */

func simpleArraySum(ar []int32) int32 {
	var sum int32
	for _, v := range ar {
		sum = sum + v
	}
	return sum
}
