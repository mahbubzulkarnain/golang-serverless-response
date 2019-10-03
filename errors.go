package tools

import (
	"github.com/aws/aws-lambda-go/events"
)

func Errors(err error, StatusCode ...int) (events.APIGatewayProxyResponse, error) {
	statusCode := 500;
	if len(StatusCode) > 0 && StatusCode[0] != statusCode {
		statusCode = StatusCode[0]
	}
	return BuildResponse(statusCode, map[string]interface{}{
		"message": err.Error(),
	})
}
