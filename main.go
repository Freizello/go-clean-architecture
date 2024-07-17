package main

import (
	"os"
	"os/signal"
	"syscall"

	"log"

	"github.com/freizello/go-clean-architecture/cmd/cron"
	"github.com/freizello/go-clean-architecture/cmd/grpc"
	"github.com/freizello/go-clean-architecture/cmd/web"
)

var (
	grpcServiceMode bool
	webServiceMode  bool
	cronServiceMode bool
)

const (
	grpcServiceModeParam = "grpc-server"
	cronServiceModeParam = "cron-server"
)

func main() {
	// todo: read command args
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGQUIT)

	var stopFunc func()
	defer func() {
		if stopFunc != nil {
			stopFunc()
		}
	}()

	switch {
	case cronServiceMode:
		go cron.Start()
	case grpcServiceMode:
		go grpc.Start()
	default:
		webServiceMode = true
		go func() {
			stopFunc = web.Start()
		}()
	}

	// will wait until terminate signal or interrupt happened
	for {
		<-c
		log.Println("terminate service...")
		if webServiceMode {
			web.GracefulStop()
		}
		os.Exit(0)
	}
}
