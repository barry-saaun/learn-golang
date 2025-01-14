package arrays

func Sum(numbers []int) int {
	var sumTotal int
	// range --> [idx, value]
	for _, number := range numbers {
		sumTotal += number
	}

	return sumTotal
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
