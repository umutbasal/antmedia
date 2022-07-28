package main

import (
	"log"
	"os"
	"time"

	apiclient "github.com/umutbasal/antmedia/client"
	"github.com/umutbasal/antmedia/client/broadcast_rest_service"

	"github.com/go-openapi/strfmt"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	// create the API client, with the transport
	client := apiclient.NewHTTPClientWithConfig(strfmt.Default, &apiclient.TransportConfig{
		Host:     os.Getenv("ANTMEDIA_HOST"),
		BasePath: os.Getenv("ANTMEDIA_BASE"),
		Schemes:  []string{"http"},
	})

	// to override the host for the default client
	// apiclient.Default.SetTransport(transport)

	// make the request to get all items
	params := broadcast_rest_service.NewGetBroadcastListParamsWithTimeout(time.Second * 10)
	params.Offset = 0
	params.Size = 10

	resp, err := client.BroadcastRestService.GetBroadcastList(params)
	if err != nil {
		log.Fatal(err)
	}
	broadcasts := resp.Payload
	for _, broadcast := range broadcasts {
		log.Printf("%s", broadcast.Name)
	}
}
