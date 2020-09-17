package main

import (
	"context"
	"flag"
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/lab5e/spanclient"
)

func main() {
	// It's always a good idea to leave authentication tokens out of the source
	// code so we use a command line parameter here.
	token := ""
	collectionID := ""
	flag.StringVar(&token, "token", "", "API token for the Span service")
	flag.StringVar(&collectionID, "collection-id", "", "The collection to query")

	flag.Parse()

	if token == "" {
		fmt.Println("Missing token parameter")
		flag.PrintDefaults()
		return
	}

	config := spanclient.NewConfiguration()
	config.Host = "127.0.0.1:8080"
	config.BasePath = ""
	config.Scheme = "http"

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

	collection, _, err := client.CollectionsApi.RetrieveCollection(ctx, collectionID)
	if err != nil {
		fmt.Println("Error retrieving collection: ", err.Error())
		return
	}

	fmt.Println("Data from collection ", collection.CollectionId)
	fmt.Println("======================================")

	options := &spanclient.ListCollectionDataOpts{
		Limit: optional.NewInt32(10),
		Start: optional.NewString("0"),
		End:   optional.NewString(fmt.Sprintf("%d", time.Now().UnixNano())),
	}
	items, _, err := client.CollectionsApi.ListCollectionData(ctx, collection.CollectionId, options)
	if err != nil {
		fmt.Println("Error retrieving data: ", err.Error())
	}
	for _, data := range items.Data {
		fmt.Println("Device ID: ", data.Device.DeviceId, " Payload: ", data.Payload, " Transport: ", data.Transport)
	}
}
