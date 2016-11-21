Add Server Message Generator
======

[![Build Status](https://travis-ci.org/FactomProject/addservermessage.svg?branch=develop)](https://travis-ci.org/FactomProject/addservermessage)

This program will assist in sending an addserver message to the network. The private key to sign the messages is currently hardcoded in the privatekey.txt file. If no file is present, one will be created with a zerohash as the private key, you must replace this value with the correct key.
## Before running
The privatekey.txt must be filled in with the correct privatekey. The correct key might be found in factomd.conf under "LocalServerPrivKey". Replace the 0's with the correct key
## To run
The program has two functions:
* send
  * Sends the messge
* show
  * Prints out the curl commands necessary to send the message

### Send
```
addservermessage send fed|audit CHAINID
```

### Show
```
addservermessage show fed|audit CHAINID
```
