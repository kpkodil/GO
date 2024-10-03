// Go поддерживает [_анонимные функции_](https://en.wikipedia.org/wiki/Anonymous_function),
// которые могут формировать <a href="https://en.wikipedia.org/wiki/Closure_(computer_science)"><em>замыкания</em></a>.
// Анонимные функции полезны, когда вы хотите определить
// функцию прямо в коде без необходимости её именования.

package main

import "fmt"

// Эта функция `intSeq` возвращает другую функцию, которую
// мы определяем анонимно в теле `intSeq`. Возвращённая функция
// _замыкается_ на переменной `i`, формируя замыкание.
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {

	// Мы вызываем `intSeq`, присваивая результат (функцию)
	// переменной `nextInt`. Это значение функции захватывает
	// своё собственное значение `i`, которое будет обновляться
	// при каждом вызове `nextInt`.
	nextInt := intSeq()

	// Убедитесь в эффекте замыкания, вызвав `nextInt`
	// несколько раз.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Чтобы подтвердить, что состояние уникально для этой
	// конкретной функции, создайте и протестируйте новую.
	newInts := intSeq()
	fmt.Println(newInts())
}
