package main

import (
	"fmt"
	"log"
	"os"
	"time"

	apiclient "github.com/umutbasal/antmedia/client"
	"github.com/umutbasal/antmedia/client/broadcast_rest_service"
	models "github.com/umutbasal/antmedia/models"

	"github.com/davecgh/go-spew/spew"
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
	params := broadcast_rest_service.NewCreateBroadcastParamsWithTimeout(time.Second * 10)
	params.Body = &models.Broadcast{
		Name: "test",
	}

	resp, err := client.BroadcastRestService.CreateBroadcast(params)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump(resp)

	// rtmp url
	fmt.Printf("RTMP: %s", resp.Payload.RtmpURL)

}
