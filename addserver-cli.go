package main

import (
	"flag"
	"github.com/FactomProject/cli"

	"fmt"
)

func main() {
	HostVal := flag.String("host", "localhost:8080", "Changing the message location to send to")

	flag.Parse()
	args := flag.Args()
	Host = *HostVal
	fmt.Println(Host)
	c := cli.New()

	c.Handle("help", Help)
	c.Handle("show", ShowAddServer)
	c.Handle("send", SendAddServer)

	c.HandleDefault(Help)
	c.Execute(args)
}
