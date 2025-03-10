package main

import (
	"github.com/Brumidori/githubAction-testes-automatizados/database"
	"github.com/Brumidori/githubAction-testes-automatizados/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
