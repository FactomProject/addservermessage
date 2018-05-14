package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/FactomProject/cli"
)

var ShowRemoveServer = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver showR fed|raudit CHAINID"
	cmd.description = "Shows a new removeserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()

		c.Handle("f", newRemoveFed)
		c.Handle("fed", newRemoveFed)
		c.Handle("federated", newRemoveFed)
		c.Handle("a", newRemoveAudit)
		c.Handle("aud", newRemoveAudit)
		c.Handle("audit", newRemoveAudit)

		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Shows a new server remove message curl command from the chainID", cmd)
	return cmd
}()

var SendRemoveServer = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver sendR fed|audit CHAINID"
	cmd.description = "Sends a new removeserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()

		c.Handle("f", sendRemoveFed)
		c.Handle("fed", sendRemoveFed)
		c.Handle("federated", sendRemoveFed)
		c.Handle("a", sendRemoveAudit)
		c.Handle("aud", sendRemoveAudit)
		c.Handle("audit", sendRemoveAudit)

		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Sends a new server remove message from the chainID", cmd)
	return cmd
}()

var newRemoveFed = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver showR fed|f CHAINID"
	cmd.description = "Shows a new federated removeserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'removeserver showR fed CHAINID'")
			return
		}
		message(args[1:], []byte{0x00}, false, false)
	}
	return cmd
}()

var newRemoveAudit = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver showR audit|a CHAINID"
	cmd.description = "Shows a new audit removeserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'removeserver showR audit CHAINID'")
			return
		}
		message(args[1:], []byte{0x01}, false, false)
	}
	return cmd
}()

var sendRemoveFed = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver sendRfed|f CHAINID"
	cmd.description = "Sends a new federated removeserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'removeserver sendR fed CHAINID'")
			return
		}
		message(args[1:], []byte{0x00}, true, false)
	}
	return cmd
}()

var sendRemoveAudit = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver sendR audit|a CHAINID"
	cmd.description = "Sends a new audit removeserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'removeserver sendRaudit CHAINID'")
			return
		}
		message(args[1:], []byte{0x01}, true, false)
	}
	return cmd
}()
