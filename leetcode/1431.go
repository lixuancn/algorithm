func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	res := make([]bool, n)
	max := 0
	for _, candy := range candies{
		if candy > max{
			max = candy
		}
	}
	for i, candy := range candies{
		res[i] = candy + extraCandies >= max
	}
	return res
}
