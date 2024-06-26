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

// checks if the ForumSearchResponseV1TopicResultsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumSearchResponseV1TopicResultsMetadata{}

// ForumSearchResponseV1TopicResultsMetadata struct for ForumSearchResponseV1TopicResultsMetadata
type ForumSearchResponseV1TopicResultsMetadata struct {
	IsSubscribed        *bool  `json:"is_subscribed,omitempty"`
	MyLatestPostInTopic *int64 `json:"my_latest_post_in_topic,omitempty"`
}

// NewForumSearchResponseV1TopicResultsMetadata instantiates a new ForumSearchResponseV1TopicResultsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumSearchResponseV1TopicResultsMetadata() *ForumSearchResponseV1TopicResultsMetadata {
	this := ForumSearchResponseV1TopicResultsMetadata{}
	return &this
}

// NewForumSearchResponseV1TopicResultsMetadataWithDefaults instantiates a new ForumSearchResponseV1TopicResultsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumSearchResponseV1TopicResultsMetadataWithDefaults() *ForumSearchResponseV1TopicResultsMetadata {
	this := ForumSearchResponseV1TopicResultsMetadata{}
	return &this
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise.
func (o *ForumSearchResponseV1TopicResultsMetadata) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumSearchResponseV1TopicResultsMetadata) GetIsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribed) {
		return nil, false
	}
	return o.IsSubscribed, true
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *ForumSearchResponseV1TopicResultsMetadata) HasIsSubscribed() bool {
	if o != nil && !IsNil(o.IsSubscribed) {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given bool and assigns it to the IsSubscribed field.
func (o *ForumSearchResponseV1TopicResultsMetadata) SetIsSubscribed(v bool) {
	o.IsSubscribed = &v
}

// GetMyLatestPostInTopic returns the MyLatestPostInTopic field value if set, zero value otherwise.
func (o *ForumSearchResponseV1TopicResultsMetadata) GetMyLatestPostInTopic() int64 {
	if o == nil || IsNil(o.MyLatestPostInTopic) {
		var ret int64
		return ret
	}
	return *o.MyLatestPostInTopic
}

// GetMyLatestPostInTopicOk returns a tuple with the MyLatestPostInTopic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumSearchResponseV1TopicResultsMetadata) GetMyLatestPostInTopicOk() (*int64, bool) {
	if o == nil || IsNil(o.MyLatestPostInTopic) {
		return nil, false
	}
	return o.MyLatestPostInTopic, true
}

// HasMyLatestPostInTopic returns a boolean if a field has been set.
func (o *ForumSearchResponseV1TopicResultsMetadata) HasMyLatestPostInTopic() bool {
	if o != nil && !IsNil(o.MyLatestPostInTopic) {
		return true
	}

	return false
}

// SetMyLatestPostInTopic gets a reference to the given int64 and assigns it to the MyLatestPostInTopic field.
func (o *ForumSearchResponseV1TopicResultsMetadata) SetMyLatestPostInTopic(v int64) {
	o.MyLatestPostInTopic = &v
}

func (o ForumSearchResponseV1TopicResultsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumSearchResponseV1TopicResultsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IsSubscribed) {
		toSerialize["is_subscribed"] = o.IsSubscribed
	}
	if !IsNil(o.MyLatestPostInTopic) {
		toSerialize["my_latest_post_in_topic"] = o.MyLatestPostInTopic
	}
	return toSerialize, nil
}

type NullableForumSearchResponseV1TopicResultsMetadata struct {
	value *ForumSearchResponseV1TopicResultsMetadata
	isSet bool
}

func (v NullableForumSearchResponseV1TopicResultsMetadata) Get() *ForumSearchResponseV1TopicResultsMetadata {
	return v.value
}

func (v *NullableForumSearchResponseV1TopicResultsMetadata) Set(val *ForumSearchResponseV1TopicResultsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableForumSearchResponseV1TopicResultsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableForumSearchResponseV1TopicResultsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumSearchResponseV1TopicResultsMetadata(val *ForumSearchResponseV1TopicResultsMetadata) *NullableForumSearchResponseV1TopicResultsMetadata {
	return &NullableForumSearchResponseV1TopicResultsMetadata{value: val, isSet: true}
}

func (v NullableForumSearchResponseV1TopicResultsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumSearchResponseV1TopicResultsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
