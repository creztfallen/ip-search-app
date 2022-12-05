package main

import (
	"ip-search/app"
	"log"
	"os"
)

func main() {
	application := app.Gerar()
	if err := application.Run(os.Args); err != nil { // permite que os comandos da linha de comando sejam interpretados e usados
		log.Fatal(err) // Fatal() para a execução do programa.
	}
}
