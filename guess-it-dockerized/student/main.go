package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var dataSlice []float64
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input!!!!")
			continue
		}
		dataSlice = append(dataSlice, num)
		if len(dataSlice) > 1 {
			lower, upper := expect(dataSlice)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}
}

func expect(data []float64) (int, int) {
	s := len(data) - 4
	if s < 0 {
		s = 0
	}
	a := average(data[s:])
	v := variance(data[s:])
	dv := math.Sqrt(v)
	i, j := ans(a, dv)
	return i, j
}

func ans(i, j float64) (int, int) {
	if j < 0 {
		j = -j
	}
	return int(i - (1.2 * j)), int(i + (1.2 * j))
}

func average(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func variance(numbers []float64) float64 {
	average := average(numbers)
	var sumOfSquares uint64 = 0
	for _, num := range numbers {
		sumOfSquares += uint64((num)-average) * uint64((num)-average)
	}
	return float64(sumOfSquares) / float64(len(numbers))
}
