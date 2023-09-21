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

// checks if the ForumPostReportModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumPostReportModelV1{}

// ForumPostReportModelV1 struct for ForumPostReportModelV1
type ForumPostReportModelV1 struct {
	ReportId *int64                   `json:"report_id,omitempty"`
	TopicId  *int64                   `json:"topic_id,omitempty"`
	Topic    *ForumTopicModelSearchV1 `json:"topic,omitempty"`
	PostId   *int64                   `json:"post_id,omitempty"`
	Post     *ForumPostModelSearchV1  `json:"post,omitempty"`
	UserId   *int64                   `json:"user_id,omitempty"`
	User     *UserModelSearchV1       `json:"user,omitempty"`
	Reason   *string                  `json:"reason,omitempty"`
}

// NewForumPostReportModelV1 instantiates a new ForumPostReportModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumPostReportModelV1() *ForumPostReportModelV1 {
	this := ForumPostReportModelV1{}
	return &this
}

// NewForumPostReportModelV1WithDefaults instantiates a new ForumPostReportModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumPostReportModelV1WithDefaults() *ForumPostReportModelV1 {
	this := ForumPostReportModelV1{}
	return &this
}

// GetReportId returns the ReportId field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetReportId() int64 {
	if o == nil || IsNil(o.ReportId) {
		var ret int64
		return ret
	}
	return *o.ReportId
}

// GetReportIdOk returns a tuple with the ReportId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetReportIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ReportId) {
		return nil, false
	}
	return o.ReportId, true
}

// HasReportId returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasReportId() bool {
	if o != nil && !IsNil(o.ReportId) {
		return true
	}

	return false
}

// SetReportId gets a reference to the given int64 and assigns it to the ReportId field.
func (o *ForumPostReportModelV1) SetReportId(v int64) {
	o.ReportId = &v
}

// GetTopicId returns the TopicId field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetTopicId() int64 {
	if o == nil || IsNil(o.TopicId) {
		var ret int64
		return ret
	}
	return *o.TopicId
}

// GetTopicIdOk returns a tuple with the TopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetTopicIdOk() (*int64, bool) {
	if o == nil || IsNil(o.TopicId) {
		return nil, false
	}
	return o.TopicId, true
}

// HasTopicId returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasTopicId() bool {
	if o != nil && !IsNil(o.TopicId) {
		return true
	}

	return false
}

// SetTopicId gets a reference to the given int64 and assigns it to the TopicId field.
func (o *ForumPostReportModelV1) SetTopicId(v int64) {
	o.TopicId = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetTopic() ForumTopicModelSearchV1 {
	if o == nil || IsNil(o.Topic) {
		var ret ForumTopicModelSearchV1
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetTopicOk() (*ForumTopicModelSearchV1, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given ForumTopicModelSearchV1 and assigns it to the Topic field.
func (o *ForumPostReportModelV1) SetTopic(v ForumTopicModelSearchV1) {
	o.Topic = &v
}

// GetPostId returns the PostId field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetPostId() int64 {
	if o == nil || IsNil(o.PostId) {
		var ret int64
		return ret
	}
	return *o.PostId
}

// GetPostIdOk returns a tuple with the PostId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetPostIdOk() (*int64, bool) {
	if o == nil || IsNil(o.PostId) {
		return nil, false
	}
	return o.PostId, true
}

// HasPostId returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasPostId() bool {
	if o != nil && !IsNil(o.PostId) {
		return true
	}

	return false
}

// SetPostId gets a reference to the given int64 and assigns it to the PostId field.
func (o *ForumPostReportModelV1) SetPostId(v int64) {
	o.PostId = &v
}

// GetPost returns the Post field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetPost() ForumPostModelSearchV1 {
	if o == nil || IsNil(o.Post) {
		var ret ForumPostModelSearchV1
		return ret
	}
	return *o.Post
}

// GetPostOk returns a tuple with the Post field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetPostOk() (*ForumPostModelSearchV1, bool) {
	if o == nil || IsNil(o.Post) {
		return nil, false
	}
	return o.Post, true
}

// HasPost returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasPost() bool {
	if o != nil && !IsNil(o.Post) {
		return true
	}

	return false
}

// SetPost gets a reference to the given ForumPostModelSearchV1 and assigns it to the Post field.
func (o *ForumPostReportModelV1) SetPost(v ForumPostModelSearchV1) {
	o.Post = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetUserId() int64 {
	if o == nil || IsNil(o.UserId) {
		var ret int64
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetUserIdOk() (*int64, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int64 and assigns it to the UserId field.
func (o *ForumPostReportModelV1) SetUserId(v int64) {
	o.UserId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetUser() UserModelSearchV1 {
	if o == nil || IsNil(o.User) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetUserOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given UserModelSearchV1 and assigns it to the User field.
func (o *ForumPostReportModelV1) SetUser(v UserModelSearchV1) {
	o.User = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ForumPostReportModelV1) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumPostReportModelV1) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ForumPostReportModelV1) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ForumPostReportModelV1) SetReason(v string) {
	o.Reason = &v
}

func (o ForumPostReportModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumPostReportModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReportId) {
		toSerialize["report_id"] = o.ReportId
	}
	if !IsNil(o.TopicId) {
		toSerialize["topic_id"] = o.TopicId
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.PostId) {
		toSerialize["post_id"] = o.PostId
	}
	if !IsNil(o.Post) {
		toSerialize["post"] = o.Post
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableForumPostReportModelV1 struct {
	value *ForumPostReportModelV1
	isSet bool
}

func (v NullableForumPostReportModelV1) Get() *ForumPostReportModelV1 {
	return v.value
}

func (v *NullableForumPostReportModelV1) Set(val *ForumPostReportModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumPostReportModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumPostReportModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumPostReportModelV1(val *ForumPostReportModelV1) *NullableForumPostReportModelV1 {
	return &NullableForumPostReportModelV1{value: val, isSet: true}
}

func (v NullableForumPostReportModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumPostReportModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}