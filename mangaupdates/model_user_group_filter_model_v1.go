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

// checks if the UserGroupFilterModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserGroupFilterModelV1{}

// UserGroupFilterModelV1 struct for UserGroupFilterModelV1
type UserGroupFilterModelV1 struct {
	GroupId   *int64  `json:"group_id,omitempty"`
	GroupName *string `json:"group_name,omitempty"`
}

// NewUserGroupFilterModelV1 instantiates a new UserGroupFilterModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserGroupFilterModelV1() *UserGroupFilterModelV1 {
	this := UserGroupFilterModelV1{}
	return &this
}

// NewUserGroupFilterModelV1WithDefaults instantiates a new UserGroupFilterModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserGroupFilterModelV1WithDefaults() *UserGroupFilterModelV1 {
	this := UserGroupFilterModelV1{}
	return &this
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *UserGroupFilterModelV1) GetGroupId() int64 {
	if o == nil || IsNil(o.GroupId) {
		var ret int64
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupFilterModelV1) GetGroupIdOk() (*int64, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *UserGroupFilterModelV1) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given int64 and assigns it to the GroupId field.
func (o *UserGroupFilterModelV1) SetGroupId(v int64) {
	o.GroupId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *UserGroupFilterModelV1) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserGroupFilterModelV1) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *UserGroupFilterModelV1) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *UserGroupFilterModelV1) SetGroupName(v string) {
	o.GroupName = &v
}

func (o UserGroupFilterModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserGroupFilterModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	if !IsNil(o.GroupName) {
		toSerialize["group_name"] = o.GroupName
	}
	return toSerialize, nil
}

type NullableUserGroupFilterModelV1 struct {
	value *UserGroupFilterModelV1
	isSet bool
}

func (v NullableUserGroupFilterModelV1) Get() *UserGroupFilterModelV1 {
	return v.value
}

func (v *NullableUserGroupFilterModelV1) Set(val *UserGroupFilterModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableUserGroupFilterModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableUserGroupFilterModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserGroupFilterModelV1(val *UserGroupFilterModelV1) *NullableUserGroupFilterModelV1 {
	return &NullableUserGroupFilterModelV1{value: val, isSet: true}
}

func (v NullableUserGroupFilterModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserGroupFilterModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
