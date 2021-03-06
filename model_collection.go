/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.1.3 factual-kahlil
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanclient
// Collection struct for Collection
type Collection struct {
	// The ID of the collection. This is assigned by the backend.
	CollectionId string `json:"collectionId,omitempty"`
	// The team ID that owns the collection. This field is required. When you create new collections the default is to use your private team ID.
	TeamId string `json:"teamId,omitempty"`
	FieldMask FieldMask `json:"fieldMask,omitempty"`
	Firmware CollectionFirmware `json:"firmware,omitempty"`
	// Tags for the collection. Tags are metadata fields that you can assign to the collection.
	Tags map[string]string `json:"tags,omitempty"`
}
