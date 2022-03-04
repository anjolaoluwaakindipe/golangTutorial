package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/anjolaoluwaakindipe/golangtutorial/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)
	goodbyeHandler := handlers.NewGoodbye(l)
	productHandler := handlers.NewProducts(l)

	sm := http.NewServeMux()

	sm.Handle("/", helloHandler)
	sm.Handle("/goodbye", goodbyeHandler)
	sm.Handle("/products", productHandler)

	s := &http.Server{
		Addr:        "localhost:9090",
		Handler:     sm,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 *
			time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal()
		}
	}()

	sigChan := make(chan os.Signal)

	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan

	l.Println("Recieved terminate, graceful shutdowm", sig)
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)

}
