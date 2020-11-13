package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/lab5e/spanclient-go"
)

func main() {
	// It's always a good idea to leave authentication tokens out of the source
	// code so we use a command line parameter here.
	token := ""
	flag.StringVar(&token, "token", "", "API token for the Span service")
	flag.Parse()

	if token == "" {
		fmt.Println("Missing token parameter")
		flag.PrintDefaults()
		return
	}

	config := spanclient.NewConfiguration()

	// Set this to true to list the requests and responses in the client. It can
	// be useful if you are wondering what is happening.
	config.Debug = false

	client := spanclient.NewAPIClient(config)

	// In the Real World this context should also have a context.WithTimeout
	// call to ensure we time out if there's no response.
	ctx := context.WithValue(context.Background(),
		spanclient.ContextAPIKey,
		spanclient.APIKey{
			Key:    token,
			Prefix: "",
		})

	collections, _, err := client.CollectionsApi.ListCollections(ctx)
	if err != nil {
		fmt.Println("Error listing collections: ", err.Error())
		return
	}

	fmt.Println("Collections and devices")
	fmt.Println("=======================")

	for _, collection := range collections.Collections {
		fmt.Println("Collection ID = ", collection.CollectionId, " Tags = ", collection.Tags)

		devices, _, err := client.DevicesApi.ListDevices(ctx, collection.CollectionId)
		if err != nil {
			fmt.Println("Error listing devices: ", err.Error())
			continue
		}

		for _, device := range devices.Devices {
			fmt.Println("   Device ID = ", device.DeviceId, " IMSI = ", device.Imsi, " IMEI = ", device.Imei, " Tags = ", device.Tags)
		}
		fmt.Println(len(devices.Devices), " devices in collection")
		fmt.Println()
	}
}
