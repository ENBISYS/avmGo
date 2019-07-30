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

type RequiredFeatures struct {
	// Postcode
	Postcode string `json:"postcode"`
	NewOrResale NewOrResale `json:"newOrResale"`
	FloorLevel FloorLevel `json:"floorLevel"`
	// Floor area (sqf)
	TotalFloorAreaInSqf int32 `json:"totalFloorAreaInSqf"`
	PropertyType PropertyType `json:"propertyType"`
	NumberOfRooms int32 `json:"numberOfRooms"`
}