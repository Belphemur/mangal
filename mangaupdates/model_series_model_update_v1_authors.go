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

// checks if the SeriesModelUpdateV1Authors type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesModelUpdateV1Authors{}

// SeriesModelUpdateV1Authors struct for SeriesModelUpdateV1Authors
type SeriesModelUpdateV1Authors struct {
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewSeriesModelUpdateV1Authors instantiates a new SeriesModelUpdateV1Authors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesModelUpdateV1Authors() *SeriesModelUpdateV1Authors {
	this := SeriesModelUpdateV1Authors{}
	return &this
}

// NewSeriesModelUpdateV1AuthorsWithDefaults instantiates a new SeriesModelUpdateV1Authors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesModelUpdateV1AuthorsWithDefaults() *SeriesModelUpdateV1Authors {
	this := SeriesModelUpdateV1Authors{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SeriesModelUpdateV1Authors) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1Authors) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SeriesModelUpdateV1Authors) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SeriesModelUpdateV1Authors) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SeriesModelUpdateV1Authors) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesModelUpdateV1Authors) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SeriesModelUpdateV1Authors) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SeriesModelUpdateV1Authors) SetType(v string) {
	o.Type = &v
}

func (o SeriesModelUpdateV1Authors) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesModelUpdateV1Authors) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSeriesModelUpdateV1Authors struct {
	value *SeriesModelUpdateV1Authors
	isSet bool
}

func (v NullableSeriesModelUpdateV1Authors) Get() *SeriesModelUpdateV1Authors {
	return v.value
}

func (v *NullableSeriesModelUpdateV1Authors) Set(val *SeriesModelUpdateV1Authors) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesModelUpdateV1Authors) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesModelUpdateV1Authors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesModelUpdateV1Authors(val *SeriesModelUpdateV1Authors) *NullableSeriesModelUpdateV1Authors {
	return &NullableSeriesModelUpdateV1Authors{value: val, isSet: true}
}

func (v NullableSeriesModelUpdateV1Authors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesModelUpdateV1Authors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


