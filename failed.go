package golang_serverless_response

import (
	"github.com/aws/aws-lambda-go/events"
)

func Failed(Message string, StatusCode ...int) (events.APIGatewayProxyResponse, error) {
	statusCode := 400;
	if len(StatusCode) > 0 && StatusCode[0] != statusCode {
		statusCode = StatusCode[0]
	}
	return BuildResponse(statusCode, map[string]interface{}{
		"message": Message,
	})
}
