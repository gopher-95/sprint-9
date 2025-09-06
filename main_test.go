package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	testList := []int{0, 1, 7}

	trueAnswer := []int{0, 1, 7}

	for i, list := range testList {
		require.Equal(t, trueAnswer[i], len(generateRandomElements(list)))
	}

}

func TestMaximum(t *testing.T) {
	//создаем слайсы, которые будут участвовать в тестировании, затронуты такие случаи, когда len(slice) == 0 и len(slice) == 1, повторяющиеся элементы и т. д.
	lists := [][]int{
		{1, 3, 5, 2, 6, 8, 1, 4},
		{3, 1, 11, 78, 92, 4},
		{17, 90},
		{10},
		{},
		{4, 4, 4, 4, 4, 4},
		{0, 0, 0, 0, 0},
		{1, 2, 3, 4, 5, 6, 7},
	}

	trueAnswer := []int{8, 92, 90, 10, 0, 4, 0, 7}

	for i, list := range lists {
		require.Equal(t, trueAnswer[i], maximum(list))
	}
}
