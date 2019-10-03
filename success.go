package tools

import (
	"github.com/aws/aws-lambda-go/events"
)

func Success(Body map[string]interface{}, StatusCode ...int) (events.APIGatewayProxyResponse, error) {
	statusCode := 200;
	if len(StatusCode) > 0 && StatusCode[0] != statusCode {
		statusCode = StatusCode[0]
	}
	return BuildResponse(statusCode, Body)
}
