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

// checks if the MiscStatsModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MiscStatsModelV1{}

// MiscStatsModelV1 struct for MiscStatsModelV1
type MiscStatsModelV1 struct {
	TotalUsers       *int64             `json:"total_users,omitempty"`
	LatestUser       *UserModelSearchV1 `json:"latest_user,omitempty"`
	TotalForumTopics *int64             `json:"total_forum_topics,omitempty"`
	TotalForumPosts  *int64             `json:"total_forum_posts,omitempty"`
}

// NewMiscStatsModelV1 instantiates a new MiscStatsModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMiscStatsModelV1() *MiscStatsModelV1 {
	this := MiscStatsModelV1{}
	return &this
}

// NewMiscStatsModelV1WithDefaults instantiates a new MiscStatsModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMiscStatsModelV1WithDefaults() *MiscStatsModelV1 {
	this := MiscStatsModelV1{}
	return &this
}

// GetTotalUsers returns the TotalUsers field value if set, zero value otherwise.
func (o *MiscStatsModelV1) GetTotalUsers() int64 {
	if o == nil || IsNil(o.TotalUsers) {
		var ret int64
		return ret
	}
	return *o.TotalUsers
}

// GetTotalUsersOk returns a tuple with the TotalUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscStatsModelV1) GetTotalUsersOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalUsers) {
		return nil, false
	}
	return o.TotalUsers, true
}

// HasTotalUsers returns a boolean if a field has been set.
func (o *MiscStatsModelV1) HasTotalUsers() bool {
	if o != nil && !IsNil(o.TotalUsers) {
		return true
	}

	return false
}

// SetTotalUsers gets a reference to the given int64 and assigns it to the TotalUsers field.
func (o *MiscStatsModelV1) SetTotalUsers(v int64) {
	o.TotalUsers = &v
}

// GetLatestUser returns the LatestUser field value if set, zero value otherwise.
func (o *MiscStatsModelV1) GetLatestUser() UserModelSearchV1 {
	if o == nil || IsNil(o.LatestUser) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.LatestUser
}

// GetLatestUserOk returns a tuple with the LatestUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscStatsModelV1) GetLatestUserOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.LatestUser) {
		return nil, false
	}
	return o.LatestUser, true
}

// HasLatestUser returns a boolean if a field has been set.
func (o *MiscStatsModelV1) HasLatestUser() bool {
	if o != nil && !IsNil(o.LatestUser) {
		return true
	}

	return false
}

// SetLatestUser gets a reference to the given UserModelSearchV1 and assigns it to the LatestUser field.
func (o *MiscStatsModelV1) SetLatestUser(v UserModelSearchV1) {
	o.LatestUser = &v
}

// GetTotalForumTopics returns the TotalForumTopics field value if set, zero value otherwise.
func (o *MiscStatsModelV1) GetTotalForumTopics() int64 {
	if o == nil || IsNil(o.TotalForumTopics) {
		var ret int64
		return ret
	}
	return *o.TotalForumTopics
}

// GetTotalForumTopicsOk returns a tuple with the TotalForumTopics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscStatsModelV1) GetTotalForumTopicsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalForumTopics) {
		return nil, false
	}
	return o.TotalForumTopics, true
}

// HasTotalForumTopics returns a boolean if a field has been set.
func (o *MiscStatsModelV1) HasTotalForumTopics() bool {
	if o != nil && !IsNil(o.TotalForumTopics) {
		return true
	}

	return false
}

// SetTotalForumTopics gets a reference to the given int64 and assigns it to the TotalForumTopics field.
func (o *MiscStatsModelV1) SetTotalForumTopics(v int64) {
	o.TotalForumTopics = &v
}

// GetTotalForumPosts returns the TotalForumPosts field value if set, zero value otherwise.
func (o *MiscStatsModelV1) GetTotalForumPosts() int64 {
	if o == nil || IsNil(o.TotalForumPosts) {
		var ret int64
		return ret
	}
	return *o.TotalForumPosts
}

// GetTotalForumPostsOk returns a tuple with the TotalForumPosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MiscStatsModelV1) GetTotalForumPostsOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalForumPosts) {
		return nil, false
	}
	return o.TotalForumPosts, true
}

// HasTotalForumPosts returns a boolean if a field has been set.
func (o *MiscStatsModelV1) HasTotalForumPosts() bool {
	if o != nil && !IsNil(o.TotalForumPosts) {
		return true
	}

	return false
}

// SetTotalForumPosts gets a reference to the given int64 and assigns it to the TotalForumPosts field.
func (o *MiscStatsModelV1) SetTotalForumPosts(v int64) {
	o.TotalForumPosts = &v
}

func (o MiscStatsModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MiscStatsModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalUsers) {
		toSerialize["total_users"] = o.TotalUsers
	}
	if !IsNil(o.LatestUser) {
		toSerialize["latest_user"] = o.LatestUser
	}
	if !IsNil(o.TotalForumTopics) {
		toSerialize["total_forum_topics"] = o.TotalForumTopics
	}
	if !IsNil(o.TotalForumPosts) {
		toSerialize["total_forum_posts"] = o.TotalForumPosts
	}
	return toSerialize, nil
}

type NullableMiscStatsModelV1 struct {
	value *MiscStatsModelV1
	isSet bool
}

func (v NullableMiscStatsModelV1) Get() *MiscStatsModelV1 {
	return v.value
}

func (v *NullableMiscStatsModelV1) Set(val *MiscStatsModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableMiscStatsModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableMiscStatsModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMiscStatsModelV1(val *MiscStatsModelV1) *NullableMiscStatsModelV1 {
	return &NullableMiscStatsModelV1{value: val, isSet: true}
}

func (v NullableMiscStatsModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMiscStatsModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
