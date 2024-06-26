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

// checks if the SeriesCommentModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeriesCommentModelUpdateV1{}

// SeriesCommentModelUpdateV1 struct for SeriesCommentModelUpdateV1
type SeriesCommentModelUpdateV1 struct {
	Subject *string `json:"subject,omitempty"`
	Content *string `json:"content,omitempty"`
	Admin *SeriesCommentModelUpdateV1Admin `json:"admin,omitempty"`
}

// NewSeriesCommentModelUpdateV1 instantiates a new SeriesCommentModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeriesCommentModelUpdateV1() *SeriesCommentModelUpdateV1 {
	this := SeriesCommentModelUpdateV1{}
	return &this
}

// NewSeriesCommentModelUpdateV1WithDefaults instantiates a new SeriesCommentModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeriesCommentModelUpdateV1WithDefaults() *SeriesCommentModelUpdateV1 {
	this := SeriesCommentModelUpdateV1{}
	return &this
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *SeriesCommentModelUpdateV1) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelUpdateV1) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *SeriesCommentModelUpdateV1) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *SeriesCommentModelUpdateV1) SetSubject(v string) {
	o.Subject = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SeriesCommentModelUpdateV1) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelUpdateV1) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SeriesCommentModelUpdateV1) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *SeriesCommentModelUpdateV1) SetContent(v string) {
	o.Content = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *SeriesCommentModelUpdateV1) GetAdmin() SeriesCommentModelUpdateV1Admin {
	if o == nil || IsNil(o.Admin) {
		var ret SeriesCommentModelUpdateV1Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeriesCommentModelUpdateV1) GetAdminOk() (*SeriesCommentModelUpdateV1Admin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *SeriesCommentModelUpdateV1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given SeriesCommentModelUpdateV1Admin and assigns it to the Admin field.
func (o *SeriesCommentModelUpdateV1) SetAdmin(v SeriesCommentModelUpdateV1Admin) {
	o.Admin = &v
}

func (o SeriesCommentModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeriesCommentModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

type NullableSeriesCommentModelUpdateV1 struct {
	value *SeriesCommentModelUpdateV1
	isSet bool
}

func (v NullableSeriesCommentModelUpdateV1) Get() *SeriesCommentModelUpdateV1 {
	return v.value
}

func (v *NullableSeriesCommentModelUpdateV1) Set(val *SeriesCommentModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableSeriesCommentModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableSeriesCommentModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeriesCommentModelUpdateV1(val *SeriesCommentModelUpdateV1) *NullableSeriesCommentModelUpdateV1 {
	return &NullableSeriesCommentModelUpdateV1{value: val, isSet: true}
}

func (v NullableSeriesCommentModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeriesCommentModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


