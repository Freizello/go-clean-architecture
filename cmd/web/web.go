package web

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

var (
	srv *http.Server
)

func Start() (StopFunc func()) {
	// TODO: init db
	// TODO: init redis
	// TODO: init 3rd party if any

	go func() {
		addr := fmt.Sprintf(":%s", "9090")
		pprofErr := http.ListenAndServe(addr, nil)
		if pprofErr != nil {
			log.Printf("pprof failed started: %+v\n", pprofErr.Error())
		}
	}()

	var handlers http.Handler
	// TODO: init handler if any

	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = "8080"
	}
	srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", envPort),
		Handler: handlers,
	}

	go func() {
		srvErr := srv.ListenAndServe()
		if srvErr != nil && srvErr != http.ErrServerClosed {
			log.Fatalln("failed start web on: ", srv.Addr, srvErr)
		}
	}()

	return func() {
		//TODO: stop func
	}
}

func GracefulStop() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))
	defer cancel()

	log.Println("shuting down web on", srv.Addr)
	err = srv.Shutdown(ctx)
	if err != nil {
		log.Fatalln("failed shutdown server", err)
	}
	log.Println("web gracefully stopped")
	return
}
