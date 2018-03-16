package apiggo

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/aws/aws-lambda-go/events"
)

// Handler turns an APIGatewayProxyRequest into a standard http.Request
func Handler(handler http.Handler, host string, proxyRequest events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//Check if the APIGatewayProxyRequest.Body is a base64 encoded string
	// If it is, decode it
	var body string
	decodedBody, err := base64.StdEncoding.DecodeString(proxyRequest.Body)
	if err != nil {
		body = proxyRequest.Body
	} else {
		body = string(decodedBody)
	}

	// Create a new *http.Request and *httptest.ResponseRecorder
	r := httptest.NewRequest(proxyRequest.HTTPMethod, proxyRequest.Path, strings.NewReader(body))
	w := httptest.NewRecorder()

	// Set the host to whatever your APIG base path is
	r.Host = host

	// Add the APIGatewayProxyRequest headers to the *http.Request
	for key, value := range proxyRequest.Headers {
		r.Header.Add(key, value)
	}

	// Create a x-url-formencoded string from QueryStringParameters
	var queryString string
	for key, value := range proxyRequest.QueryStringParameters {
		queryString = fmt.Sprintf("%s%s=%s&", queryString, key, value)
	}
	queryString = strings.TrimSuffix(queryString, "&")

	// Create a new url.URL
	rURL := url.URL{
		Scheme:     "https",
		Opaque:     "",
		User:       nil,
		Host:       host,
		Path:       proxyRequest.Path,
		RawPath:    proxyRequest.Path,
		ForceQuery: false,
		RawQuery:   queryString,
		Fragment:   "",
	}

	r.URL = &rURL

	handler.ServeHTTP(w, r)

	// Add the response headers to a map that can be used in the APIGatewayProxyResponse
	// It is important to do this AFTER receiving requests
	respHeaders := map[string]string{}

	for key, value := range w.Header() {
		respHeaders[key] = strings.Join(value, "")
	}

	proxyResp := events.APIGatewayProxyResponse{
		Body:       w.Body.String(),
		Headers:    respHeaders,
		StatusCode: w.Code,
	}
	return proxyResp, nil
}
