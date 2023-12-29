package arraysslices

func Sum(arr []int) (Sum int) {
	for _, val := range arr {
		Sum += val
	}
	return
}

func SumAll(sample ...[]int) []int {
	var sum []int
	for _, set := range sample {
		sum = append(sum, Sum(set))
	}
	return sum
}

func SumAllTails(sample ...[]int) []int {
	var sum []int
	for _, set := range sample {
		if len(set) == 0 {
			sum = append(sum, 0)
			continue
		}
		sum = append(sum, Sum(set[1:]))
	}
	return sum
}
