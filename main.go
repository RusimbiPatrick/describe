package main

import (
    "fmt"
    "github.com/RusimbiPatrick/describe/describe"
)

func main() {
    data := []float64{1, 2, 3, 4, 5}
    fmt.Println("Mean:", describe.Mean(data))
    fmt.Println("Median:", describe.Median(data))
    fmt.Println("Variance:", describe.Variance(data, false))
    fmt.Println("StdDev:", describe.StdDev(data, false))
}
