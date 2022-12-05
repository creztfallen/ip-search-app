package main

import (
	"fmt"
	"ip-search/app"
	"log"
	"os"
)

func main() {
	fmt.Println("Ponto de partida")

	application := app.Gerar()
	if err := application.Run(os.Args); err != nil { // permite que os comandos da linha de comando sejam interpretados e usados
		log.Fatal(err) // Fatal() para a execução do programa.
	}
}
