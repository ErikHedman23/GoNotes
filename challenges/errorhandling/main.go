package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

// Config holds configuration
type Config struct {
	//configuration fields go here (redacted)
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can't open configuration file")
	}

	defer file.Close()

	cfg := &Config{}

	return cfg, nil
}

func setupLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 06)
	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setupLogging()

	cfg, err := readConfig("./errorhandling/hello.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %+v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
