package main

import (
	// "fmt"
	// "net/http"

	"flag"

	"github.com/odysseymorphey/vkTestRESTAPI/internal/application"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "config file")
	flag.Parse()
	
	a := &application.Application{}

	a.Build(configPath)
}
