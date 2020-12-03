package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
	"time"

	"github.com/antihax/optional"
	"github.com/lab5e/spanclient-go/v4"
)

// This example shows you how to retrieve payloads in a time interval. It's
// very similar to the paging example but we have added additional checks for
// the Received field on the returned data.

func main() {

	// It's always a good idea to leave authentication tokens out of the source
	// code so we use a command line parameter here.
	token := ""
	collectionID := ""
	start := int64(0)
	end := int64(time.Now().UnixNano() / int64(time.Millisecond)) // Right now in ms since epoch

	flag.StringVar(&token, "token", "", "API token for the Span service")
	flag.StringVar(&collectionID, "collection-id", "", "The collection to query")
	flag.Int64Var(&start, "start", start, "Start of the time interval (ms since epoch), ie the oldest payload")
	flag.Int64Var(&end, "end", end, "End of the time interval (ms since epoch), ie the newest payload")
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
	fmt.Fprintln(tw, "Device ID\tTime\tTransport\tPayload\tReceived")

	rows := 0
	lastMessageID := ""

	for {
		if rows > 100 {
			fmt.Println("Reached limit, stopping paging")
			return
		}

		var options *spanclient.ListCollectionDataOpts

		if rows == 0 {
			// The first time we request data we use the start and end parameters
			options = &spanclient.ListCollectionDataOpts{
				Limit: optional.NewInt32(10),
				Start: optional.NewString(fmt.Sprintf("%d", start)),
				End:   optional.NewString(fmt.Sprintf("%d", end)),
			}
		} else {
			// The subsequent calls uses the "offset" parameter to page through
			// the payloads
			options = &spanclient.ListCollectionDataOpts{
				Limit:  optional.NewInt32(10),
				Offset: optional.NewString(lastMessageID),
			}
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
			// If this is the end of the interval we'll discard the rest of the results
			ms, ts := receivedToTime(data.Received)
			if ms < start {
				tw.Flush()
				fmt.Printf("End of interval, %d rows returned\n", rows)
				return
			}
			fmt.Fprintf(tw, "%s\t%s\t%s\t%d\t%s\n", data.Device.DeviceId, ts, data.Transport, len(data.Payload), data.Received)
			// The returned payloads are sorted with the newest payload first. When lastMessageID
			// is updated here we'll have the oldest message ID set and ready to be used the next
			// time we request a chunk of messages.
			lastMessageID = data.MessageId
			rows++
		}
		tw.Flush()
	}
}

func receivedToTime(ts string) (int64, time.Time) {
	r, err := strconv.ParseInt(ts, 10, 63)
	if err != nil {
		return time.Now().UnixNano() / int64(time.Millisecond), time.Now()
	}
	return r, time.Unix(0, r*int64(time.Millisecond))
}
