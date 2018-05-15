

# to encrypt the key
# echo -n "key here" | gpg -ca > enckey.txt

fullnode=$(head -n 1 fullnode.txt)
identity=$(head -n 1 identity.txt)

echo "using:"
echo $fullnode
echo $identity

masterkey=$(gpg -o- -q enckey.txt)
#echo $masterkey
echo RELOADAGENT | gpg-connect-agent

#change to send or sendR to update the network
addservermessage -host=$fullnode showR audit $identity $masterkey
