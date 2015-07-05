package math

import "testing"

type testpair struct {
	values []float64
	result float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

var testMax = []testpair{
	{[]float64{1, 2, 3, 4, 5, 6}, 6},
	{[]float64{1, 2, 3, 109, 5, 6}, 109},
	{[]float64{1, 2, 3, 4, 5, 6}, 6},
	{[]float64{}, 0},
}

func TestMax(t *testing.T) {
	for _, pair := range testMax {
		v := Max(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}

var testMin = []testpair{
	{[]float64{1, 2, 3, 4, 5, 6}, 1},
	{[]float64{324, 2, 3, 109, 5, 6}, 2},
	{[]float64{5, 6}, 5},
	{[]float64{}, 0},
}

func TestMin(t *testing.T) {
	for _, pair := range testMin {
		v := Min(pair.values)
		if v != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", v,
			)
		}
	}
}
