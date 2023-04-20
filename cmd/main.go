package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"refactoring/cmd/config"
	"refactoring/container"
	"refactoring/router"
)

func main() {
	if err := mainNotExit(); err != nil {
		log.Fatalf("fatal error: %s", err)
	}
}

func mainNotExit() error {
	log.Println("Starting config initialization...")
	confFlag := flag.String("conf", "", "config file yaml")
	flag.Parse()

	confString := *confFlag
	if confString == "" {
		return fmt.Errorf(" 'conf' flag is require")
	}

	config, err := config.Parse(confString)
	if err != nil {
		return err
	}

	log.Println("Starting container initialization...")
	container.InitContainer(config)

	log.Println("Starting router initialization...")

	_ = context.Background()

	router, err := router.MakeRouter()
	if err != nil {
		return err
	}

	log.Println("Listen and serve..")

	return http.ListenAndServe(
		config.AppPort,
		router,
	)
}
