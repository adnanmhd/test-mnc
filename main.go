package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"test-mnc/router"
	"test-mnc/util"
	"time"
)

func main() {
	// enable this code to answer test tahap 1
	//test_tahap_1.SolutionOne()
	//test_tahap_1.SolutionTwo()
	//test_tahap_1.SolutionFour()

	defer func() {
		if panicMessage := recover(); panicMessage != nil {
			fmt.Printf("Panic '%v' captured ", panicMessage)
		}
	}()
	//test1.SolutionFour()
	appConfig := util.InitConfig()

	router := router.SetupRouter(appConfig)

	// Start server
	serverPort := fmt.Sprintf(":%s", appConfig.Port)
	router.Server.Addr = serverPort

	go func() {
		if err := router.Start(serverPort); err != nil {
			router.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := router.Shutdown(ctx); err != nil {
		router.Logger.Fatal(err)
	}
}
