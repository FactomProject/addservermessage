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

```
addservermessage -host=HOST:8088 send|sendR|show f|a CHAINID PRIVATEKEY
```

HOST = I.P. address of node to which the promotion message is to be sent (can be any node)

NOTE: promotion message can be sent to any node, but removal message must be sent to a federated server

send actually sends a promotion message  
sendR actually sends a removal message  
show just prints the curl commands that would do the promotion so that you can take it elsewhere and send it

f = promote to federated server  
a = promote to audit server 

NOTE: a federated server can be 'promoted' to an audit server


The file:

```~/factom/factom-ansible/identities/identities.yml```

is a number of entries like this:

  ```default:  
      ip: "0.5"  
      IdentityChainID: 38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9
      ServerPrivKey: 4c38c65dc5cdad68f13b74789d3ffb1f3d63e335610868c9b90766553448d26d
      ServerPublicKey: cc1985cdfae4e32b5a454dfda8ce5e1361558482684f3367649c3ad852c8e31a
  i10:
       ip: "0.10"
       IdentityChainID: 888888367795422bb2b15bae1af83396a94efa1cecab8cd171197eabd4b4bf9b
       ServerPrivKey: 5319e60d236893ed32e0a863b2877521d2d59ce7cf644deca93bf5a065af87cd
       ServerPublicKey: ad6f634018389a29da51586ef69f747bb4608d29e19c40242f1b1c3bd4cede16```

CHAINID = the IdentityChainID of the node to be promoted as listed in this file  
PRIVATEKEY = the ServerPrivKey of the node to which the message is sent  

e.g.

To tell default to add i10 as a federated server:   
```addservermessage -host=10.41.0.5:8088 send f 888888367795422bb2b15bae1af83396a94efa1cecab8cd171197eabd4b4bf9b 4c38c65dc5cdad68f13b74789d3ffb1f3d63e335610868c9b90766553448d26d```

To tell default to remove itself as an audit server:   
```addservermessage -host=10.41.0.5:8088 sendR a 38bab1455b7bd7e5efd15c53c777c79d0c988e9210f1da49a99d95b3a6417be9 4c38c65dc5cdad68f13b74789d3ffb1f3d63e335610868c9b90766553448d26d```

NOTE: when servers are reconfigured via addservermessage, the server count is actually changed to the new value. 
This means that if a federated server is 'promoted' to an audit server, no audit server will automatically be promoted to federated server to take its place.
It also means that a network adjusted via addservermessage to have less federated servers than its original quorum will not stall.


