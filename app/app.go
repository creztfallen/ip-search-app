package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação cli pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	return app
}
