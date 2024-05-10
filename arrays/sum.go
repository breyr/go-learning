package arrays

// go slices do not encode the size of the collection and can have any size
// arrays have a fixed length

func Sum(numbers []int) int {
	sum := 0
	// _ blank identifier, here we ignore the index
	for _, number := range numbers {
		sum += number
	}
	return sum
}
