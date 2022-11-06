// Цели обучения:
// Практика работы с arr, for или range, min, max

// Узнайте, как Golang обрабатывает основные типы

// Инструкция:
// Прочитайте документацию Golang о том, как реализовать минимальные
//  и максимальные значения.

// Задача:
// Дан массив из пяти целых чисел.
// Найдите максимально и минимально возможную сумму, которую можно
// вычислить, суммируя ровно 4 из 5 элементов массива.
// Выведите соответствующие значения на стандартный вывод в «минимум: x,
//  максимум: y».

// Напишите модульные тесты, чтобы проверить, правильно ли работает ваш код с разными массивами
// (не менее пяти наборов данных). Для тестов используйте библиотеку testify (https://github.com/stretchr/testify)

// * Попробуйте реализовать без sort()

// Использовать цикл: берем первый элемент как начальный и затем в цикле сравниваем и если больше, то записываем в максимально значение
// Ниже реализация двух вариантов на Golang:

package main

import "fmt"

func main() {
	var arr = [5]int{10, 11, 12, 13, 14}
	min, max := findMinAndMax(arr)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
}
func findMinAndMax(arr [5]int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

// 	// max := arr[0]
// 	var min float64 = 0

// 	arr := [5]float64{10, 11, 12, 13, 14}
// 	for _, value := range arr {
// 		// fmt.Printf("%d\n", value)

// 		if value > min {
// 			min += value
// 			fmt.Println("min", min)
// 		}

// 		// if max < value {
// 		// 	max = value

// 	}
// 	// if min > value {
// 	// 	min = value
// 	// 	fmt.Println("min", min)
// 	// }
// }
