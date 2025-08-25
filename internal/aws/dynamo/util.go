package dynamo

import (
	"log"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func GetKey(t ItemDynamo) map[string]types.AttributeValue {
	service, err := attributevalue.Marshal(t.Service)
	if err != nil {
		log.Fatal(err)
	}

	statusCode, err := attributevalue.Marshal(t.StatusCode)
	if err != nil {
		log.Fatal(err)
	}

	return map[string]types.AttributeValue{"service": service, "statusCode": statusCode}

}
