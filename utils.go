package main

import (
	"strconv"
)

// convert an array of string to an array of int
func AtoiArray(input []string) ([]int, error) {
	output := make([]int, len(input))
	for i := 0; i < len(input); i++ {
		num, err := strconv.Atoi(input[i])
		if err != nil {
			return nil, err
		}
		output[i] = num
	}
	return output, nil
}

// reverse an array
func ReverseArray(input []int) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}

// get the Least Common Multiple of A and B
func LCM(a, b int) int {
	for b != 0 {
		remainder := a % b
		a = b
		b = remainder
	}
	return a
}

// detect overlap between a range AA' and a range BB' using their start/end as coordinates
func checkIntersect(aStart int, aEnd int, bStart int, bEnd int) (int, string) {
	if aEnd < bStart || bEnd < aStart {
		if aEnd == bStart-1 {
			return 7, "A glued to B by the left but no intersect"
		} else if bEnd == aStart-1 {
			return 8, "A glued to B by the right but no intersect"
		} else {
			return 0, "no overlap and no glued data"
		}
	} else if aStart >= bStart && aEnd <= bEnd {
		return 1, "B fully overlap A"
	} else if aStart <= bStart && aEnd <= bEnd {
		return 2, "single point right"
	} else if aEnd >= bEnd && aStart >= bEnd {
		return 3, "single point left"
	} else if aStart < bStart && aEnd <= bEnd {
		return 4, "A overlap B from the left"
	} else if aStart >= bStart && aEnd > bEnd {
		return 5, "A overlap B from the right"
	} else if bStart > aStart && bEnd < aEnd {
		return 6, "A fully overlap B"
	}
	return -1, "ERROR case not handled"
}
