package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	// TODO: Routing here

	if err := e.Start(":8080"); err != nil {
		fmt.Println(err.Error())
		return
	}

	// Setup graceful shutdown with timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c

	// Echo shutdown
	if err := e.Shutdown(ctx); err != nil {
		fmt.Println(err.Error())
		return
	}

}
