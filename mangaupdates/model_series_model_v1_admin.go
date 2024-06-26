/*
MangaUpdates API

This API powers our website. Most API functions are public and do not require an account. For user-based functions, you must register an account. Currently, user registration must be done through our main website and is disabled via this API.  Download OpenAPI Specification: [openapi.yaml](openapi.yaml)  Please contact us at the following emails if you have questions:  * lambchopsil@mangaupdates.com * manick@mangaupdates.com  ## Warranties  MangaUpdates makes no warranties about service availability, correctness of the data, or anything else. The service is provided as is, and may change at any time.  ## Acceptable Use Policy  * You will credit MangaUpdates when using data provided by this API. * You will use reasonable spacing between requests so as not to overwhelm the MangaUpdates servers, and employ caching mechanisms when accessing data. * You will NOT use MangaUpdates data or API in a way that will:     * Deceive or defraud users     * Assist or perform an illegal action     * Create spam     * Damage the database  We reserve the right to change this policy at any time.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the SeriesModelV1Admin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelV1Admin{}

// SeriesModelV1Admin struct for SeriesModelV1Admin
type SeriesModelV1Admin struct {
	AddedBy *UserModelSearchV1 `json:"added_by,omitempty"`
	Approved *bool `json:"approved,omitempty"`
}

// NewSeriesModelV1Admin instantiates a new SeriesModelV1Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelV1Admin() *SeriesModelV1Admin {
	this := SeriesModelV1Admin{}
	return &this
}

// NewSeriesModelV1AdminWithDefaults instantiates a new SeriesModelV1Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelV1AdminWithDefaults() *SeriesModelV1Admin {
	this := SeriesModelV1Admin{}
	return &this
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *SeriesModelV1Admin) GetAddedBy() UserModelSearchV1 {
	if o == nil || IsNil(o.AddedBy) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Admin) GetAddedByOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *SeriesModelV1Admin) HasAddedBy() bool {
	if o != nil && !IsNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given UserModelSearchV1 and assigns it to the AddedBy field.
func (o *SeriesModelV1Admin) SetAddedBy(v UserModelSearchV1) {
	o.AddedBy = &v
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *SeriesModelV1Admin) GetApproved() bool {
	if o == nil || IsNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelV1Admin) GetApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *SeriesModelV1Admin) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *SeriesModelV1Admin) SetApproved(v bool) {
	o.Approved = &v
}

func (o SeriesModelV1Admin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelV1Admin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddedBy) {
		toSerialize["added_by"] = o.AddedBy
	}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	return toSerialize, nil
}

type NullableSeriesModelV1Admin struct {
	value *SeriesModelV1Admin
	isSet bool
}

func (v NullableSeriesModelV1Admin) Get() *SeriesModelV1Admin {
	return v.value
}

func (v *NullableSeriesModelV1Admin) Set(val *SeriesModelV1Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelV1Admin) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelV1Admin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelV1Admin(val *SeriesModelV1Admin) *NullableSeriesModelV1Admin {
	return &NullableSeriesModelV1Admin{value: val, isSet: true}
}

func (v NullableSeriesModelV1Admin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelV1Admin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


