package app

import (
	"log"
	"github.com/urfave/cli"
	"net"
	"fmt"
)

//Retorna app de linha de comanado pronta para ser executada
func Generate() *cli.App{
	app := cli.NewApp()
	app.Name = "Aplicação de Linha de Comando"
	app.Usage = "Busca por IP e nomes de Servidores na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "dev.book.com.br",
		},
	}
	app.Commands = []cli.Command{
		{
			
			Name : "ip",
			Usage: "Busca por ip de endereços na internet",
			Flags: flags,
			Action: searchIP,
		},
		{
			Name: "server",
			Usage: "Busca por Servidores na Internet",
			Flags: flags,
			Action : searchServer,
		},
	}
	return app
}

func searchIP(c *cli.Context){
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil{
		log.Fatal(erro)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServer(c *cli.Context){
	host := c.String("host")

	servers, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, server := range servers{
		fmt.Println(server.Host)
	}
}

