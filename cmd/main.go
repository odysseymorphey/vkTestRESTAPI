package main

import (
	"fmt"
	"net/http"

	"github.com/odysseymorphey/vkTestRESTAPI/internal/application"
)

func main() {
	a := &application.Application{}

	a.Build("/nahui")
	a.GetCFG()
	
	a.Run()
}
