package opcclient

/*
#include <stdio.h>
#include "../open62541.c"
*/
import "C"

type OPCClient struct {
	client *C.UA_Client
}

func (c *OPCClient) ReadObject(objectName string) interface{} {

	return 0
}

func (c *OPCClient) Close() {

	return
}

type Node struct {
	Namespace   int
	NodeId      int
	DisplayName string
	BrowserName string
}
