package describe

import (
	"testing"
)

func testMean(t *testing.T) {
	data := []float64{1,2,3,4,5,6,7}
	got := Mean(data)
	want := 3.0
	if got != want {
		t.Errorf("Mean() = %v, want %v", got, want)
	}
}

func testMedian(t *testing.T) {
	data := []float64{1,2,3,4,5,6,7}
	got := Median(data)
	want := 3.0
	if got != want {
		t.Errorf("Median() = %v, ant %v", got, want)
	}
}

func TestVariance(t *testing.T){
	data := []float64{1, 2, 3, 4, 5}
	got := Variance(data)
	want := 2.0 
	if got != want {
		t.Errorf("Variance() = %v, want %v", got, want)
	}
}

func TestStdDev(t *testing.T) {
    data := []float64{1, 2, 3, 4, 5}
    got := StdDev(data)
    want := 1.4142135623730951 
    if got != want {
        t.Errorf("StdDev() = %v, want %v", got, want)
    }
}