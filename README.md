apiggo allows you to easily port your existing Go request handlers to AWS Lambda
functions by eliminating the need to work directly with APIGatewayProxyRequest.

## Usage

The Handler function takes an http.Handler as its first argument and so you can
use any multiplexer that conforms to http.Handler. The example below is using
"github.com/gorilla/mux" The second argument is the host name. You can leave
this blank or put the base path you are using on your AWS API Gateway. Lastly
you pass it the APIGatewayProxyRequest so it can get all of that good data from
it.

```
package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/clevengermatt/apiggo"
	"github.com/gorilla/mux"
)

func main() {

	lambda.Start(func(pr events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		router := mux.NewRouter()
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "Hello, world!")
		})

		return apiggo.Handler(router, "example.com", pr)
	})
}
```

## Built With

* [Go](https://golang.org) - The Go Programming Language
* [AWS Lambda](https://github.com/aws/aws-lambda-go) - Libraries, samples and
  tools to help Go developers develop AWS Lambda functions.

## Authors

* **Matt Clevenger** - _Initial work_ -
  [clevengermatt](https://github.com/clevengermatt)

See also the list of
[contributors](https://github.com/clevengermatt/apiggo/contributors) who
participate in this project.
