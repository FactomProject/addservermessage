package main

import (
	"bufio"
	"encoding/hex"
	"errors"
	//"fmt"
	"os"
	"strings"
)

// Reads the private key from the file
func GetPrivateKey() (*[64]byte, error) {
	var key string

	file, err := os.Open("privatekey.txt")
	if err != nil {
		return nil, err
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