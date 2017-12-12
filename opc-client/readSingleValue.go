package opcclient

/*
#include <stdio.h>
#include "../open62541.c"

int read_single_value(UA_Client* client, int spaceIdx, char indentifier[512]){
    UA_Int32 value = 0;
    printf("\nReading the value of node (spaceIdx, \"the.answer\"):\n");
    UA_Variant *val = UA_Variant_new();
    UA_StatusCode retval = UA_Client_readValueAttribute(client, UA_NODEID_STRING(spaceIdx, "the.answer"), val);
    if(retval == UA_STATUSCODE_GOOD && UA_Variant_isScalar(val) &&
       val->type == &UA_TYPES[UA_TYPES_INT32]) {
            value = *(UA_Int32*)val->data;
            printf("the value is: %i\n", value);
            return value;
    }
    UA_Variant_delete(val);
    return 0;
}

*/
import "C"

func (c *OPCClient) ReadSingleValue(spaceIdx int, indentifier, valueType string) interface{} {
	value := C.read_single_value(c.client, spaceIdx, C.CString(indentifier))
	return value
}
