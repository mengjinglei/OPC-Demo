package main

import (
	"fmt"
	"strings"
)

/*
#include <stdio.h>
#include "../open62541.c"

typedef struct{
    int Namespace;
    int NodeId;
    char DisplayName[512];
    char BrowserName[512];
}Node;

UA_Client* create_client(char url[512]){
    UA_Client *client = UA_Client_new(UA_ClientConfig_standard);
    UA_StatusCode retval = UA_Client_connect(client, url);
    if(retval != UA_STATUSCODE_GOOD) {
        UA_Client_delete(client);
        return NULL;
    }
    return client;
}

int read_answer(UA_Client* client){
    UA_Int32 value = 0;
    printf("\nReading the value of node (1, \"the.answer\"):\n");
    UA_Variant *val = UA_Variant_new();
    UA_StatusCode retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(1, "the.answer"), val);
    if(retval == UA_STATUSCODE_GOOD && UA_Variant_isScalar(val) &&
       val->type == &UA_TYPES[UA_TYPES_INT32]) {
            value = *(UA_Int32*)val->data;
            printf("the value is: %i\n", value);
            return value;
    }
    UA_Variant_delete(val);
    return 0;
}

int read_single_value(UA_Client* client, int spaceIdx, char indentifier[512]){
    UA_Int32 value = 0;
    printf("\nReading the value of node (%d, \"%s\"):\n",spaceIdx,indentifier);
    UA_Variant *val = UA_Variant_new();
    UA_StatusCode retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(spaceIdx, indentifier), val);
    if(retval == UA_STATUSCODE_GOOD && UA_Variant_isScalar(val) &&
       val->type == &UA_TYPES[UA_TYPES_INT32]) {
            value = *(UA_Int32*)val->data;
            printf("the value is: %i\n", value);
            return value;
    }
    UA_Variant_delete(val);
    return 0;
}

int get_server_state(UA_Client* client){

    UA_Variant value;
    UA_Variant_init(&value);

    const UA_NodeId nodeId =
        UA_NODEID_NUMERIC(0, UA_NS0ID_SERVER_SERVERSTATUS_STATE);

    UA_StatusCode retval = UA_Client_readValueAttribute(client, nodeId, &value);
    if(retval == UA_STATUSCODE_GOOD &&
    UA_Variant_isScalar(&value)) {
        UA_ServerState server_state = *(UA_ServerState*)value.data;
        printf("server state is: %d\n", (int)server_state);
        return server_state;
    }

    UA_Variant_deleteMembers(&value);
    return 1;
}

Node* browser_node(UA_Client* client, int spaceIdx){
     printf("Browsing nodes in objects folder:\n");
     UA_BrowseRequest bReq;
     UA_BrowseRequest_init(&bReq);
     bReq.requestedMaxReferencesPerNode = 0;
     bReq.nodesToBrowse = UA_BrowseDescription_new();
     bReq.nodesToBrowseSize = 1;
     bReq.nodesToBrowse[0].nodeId = UA_NODEID_NUMERIC(spaceIdx, UA_NS0ID_OBJECTSFOLDER);
     bReq.nodesToBrowse[0].resultMask = UA_BROWSERESULTMASK_ALL;
     UA_BrowseResponse bResp = UA_Client_Service_browse(client, bReq);
     printf("%-9s %-16s %-16s %-16s\n", "NAMESPACE", "NODEID", "BROWSE NAME", "DISPLAY NAME");
     for (size_t i = 0; i < bResp.resultsSize; ++i) {
         for (size_t j = 0; j < bResp.results[i].referencesSize; ++j) {
             UA_ReferenceDescription *ref = &(bResp.results[i].references[j]);
             if(ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_NUMERIC) {
                 printf("%-9d %-16d %-16.*s %-16.*s\n", ref->nodeId.nodeId.namespaceIndex,
                        ref->nodeId.nodeId.identifier.numeric, (int)ref->browseName.name.length,
                        ref->browseName.name.data, (int)ref->displayName.text.length,
                        ref->displayName.text.data);
             } else if(ref->nodeId.nodeId.identifierType == UA_NODEIDTYPE_STRING) {
                 printf("%-9d %-16.*s %-16.*s %-16.*s\n", ref->nodeId.nodeId.namespaceIndex,
                        (int)ref->nodeId.nodeId.identifier.string.length,
                        ref->nodeId.nodeId.identifier.string.data,
                        (int)ref->browseName.name.length, ref->browseName.name.data,
                        (int)ref->displayName.text.length, ref->displayName.text.data);
             }
         }
     }
     UA_BrowseRequest_deleteMembers(&bReq);
     UA_BrowseResponse_deleteMembers(&bResp);
     return NULL;
}
*/
import "C"

type OPCClient struct {
	client *C.UA_Client
	url    string
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

func NewClient(host string, port int) *OPCClient {
	if !strings.HasPrefix(host, "opc.tcp://") {
		host = "opc.tcp://" + host
	}
	url := fmt.Sprintf("%s:%d", host, port)

	opcClient := C.create_client(C.CString(url))
	client := &OPCClient{
		client: opcClient,
		url:    url,
	}
	return client
}

func (c *OPCClient) ReadSingleValue(spaceIdx int, indentifier string) interface{} {
	value := C.read_single_value(c.client, C.int(spaceIdx), C.CString(indentifier))
	return value
}

func (c *OPCClient) BrowserNode(nodeId int) (ret []Node) {
	ret = make([]Node, 0)
	C.browser_node(c.client, C.int(nodeId))
	return
}

// typedef struct {
//     UA_DateTime startTime;
//     UA_DateTime currentTime;
//     UA_ServerState state;
//     UA_BuildInfo buildInfo;
//     UA_UInt32 secondsTillShutdown;
//     UA_LocalizedText shutdownReason;
// } UA_ServerStatusDataType;

type ServerInfo struct {
	startTime           int64
	currentTime         int64
	state               string
	secondsTillShutdown int64
	shutdownReason      string
}

func (c *OPCClient) ServerInfo(nodeId int) ServerInfo {

	return ServerInfo{}
}

func (c *OPCClient) Ping() error {
	ret := C.get_server_state(c.client)
	if ret != 0 {
		err := fmt.Errorf("no endpoint found")
		return err
	}
	return nil
}
