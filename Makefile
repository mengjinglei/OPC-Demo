build:
		go build myClient.go

client:
		go build myClient.go

server:
		gcc -std=c99 myServer.c -o myServer

answer:
		gcc -std=c99 answerServer.c open62541.c -o answerServer
		gcc -std=c99 answerClient.c open62541.c -o answerClient
