package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrBadRequest = errors.New("bad request")
	ErrNoLocation = errors.New("no location provided")
)

// Request is what a user sends to the service to get the current time
// in a particular location.
type Request struct {

	// Location is the name of a time zone location from the
	// IANA Time Zone database. E.g., "America/New_York"
	Location string `json:"location"`

	// Format is an optional Go date format to specify how
	// the time string in the response should be formatted.
	Format string `json:"format"`
}

const defaultFormat = time.UnixDate

type Response struct {
	Location string
	Time     string
}

// All main needs to do is tell AWS Lambda what function will handle requests.
func main() {
	lambda.Start(Handler)
}

// Handler is invoked by the AWS Lambda API Gateway Proxy.
// This handler gets the current time in a requested time zone,
// after unpacking and validating the request.
func Handler(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s", r.RequestContext.RequestID)
	if len(r.Body) == 0 {
		return badRequest("body is empty")
	}
	var req Request
	if err := json.Unmarshal([]byte(r.Body), &req); err != nil {
		return badRequest("error unmarshaling request: " + err.Error())
	}
	if len(req.Location) == 0 {
		return badRequest("no location provided")
	}
	if len(req.Format) == 0 {
		req.Format = defaultFormat
	}
	t, err := getTimeInLocation(time.Now(), req.Location)
	if err != nil {
		return badRequest(err.Error())
	}
	resp, err := genResponsePayload(t, req.Format)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return resp, nil
}

func getTimeInLocation(t time.Time, locstr string) (time.Time, error) {
	loc, err := time.LoadLocation(locstr)
	if err != nil {
		return time.Time{}, err
	}
	return t.In(loc), nil
}

func genResponsePayload(t time.Time, format string) (events.APIGatewayProxyResponse, error) {
	r := Response{
		Location: t.Location().String(),
		Time:     t.Format(format),
	}
	rbytes, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(rbytes),
		StatusCode: http.StatusOK,
	}, nil
}

func badRequest(logmsg string) (events.APIGatewayProxyResponse, error) {
	log.Printf("%v: %s", ErrBadRequest, logmsg)
	return events.APIGatewayProxyResponse{
		Body:       http.StatusText(http.StatusBadRequest),
		StatusCode: http.StatusBadRequest,
	}, nil
}
