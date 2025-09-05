package main

import "testing"

// Пишите тесты в этом файле
func TestGenerateRandomElements(t *testing.T) {
	//проверка на возвращаемую длину слайса в зависимости от размера size
	zeroElementSlice := generateRandomElements(0)
	if len(zeroElementSlice) != 0 {
		t.Fatal("длина слайса не равна нулю!")
	}

	oneElementSlice := generateRandomElements(1)
	if len(oneElementSlice) != 1 {
		t.Fatal("длина слайса не равна единице!")
	}

	sevenElementsSlice := generateRandomElements(7)
	if len(sevenElementsSlice) != 7 {
		t.Fatal("длина слайса не равна семи!")
	}
	result := generateRandomElements(20)
	for _, number := range result {
		if number > 100 {
			t.Fatal("сгенерированные числа не входят в сотню") //в функции generateRandomElements поставлено ограничение на генерированное рандомное число - оно не больше 100 должно быть
		}
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
		if maximum(list) != trueAnswer[i] {
			t.Fatal(i, ":", maximum(list), "!=", trueAnswer[i])
		}
	}
}
