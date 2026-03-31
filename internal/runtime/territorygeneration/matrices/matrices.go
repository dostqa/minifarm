// Package matrices содержит описание типа Matrix - матриц.
// А также предоставляет дополнительные функции для работы с Matrix.
package matrices

import (
	"fmt"
	"minifarm/internal/runtime/territorygeneration/worldtypes"
)

type WOType = worldtypes.WorldObjectType

// NewMatrix возвращает пустую матрицу m-строк на n-столбцов
func NewMatrix(m, n int) Matrix {
	matrix := make(Matrix, m)
	for i := range m {
		matrix[i] = make([]WOType, n)
	}
	return matrix
}

type Matrix [][]WOType

// M возвращает количество строк в матрице
func (matrix Matrix) M() int {
	return len(matrix)
}

// N возвращает количество столбцов в матрице
func (matrix Matrix) N() int {
	return len(matrix[0])
}

// Show выводит матрицу в консоль
func (matrix Matrix) Show() {
	// Выводим номера столбцов
	fmt.Printf("    ")
	for col := range matrix[0] {
		fmt.Printf("%d  ", col+1)
	}
	fmt.Println()

	// Выводим каждую строку с номером
	var i int
	for m := range matrix {
		i += 1
		fmt.Printf("%d  ", i+1)
		for n := range matrix[m] {
			if matrix[n][m] == 0 {
				fmt.Printf("  ") // Два пробела для выравнивания
			} else {
				fmt.Printf("%v  ", matrix[n][m])
			}
		}
		fmt.Println()
	}
}
