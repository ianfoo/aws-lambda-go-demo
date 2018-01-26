# aws-lambda-go-demo
This is a demonstration of native AWS Lambda Golang support.

Given a JSON request with a `location` and optional `format` field, it will
return the time in the specified time zone.

`location` is a time zone name from the [IANA Time Zone
Database](https://www.iana.org/time-zones).

`format` is a [Golang time format](https://godoc.org/time#Time.Format). If no
time format is specified, the service will default to the
[`time.UnixDate`](https://godoc.org/time#pkg-constants) time format.

# Credits

This was completed by following the first portion of the [AWS Lambda Golang
introductory blog
post](https://aws.amazon.com/blogs/compute/announcing-go-support-for-aws-lambda/),
with a handler that performed a slightly different funciton than string
concatenation (just because).

# Notes

While the AWS Lambda Go SDK is advanced enough to automatically marshal and
unmarshal to and from custom request and response types (within certain
limits), I opted to use the raw `events.APIGatewayProxyRequest` and
`events.APIGatewayProxyResponse` to allow for greater granularity in the
handler. For example, logging the request ID (as is done in the initial example
in the blog post referenced above) and controlling the status code when errors
occur.

# TODO
* Format errors as JSON.
* Get this set up with continuous deployment as described in the rest of the
  referenced AWS blog post.
