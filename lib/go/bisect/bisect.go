package bisect

// START HERE

// Finds i such that arr[i:end] contains all points >= target.  len(arr) if all pts are < target
func bisect_left(arr []int, targ int) int {
	l, u := -1, len(arr)
	for u-l > 1 {
		m := (u + l) >> 1
		if arr[m] < targ {
			l = m
		} else {
			u = m
		}
	}
	return u
}

// Finds i such that arr[i:end] contains all points > target.  len(arr) if all pts are <= target
func bisect_right(arr []int, targ int) int {
	l, u := -1, len(arr)
	for u-l > 1 {
		m := (u + l) >> 1
		if arr[m] <= targ {
			l = m
		} else {
			u = m
		}
	}
	return u
}
