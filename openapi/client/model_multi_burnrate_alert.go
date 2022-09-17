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

// MultiBurnrateAlert struct for MultiBurnrateAlert
type MultiBurnrateAlert struct {
	Labels   map[string]string `json:"labels"`
	Severity string            `json:"severity"`
	For      int64             `json:"for"`
	Factor   float64           `json:"factor"`
	Short    Burnrate          `json:"short"`
	Long     Burnrate          `json:"long"`
	State    string            `json:"state"`
}

// NewMultiBurnrateAlert instantiates a new MultiBurnrateAlert object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultiBurnrateAlert(labels map[string]string, severity string, for_ int64, factor float64, short Burnrate, long Burnrate, state string) *MultiBurnrateAlert {
	this := MultiBurnrateAlert{}
	this.Labels = labels
	this.Severity = severity
	this.For = for_
	this.Factor = factor
	this.Short = short
	this.Long = long
	this.State = state
	return &this
}

// NewMultiBurnrateAlertWithDefaults instantiates a new MultiBurnrateAlert object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultiBurnrateAlertWithDefaults() *MultiBurnrateAlert {
	this := MultiBurnrateAlert{}
	return &this
}

// GetLabels returns the Labels field value
func (o *MultiBurnrateAlert) GetLabels() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetLabelsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *MultiBurnrateAlert) SetLabels(v map[string]string) {
	o.Labels = v
}

// GetSeverity returns the Severity field value
func (o *MultiBurnrateAlert) GetSeverity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetSeverityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *MultiBurnrateAlert) SetSeverity(v string) {
	o.Severity = v
}

// GetFor returns the For field value
func (o *MultiBurnrateAlert) GetFor() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.For
}

// GetForOk returns a tuple with the For field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetForOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.For, true
}

// SetFor sets field value
func (o *MultiBurnrateAlert) SetFor(v int64) {
	o.For = v
}

// GetFactor returns the Factor field value
func (o *MultiBurnrateAlert) GetFactor() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Factor
}

// GetFactorOk returns a tuple with the Factor field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetFactorOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Factor, true
}

// SetFactor sets field value
func (o *MultiBurnrateAlert) SetFactor(v float64) {
	o.Factor = v
}

// GetShort returns the Short field value
func (o *MultiBurnrateAlert) GetShort() Burnrate {
	if o == nil {
		var ret Burnrate
		return ret
	}

	return o.Short
}

// GetShortOk returns a tuple with the Short field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetShortOk() (*Burnrate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Short, true
}

// SetShort sets field value
func (o *MultiBurnrateAlert) SetShort(v Burnrate) {
	o.Short = v
}

// GetLong returns the Long field value
func (o *MultiBurnrateAlert) GetLong() Burnrate {
	if o == nil {
		var ret Burnrate
		return ret
	}

	return o.Long
}

// GetLongOk returns a tuple with the Long field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetLongOk() (*Burnrate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Long, true
}

// SetLong sets field value
func (o *MultiBurnrateAlert) SetLong(v Burnrate) {
	o.Long = v
}

// GetState returns the State field value
func (o *MultiBurnrateAlert) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *MultiBurnrateAlert) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *MultiBurnrateAlert) SetState(v string) {
	o.State = v
}

func (o MultiBurnrateAlert) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["for"] = o.For
	}
	if true {
		toSerialize["factor"] = o.Factor
	}
	if true {
		toSerialize["short"] = o.Short
	}
	if true {
		toSerialize["long"] = o.Long
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableMultiBurnrateAlert struct {
	value *MultiBurnrateAlert
	isSet bool
}

func (v NullableMultiBurnrateAlert) Get() *MultiBurnrateAlert {
	return v.value
}

func (v *NullableMultiBurnrateAlert) Set(val *MultiBurnrateAlert) {
	v.value = val
	v.isSet = true
}

func (v NullableMultiBurnrateAlert) IsSet() bool {
	return v.isSet
}

func (v *NullableMultiBurnrateAlert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultiBurnrateAlert(val *MultiBurnrateAlert) *NullableMultiBurnrateAlert {
	return &NullableMultiBurnrateAlert{value: val, isSet: true}
}

func (v NullableMultiBurnrateAlert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultiBurnrateAlert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
