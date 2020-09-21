package spanclient

import "encoding/base64"

// PayloadBytes returns the payload for an OutputDataMessage as a byte array
func (d *OutputDataMessage) PayloadBytes() ([]byte, error) {
	return base64.StdEncoding.DecodeString(d.Payload)
}
