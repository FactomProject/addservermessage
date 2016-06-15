package main

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/FactomProject/cli"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factom"
	"github.com/FactomProject/factomd/common/interfaces"
)

// Number of 0x88 bytes needed to match
var proofOfWorkLength int = 1

// No signiture on sent messages
var sigRequired bool = true

/********************************
 *          Cli Control         *
 ********************************/
var ShowAddServer = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver show fed|audit CHAINID"
	cmd.description = "Shows a new addserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()

		c.Handle("f", newFed)
		c.Handle("fed", newFed)
		c.Handle("federated", newFed)
		c.Handle("a", newAudit)
		c.Handle("aud", newAudit)
		c.Handle("audit", newAudit)

		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Shows a new server message curl command from the chainID", cmd)
	return cmd
}()

var SendAddServer = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver send fed|audit CHAINID"
	cmd.description = "Sends a new addserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		args = flag.Args()
		c := cli.New()

		c.Handle("f", sendFed)
		c.Handle("fed", sendFed)
		c.Handle("federated", sendFed)
		c.Handle("a", sendAudit)
		c.Handle("aud", sendAudit)
		c.Handle("audit", sendAudit)

		c.HandleDefaultFunc(func(args []string) {
			fmt.Println(cmd.helpMsg)
		})
		c.Execute(args)
	}
	Help.Add("Sends a new server message from the chainID", cmd)
	return cmd
}()

var newFed = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver show fed|f CHAINID"
	cmd.description = "Shows a new federated addserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'addserver show fed CHAINID'")
			return
		}
		message(args[1], []byte{0x00}, false)
	}
	return cmd
}()

var newAudit = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver show audit|a CHAINID"
	cmd.description = "Shows a new audit addserver message curl command from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'addserver show audit CHAINID'")
			return
		}
		message(args[1], []byte{0x01}, false)
	}
	return cmd
}()

var sendFed = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver send fed|f CHAINID"
	cmd.description = "Sends a new federated addserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'addserver send fed CHAINID'")
			return
		}
		message(args[1], []byte{0x00}, true)
	}
	return cmd
}()

var sendAudit = func() *cliCmd {
	cmd := new(cliCmd)
	cmd.helpMsg = "addserver send audit|a CHAINID"
	cmd.description = "Sends a new audit addserver message from the chainID."
	cmd.execFunc = func(args []string) {
		os.Args = args
		flag.Parse()
		if len(args) < 2 {
			fmt.Println("No chainID given, 'addserver send audit CHAINID'")
			return
		}
		message(args[1], []byte{0x01}, true)
	}
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/
// Marshal order
// Byte[0] : 0x21
// Timestamp
// ChainID
// SeverType
// Signiture
func message(chainID string, serverType []byte, send bool) {
	priv, err := GetPrivateKey()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	identityChainIDPrefix := "888888"
	if strings.Compare(chainID[:proofOfWorkLength*2], identityChainIDPrefix[:proofOfWorkLength*2]) != 0 {
		fmt.Println("Error: Invalid identity chain id prefix")
		return
	} else if len(chainID) != 64 {
		fmt.Println("Error: Invalid identity chain id length")
		return
	}

	chain, err := hex.DecodeString(chainID)
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}

	buf := new(bytes.Buffer)

	// Message Type
	buf.Write([]byte{0x15})

	// Timestamp
	t := interfaces.NewTimeStampNow()
	data, err := t.MarshalBinary()
	if err != nil {
		fmt.Println("Error: " + err.Error())
		return
	}
	buf.Write(data)

	// ChainID
	buf.Write(chain)

	// Server type (0 or 1)
	buf.Write(serverType)

	// Signiture
	upToSig := buf.Bytes()
	noSig := hex.EncodeToString(upToSig[:])
	//curlNoSig := toCurl(noSig)

	sig := ed.Sign(priv, upToSig)
	pub := ed.GetPublicKey(priv)

	newBuf := new(bytes.Buffer)
	newBuf.Write(upToSig)
	newBuf.Write(pub[:])
	newBuf.Write(sig[:])
	message := newBuf.Bytes()
	withSig := hex.EncodeToString(message[:])
	curlWithSig := toCurl(withSig)

	if send == false {
		//PrintHeader("Curl command without Signiture")
		//fmt.Println(curlNoSig)
		PrintHeader("Send addserver message") //with Signiture
		fmt.Println(curlWithSig)
		//fmt.Println()
	} else {
		var resp *factom.SendRawMessageResponse
		var err error
		if sigRequired {
			resp, err = factom.SendRawMsg(withSig)
		} else {
			resp, err = factom.SendRawMsg(noSig)
		}
		if err != nil {
			fmt.Println("Message not send, Error: " + err.Error())
		}
		fmt.Println(resp.Message)
	}

}

func toCurl(str string) string {
	mes := "curl -X POST --data '{\"jsonrpc\": \"2.0\", \"id\": 0, \"params\": {\"message\":\"" + str + "\"}, \"method\": \"send-raw-message\"}' -H 'content-type:text/plain;' http://localhost:8088/v2"
	return mes
}
