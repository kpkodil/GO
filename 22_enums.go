// _Перечисляемые типы_ (enums) — это особый случай
// [суммовых типов](https://en.wikipedia.org/wiki/Algebraic_data_type).
// Перечисляемый тип — это тип, который имеет фиксированное количество возможных
// значений, каждое из которых имеет свое уникальное имя. В Go нет
// отдельного типа enum как особой функции языка, но перечисления
// легко реализовать с помощью существующих идиом языка.

package main

import "fmt"

// Наш перечисляемый тип `ServerState` имеет базовый тип `int`.
type ServerState int

// Возможные значения для `ServerState` определены как
// константы. Специальное ключевое слово [iota](https://go.dev/ref/spec#Iota)
// автоматически генерирует последовательные значения констант; в этом
// случае 0, 1, 2 и так далее.
const (
	StateIdle = iota
	StateConnected
	StateError
	StateRetrying
)

// Реализовав интерфейс [fmt.Stringer](https://pkg.go.dev/fmt#Stringer),
// значения `ServerState` можно вывести на печать или преобразовать
// в строки.
//
// Это может быть утомительно, если существует много возможных значений. В таких
// случаях можно использовать [утилиту stringer](https://pkg.go.dev/golang.org/x/tools/cmd/stringer)
// вместе с `go:generate` для автоматизации этого процесса. См. [этот пост](https://eli.thegreenplace.net/2021/a-comprehensive-guide-to-go-generate)
// для более подробного объяснения.
var stateName = map[ServerState]string{
	StateIdle:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {
	ns := transition(StateIdle)
	fmt.Println(ns)
	// Если у нас есть значение типа `int`, мы не можем передать его в `transition` -
	// компилятор выдаст ошибку несоответствия типов. Это обеспечивает некоторую степень
	// безопасности типов на этапе компиляции для перечислений.

	ns2 := transition(ns)
	fmt.Println(ns2)
}

// transition имитирует переход состояния для
// сервера; принимает текущее состояние и возвращает
// новое состояние.
func transition(s ServerState) ServerState {
	switch s {
	case StateIdle:
		return StateConnected
	case StateConnected, StateRetrying:
		// Предположим, что мы проверяем некоторые условия здесь, чтобы
		// определить следующее состояние...
		return StateIdle
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknown state: %s", s))
	}
}
