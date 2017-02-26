package numbers

/*Merge takes two arrays a, b which are sorted and
returns result of merging a, b after removing the duplicates.
The merged array is also in sorted order*/
func Merge(a []int, b []int, result chan<- []int) {
	c := make([]int, len(a)+len(b))

	aIndex, bIndex, cIndex := 0, 0, 0

	for aIndex < len(a) && bIndex < len(b) && cIndex < len(a)+len(b) {
		if a[aIndex] <= b[bIndex] {
			if cIndex == 0 || c[cIndex-1] != a[aIndex] {
				c[cIndex] = a[aIndex]
				cIndex++
			}

			aIndex++
		} else {
			if cIndex == 0 || c[cIndex-1] != b[bIndex] {
				c[cIndex] = b[bIndex]
				cIndex++
			}

			bIndex++
		}
	}

	for aIndex < len(a) {
		if cIndex == 0 || c[cIndex-1] != a[aIndex] {
			c[cIndex] = a[aIndex]
			cIndex++
		}

		aIndex++
	}

	for bIndex < len(b) {
		if cIndex == 0 || c[cIndex-1] != b[bIndex] {
			c[cIndex] = b[bIndex]
			cIndex++
		}

		bIndex++
	}

	result <- c[:cIndex]
}
