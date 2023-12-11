package main

import (
	"reflect"
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

// check if an item exist in a slice
// input: slice, item
// output: bool
func SliceContains(slice interface{}, item interface{}) bool {

	// check that input slice is of type Slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		panic("The first argument must be a slice")
	}

	// Type check if input item type matches input slice type
	if sliceValue.Type().Elem() != reflect.TypeOf(item) {
		panic("The item type does not match the slice element type")
	}

	// look for the item in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		if reflect.DeepEqual(sliceValue.Index(i).Interface(), item) {
			return true
		}
	}

	return false
}

// insert an item in a slice at a specific position
// input slice pointer, item, index
// output: modify directly the provided slice
func SliceInsertAt(slicePtr interface{}, item interface{}, index int) {

	// check if provided slice pointer is of correct type
	sliceValue := reflect.ValueOf(slicePtr)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		panic("The first argument must be a pointer to a slice")
	}

	// Type check if input item type matches input slice type
	slice := sliceValue.Elem()
	if slice.Type().Elem() != reflect.TypeOf(item) {
		panic("The item type does not match the slice element type")
	}

	// check if the index provided isn't out of bounds
	if index < 0 || index > slice.Len() {
		panic("Index is out of range")
	}

	sliceLen := slice.Len()
	slice = reflect.Append(slice, reflect.Zero(slice.Type().Elem())) // Append a zero value to extend the slice

	// Shift elements after the index to the right by one position
	for i := sliceLen; i > index; i-- {
		slice.Index(i).Set(slice.Index(i - 1))
	}

	// Insert the new item at the specified index
	slice.Index(index).Set(reflect.ValueOf(item))

	// Update the pointer to the modified slice
	sliceValue.Elem().Set(slice)
}

// compare 2 slices and check if they are equal
func SlicesEqual(slice1, slice2 interface{}) bool {
	sliceValue1 := reflect.ValueOf(slice1)
	sliceValue2 := reflect.ValueOf(slice2)

	if sliceValue1.Kind() != reflect.Slice || sliceValue2.Kind() != reflect.Slice {
		return false
	}

	if sliceValue1.Type() != sliceValue2.Type() {
		return false
	}

	if sliceValue1.Len() != sliceValue2.Len() {
		return false
	}

	for i := 0; i < sliceValue1.Len(); i++ {
		if !reflect.DeepEqual(sliceValue1.Index(i).Interface(), sliceValue2.Index(i).Interface()) {
			return false
		}
	}

	return true
}
