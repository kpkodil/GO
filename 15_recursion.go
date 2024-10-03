// Go поддерживает
// <a href="https://en.wikipedia.org/wiki/Recursion_(computer_science)"><em>рекурсивные функции</em></a>.
// Вот классический пример.

package main

import "fmt"

// Эта функция `fact` вызывает саму себя до тех пор, пока не достигнет
// базового случая `fact(0)`.
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(fact(7))

	// Замыкания также могут быть рекурсивными, но это требует,
	// чтобы замыкание было явно объявлено с типом `var`
	// до того, как оно будет определено.
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		// Поскольку `fib` была ранее объявлена в `main`, Go
		// знает, какую функцию вызвать с помощью `fib` здесь.
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
