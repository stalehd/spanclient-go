/*
 * The Span API
 *
 * API for device, collection, output and firmware management
 *
 * API version: 4.0.11 ambulant-epsie
 * Contact: dev@lab5e.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spanclient
// DataDumpResponse struct for DataDumpResponse
type DataDumpResponse struct {
	Collections []DumpedCollection `json:"collections,omitempty"`
}
