package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	//проверка на размер слайса
	if size == 0 || size <= 0 {
		fmt.Println("incorrect value of size")
		return nil
	}
	//иницилизируем слайс размером size
	slice := make([]int, size)
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	for i := 0; i < len(slice); i++ {
		//максимальное сгенерированное число - 100
		randomNumber := r.Intn(101)
		//заполняем слайс сгенерированными случайными числами
		slice[i] = randomNumber
	}
	return slice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	} else if len(data) == 1 {
		return data[0]
	} else {
		//нахождение максимального числа в слайсе
		var maxNumber int = 0
		for _, value := range data {
			if value > maxNumber {
				maxNumber = value
			} else if value < maxNumber {
				value = maxNumber
			}
		}
		return maxNumber
	}

}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {

	//создаем WaitGroup
	var wg sync.WaitGroup

	//слайс максимумов
	finalSlice := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		//количество элементов в каждом кусочке слайса
		elementInEachPart := len(data) / CHUNKS
		//начальный индекс
		startIndex := i * elementInEachPart
		//конечный индекс среза
		endIndex := startIndex + elementInEachPart
		//один из восьми слайсов
		partOfData := data[startIndex:endIndex]
		//добавляем с помощью WaitGroup одну горутину
		wg.Add(1)
		go func(slice []int) {
			defer wg.Done()
			maxNumber := maximum(slice)
			finalSlice[i] = maxNumber
		}(partOfData)

	}

	wg.Wait()

	return maximum(finalSlice)

}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	// ваш код здесь
	bigSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	now := time.Now()
	max := maximum(bigSlice)
	elapsed := time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков", CHUNKS)
	now = time.Now()
	max = maxChunks(bigSlice)
	elapsed = time.Since(now)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
