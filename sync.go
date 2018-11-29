//
// sync and topology primitives
//
package main

import(
	"fmt"
	"strings"
	"time"
)

const TIMEOUT = 800

// минимальная конфигурация

//
// main driver
//
func main() {
	//начало
	done := make(chan struct{finished bool})
	workerDone := make(chan struct{})

	// какие-то данные
	data := []string{
		"The yellow fish swims",
		"Blue is the water in a crystal glass",
		"Blue are the lilacs that grow in the grass",
		"Blue is delicious blueberry pie",
		"Blue are the sparkles in my cats eyes",
	}

	// одиночный 
	go func() {
		//сигнал синхронизации
		defer close(workerDone)
		// или 
		// workerDone<- struct{finished boo}{true}

		// восстановление после сбоев
		defer func() {
			//
			// To be able to recover from an unwinding panic sequence, 
			// the code must make a deferred call to the recover function.
			//
			if r := recover(); r != nil {
				fmt.Println("[W] Reader FAULT")
			}
		}()
		fmt.Printf(" * Reader started")
		for word := range wordsGenerator(data) {
			fmt.Println("word:", word)
		}
		fmt.Println()
		fmt.Println(" * Reader terminated")
	}()

	// корневой супревизор
	go func() {
		//сигнал синхронизации
		defer close(done)
		// восстановление после сбоев
		defer func() {
			//
			// To be able to recover from an unwinding panic sequence, 
			// the code must make a deferred call to the recover function.
			//
			if r := recover(); r != nil {
				fmt.Println("[W] Root Supervisor FAULT")
			}
		}()
		fmt.Printf(" * Root Supervisor started")

		// более общий вариант
		// обработка множественных каналов
		//for {
		//	select {
		//	case <-workerDone: // воркер сигналит об окончании работы
		//	case <-groupDone: // группа сигналит об окончании работы
		//	case <-workerFault: //<...> например, отказ воркера(упал или ещё чего)
		//	case <-time.After(TIMEOUT * time.Microsecond): // timeout...
		//	}
		//}

		for range workerDone {
		//for worker := range workerDone {
			//fmt.Println(" * Worker data:", worker)
			// или другой вариант
			// workersDoneCount++
			//if workersDoneCount >= WORKERS {
				// остановить цикл и отправить сигнал о завершении работы
			//	close(workerDone)
			//}
		}
		// данные 
		done<- struct{finished bool}{true}
		fmt.Println(" * Root Supervisor terminated")
	}()

	//контроллер сихронизатор приложения с таймаутом для ликвидации неопределённостей
	select {
	case value := <-done:
		// окончание
		fmt.Println("done:",value.finished)
	case <-time.After(TIMEOUT * time.Microsecond):
		// таймаут
		fmt.Println("timeout")
	}
} // eof main

// generator
func wordsGenerator(data []string) <-chan string {
  outChan := make(chan string)
  go func(c chan string) {
    defer close(c)
    for _, line := range data {
      words := strings.Split(line, " ")
      for _, word := range words {
        c<- word
      }
    }
  }(outChan)
  return outChan
}
