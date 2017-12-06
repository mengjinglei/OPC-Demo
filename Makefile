build:
		go build myClient.go

client:
		go build myClient.go

server:
		gcc -std=c99 myServer.c -o myServer