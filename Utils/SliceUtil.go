package utils

import (
	"errors"
	"math/rand"
	"reflect"
	"time"
)

// [Action]: check if an item exist in a slice
// [Input]: slice, item
// [Output]: bool, error
func SliceContains(slice interface{}, item interface{}) (bool, error) {

	// check that input slice is of type Slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return false,
			errors.New("SliceContains: The first argument must be a slice")
	}

	// Type check if input item type matches input slice type
	if sliceValue.Type().Elem() != reflect.TypeOf(item) {
		return false,
			errors.New("SliceContains: The item type does not match the slice element type")
	}

	// look for the item in the slice
	for i := 0; i < sliceValue.Len(); i++ {
		if reflect.DeepEqual(sliceValue.Index(i).Interface(), item) {
			return true, nil
		}
	}

	return false, nil
}

// [Action]: cut a slice into x parts
// [Input]: slice, parts
// [Output]: slice of slice, error
func SliceArray(slice interface{}, parts int) ([]interface{}, error) {
	//check if slice if of type slice
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return []interface{}{},
			errors.New("SliceArray: input is not a slice")
	}

	//check if part is negative
	if parts <= 0 {
		return []interface{}{},
			errors.New("SliceArray: parts need to be a positive int")
	}

	//check if slice has enough items to be cut in x parts
	sliceLen := sliceValue.Len()
	if sliceLen < parts {
		return []interface{}{},
			errors.New("SliceArray: not enough elements in the array to be cut in the requested amount of parts")
	}

	//calculate the lenght of each part
	partSize := sliceLen / parts
	remaining := sliceLen % parts

	//create the parts
	result := make([]interface{}, parts)
	start := 0
	for i := 0; i < parts; i++ {
		end := start + partSize
		if remaining > 0 {
			end++
			remaining--
		}
		result[i] = sliceValue.Slice(start, end).Interface()
		start = end
	}

	return result, nil
}

// [Action]: cut a slice at every occurence of a specified item
// [Input]: slice, item
// [Output]: slice of slice, error
func SliceCutByItem(slice interface{}, item interface{}) ([][]interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return [][]interface{}{}, errors.New("SliceByItem: input is not a slice")
	}

	var result [][]interface{}
	part := []interface{}{}

	for i := 0; i < sliceValue.Len(); i++ {
		val := sliceValue.Index(i).Interface()
		if val == item {
			if len(part) > 0 {
				result = append(result, part)
				part = []interface{}{}
			}
		} else {
			part = append(part, val)
		}
	}

	if len(part) > 0 {
		result = append(result, part)
	}

	return result, nil
}

// [Action]: insert an item in a slice at a specific position
// [Input]: slice pointer , index, item
// [Output]: error
func SliceInsertAt(slicePtr interface{}, index int, item interface{}) error {
	sliceValue := reflect.ValueOf(slicePtr)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("SliceInsertAt: input is not a pointer to a slice")
	}

	sliceValue = sliceValue.Elem()

	// Check index validity
	if index < 0 || index > sliceValue.Len() {
		return errors.New("SliceInsertAt: index is out of range")
	}

	// Check item type compatibility
	if reflect.TypeOf(item) != sliceValue.Type().Elem() {
		return errors.New("SliceInsertAt: item type doesn't match slice element type")
	}

	// Create a new slice to hold the modified elements
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Len()+1)

	// Copy elements before the index
	for i := 0; i < index; i++ {
		newSlice.Index(i).Set(sliceValue.Index(i))
	}

	// Set the new item at the specified index
	newSlice.Index(index).Set(reflect.ValueOf(item))

	// Copy elements after the index
	for i := index; i < sliceValue.Len(); i++ {
		newSlice.Index(i + 1).Set(sliceValue.Index(i))
	}

	// Update original slice
	sliceValue.Set(newSlice)

	return nil
}

// [Action]: insert a slice in a slice at a specific position
// [Input]: slice pointer , index, slice to insert
// [Output]: error
func SliceInsertSliceAt(slice interface{}, index int, insertSlice interface{}) error {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("SliceInsertSliceAt: input is not a pointer to a slice")
	}

	sliceValue = sliceValue.Elem()

	// Check index validity
	if index < 0 || index > sliceValue.Len() {
		return errors.New("SliceInsertSliceAt: index is out of range")
	}

	insertSliceValue := reflect.ValueOf(insertSlice)
	if insertSliceValue.Kind() != reflect.Slice {
		return errors.New("SliceInsertSliceAt: insertSlice is not a slice")
	}

	// Check item type compatibility
	sliceElemType := sliceValue.Type().Elem()
	insertSliceElemType := insertSliceValue.Type().Elem()
	if sliceElemType != insertSliceElemType {
		return errors.New("SliceInsertSliceAt: insertSlice element type doesn't match slice element type")
	}

	// Create a new slice to hold the modified elements
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+insertSliceValue.Len(), sliceValue.Len()+insertSliceValue.Len())

	// Copy elements before the index
	for i := 0; i < index; i++ {
		newSlice.Index(i).Set(sliceValue.Index(i))
	}

	// Copy insertSlice elements
	for i := 0; i < insertSliceValue.Len(); i++ {
		newSlice.Index(index + i).Set(insertSliceValue.Index(i))
	}

	// Copy elements after the index
	for i := index; i < sliceValue.Len(); i++ {
		newSlice.Index(i + insertSliceValue.Len()).Set(sliceValue.Index(i))
	}

	// Update original slice
	sliceValue.Set(newSlice)

	return nil
}

