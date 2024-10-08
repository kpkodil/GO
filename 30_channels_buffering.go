// По умолчанию каналы _небуферизованы_, это означает, что они
// будут принимать данные (`chan <-`) только если есть
// соответствующее получение (`<- chan`), готовое получить
// отправленное значение. _Буферизованные каналы_ могут принимать
// ограниченное количество значений без необходимости
// немедленного получения.

package main

import "fmt"

func main() {

	// Здесь мы создаем канал строк с буфером на 2 значения.
	messages := make(chan string, 2)

	// Поскольку этот канал буферизован, мы можем отправить
	// эти значения в канал, даже если в данный момент
	// нет соответствующих операций получения.
	messages <- "buffered"
	messages <- "channel"

	// Позже мы можем получить эти два значения как обычно.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// Пояснение:
// Небуферизованные каналы: Обычно каналы в Go блокируют выполнение программы при отправке значения, пока не будет готов получатель. Это предотвращает накопление данных в канале без их обработки. Таким образом, отправка и получение происходят синхронно.

// Буферизованные каналы: Когда канал буферизован (в данном случае размер буфера равен 2), вы можете отправлять несколько значений без немедленного получения. Здесь мы отправляем два значения в канал, и программа не блокируется, даже если операция получения происходит позже.

// Отправка и получение: Мы сначала отправляем строки "buffered" и "channel" в канал. Несмотря на отсутствие получателя в этот момент, канал сохраняет значения, благодаря буферу. Позже мы можем получить эти значения и вывести их с помощью fmt.Println().

// Как это работает:
// Мы создали буферизованный канал с ёмкостью 2, что позволяет отправить два значения, не ожидая немедленного получения.
// После отправки данных мы можем в любой момент считать эти значения из канала, что делает работу с каналами более гибкой, особенно если обработка данных происходит в асинхронном режиме.
// Буферизованные каналы полезны, когда нужно сохранять данные до их последующей обработки.
