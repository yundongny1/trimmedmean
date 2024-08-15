package trimmedmean

import (
	"fmt"
	"math"
	"testing"
)

func floatsAreEqual(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

func TestTrimmedMeanFloat(t *testing.T) {
	want := 0.54695
	x := []float64{0.837, 0.483, 0.627, 0.748, 0.374, 0.284, 0.627, 0.193, 0.927, 0.938,
		0.472, 0.394, 0.563, 0.849, 0.293, 0.583, 0.937, 0.273, 0.817, 0.638,
		0.583, 0.738, 0.293, 0.849, 0.274, 0.374, 0.593, 0.482, 0.392, 0.718,
		0.927, 0.274, 0.384, 0.938, 0.273, 0.847, 0.194, 0.372, 0.849, 0.638,
		0.294, 0.483, 0.748, 0.583, 0.927, 0.293, 0.483, 0.192, 0.748, 0.839,
		0.483, 0.192, 0.627, 0.748, 0.372, 0.837, 0.294, 0.284, 0.749, 0.483,
		0.928, 0.748, 0.583, 0.937, 0.284, 0.384, 0.748, 0.927, 0.483, 0.293,
		0.294, 0.627, 0.748, 0.837, 0.384, 0.273, 0.847, 0.372, 0.384, 0.293,
		0.948, 0.372, 0.849, 0.382, 0.749, 0.284, 0.839, 0.483, 0.192, 0.927,
		0.384, 0.749, 0.928, 0.372, 0.483, 0.294, 0.372, 0.748, 0.927, 0.382}

	got, err := TrimmedMeanFloat(x, 0.3)
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
	}

	if !floatsAreEqual(want, got, 1e-6) {
		t.Errorf("Expected %f, but got %f", want, got)
	}
}

func TestTrimmedMeanInt(t *testing.T) {
	want := 57.357143
	y := []int{52, 93, 15, 72, 61, 21, 83, 87, 75, 75, 88, 24, 3, 22, 53, 2, 69, 70, 79, 47}

	got, err := TrimmedMeanInt(y, 0.15)
	if err != nil {
		// Handle the error
		fmt.Println("Error:", err)
	}

	if !floatsAreEqual(want, got, 1e-6) {
		t.Errorf("Expected %f, but got %f", want, got)
	}
}
