package main

func main() {
	numbers := []int{0, 50}

	print(Average(numbers))
}

func Average(numbers []int) int {
	var sum int
	for _, n := range numbers {
		sum += n
	}
	return sum/len(numbers)
}


func Max(numbers []int) int {
	var max int
	for idx, n := range numbers {
		if idx == 0 || n < max {
			max = n
		}
	}
	return max
}
