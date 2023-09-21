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

// checks if the ForumTopicModelSearchV1Forum type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumTopicModelSearchV1Forum{}

// ForumTopicModelSearchV1Forum struct for ForumTopicModelSearchV1Forum
type ForumTopicModelSearchV1Forum struct {
	ForumId   *int64  `json:"forum_id,omitempty"`
	ForumName *string `json:"forum_name,omitempty"`
}

// NewForumTopicModelSearchV1Forum instantiates a new ForumTopicModelSearchV1Forum object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumTopicModelSearchV1Forum() *ForumTopicModelSearchV1Forum {
	this := ForumTopicModelSearchV1Forum{}
	return &this
}

// NewForumTopicModelSearchV1ForumWithDefaults instantiates a new ForumTopicModelSearchV1Forum object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumTopicModelSearchV1ForumWithDefaults() *ForumTopicModelSearchV1Forum {
	this := ForumTopicModelSearchV1Forum{}
	return &this
}

// GetForumId returns the ForumId field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1Forum) GetForumId() int64 {
	if o == nil || IsNil(o.ForumId) {
		var ret int64
		return ret
	}
	return *o.ForumId
}

// GetForumIdOk returns a tuple with the ForumId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1Forum) GetForumIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ForumId) {
		return nil, false
	}
	return o.ForumId, true
}

// HasForumId returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1Forum) HasForumId() bool {
	if o != nil && !IsNil(o.ForumId) {
		return true
	}

	return false
}

// SetForumId gets a reference to the given int64 and assigns it to the ForumId field.
func (o *ForumTopicModelSearchV1Forum) SetForumId(v int64) {
	o.ForumId = &v
}

// GetForumName returns the ForumName field value if set, zero value otherwise.
func (o *ForumTopicModelSearchV1Forum) GetForumName() string {
	if o == nil || IsNil(o.ForumName) {
		var ret string
		return ret
	}
	return *o.ForumName
}

// GetForumNameOk returns a tuple with the ForumName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumTopicModelSearchV1Forum) GetForumNameOk() (*string, bool) {
	if o == nil || IsNil(o.ForumName) {
		return nil, false
	}
	return o.ForumName, true
}

// HasForumName returns a boolean if a field has been set.
func (o *ForumTopicModelSearchV1Forum) HasForumName() bool {
	if o != nil && !IsNil(o.ForumName) {
		return true
	}

	return false
}

// SetForumName gets a reference to the given string and assigns it to the ForumName field.
func (o *ForumTopicModelSearchV1Forum) SetForumName(v string) {
	o.ForumName = &v
}

func (o ForumTopicModelSearchV1Forum) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumTopicModelSearchV1Forum) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ForumId) {
		toSerialize["forum_id"] = o.ForumId
	}
	if !IsNil(o.ForumName) {
		toSerialize["forum_name"] = o.ForumName
	}
	return toSerialize, nil
}

type NullableForumTopicModelSearchV1Forum struct {
	value *ForumTopicModelSearchV1Forum
	isSet bool
}

func (v NullableForumTopicModelSearchV1Forum) Get() *ForumTopicModelSearchV1Forum {
	return v.value
}

func (v *NullableForumTopicModelSearchV1Forum) Set(val *ForumTopicModelSearchV1Forum) {
	v.value = val
	v.isSet = true
}

func (v NullableForumTopicModelSearchV1Forum) IsSet() bool {
	return v.isSet
}

func (v *NullableForumTopicModelSearchV1Forum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumTopicModelSearchV1Forum(val *ForumTopicModelSearchV1Forum) *NullableForumTopicModelSearchV1Forum {
	return &NullableForumTopicModelSearchV1Forum{value: val, isSet: true}
}

func (v NullableForumTopicModelSearchV1Forum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumTopicModelSearchV1Forum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}