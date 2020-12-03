package main

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/antihax/optional"
	"github.com/lab5e/spanclient-go/v4"
)

// This example demonstrates how to page through the data/payloads from devices.
// The paging uses the "offset" parameter to page through the data.
//
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

	// Set this to true to list the requests and responses in the client. It can
	// be useful if you are wondering what is happening.
	config.Debug = false

	client := spanclient.NewAPIClient(config)

	// In the Real World this context should also have a context.WithTimeout
	// call to ensure we time out if there's no response.
	ctx := spanclient.NewAuthContext(token)

	collection, _, err := client.CollectionsApi.RetrieveCollection(ctx, collectionID)
	if err != nil {
		fmt.Println("Error retrieving collection: ", err.Error())
		return
	}

	fmt.Println("Data from collection ", collection.CollectionId)
	fmt.Println("======================================")

	tw := tabwriter.NewWriter(os.Stderr, 12, 4, 2, ' ', 0)
	fmt.Fprintln(tw, "Device ID\tTransport\tPayload")

	rows := 0
	lastMessageID := ""
	for {
		if rows > 100 {
			fmt.Println("Reached limit, stopping paging")
			return
		}

		// The first time the options are set the "offset" parameter is set to
		// an empty string and you'll get the most recent messages back from the
		// service.
		options := &spanclient.ListCollectionDataOpts{
			Limit:  optional.NewInt32(10),
			Offset: optional.NewString(lastMessageID),
		}

		items, _, err := client.CollectionsApi.ListCollectionData(ctx, collection.CollectionId, options)

		if err != nil {
			tw.Flush()
			fmt.Println("Error retrieving data: ", err.Error())
			return
		}

		if len(items.Data) == 0 {
			tw.Flush()
			fmt.Println("Zero rows returned, exiting.")
			return
		}

		for _, data := range items.Data {
			fmt.Fprintf(tw, "%s\t%s\t%s\n", data.Device.DeviceId, data.Transport, data.Payload)
			// The returned payloads are sorted with the newest payload first. When lastMessageID
			// is updated here we'll have the oldest message ID set and ready to be used the next
			// time we request a chunk of messages.
			lastMessageID = data.MessageId
			rows++
		}
		tw.Flush()
	}
}
