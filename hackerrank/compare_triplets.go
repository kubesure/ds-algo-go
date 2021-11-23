package hackerrank

/*
 * Complete the 'compareTriplets' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func compareTriplets(a []int32, b []int32) []int32 {
	var ascore, bscore int32
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			ascore++
		} else if a[i] < b[i] {
			bscore++
		}
	}
	return []int32{ascore, bscore}

}
