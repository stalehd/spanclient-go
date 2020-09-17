/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 3.0.0
 * Contact: dev@lab5e.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package spanclient
// DumpedCollection struct for DumpedCollection
type DumpedCollection struct {
	Collection Collection `json:"collection,omitempty"`
	Devices []DumpedDevice `json:"devices,omitempty"`
	Outputs []Output `json:"outputs,omitempty"`
}