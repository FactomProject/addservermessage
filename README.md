Add Server Message Generator
======
This program will assist in sending an addserver message to the network. The private key to sign the messages is currently hardcoded in the privatekey.txt file. Currently does not sign the sent messages, but if using 'show', it will show the signed message curl command.
## Before running
The privatekey.txt must be filled in with the correct privatekey. It might be found in factomd.conf under "LocalServerPrivKey". Replace the 0's with the correct key
## To run
The program has two functions:
* send
  * Sends the messge
* show
  * Prints out the curl commands necessary to send the message

### Send
* Will currently send unsigned message

```
addservermessage send fed|audit CHAINID
```

### Show
```
addservermessage show fed|audit CHAINID
```
