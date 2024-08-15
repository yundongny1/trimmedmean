package trimmedmean

import (
	"errors"
	"math"
	"sort"
)

func TrimmedMeanFloat(numbers []float64, trimValue float64) (float64, error) {
	sort.Float64s(numbers)
	sum := 0.0

	if trimValue < 0.5 {
		to_trim := uint(math.Floor(float64(trimValue * float64(len(numbers)))))

		trimmedValues := numbers[to_trim : len(numbers)-int(to_trim)]

		for _, number := range trimmedValues {
			sum += number
		}
		return sum / float64(len(trimmedValues)), nil
	} else {
		return 0.0, errors.New("invalid parameter values")
	}
}

func TrimmedMeanInt(numbers []int, trimValue float64) (float64, error) {
	sort.Ints(numbers)
	sum := 0

	if trimValue < 0.5 {
		to_trim := uint(math.Floor(float64(trimValue * float64(len(numbers)))))

		trimmedValues := numbers[to_trim : len(numbers)-int(to_trim)]

		for _, number := range trimmedValues {
			sum += number
		}
		return float64(sum) / float64(len(trimmedValues)), nil
	} else {
		return 0.0, errors.New("invalid parameter values")
	}
}
