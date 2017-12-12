package main

import "fmt"
/*
#include <stdio.h>
#include "open62541.c"

UA_Client* create_client(int port){
    UA_Client *client = UA_Client_new(UA_ClientConfig_standard);
    UA_StatusCode retval = UA_Client_connect(client, "opc.tcp://localhost:16664");
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

int main_loop(void) {
    UA_Client *client = UA_Client_new(UA_ClientConfig_standard);
    UA_StatusCode retval = UA_Client_connect(client, "opc.tcp://localhost:16664");
    if(retval != UA_STATUSCODE_GOOD) {
        UA_Client_delete(client);
        return (int)retval;
    }

    UA_Int32 value = 0;
    printf("\nReading the value of node (1, \"the.answer\"):\n");
    UA_Variant *val = UA_Variant_new();
    retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(1, "the.answer"), val);
    if(retval == UA_STATUSCODE_GOOD && UA_Variant_isScalar(val) &&
       val->type == &UA_TYPES[UA_TYPES_INT32]) {
            value = *(UA_Int32*)val->data;
            printf("the value is: %i\n", value);
    }
    UA_Variant_delete(val);
    UA_Client_delete(client); 
    return UA_STATUSCODE_GOOD;
}

// int main(void){
//     UA_Client *client = create_client(16664);
//     int value = read_answer(client);
//     printf("\nthe answer is %d\n",value);
//     UA_Client_delete(client);
//     return UA_STATUSCODE_GOOD;
// }
*/
import "C"

func main(){
	var client *C.UA_Client;
	client = C.create_client(16664);
	value := C.read_answer(client);
	fmt.Printf("\nthe answer is %d\n",value)
	C.UA_Client_delete(client);
}