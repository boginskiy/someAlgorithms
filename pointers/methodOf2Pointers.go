package main

import "fmt"

func MovingAverage(arr []int, k int) []float64 {
	result := []float64{}
	sum := 0

	// Первое значение вычисляем стандартным способом
	for _, v := range arr[:k] {
		sum += v
	}
	result = append(result, float64(sum/k))

	for i := 0; i < len(arr)-3; i++ {
		sum -= arr[i]
		sum += arr[i+k]
		result = append(result, float64(sum)/float64(k))
	}
	return result
}

func main() {
	arr := []int{4, 3, 8, 1, 5, 6, 3}
	k := 3

	res := MovingAverage(arr, k)
	fmt.Println(res) // [5 4 4.666666666666667 4 4.666666666666667]
}
