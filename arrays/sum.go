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

// variadic function, take a variable number of arguments
func SumAll(numbersToSum ...[]int) []int {
	numOfSlices := len(numbersToSum)
	// make allows you to create a slice with a starting capacity
	sums := make([]int, numOfSlices)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// [low:high], [low:] -> until the end
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
