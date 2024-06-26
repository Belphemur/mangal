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

// checks if the GroupsModelSearchV1SocialIrc type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsModelSearchV1SocialIrc{}

// GroupsModelSearchV1SocialIrc struct for GroupsModelSearchV1SocialIrc
type GroupsModelSearchV1SocialIrc struct {
	Channel *string `json:"channel,omitempty"`
	Server *string `json:"server,omitempty"`
}

// NewGroupsModelSearchV1SocialIrc instantiates a new GroupsModelSearchV1SocialIrc object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsModelSearchV1SocialIrc() *GroupsModelSearchV1SocialIrc {
	this := GroupsModelSearchV1SocialIrc{}
	return &this
}

// NewGroupsModelSearchV1SocialIrcWithDefaults instantiates a new GroupsModelSearchV1SocialIrc object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsModelSearchV1SocialIrcWithDefaults() *GroupsModelSearchV1SocialIrc {
	this := GroupsModelSearchV1SocialIrc{}
	return &this
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *GroupsModelSearchV1SocialIrc) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1SocialIrc) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *GroupsModelSearchV1SocialIrc) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *GroupsModelSearchV1SocialIrc) SetChannel(v string) {
	o.Channel = &v
}

// GetServer returns the Server field value if set, zero value otherwise.
func (o *GroupsModelSearchV1SocialIrc) GetServer() string {
	if o == nil || IsNil(o.Server) {
		var ret string
		return ret
	}
	return *o.Server
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1SocialIrc) GetServerOk() (*string, bool) {
	if o == nil || IsNil(o.Server) {
		return nil, false
	}
	return o.Server, true
}

// HasServer returns a boolean if a field has been set.
func (o *GroupsModelSearchV1SocialIrc) HasServer() bool {
	if o != nil && !IsNil(o.Server) {
		return true
	}

	return false
}

// SetServer gets a reference to the given string and assigns it to the Server field.
func (o *GroupsModelSearchV1SocialIrc) SetServer(v string) {
	o.Server = &v
}

func (o GroupsModelSearchV1SocialIrc) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsModelSearchV1SocialIrc) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Server) {
		toSerialize["server"] = o.Server
	}
	return toSerialize, nil
}

type NullableGroupsModelSearchV1SocialIrc struct {
	value *GroupsModelSearchV1SocialIrc
	isSet bool
}

func (v NullableGroupsModelSearchV1SocialIrc) Get() *GroupsModelSearchV1SocialIrc {
	return v.value
}

func (v *NullableGroupsModelSearchV1SocialIrc) Set(val *GroupsModelSearchV1SocialIrc) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsModelSearchV1SocialIrc) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsModelSearchV1SocialIrc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsModelSearchV1SocialIrc(val *GroupsModelSearchV1SocialIrc) *NullableGroupsModelSearchV1SocialIrc {
	return &NullableGroupsModelSearchV1SocialIrc{value: val, isSet: true}
}

func (v NullableGroupsModelSearchV1SocialIrc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsModelSearchV1SocialIrc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


