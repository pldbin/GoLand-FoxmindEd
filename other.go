//  FoxmindED курс з 23-10-2022

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//  ПЕРШЕ завдання  23-10-2022  -  31-10-2022

func main1() {
	startTime := time.Now().Unix() // Отримання часу в даний момент часу
	fmt.Printf("Hello World\n")
	sigs := make(chan os.Signal, 1) // Створення каналу з об'ємом буферу 1

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Підключення каналу за допомогою функції Notify пакету "os/signal"

	go func() { // Створення горутини як анонімної функції
		<-sigs // Очікування отримання сигналу(об'єкт os.Signal) від каналу sigs

		fmt.Printf("Stopped by the user after %d seconds\n", time.Now().Unix()-startTime)
		// Вираховування часу роботи програми з моменту отримання значення  startTime

		os.Exit(0) // Завершення роботи
	}()

	time.Sleep(10 * time.Second) // Очікування 10 секунд
	fmt.Printf("Goodbye world\n")
	os.Exit(1) // Завершення роботи програми з кодом виходу(помилки) 1
}

func main11() {
	start := time.Now()
	fmt.Println("Hello world")

	timer1 := time.NewTimer(10 * time.Second)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		for {
			signal := <-sigs
			fmt.Println()
			if signal == syscall.SIGINT {
				fmt.Println("Got CTRL+C signal")
				elapsed := time.Since(start)
				fmt.Printf("Stopped by the user after %.2f seconds\n", elapsed.Seconds())
				fmt.Println("Closing.")
				os.Exit(0)
			}
		}
	}()

	<-timer1.C
	fmt.Println("Goodbye world")
}

//  ДРУГЕ завдання  31-10-2022  -  __-__-____

func main2() {

}
