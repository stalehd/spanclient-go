package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lab5e/spanclient-go/v4"
)

func main() {
	// It's always a good idea to leave authentication tokens out of the source
	// code so we use a command line parameter here.
	token := ""
	collectionID := ""
	deviceID := ""
	flag.StringVar(&token, "token", "", "API token for the Span service")
	flag.StringVar(&collectionID, "collection-id", "", "The collection to query")
	flag.StringVar(&deviceID, "device-id", "", "The device to query")

	flag.Parse()

	if token == "" {
		fmt.Println("Missing token parameter")
		flag.PrintDefaults()
		return
	}
	if collectionID == "" {
		fmt.Println("Needs a collection ID")
		flag.PrintDefaults()
		return
	}

	config := spanclient.NewConfiguration()

	// Set this to true to list the requests and responses in the client. It can
	// be useful if you are wondering what is happening.
	config.Debug = false

	// In the Real World this context should also have a context.WithTimeout
	// call to ensure we time out if there's no response.
	ctx := context.WithValue(context.Background(),
		spanclient.ContextAPIKey,
		spanclient.APIKey{
			Key:    token,
			Prefix: "",
		})

	if deviceID != "" {
		fmt.Println("Will read device data stream from device ", deviceID)
		ds, err := spanclient.NewDeviceDataStream(ctx, config, collectionID, deviceID)
		if err != nil {
			fmt.Println("Error connecting data stream: ", err.Error())
			return
		}
		readDataStream(ds)
		return
	}

	fmt.Println("Will read collection data stream from collection ", collectionID)
	ds, err := spanclient.NewCollectionDataStream(ctx, config, collectionID)
	if err != nil {
		fmt.Println("Error connecting data stream: ", err.Error())
		return
	}
	readDataStream(ds)
}

func readDataStream(ds spanclient.DataStream) {
	defer ds.Close()
	fmt.Println("Waiting for data...")
	for {
		msg, err := ds.Recv()
		if err != nil {
			fmt.Println("Error reading message: ", err.Error())
			return
		}
		fmt.Printf("%+v\n", msg)
	}
}
