/*
 * AVM
 *
 * This is api for AVM (automated valuation machine)
 *
 * API version: 1.0.0
 * Contact: info@enbisys.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package avm

// A Problem Details object (RFC 7807)
type Problem struct {
	// A relative URI reference that identifies the problem type. When dereferenced, it SHOULD provide human-readable documentation for the problem type (e.g. using HTML).
	Type string `json:"type,omitempty"`
	// A short, summary of the problem type. Written in english and readable for engineers (usually not suited for non technical stakeholders and not localized)
	Title string `json:"title,omitempty"`
	// The HTTP status code generated by the origin server for this occurrence of the problem.
	Status int32 `json:"status,omitempty"`
	// A human-readable explanation specific to this occurrence of the problem
	Detail string `json:"detail,omitempty"`
	// A URI reference that identifies the specific occurrence of the problem. It may or may not yield further information if dereferenced.
	Instance string `json:"instance,omitempty"`
}
