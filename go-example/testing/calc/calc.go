package calc

import "math"

func Sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}
	return result
}

func Average(nums ...int) float64 {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	result := float64(sum) / float64(len(nums))
	pow := math.Pow(10, float64(2))
	digit := pow * result
	round := math.Floor(digit)
	return round / pow
}
