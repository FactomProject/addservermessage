package main

import (
	"flag"
	"github.com/FactomProject/cli"
)

func main() {
	flag.Parse()
	args := flag.Args()

	c := cli.New()

	c.Handle("help", Help)
	c.Handle("show", ShowAddServer)
	c.Handle("send", SendAddServer)

	c.HandleDefault(Help)
	c.Execute(args)
}
