package __1_writehelloworld

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WriteHelloworld() {
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
