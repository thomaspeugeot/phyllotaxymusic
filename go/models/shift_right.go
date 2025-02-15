package models

// shiftRight returns a new slice whose elements are rotated by k to the right.
//
// Example:
//
//	input = [p0, p1, p2, p3]
//	k = 1 -> output = [p3, p0, p1, p2]
func shiftRight[T any](input []*T, k int) []*T {
	n := len(input)

	// Allocate a new slice so we don't reuse the original backing array
	shifted := make([]*T, 0, n)

	// Ensure k is within the range [0, n)
	k = k % n
	if k < 0 {
		k += n
	}

	// Append the last k elements, then the first n-k elements
	shifted = append(shifted, input[n-k:]...)
	shifted = append(shifted, input[:n-k]...)
	return shifted
}
