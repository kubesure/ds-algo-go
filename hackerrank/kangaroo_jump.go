package hackerrank

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {

	if x2 > x1 && v2 > v1 {
		return "NO"
	} else {
		for {
			x2 = x2 + v2
			x1 = x1 + v1

			if x1 == x2 {
				return "YES"
			}

			if x1 > x2 {
				return "NO"
			}
		}
	}
}
