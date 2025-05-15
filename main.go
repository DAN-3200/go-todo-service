package main

import (
	core "app/internal/infrastructure"
)

func main() {
	println("Executando o servidor... \n")
	core.RunServerAPI()
}
