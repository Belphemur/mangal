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

// checks if the ListsSeriesModelV1Status type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListsSeriesModelV1Status{}

// ListsSeriesModelV1Status struct for ListsSeriesModelV1Status
type ListsSeriesModelV1Status struct {
	Volume  *int64 `json:"volume,omitempty"`
	Chapter *int64 `json:"chapter,omitempty"`
}

// NewListsSeriesModelV1Status instantiates a new ListsSeriesModelV1Status object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListsSeriesModelV1Status() *ListsSeriesModelV1Status {
	this := ListsSeriesModelV1Status{}
	return &this
}

// NewListsSeriesModelV1StatusWithDefaults instantiates a new ListsSeriesModelV1Status object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListsSeriesModelV1StatusWithDefaults() *ListsSeriesModelV1Status {
	this := ListsSeriesModelV1Status{}
	return &this
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ListsSeriesModelV1Status) GetVolume() int64 {
	if o == nil || IsNil(o.Volume) {
		var ret int64
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsSeriesModelV1Status) GetVolumeOk() (*int64, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ListsSeriesModelV1Status) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given int64 and assigns it to the Volume field.
func (o *ListsSeriesModelV1Status) SetVolume(v int64) {
	o.Volume = &v
}

// GetChapter returns the Chapter field value if set, zero value otherwise.
func (o *ListsSeriesModelV1Status) GetChapter() int64 {
	if o == nil || IsNil(o.Chapter) {
		var ret int64
		return ret
	}
	return *o.Chapter
}

// GetChapterOk returns a tuple with the Chapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListsSeriesModelV1Status) GetChapterOk() (*int64, bool) {
	if o == nil || IsNil(o.Chapter) {
		return nil, false
	}
	return o.Chapter, true
}

// HasChapter returns a boolean if a field has been set.
func (o *ListsSeriesModelV1Status) HasChapter() bool {
	if o != nil && !IsNil(o.Chapter) {
		return true
	}

	return false
}

// SetChapter gets a reference to the given int64 and assigns it to the Chapter field.
func (o *ListsSeriesModelV1Status) SetChapter(v int64) {
	o.Chapter = &v
}

func (o ListsSeriesModelV1Status) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListsSeriesModelV1Status) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Chapter) {
		toSerialize["chapter"] = o.Chapter
	}
	return toSerialize, nil
}

type NullableListsSeriesModelV1Status struct {
	value *ListsSeriesModelV1Status
	isSet bool
}

func (v NullableListsSeriesModelV1Status) Get() *ListsSeriesModelV1Status {
	return v.value
}

func (v *NullableListsSeriesModelV1Status) Set(val *ListsSeriesModelV1Status) {
	v.value = val
	v.isSet = true
}

func (v NullableListsSeriesModelV1Status) IsSet() bool {
	return v.isSet
}

func (v *NullableListsSeriesModelV1Status) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListsSeriesModelV1Status(val *ListsSeriesModelV1Status) *NullableListsSeriesModelV1Status {
	return &NullableListsSeriesModelV1Status{value: val, isSet: true}
}

func (v NullableListsSeriesModelV1Status) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListsSeriesModelV1Status) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
