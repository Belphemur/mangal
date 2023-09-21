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

// checks if the ForumForumModelUpdateV1Admin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumForumModelUpdateV1Admin{}

// ForumForumModelUpdateV1Admin struct for ForumForumModelUpdateV1Admin
type ForumForumModelUpdateV1Admin struct {
	Locked *bool `json:"locked,omitempty"`
	Hidden *bool `json:"hidden,omitempty"`
	VerifyAge *bool `json:"verify_age,omitempty"`
}

// NewForumForumModelUpdateV1Admin instantiates a new ForumForumModelUpdateV1Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumForumModelUpdateV1Admin() *ForumForumModelUpdateV1Admin {
	this := ForumForumModelUpdateV1Admin{}
	return &this
}

// NewForumForumModelUpdateV1AdminWithDefaults instantiates a new ForumForumModelUpdateV1Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumForumModelUpdateV1AdminWithDefaults() *ForumForumModelUpdateV1Admin {
	this := ForumForumModelUpdateV1Admin{}
	return &this
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *ForumForumModelUpdateV1Admin) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumForumModelUpdateV1Admin) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *ForumForumModelUpdateV1Admin) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *ForumForumModelUpdateV1Admin) SetLocked(v bool) {
	o.Locked = &v
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *ForumForumModelUpdateV1Admin) GetHidden() bool {
	if o == nil || IsNil(o.Hidden) {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumForumModelUpdateV1Admin) GetHiddenOk() (*bool, bool) {
	if o == nil || IsNil(o.Hidden) {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *ForumForumModelUpdateV1Admin) HasHidden() bool {
	if o != nil && !IsNil(o.Hidden) {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *ForumForumModelUpdateV1Admin) SetHidden(v bool) {
	o.Hidden = &v
}

// GetVerifyAge returns the VerifyAge field value if set, zero value otherwise.
func (o *ForumForumModelUpdateV1Admin) GetVerifyAge() bool {
	if o == nil || IsNil(o.VerifyAge) {
		var ret bool
		return ret
	}
	return *o.VerifyAge
}

// GetVerifyAgeOk returns a tuple with the VerifyAge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumForumModelUpdateV1Admin) GetVerifyAgeOk() (*bool, bool) {
	if o == nil || IsNil(o.VerifyAge) {
		return nil, false
	}
	return o.VerifyAge, true
}

// HasVerifyAge returns a boolean if a field has been set.
func (o *ForumForumModelUpdateV1Admin) HasVerifyAge() bool {
	if o != nil && !IsNil(o.VerifyAge) {
		return true
	}

	return false
}

// SetVerifyAge gets a reference to the given bool and assigns it to the VerifyAge field.
func (o *ForumForumModelUpdateV1Admin) SetVerifyAge(v bool) {
	o.VerifyAge = &v
}

func (o ForumForumModelUpdateV1Admin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumForumModelUpdateV1Admin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.Hidden) {
		toSerialize["hidden"] = o.Hidden
	}
	if !IsNil(o.VerifyAge) {
		toSerialize["verify_age"] = o.VerifyAge
	}
	return toSerialize, nil
}

type NullableForumForumModelUpdateV1Admin struct {
	value *ForumForumModelUpdateV1Admin
	isSet bool
}

func (v NullableForumForumModelUpdateV1Admin) Get() *ForumForumModelUpdateV1Admin {
	return v.value
}

func (v *NullableForumForumModelUpdateV1Admin) Set(val *ForumForumModelUpdateV1Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableForumForumModelUpdateV1Admin) IsSet() bool {
	return v.isSet
}

func (v *NullableForumForumModelUpdateV1Admin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumForumModelUpdateV1Admin(val *ForumForumModelUpdateV1Admin) *NullableForumForumModelUpdateV1Admin {
	return &NullableForumForumModelUpdateV1Admin{value: val, isSet: true}
}

func (v NullableForumForumModelUpdateV1Admin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumForumModelUpdateV1Admin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

