package main

import (
	"fmt"
	"time"
)

// Функция, которую мы будем запускать в goroutine
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// Синхронный вызов функции
	f("direct")

	// Запуск функции в новой goroutine
	go f("goroutine")

	// Запуск анонимной функции в новой goroutine
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Ожидание завершения goroutines
	// Использование time.Sleep для ожидания не является надежным методом
	// для синхронизации goroutines и используется здесь только для примера.
	time.Sleep(time.Second)
	fmt.Println("done")
}

// В Go goroutines позволяют запускать функции параллельно и эффективно использовать многозадачность. Они легковесные и управляются Go-рантаймом, что упрощает параллельное выполнение задач. Вот как можно использовать и управлять goroutines на практике:

// Объяснение
// Синхронный вызов:

// f("direct") выполняется в основном потоке программы.
// Асинхронные вызовы:

// go f("goroutine") запускает f в отдельной goroutine, позволяя ей работать параллельно.
// go func(msg string) {...}("going") запускает анонимную функцию в отдельной goroutine.
// Ожидание:

// time.Sleep(time.Second) используется для ожидания завершения выполнения goroutines. Это простое решение для примера, но в реальных приложениях рекомендуется использовать sync.WaitGroup для надежного ожидания завершения.

// Более Надежный Подход
// Для надежного ожидания завершения всех goroutines, можно использовать sync.WaitGroup:

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// // Функция, которую мы будем запускать в goroutine
// func f(from string, wg *sync.WaitGroup) {
// 	defer wg.Done() // Уменьшить счетчик после завершения
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(from, ":", i)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// Добавить счетчик goroutines
// 	wg.Add(2)

// 	// Запуск функции в новой goroutine
// 	go f("goroutine", &wg)

// 	// Запуск анонимной функции в новой goroutine
// 	go func(msg string) {
// 		defer wg.Done() // Уменьшить счетчик после завершения
// 		fmt.Println(msg)
// 	}("going")

// 	// Ожидание завершения всех goroutines
// 	wg.Wait()
// 	fmt.Println("done")
// }

// Заключение
// Использование goroutines в Go позволяет легко и эффективно выполнять задачи параллельно. Для правильного управления завершением goroutines рекомендуется использовать sync.WaitGroup или другие синхронизационные механизмы, чтобы избежать потенциальных проблем и ошибок в реальных приложениях.
