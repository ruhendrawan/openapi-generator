/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	"bytes"
	"encoding/json"
)

// HealthCheckResult Just a string to inform instance is up and running. Make it nullable in hope to get it as pointer in generated model.
type HealthCheckResult struct {
	NullableMessage *NullableString `json:"NullableMessage,omitempty"`
}

// NewHealthCheckResult instantiates a new HealthCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthCheckResult() *HealthCheckResult {
    this := HealthCheckResult{}
    return &this
}

// NewHealthCheckResultWithDefaults instantiates a new HealthCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthCheckResultWithDefaults() *HealthCheckResult {
    this := HealthCheckResult{}
    return &this
}

// GetNullableMessage returns the NullableMessage field value if set, zero value otherwise.
func (o *HealthCheckResult) GetNullableMessage() NullableString {
	if o == nil || o.NullableMessage == nil {
		var ret NullableString
		return ret
	}
	return *o.NullableMessage
}

// GetNullableMessageOk returns a tuple with the NullableMessage field value if set, zero value otherwise
// and a boolean to check if the value has been set.
func (o *HealthCheckResult) GetNullableMessageOk() (NullableString, bool) {
	if o == nil || o.NullableMessage == nil {
		var ret NullableString
		return ret, false
	}
	return *o.NullableMessage, true
}

// HasNullableMessage returns a boolean if a field has been set.
func (o *HealthCheckResult) HasNullableMessage() bool {
	if o != nil && o.NullableMessage != nil {
		return true
	}

	return false
}

// SetNullableMessage gets a reference to the given NullableString and assigns it to the NullableMessage field.
func (o *HealthCheckResult) SetNullableMessage(v NullableString) {
	o.NullableMessage = &v
}

type NullableHealthCheckResult struct {
	Value HealthCheckResult
	ExplicitNull bool
}

func (v NullableHealthCheckResult) MarshalJSON() ([]byte, error) {
    switch {
    case v.ExplicitNull:
        return []byte("null"), nil
    default:
		return json.Marshal(v.Value)
	}
}

func (v *NullableHealthCheckResult) UnmarshalJSON(src []byte) error {
	if bytes.Equal(src, []byte("null")) {
		v.ExplicitNull = true
		return nil
	}

	return json.Unmarshal(src, &v.Value)
}
