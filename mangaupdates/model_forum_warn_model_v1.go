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

// checks if the ForumWarnModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumWarnModelV1{}

// ForumWarnModelV1 struct for ForumWarnModelV1
type ForumWarnModelV1 struct {
	UserId     *int64             `json:"user_id,omitempty"`
	TimeAdded  *TimeV1            `json:"time_added,omitempty"`
	Level      int64              `json:"level"`
	Reason     string             `json:"reason"`
	SendReason *bool              `json:"send_reason,omitempty"`
	ByUser     *UserModelSearchV1 `json:"by_user,omitempty"`
}

// NewForumWarnModelV1 instantiates a new ForumWarnModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumWarnModelV1(level int64, reason string) *ForumWarnModelV1 {
	this := ForumWarnModelV1{}
	this.Level = level
	this.Reason = reason
	return &this
}

// NewForumWarnModelV1WithDefaults instantiates a new ForumWarnModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumWarnModelV1WithDefaults() *ForumWarnModelV1 {
	this := ForumWarnModelV1{}
	return &this
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ForumWarnModelV1) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ForumWarnModelV1) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *ForumWarnModelV1) SetUserId(v int64) {
	o.UserId = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *ForumWarnModelV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *ForumWarnModelV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *ForumWarnModelV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

// GetLevel returns the Level field value
func (o *ForumWarnModelV1) GetLevel() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Level
}

// GetLevelOk returns a tuple with the Level field value
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetLevelOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Level, true
}

// SetLevel sets field value
func (o *ForumWarnModelV1) SetLevel(v int64) {
	o.Level = v
}

// GetReason returns the Reason field value
func (o *ForumWarnModelV1) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ForumWarnModelV1) SetReason(v string) {
	o.Reason = v
}

// GetSendReason returns the SendReason field value if set, zero value otherwise.
func (o *ForumWarnModelV1) GetSendReason() bool {
	if o == nil || IsNil(o.SendReason) {
		var ret bool
		return ret
	}
	return *o.SendReason
}

// GetSendReasonOk returns a tuple with the SendReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetSendReasonOk() (*bool, bool) {
	if o == nil || IsNil(o.SendReason) {
		return nil, false
	}
	return o.SendReason, true
}

// HasSendReason returns a boolean if a field has been set.
func (o *ForumWarnModelV1) HasSendReason() bool {
	if o != nil && !IsNil(o.SendReason) {
		return true
	}

	return false
}

// SetSendReason gets a reference to the given bool and assigns it to the SendReason field.
func (o *ForumWarnModelV1) SetSendReason(v bool) {
	o.SendReason = &v
}

// GetByUser returns the ByUser field value if set, zero value otherwise.
func (o *ForumWarnModelV1) GetByUser() UserModelSearchV1 {
	if o == nil || IsNil(o.ByUser) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.ByUser
}

// GetByUserOk returns a tuple with the ByUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumWarnModelV1) GetByUserOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.ByUser) {
		return nil, false
	}
	return o.ByUser, true
}

// HasByUser returns a boolean if a field has been set.
func (o *ForumWarnModelV1) HasByUser() bool {
	if o != nil && !IsNil(o.ByUser) {
		return true
	}

	return false
}

// SetByUser gets a reference to the given UserModelSearchV1 and assigns it to the ByUser field.
func (o *ForumWarnModelV1) SetByUser(v UserModelSearchV1) {
	o.ByUser = &v
}

func (o ForumWarnModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumWarnModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	toSerialize["level"] = o.Level
	toSerialize["reason"] = o.Reason
	if !IsNil(o.SendReason) {
		toSerialize["send_reason"] = o.SendReason
	}
	if !IsNil(o.ByUser) {
		toSerialize["by_user"] = o.ByUser
	}
	return toSerialize, nil
}

type NullableForumWarnModelV1 struct {
	value *ForumWarnModelV1
	isSet bool
}

func (v NullableForumWarnModelV1) Get() *ForumWarnModelV1 {
	return v.value
}

func (v *NullableForumWarnModelV1) Set(val *ForumWarnModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumWarnModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumWarnModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumWarnModelV1(val *ForumWarnModelV1) *NullableForumWarnModelV1 {
	return &NullableForumWarnModelV1{value: val, isSet: true}
}

func (v NullableForumWarnModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumWarnModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
