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
	// Implement the shutdown logic for both servers here
	// For instance, you can close all connections and clean up resources
	fmt.Println("Shutting down servers...")
	email.ShutdownServer()
	web.ShutdownServer()
}

func main() {
	var wg sync.WaitGroup

	// Start the email server
	wg.Add(1)
	go func() {
		defer wg.Done()
		email.StartServer()
	}()

	// Start the web server
	wg.Add(1)
	go func() {
		defer wg.Done()
		web.StartServer()
	}()

	// Listen for an interrupt signal (e.g., when you press Ctrl+C)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until an interrupt signal is received
	<-c
	shutdown()

	// Wait for both servers to finish
	wg.Wait()
}
