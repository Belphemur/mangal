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

// checks if the ConvoModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvoModelV1{}

// ConvoModelV1 struct for ConvoModelV1
type ConvoModelV1 struct {
	ConvoId    *int64  `json:"convo_id,omitempty"`
	Topic      *string `json:"topic,omitempty"`
	AuthorId   *int64  `json:"author_id,omitempty"`
	AuthorName *string `json:"author_name,omitempty"`
	TimeAdded  *TimeV1 `json:"time_added,omitempty"`
	LastEdit   *TimeV1 `json:"last_edit,omitempty"`
}

// NewConvoModelV1 instantiates a new ConvoModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvoModelV1() *ConvoModelV1 {
	this := ConvoModelV1{}
	return &this
}

// NewConvoModelV1WithDefaults instantiates a new ConvoModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvoModelV1WithDefaults() *ConvoModelV1 {
	this := ConvoModelV1{}
	return &this
}

// GetConvoId returns the ConvoId field value if set, zero value otherwise.
func (o *ConvoModelV1) GetConvoId() int64 {
	if o == nil || IsNil(o.ConvoId) {
		var ret int64
		return ret
	}
	return *o.ConvoId
}

// GetConvoIdOk returns a tuple with the ConvoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetConvoIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ConvoId) {
		return nil, false
	}
	return o.ConvoId, true
}

// HasConvoId returns a boolean if a field has been set.
func (o *ConvoModelV1) HasConvoId() bool {
	if o != nil && !IsNil(o.ConvoId) {
		return true
	}

	return false
}

// SetConvoId gets a reference to the given int64 and assigns it to the ConvoId field.
func (o *ConvoModelV1) SetConvoId(v int64) {
	o.ConvoId = &v
}

// GetTopic returns the Topic field value if set, zero value otherwise.
func (o *ConvoModelV1) GetTopic() string {
	if o == nil || IsNil(o.Topic) {
		var ret string
		return ret
	}
	return *o.Topic
}

// GetTopicOk returns a tuple with the Topic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetTopicOk() (*string, bool) {
	if o == nil || IsNil(o.Topic) {
		return nil, false
	}
	return o.Topic, true
}

// HasTopic returns a boolean if a field has been set.
func (o *ConvoModelV1) HasTopic() bool {
	if o != nil && !IsNil(o.Topic) {
		return true
	}

	return false
}

// SetTopic gets a reference to the given string and assigns it to the Topic field.
func (o *ConvoModelV1) SetTopic(v string) {
	o.Topic = &v
}

// GetAuthorId returns the AuthorId field value if set, zero value otherwise.
func (o *ConvoModelV1) GetAuthorId() int64 {
	if o == nil || IsNil(o.AuthorId) {
		var ret int64
		return ret
	}
	return *o.AuthorId
}

// GetAuthorIdOk returns a tuple with the AuthorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetAuthorIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AuthorId) {
		return nil, false
	}
	return o.AuthorId, true
}

// HasAuthorId returns a boolean if a field has been set.
func (o *ConvoModelV1) HasAuthorId() bool {
	if o != nil && !IsNil(o.AuthorId) {
		return true
	}

	return false
}

// SetAuthorId gets a reference to the given int64 and assigns it to the AuthorId field.
func (o *ConvoModelV1) SetAuthorId(v int64) {
	o.AuthorId = &v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *ConvoModelV1) GetAuthorName() string {
	if o == nil || IsNil(o.AuthorName) {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetAuthorNameOk() (*string, bool) {
	if o == nil || IsNil(o.AuthorName) {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *ConvoModelV1) HasAuthorName() bool {
	if o != nil && !IsNil(o.AuthorName) {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *ConvoModelV1) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *ConvoModelV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *ConvoModelV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *ConvoModelV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

// GetLastEdit returns the LastEdit field value if set, zero value otherwise.
func (o *ConvoModelV1) GetLastEdit() TimeV1 {
	if o == nil || IsNil(o.LastEdit) {
		var ret TimeV1
		return ret
	}
	return *o.LastEdit
}

// GetLastEditOk returns a tuple with the LastEdit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoModelV1) GetLastEditOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.LastEdit) {
		return nil, false
	}
	return o.LastEdit, true
}

// HasLastEdit returns a boolean if a field has been set.
func (o *ConvoModelV1) HasLastEdit() bool {
	if o != nil && !IsNil(o.LastEdit) {
		return true
	}

	return false
}

// SetLastEdit gets a reference to the given TimeV1 and assigns it to the LastEdit field.
func (o *ConvoModelV1) SetLastEdit(v TimeV1) {
	o.LastEdit = &v
}

func (o ConvoModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvoModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ConvoId) {
		toSerialize["convo_id"] = o.ConvoId
	}
	if !IsNil(o.Topic) {
		toSerialize["topic"] = o.Topic
	}
	if !IsNil(o.AuthorId) {
		toSerialize["author_id"] = o.AuthorId
	}
	if !IsNil(o.AuthorName) {
		toSerialize["author_name"] = o.AuthorName
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	if !IsNil(o.LastEdit) {
		toSerialize["last_edit"] = o.LastEdit
	}
	return toSerialize, nil
}

type NullableConvoModelV1 struct {
	value *ConvoModelV1
	isSet bool
}

func (v NullableConvoModelV1) Get() *ConvoModelV1 {
	return v.value
}

func (v *NullableConvoModelV1) Set(val *ConvoModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableConvoModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableConvoModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvoModelV1(val *ConvoModelV1) *NullableConvoModelV1 {
	return &NullableConvoModelV1{value: val, isSet: true}
}

func (v NullableConvoModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvoModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
