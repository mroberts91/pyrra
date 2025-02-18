/*
Pyrra

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.3.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ObjectiveStatusAvailability struct for ObjectiveStatusAvailability
type ObjectiveStatusAvailability struct {
	Percentage float64 `json:"percentage"`
	Total      float64 `json:"total"`
	Errors     float64 `json:"errors"`
}

// NewObjectiveStatusAvailability instantiates a new ObjectiveStatusAvailability object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectiveStatusAvailability(percentage float64, total float64, errors float64) *ObjectiveStatusAvailability {
	this := ObjectiveStatusAvailability{}
	this.Percentage = percentage
	this.Total = total
	this.Errors = errors
	return &this
}

// NewObjectiveStatusAvailabilityWithDefaults instantiates a new ObjectiveStatusAvailability object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectiveStatusAvailabilityWithDefaults() *ObjectiveStatusAvailability {
	this := ObjectiveStatusAvailability{}
	return &this
}

// GetPercentage returns the Percentage field value
func (o *ObjectiveStatusAvailability) GetPercentage() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Percentage
}

// GetPercentageOk returns a tuple with the Percentage field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusAvailability) GetPercentageOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Percentage, true
}

// SetPercentage sets field value
func (o *ObjectiveStatusAvailability) SetPercentage(v float64) {
	o.Percentage = v
}

// GetTotal returns the Total field value
func (o *ObjectiveStatusAvailability) GetTotal() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusAvailability) GetTotalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *ObjectiveStatusAvailability) SetTotal(v float64) {
	o.Total = v
}

// GetErrors returns the Errors field value
func (o *ObjectiveStatusAvailability) GetErrors() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *ObjectiveStatusAvailability) GetErrorsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *ObjectiveStatusAvailability) SetErrors(v float64) {
	o.Errors = v
}

func (o ObjectiveStatusAvailability) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["percentage"] = o.Percentage
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableObjectiveStatusAvailability struct {
	value *ObjectiveStatusAvailability
	isSet bool
}

func (v NullableObjectiveStatusAvailability) Get() *ObjectiveStatusAvailability {
	return v.value
}

func (v *NullableObjectiveStatusAvailability) Set(val *ObjectiveStatusAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectiveStatusAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectiveStatusAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectiveStatusAvailability(val *ObjectiveStatusAvailability) *NullableObjectiveStatusAvailability {
	return &NullableObjectiveStatusAvailability{value: val, isSet: true}
}

func (v NullableObjectiveStatusAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectiveStatusAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
