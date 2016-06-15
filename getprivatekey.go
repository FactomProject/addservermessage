package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	"os"
	"strings"
)

// Reads the private key from the file
func GetPrivateKey() (*[64]byte, error) {
	var key string

	file, err := os.Open("privatekey.txt")
	if err != nil {
		file, err = os.Create("privatekey.txt")
		file.WriteString("************** Private Key to sign addserver message ***********\n0000000000000000000000000000000000000000000000000000000000000000\n*********** Insert your own provate key ************************")
		file.Close()
		return nil, errors.New("Privatekey.txt does not exist, file created but needs a private key to be inserted")
		if err != nil {
			return nil, err
		}
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := scanner.Text()
		if strings.HasPrefix(str, "*") {

		} else {
			key = scanner.Text()
			key = strings.TrimRight(key, "\n")
			break
		}
	}
	if len(key) != 64 {
		return nil, errors.New("Private key of invalid length")
	} else if strings.Compare(key, "0000000000000000000000000000000000000000000000000000000000000000") == 0 {
		return nil, errors.New("Private key is all 0s. Please replace the key with in 'privatekey.txt'.")
	} else {
		h, err := hex.DecodeString(key)
		if err != nil {
			return nil, err
		}

		var ret [64]byte
		copy(ret[:], h[:])
		return &ret, nil
	}
}
