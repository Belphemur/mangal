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

// checks if the SeriesHistoryModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesHistoryModelV1{}

// SeriesHistoryModelV1 struct for SeriesHistoryModelV1
type SeriesHistoryModelV1 struct {
	ChangeId  *int64  `json:"change_id,omitempty"`
	Username  *string `json:"username,omitempty"`
	Action    *string `json:"action,omitempty"`
	Changed   *string `json:"changed,omitempty"`
	TimeAdded *TimeV1 `json:"time_added,omitempty"`
}

// NewSeriesHistoryModelV1 instantiates a new SeriesHistoryModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesHistoryModelV1() *SeriesHistoryModelV1 {
	this := SeriesHistoryModelV1{}
	return &this
}

// NewSeriesHistoryModelV1WithDefaults instantiates a new SeriesHistoryModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesHistoryModelV1WithDefaults() *SeriesHistoryModelV1 {
	this := SeriesHistoryModelV1{}
	return &this
}

// GetChangeId returns the ChangeId field value if set, zero value otherwise.
func (o *SeriesHistoryModelV1) GetChangeId() int64 {
	if o == nil || IsNil(o.ChangeId) {
		var ret int64
		return ret
	}
	return *o.ChangeId
}

// GetChangeIdOk returns a tuple with the ChangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesHistoryModelV1) GetChangeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ChangeId) {
		return nil, false
	}
	return o.ChangeId, true
}

// HasChangeId returns a boolean if a field has been set.
func (o *SeriesHistoryModelV1) HasChangeId() bool {
	if o != nil && !IsNil(o.ChangeId) {
		return true
	}

	return false
}

// SetChangeId gets a reference to the given int64 and assigns it to the ChangeId field.
func (o *SeriesHistoryModelV1) SetChangeId(v int64) {
	o.ChangeId = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *SeriesHistoryModelV1) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesHistoryModelV1) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *SeriesHistoryModelV1) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *SeriesHistoryModelV1) SetUsername(v string) {
	o.Username = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SeriesHistoryModelV1) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesHistoryModelV1) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SeriesHistoryModelV1) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SeriesHistoryModelV1) SetAction(v string) {
	o.Action = &v
}

// GetChanged returns the Changed field value if set, zero value otherwise.
func (o *SeriesHistoryModelV1) GetChanged() string {
	if o == nil || IsNil(o.Changed) {
		var ret string
		return ret
	}
	return *o.Changed
}

// GetChangedOk returns a tuple with the Changed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesHistoryModelV1) GetChangedOk() (*string, bool) {
	if o == nil || IsNil(o.Changed) {
		return nil, false
	}
	return o.Changed, true
}

// HasChanged returns a boolean if a field has been set.
func (o *SeriesHistoryModelV1) HasChanged() bool {
	if o != nil && !IsNil(o.Changed) {
		return true
	}

	return false
}

// SetChanged gets a reference to the given string and assigns it to the Changed field.
func (o *SeriesHistoryModelV1) SetChanged(v string) {
	o.Changed = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *SeriesHistoryModelV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesHistoryModelV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *SeriesHistoryModelV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *SeriesHistoryModelV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

func (o SeriesHistoryModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesHistoryModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChangeId) {
		toSerialize["change_id"] = o.ChangeId
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.Changed) {
		toSerialize["changed"] = o.Changed
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	return toSerialize, nil
}

type NullableSeriesHistoryModelV1 struct {
	value *SeriesHistoryModelV1
	isSet bool
}

func (v NullableSeriesHistoryModelV1) Get() *SeriesHistoryModelV1 {
	return v.value
}

func (v *NullableSeriesHistoryModelV1) Set(val *SeriesHistoryModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesHistoryModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesHistoryModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesHistoryModelV1(val *SeriesHistoryModelV1) *NullableSeriesHistoryModelV1 {
	return &NullableSeriesHistoryModelV1{value: val, isSet: true}
}

func (v NullableSeriesHistoryModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesHistoryModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
