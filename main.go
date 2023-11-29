package main

import (
	"fmt"
	"github.com/williamsbarriquero/go-rest-api/models"
	"github.com/williamsbarriquero/go-rest-api/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "P1", Historia: "h1"},
		{Nome: "P2", Historia: "h2"},
	}
	fmt.Println("Iniciando servidor Rest com Go!")
	routes.HandleRequest()
}
