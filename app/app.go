package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação cli pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "CLI App"
	app.Usage = "Busca IPs e nomes de servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{{
		Name:   "ip",
		Usage:  "Busca IPs de endereços na internet.",
		Flags:  flags,
		Action: fetchIps,
	},
		{
			Name:   "servers",
			Usage:  "Busca o nome dos servidores",
			Flags:  flags,
			Action: fetchServers,
		}}

	return app
}

func fetchIps(c *cli.Context) error {
	host := c.String("host") // retorna o valor após a flag host

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}

func fetchServers(c *cli.Context) error {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
	return nil
}
