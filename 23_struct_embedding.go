// Go поддерживает _встраивание_ структур и интерфейсов
// для выражения более бесшовной _композиции_ типов.
// Это не следует путать с [`//go:embed`](embed-directive), который является
// директивой Go, введенной в версии 1.16+ для встраивания
// файлов и папок в бинарный файл приложения.

package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

// `container` _встраивает_ `base`. Встраивание выглядит
// как поле без имени.
type container struct {
	base
	str string
}

func main() {

	// При создании структур с помощью литералов, нам нужно
	// явно инициализировать встраивание; здесь
	// встроенный тип служит как имя поля.
	co := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	// Мы можем получить доступ к полям `base` непосредственно через `co`,
	// например, `co.num`.
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// В качестве альтернативы, мы можем указать полный путь, используя
	// имя встроенного типа.
	fmt.Println("also num:", co.base.num)

	// Поскольку `container` встраивает `base`, методы `base`
	// также становятся методами `container`. Здесь
	// мы вызываем метод, который был встроен из `base`
	// непосредственно через `co`.
	fmt.Println("describe:", co.describe())

	type describer interface {
		describe() string
	}

	// Встраивание структур с методами может использоваться для предоставления
	// реализаций интерфейсов другим структурам. Здесь
	// мы видим, что `container` теперь реализует
	// интерфейс `describer`, потому что он встраивает `base`.
	var d describer = co
	fmt.Println("describer:", d.describe())
}
