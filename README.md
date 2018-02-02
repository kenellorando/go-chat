# go-chat
Currently two separate client and server programs which demonstrate TCP data transfer over a network. 

To change the destination address, change the `const` "`LADDR`" value to a valid network address and port. For sending to a server on a foreign machine, modification to server's firewall may be necessary.

## Starting
The server needs to be started first.
```
server/
$ go run main.go
client/
$ go run main.go
```
