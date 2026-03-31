// Package randomness предоставляет инструменты для работы со случайными числами.
// А также со срезами случайных чисел, в том числе с учётом весов

package randomness

import (
	"math"
	"math/rand"
)

// NewSequenceSliceOneStart создает и возвращает срез целых чисел от 1 до n включительно.
func NewSequenceSliceOneStart(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i + 1
	}

	return slice
}

// NewSequenceSliceZeroStart создает и возвращает срез целых чисел от 0 до n включительно.
func NewSequenceSliceZeroStart(n int) []int {
	slice := make([]int, n)
	for i := range slice {
		slice[i] = i
	}

	return slice
}

// NewSymAbsSlice создает симметричный срез чисел, где минимальное значение находится в центре,
// а значения увеличиваются к краям;
// n - количество элементов в срезе.
func NewSymAbsSlice(n int) []int {
	if n <= 0 {
		return nil
	}
	slice := make([]int, n)
	mid := float64(n-1) / 2.0 // центр (для n=6 -> 2.5)
	for i := 0; i < n; i++ {
		dist := int(math.Abs(float64(i)-mid) + 0.5) // округлённое расстояние
		slice[i] = dist + 1
	}
	return slice
}

// NewCumulativeSlice создает кумулятивный срез на основе переданных весов.
// Каждый элемент равен сумме всех предыдущих весов, включая текущий.
func NewCumulativeSlice(weights []int) []int {
	j := 0
	slice := make([]int, cap(weights))
	for i := range weights {
		slice[i] = weights[i] + j
		j += weights[i]
	}

	return slice
}

// GetRandomNum возвращает случайное число в диапазоне от 0 до max включительно.
func GetRandomNum(max int) int {
	return rand.Intn(max + 1)
}

// FindIndexByCumSum находит индекс, который соответсвует куммулятивной сумме.
func FindIndexByCumSum(cum []int, target int) int {
	for i, v := range cum {
		if v >= target {
			return i
		}
	}
	return -1
}
