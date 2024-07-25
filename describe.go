package describe

import(
	"math"
	"sort"
)

func Mean(data []float64) float64{
	sum :=0.0
	for _, value := range data{
		sum += value
	}
	return sum /float64(len(data))
}


func Median(data []float64) float64 {
	sorted := make([]float64, len(data))
	copy(sorted, data)
	sort.Float64s(sorted)

	n:= len(sorted)
	if n %2 == 0 {
		return (sorted[n/2-1] + sorted[n/2]) / 2.0
	}
	return sorted[n/2]
}

func Variance(data []float64, isSample bool) float64 {
	mean := Mean(data)
	sum := 0.0
	for _, value := range data {
		diff := value - mean
		sum += diff * diff
	}
	if isSample {
		// Sample variance
		return sum / float64(len(data) - 1)
	}

	// Population variance
	return sum / float64(len(data))
	
	
	
}

func StdDev(data []float64, isSample  bool) float64 {
    return math.Sqrt(Variance(data, isSample))
}