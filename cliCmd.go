package main

type cliCmd struct {
	execFunc    func([]string)
	helpMsg     string
	description string
}

func (c *cliCmd) Execute(args []string) {
	c.execFunc(args)
}
