package tools

import (
	. "bytes"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
)

func Stringify(input map[string]interface{}) string {
	var buf Buffer
	resParser, errParser := json.Marshal(input)
	if errParser != nil {
		return errParser.Error()
	}
	json.HTMLEscape(&buf, resParser)
	return buf.String()
}

func BuildResponse(StatusCode int, Body map[string]interface{}) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: StatusCode,
		Headers: map[string]string{
			"Access-Control-Allow-Credentials": "true",
			"Access-Control-Allow-Origin":      "*",
		},
		MultiValueHeaders: nil,
		Body:              Stringify(Body),
		IsBase64Encoded:   false,
	}, nil
}
