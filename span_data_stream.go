package spanclient

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/websocket"
)

//
//Copyright 2020 Lab5e AS
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//

// DataStream is a stream of data from the Span service. The data stream uses
// a websocket under the hood.
type DataStream interface {
	Recv() (OutputDataMessage, error)
	Close() error
}

// NewCollectionDataStream creates a live data stream for devices in a
// collection. The context must contain the appropriate credentials.
func NewCollectionDataStream(ctx context.Context, config *Configuration, collectionID string) (DataStream, error) {
	wsURL, err := remapURL(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s/collections/%s/from", wsURL, collectionID)
	log.Printf("### urlStr='%s'", urlStr)
	return newDataStream(ctx, urlStr)
}

// NewDeviceDataStream creates a live data stream for a single device. The
// context must contain the appropriate credentials.
func NewDeviceDataStream(ctx context.Context, config *Configuration, collectionID, deviceID string) (DataStream, error) {
	wsURL, err := remapURL(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s/collections/%s/devices/%s/from", wsURL, collectionID, deviceID)
	log.Printf("### urlStr='%s'", urlStr)
	return newDataStream(ctx, urlStr)
}

// remapURL is a (hopefully) a temporary fix to get around the
// confusion about what Scheme, Host and BasePath are supposed to be
// in Configuration
func remapURL(config *Configuration) (string, error) {
	url, err := url.Parse(config.BasePath)
	if err != nil {
		return "", err
	}

	scheme := ""

	switch url.Scheme {
	case "http":
		scheme = "ws"
	case "https":
		scheme = "wss"
	default:
		scheme = "wss"
	}

	url.Scheme = scheme

	return url.String(), nil
}

func newDataStream(ctx context.Context, urlStr string) (DataStream, error) {
	header := http.Header{}

	t := ctx.Value(ContextAPIKey)
	if t == nil {
		return nil, errors.New("no auth token in context")
	}

	key, ok := t.(APIKey)
	if !ok {
		return nil, errors.New("not a token in context")
	}
	header.Add("X-API-Token", key.Key)
	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial(urlStr, header)
	if err != nil {
		return nil, fmt.Errorf("Error dialing websocket: %v", err)
	}
	return &wsDataStream{ws}, nil
}

type wsDataStream struct {
	ws *websocket.Conn
}

func (d *wsDataStream) Recv() (OutputDataMessage, error) {
	for {
		m := OutputDataMessage{}
		err := d.ws.ReadJSON(&m)
		if err != nil {
			return OutputDataMessage{}, err
		}
		if m.Type == DATA {
			return m, nil
		}
	}
}

func (d *wsDataStream) Close() error {
	return d.ws.Close()
}
