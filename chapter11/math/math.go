package math

func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

func Max(xs []float64) float64 {
	max := float64(0)
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

func Min(xs []float64) float64 {
	if len(xs) == 0 {
		return 0
	}
	min := xs[0]
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}
