package main

import "fmt"

func main() {
	buffer()
}

func buffer() {
	ch := make(chan int, 2) // создал канал, максимальный буфер канала = 2
	fmt.Printf("Длина: %d, Каппасити: %d\n", len(ch), cap(ch))
	// Длина: 0, Каппасити: 2

	ch <- 3
	ch <- 2 // передал в канал значение 2 горутина их приняла.
	fmt.Printf("Длина: %d, Каппасити: %d\n", len(ch), cap(ch))
	//Длина: 1, Каппасити: 2

	// как только буфер заполняется - выдается ошибка: deadlock
	//fatal error: all goroutines are asleep - deadlock!
	ch <- 4

}
