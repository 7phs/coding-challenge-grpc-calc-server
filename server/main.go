package main

import (
	"github.com/7phs/coding-challenge-grpc-calc-server/log"
		"os"
	"os/signal"
	"github.com/7phs/coding-challenge-grpc-calc-server/config"
	)

func main() {
	log.SetLogLevel(log.ALL)
	log.SetPrefix("[CALCSERVICE]:")

	// 1. Read and checking config
	log.Info("starting")
	log.Info("read config")
	conf, err := config.ParseConfig()
	if err != nil {
		log.Error("failed to get configuration parameters: ", err)
		return
	}
	// 2. Configure logger
	log.SetLogLevel(conf.LogLevel())

	// 3. Init server
	log.Info("server initialization")
	server, err := NewServer(conf)
	if err != nil {
		log.Error("failed to create a server: ", err)
		return
	}

	// 4. Run a server
	log.Info("server running on ", conf.Address())
	server.Run()

	// 5. wait for Ctrl+C or server error
	interrupt := make(chan os.Signal, 2)
	signal.Notify(interrupt, os.Interrupt) // CTRL-C
	select {
		case <-interrupt:
			// shut down server
			log.Info("server shutdown")
			server.Shutdown()

		case <-server.WaitError():
			log.Error("error while a server running: ", server.Error())
	}

	log.Info("server waiting for finish")
	server.Wait()

	log.Info("finished")
}
