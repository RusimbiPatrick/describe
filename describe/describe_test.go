package describe

import (
	"testing"
)

func TestMean(t *testing.T) {
	data := []float64{1,2,3,4,5,6,7}
	got := Mean(data)
	want := 3.0
	if got != want {
		t.Errorf("Mean() = %v, want %v", got, want)
	}
}

func TestMedian(t *testing.T) {
	data := []float64{1,2,3,4,5,6,7}
	got := Median(data)
	want := 3.0
	if got != want {
		t.Errorf("Median() = %v, ant %v", got, want)
	}
}

func TestPopulationVariance(t *testing.T){
	data := []float64{1, 2, 3, 4, 5}
	got := Variance(data, false)
	want := 2.0 
	if got != want {
		t.Errorf("Variance() = %v, want %v", got, want)
	}
}

func TestSampleVariance(t *testing.T){
	data := []float64{1, 2, 3, 4, 5}
	got := Variance(data, true)
	want := 2.5 
	if got != want {
		t.Errorf("Variance() = %v, want %v", got, want)
	}
}
func TestPopulationStdDev(t *testing.T) {
    data := []float64{1, 2, 3, 4, 5}
    got := StdDev(data, false)
    want := 1.4142135623730951 
    if got != want {
        t.Errorf("StdDev() = %v, want %v", got, want)
    }
}

func TestSampleStdDev(t *testing.T) {
    data := []float64{1, 2, 3, 4, 5}
    got := StdDev(data, true)
    want := 1.5811388300841898
    if got != want {
        t.Errorf("StdDev() = %v, want %v", got, want)
    }
}