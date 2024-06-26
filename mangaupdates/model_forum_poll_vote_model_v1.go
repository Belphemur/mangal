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

// checks if the ForumPollVoteModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumPollVoteModelV1{}

// ForumPollVoteModelV1 struct for ForumPollVoteModelV1
type ForumPollVoteModelV1 struct {
	ChoiceId *int64 `json:"choice_id,omitempty"`
}

// NewForumPollVoteModelV1 instantiates a new ForumPollVoteModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumPollVoteModelV1() *ForumPollVoteModelV1 {
	this := ForumPollVoteModelV1{}
	return &this
}

// NewForumPollVoteModelV1WithDefaults instantiates a new ForumPollVoteModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumPollVoteModelV1WithDefaults() *ForumPollVoteModelV1 {
	this := ForumPollVoteModelV1{}
	return &this
}

// GetChoiceId returns the ChoiceId field value if set, zero value otherwise.
func (o *ForumPollVoteModelV1) GetChoiceId() int64 {
	if o == nil || IsNil(o.ChoiceId) {
		var ret int64
		return ret
	}
	return *o.ChoiceId
}

// GetChoiceIdOk returns a tuple with the ChoiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPollVoteModelV1) GetChoiceIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ChoiceId) {
		return nil, false
	}
	return o.ChoiceId, true
}

// HasChoiceId returns a boolean if a field has been set.
func (o *ForumPollVoteModelV1) HasChoiceId() bool {
	if o != nil && !IsNil(o.ChoiceId) {
		return true
	}

	return false
}

// SetChoiceId gets a reference to the given int64 and assigns it to the ChoiceId field.
func (o *ForumPollVoteModelV1) SetChoiceId(v int64) {
	o.ChoiceId = &v
}

func (o ForumPollVoteModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumPollVoteModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChoiceId) {
		toSerialize["choice_id"] = o.ChoiceId
	}
	return toSerialize, nil
}

type NullableForumPollVoteModelV1 struct {
	value *ForumPollVoteModelV1
	isSet bool
}

func (v NullableForumPollVoteModelV1) Get() *ForumPollVoteModelV1 {
	return v.value
}

func (v *NullableForumPollVoteModelV1) Set(val *ForumPollVoteModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumPollVoteModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumPollVoteModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumPollVoteModelV1(val *ForumPollVoteModelV1) *NullableForumPollVoteModelV1 {
	return &NullableForumPollVoteModelV1{value: val, isSet: true}
}

func (v NullableForumPollVoteModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumPollVoteModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
