Add Server Message Generator
======
This program will assist in sending an addserver message to the network. The private key to sign the messages is currently hardcoded in the privatekey.txt file. If no file is present, one will be created with a zerohash as the private key, you must replace this value with the correct key.
## Before running
The privatekey.txt must be filled in with the correct privatekey. The correct key might be found in factomd.conf under "LocalServerPrivKey". Replace the 0's with the correct key
## To run
The program has two functions:
* send
  * Sends the messge
* show
  * Prints out the curl commands necessary to send the message

## Optional Parameters
* Not giving a private key will use the private key file in the active directory. If there is none it will create a file, this allows you to not give a private key parameter.
* -host=HOST: By default it looks at localhost:8088, changing this will send the message to a different host.

## Add Messages
### Send
```
addservermessage -host=HOST send fed|audit CHAINID PRIVATEKEY
```

### Show
```
addservermessage -host=HOST show fed|audit CHAINID PRIVATEKEY
```

## Remove Messages
### Send
```
addservermessage -host=HOST sendR fed|audit CHAINID PRIVATEKEY
```

### Show
```
addservermessage -host=HOST showR fed|audit CHAINID PRIVATEKEY
```
