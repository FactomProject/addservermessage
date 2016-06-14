Add Server Message Generator
======
This program will assist in sending an addserver message to the network. The private key to sign the messages is currently hardcoded in the privatekey.txt file. Currently does not sign the sent messages, but if using 'show', it will show the signed message curl command.

### To get commands
```
addserver help
```

## To run
The program has two functions:
* send
  * Sends the messge
* show
  * Prints out the curl commands necessary to send the message

### Send
* Will currently send unsigned message

```
addserver send fed|audit CHAINID
```

### Show
```
addserver show fed|audit CHAINID
```
