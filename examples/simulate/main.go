package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lab5e/spanclient-go/v4"
)

const (
	iterations  = 100
	deviceCount = 10
	stepSleep   = 100 * time.Millisecond
	imsiOffset  = 50000
	numQueries  = 10
)

func main() {
	setIMSIOffset(imsiOffset)
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
	config.BasePath = "http://bob.localdomain/span"
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

	errors := 0
	for i := 0; i < iterations; i++ {
		// Create a collection
		collection, _, err := client.CollectionsApi.CreateCollection(ctx, spanclient.Collection{})
		if err != nil {
			log.Printf("Got error creating collection: %v", err)
			errors++
			time.Sleep(1 * time.Second)
			continue
		}
		log.Printf("Created collection with ID %s", collection.CollectionId)
		time.Sleep(stepSleep)

		log.Printf("Creating devices")
		// Create 10 devices
		devicesToDelete := make([]string, 0)
		for devNum := 0; devNum < deviceCount; devNum++ {
			device, _, err := client.DevicesApi.CreateDevice(ctx,
				collection.CollectionId,
				spanclient.Device{
					Imsi: nextIMSI(),
					Imei: nextIMEI(),
				})
			if err != nil {
				log.Printf("Error creating device for collection %s: %v", collection.CollectionId, err)
				errors++
				break
			}
			devicesToDelete = append(devicesToDelete, device.DeviceId)
			time.Sleep(stepSleep)
		}

		log.Println("Querying devices")
		// Query all devices N times
		for queryNum := 0; queryNum < numQueries; queryNum++ {
			log.Printf("Query %d of %d...", queryNum+1, numQueries)
			for _, id := range devicesToDelete {
				_, _, err = client.DevicesApi.RetrieveDevice(ctx, collection.CollectionId, id)
				if err != nil {
					log.Printf("Error querying device %s for colleciton %s: %v", id, collection.CollectionId, err)
					time.Sleep(1 * time.Second)
					errors++
				}
				time.Sleep(stepSleep)
			}
		}

		log.Println("Removing devices")
		// Remove devices
		for _, id := range devicesToDelete {
			_, _, err := client.DevicesApi.DeleteDevice(ctx, collection.CollectionId, id)
			if err != nil {
				log.Printf("Error removing device %s for collection %s: %v", id, collection.CollectionId, err)
				time.Sleep(1 * time.Second)
				errors++
			}
			time.Sleep(stepSleep)
		}

		log.Println("Removing collection")
		// Remove collection
		_, _, err = client.CollectionsApi.DeleteCollection(ctx, collection.CollectionId)
		if err != nil {
			log.Printf("Got error creating collection: %v", err)
			errors++
			time.Sleep(1 * time.Second)
			continue
		}
		time.Sleep(stepSleep)
	}
	log.Printf("Completed with %d error(s)", errors)
}
