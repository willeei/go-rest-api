package main

import (
	"fmt"
	"github.com/williamsbarriquero/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando servidor Rest com Go!")
	routes.HandleRequest()
}
