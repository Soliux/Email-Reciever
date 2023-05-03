package main

import (
	"email-reciever/email"
	"email-reciever/web"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func shutdown() {
	fmt.Println("Shutting down servers...")
	email.ShutdownServer()
	web.ShutdownServer()
	os.Exit(0)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		email.StartServer()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		web.StartServer()
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	shutdown()

	wg.Wait()
}
