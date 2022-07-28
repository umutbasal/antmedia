package main

import (
	antmedia "antmedia/client"
	"fmt"
	"log"
)

func main() {
	resp, err := antmedia.Default.BroadcastRestService.GetBroadcastList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", resp.Payload)
}