// [Action]: remove from a slice X elements starting at Y position
// [Input]: slice pointer , startIndex, count
// [Output]: error
func SliceRemoveAt(slice interface{}, startIndex int, count int) error {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("SliceRemoveAt: input is not a pointer to a slice")
	}

	sliceValue = sliceValue.Elem()

	// Check startIndex validity
	if startIndex < 0 || startIndex >= sliceValue.Len() {
		return errors.New("SliceRemoveAt: startIndex is out of range")
	}

	// Check count validity
	if count <= 0 || startIndex+count > sliceValue.Len() {
		return errors.New("SliceRemoveAt: count is out of range or negative")
	}

	// Create a new slice to hold the modified elements
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()-count, sliceValue.Len()-count)

	// Copy elements before the startIndex
	for i := 0; i < startIndex; i++ {
		newSlice.Index(i).Set(sliceValue.Index(i))
	}

	// Copy elements after the removed section
	for i := startIndex + count; i < sliceValue.Len(); i++ {
		newSlice.Index(i - count).Set(sliceValue.Index(i))
	}

	// Update original slice
	sliceValue.Set(newSlice)

	return nil
}

// [Action]: find the first index matching provided item
// [Input]: slice, item
// [Output]: index, error
func SliceFirstIndexOf(slice interface{}, item interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return -1, errors.New("SliceFirstIndexOf: input is not a slice")
	}

	for i := 0; i < sliceValue.Len(); i++ {
		currentItem := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(currentItem, item) {
			return i, nil
		}
	}

	return -1, errors.New("SliceFirstIndexOf: item not found in slice")
}

// [Action]: find the last index matching provided item
// [Input]: slice, item
// [Output]: index, error
func SliceLastIndexOf(slice interface{}, item interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return -1, errors.New("SliceLastIndexOf: input is not a slice")
	}

	lastIndex := -1

	for i := 0; i < sliceValue.Len(); i++ {
		currentItem := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(currentItem, item) {
			lastIndex = i
		}
	}

	if lastIndex != -1 {
		return lastIndex, nil
	}

	return -1, errors.New("SliceLastIndexOf: item not found in slice")
}

// [Action]: find all Indices matching provided item
// [Input]: slice, item
// [Output]: slice with all the int index, error
func SliceFindAllIndicesOf(slice interface{}, item interface{}) ([]int, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("SliceAllIndicesOf: input is not a slice")
	}

	indices := make([]int, 0)

	for i := 0; i < sliceValue.Len(); i++ {
		currentItem := sliceValue.Index(i).Interface()
		if reflect.DeepEqual(currentItem, item) {
			indices = append(indices, i)
		}
	}

	if len(indices) > 0 {
		return indices, nil
	}

	return nil, errors.New("SliceAllIndicesOf: item not found in slice")
}

// [Action]: reverse all the elements of a slice
// [Input]: slice
// [Output]: error
func SliceReverse(slice interface{}) error {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("SliceReverse: input is not a pointer to a slice")
	}

	sliceValue = sliceValue.Elem()

	for i, j := 0, sliceValue.Len()-1; i < j; i, j = i+1, j-1 {
		temp := reflect.New(sliceValue.Index(i).Type()).Elem()
		temp.Set(sliceValue.Index(i))
		sliceValue.Index(i).Set(sliceValue.Index(j))
		sliceValue.Index(j).Set(temp)
	}

	return nil
}

// [Action]: randomize the position of all the elements of a slice
// [Input]: slice
// [Output]: error
func SliceShuffle(slice interface{}) error {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Ptr || sliceValue.Elem().Kind() != reflect.Slice {
		return errors.New("SliceShuffle: input is not a pointer to a slice")
	}

	sliceValue = sliceValue.Elem()
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := sliceValue.Len() - 1; i > 0; i-- {
		j := random.Intn(i + 1)
		temp := reflect.New(sliceValue.Index(i).Type()).Elem()
		temp.Set(sliceValue.Index(i))
		sliceValue.Index(i).Set(sliceValue.Index(j))
		sliceValue.Index(j).Set(temp)
	}

	return nil
}

// [Action]: get all unique values of elements of a slice
// [Input]: slice
// [Output]: error
func SliceUnique(slice interface{}) ([]interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("SliceUnique: input is not a slice")
	}

	uniqueMap := make(map[interface{}]struct{})
	uniqueSlice := make([]interface{}, 0)

	for i := 0; i < sliceValue.Len(); i++ {
		item := sliceValue.Index(i).Interface()
		if _, exists := uniqueMap[item]; !exists {
			uniqueMap[item] = struct{}{}
			uniqueSlice = append(uniqueSlice, item)
		}
	}

	return uniqueSlice, nil
}

// [Action]: take in a slice and return an array of matching len and cap
// [Input]: slice
// [Output]: array, error
func SliceConvertToArray(slice interface{}) (interface{}, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("SliceConvertToArray: input is not a slice")
	}

	arrayType := reflect.ArrayOf(sliceValue.Len(), sliceValue.Type().Elem())
	newArray := reflect.New(arrayType).Elem()

	for i := 0; i < sliceValue.Len(); i++ {
		newArray.Index(i).Set(sliceValue.Index(i))
	}

	return newArray.Interface(), nil
}

// [Action]: take in a slice and  an item, return number of occurence of this item in the array
// [Input]: slice, item
// [Output]: int, error
func SliceItemOccurence(slice interface{}, item interface{}) (int, error) {
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return -1, errors.New("SliceItemOccurence: input is not a slice")
	}

	// Check item type compatibility
	if reflect.TypeOf(item) != sliceValue.Type().Elem() {
		return -1, errors.New("SliceItemOccurence: item type doesn't match slice element type")
	}

	count := 0
	for i := 0; i < sliceValue.Len(); i++ {
		if sliceValue.Index(i) == item {
			count++
		}
	}
	return count, nil
}
