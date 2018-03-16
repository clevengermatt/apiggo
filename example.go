package apiggo

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
