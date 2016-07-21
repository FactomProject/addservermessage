package main

import (
	"flag"
	"github.com/FactomProject/cli"
)

func main() {
	HostVal := flag.String("host", "localhost:8088", "Changing the message location to send to")

	flag.Parse()
	args := flag.Args()
	Host = *HostVal
	c := cli.New()

	c.Handle("help", Help)
	c.Handle("show", ShowAddServer)
	c.Handle("send", SendAddServer)
	c.Handle("showR", ShowRemoveServer)
	c.Handle("sendR", SendRemoveServer)

	c.HandleDefault(Help)
	c.Execute(args)
}
