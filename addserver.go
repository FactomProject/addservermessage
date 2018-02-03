package main

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/FactomProject/cli"
	ed "github.com/FactomProject/ed25519"
	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/primitives"
)

// Number of 0x88 bytes needed to match
var proofOfWorkLength int = 3

// Signiture on sent messages
var sigRequired bool = true

var Host string

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
		message(args[1:], []byte{0x00}, false, true)
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
		message(args[1:], []byte{0x01}, false, true)
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
		message(args[1:], []byte{0x00}, true, true)
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
		message(args[1:], []byte{0x01}, true, true)
	}
	return cmd
}()

/********************************
 *        CLI Functions         *
 ********************************/

type messageRequest struct {
	Message string `json:"message"`
}

func message(args []string, serverType []byte, send bool, add bool) {
	chainID := args[0]
	var priv *[64]byte
	var err error
	if len(args) > 1 {
		h, err := hex.DecodeString(args[1])
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}

		var ret [64]byte
		copy(ret[:], h[:])
		priv = &ret
	} else {
		priv, err = GetPrivateKey()
		if err != nil {
			fmt.Println("Error: " + err.Error())
			return
		}
	}

	identityChainIDPrefix := "888888"
	if strings.Compare(chainID[:proofOfWorkLength*2], identityChainIDPrefix[:proofOfWorkLength*2]) != 0 {
		//fmt.Println("Error: Invalid identity chain id prefix")
		//return
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
	if add {
		buf.Write([]byte{constants.ADDSERVER_MSG})
	} else {
		buf.Write([]byte{constants.REMOVESERVER_MSG})
	}

	// Timestamp
	t := primitives.NewTimestampNow()
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

	pub := ed.GetPublicKey(priv)
	sig := ed.Sign(priv, upToSig)

	is := ed.VerifyCanonical(ed.GetPublicKey(priv), upToSig, sig)
	if !is {
		fmt.Println("Message did not verify, try again")
		return
	}

	newBuf := new(bytes.Buffer)
	newBuf.Write(upToSig)
	newBuf.Write((*pub)[:])
	newBuf.Write((*sig)[:])
	message := newBuf.Bytes()

	withSig := hex.EncodeToString(message[:])
	paramWS := messageRequest{Message: withSig}
	curlWithSig := toCurl(withSig)
	paramNS := messageRequest{Message: noSig}

	if send == false {
		//PrintHeader("Curl command without Signiture")
		//fmt.Println(curlNoSig)
		PrintHeader("Send addserver message") //with Signiture
		fmt.Println(curlWithSig)
		//fmt.Println()
	} else {
		var resp *JSON2Response
		var err error
		if sigRequired {
			req := NewJSON2Request("send-raw-message", 0, paramWS)
			resp, err = factomdRequest(req)
			//resp, err = factom.SendRawMsg(withSig)
		} else {
			req := NewJSON2Request("send-raw-message", 0, paramNS)
			resp, err = factomdRequest(req)
			//resp, err = factom.SendRawMsg(noSig)
		}
		if err != nil {
			fmt.Println("Message not send, Error: " + err.Error())
			return
		}
		fmt.Println(string(resp.Result))
	}

}

type JSON2Request struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Params  json.RawMessage `json:"params,omitempty"`
	Method  string          `json:"method,omitempty"`
}

func NewJSON2Request(method string, id, params interface{}) *JSON2Request {
	j := new(JSON2Request)
	j.JSONRPC = "2.0"
	j.ID = id
	if b, err := json.Marshal(params); err == nil {
		j.Params = b
	}
	j.Method = method
	return j
}

func factomdRequest(req *JSON2Request) (*JSON2Response, error) {
	j, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(
		fmt.Sprintf("http://%s/v2", Host),
		"application/json",
		bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	r := NewJSON2Response()
	if err := json.Unmarshal(body, r); err != nil {
		return nil, err
	}

	return r, nil
}

type JSON2Response struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      interface{}     `json:"id"`
	Error   string          `json:"error,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func NewJSON2Response() *JSON2Response {
	j := new(JSON2Response)
	j.JSONRPC = "2.0"
	return j
}

func toCurl(str string) string {
	mes := "curl -X POST --data '{\"jsonrpc\": \"2.0\", \"id\": 0, \"params\": {\"message\":\"" + str + "\"}, \"method\": \"send-raw-message\"}' -H 'content-type:text/plain;' http://" + Host + "/v2"
	return mes
}
