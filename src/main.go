package main

import (
	"loginpage/srv"
	"os"
	"os/signal"
)

func main() {
	srv.Init()

	// shutdown
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, os.Interrupt)
	<-shutdownCh
}
