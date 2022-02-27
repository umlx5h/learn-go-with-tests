package main

func Sum(numbers []int) int {
	var total int
	for _, n := range numbers {
		total += n
	}

	return total
}

func SumAll(numbersToSum ...[]int) []int {
	// sums := make([]int, 0, len(numbersToSum))
	var sums []int

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(slices ...[]int) []int {
	var tails []int

	for _, slice := range slices {
		if len(slice) == 0 {
			tails = append(tails, 0)
			continue
		}
		tails = append(tails, slice[len(slice)-1])
	}

	return tails
}
