package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var data [12501]int
	i, c := 0, 0
	for {

		_, err := fmt.Scanf("%d\n", &i)
		if err != nil {
			fmt.Println(err)
			return
		}
		if i >= 0 && i <= 400 {
			data[c] = i
			var a, b int = i, i
			if c > 10 {
				a, b = expect(i, data[:c+1])
				// } else {
				// 	a, b = i, i
			}
			fmt.Printf("%d %d\n", a-10, b+10)
			// fmt.Printf("                                 %t", good(data[c], i, int(f)))
		} else {
			data[c] = 150
			fmt.Printf("%d %d\n", int(median(data[:c+1])), int(median(data[:c+1])))
		}
		c++
	}
}

//	func good(a, b, c int) bool {
//		return (b-c <= a && b+c >= a)
//	}
// var a float64

func expect(i int, data []int) (int, int) {
	a := average(data)
	// a := median(data)
	// if i >= 20 {
	// m := median(data[len(data)-9:])
	// }
	v := variance(data)
	dv := math.Sqrt(v)
	// i, j := ans(float64(i), a, 2.0)
	i, j := ans(a, dv, 5.0)
	return i, j
	// return int(a), int(a)
}

func ans(i, j, n float64) (int, int) {
	if i-j >= n {
		// i = i - 3
		j = i - n
	} else if i-j <= n {
		// i = i + 3
		j = i + n
	}
	if i < j {
		return int(i), int(j)
	}
	return int(j), int(i)
}

func average(numbers []int) float64 {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return float64(sum) / float64(len(numbers))
}

func median(numbers []int) float64 {
	n := len(numbers)
	// sortedNumbers := make([]int, n)
	sort.Ints(numbers)
	if n%2 == 1 {
		return float64(numbers[n/2])
	}
	return float64(numbers[n/2-1]+numbers[n/2]) / 2.0
}

func variance(numbers []int) float64 {
	average := average(numbers)
	sumOfSquares := 0.0
	for _, num := range numbers {
		sumOfSquares += (float64(num) - average) * (float64(num) - average)
	}
	return sumOfSquares / float64(len(numbers))
}

func standardDeviation(numbers []int) float64 {
	return math.Sqrt(variance(numbers))
}
