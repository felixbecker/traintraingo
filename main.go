package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello train train go!")
	})

	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	go func() {
		sig := <-gracefulStop
		fmt.Printf("Received signal '%+v' to stop the process gracefully\n", sig)
		os.Exit(0)
	}()

	fmt.Println("Started server on localhost:8080. ctl+c to stop")
	http.ListenAndServe(":8080", nil)

}
