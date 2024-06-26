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

// checks if the ForumCategoryModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumCategoryModelUpdateV1{}

// ForumCategoryModelUpdateV1 struct for ForumCategoryModelUpdateV1
type ForumCategoryModelUpdateV1 struct {
	Name     *string `json:"name,omitempty"`
	Position *int64  `json:"position,omitempty"`
}

// NewForumCategoryModelUpdateV1 instantiates a new ForumCategoryModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumCategoryModelUpdateV1() *ForumCategoryModelUpdateV1 {
	this := ForumCategoryModelUpdateV1{}
	return &this
}

// NewForumCategoryModelUpdateV1WithDefaults instantiates a new ForumCategoryModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumCategoryModelUpdateV1WithDefaults() *ForumCategoryModelUpdateV1 {
	this := ForumCategoryModelUpdateV1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ForumCategoryModelUpdateV1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumCategoryModelUpdateV1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ForumCategoryModelUpdateV1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ForumCategoryModelUpdateV1) SetName(v string) {
	o.Name = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ForumCategoryModelUpdateV1) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumCategoryModelUpdateV1) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ForumCategoryModelUpdateV1) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *ForumCategoryModelUpdateV1) SetPosition(v int64) {
	o.Position = &v
}

func (o ForumCategoryModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumCategoryModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableForumCategoryModelUpdateV1 struct {
	value *ForumCategoryModelUpdateV1
	isSet bool
}

func (v NullableForumCategoryModelUpdateV1) Get() *ForumCategoryModelUpdateV1 {
	return v.value
}

func (v *NullableForumCategoryModelUpdateV1) Set(val *ForumCategoryModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableForumCategoryModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableForumCategoryModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumCategoryModelUpdateV1(val *ForumCategoryModelUpdateV1) *NullableForumCategoryModelUpdateV1 {
	return &NullableForumCategoryModelUpdateV1{value: val, isSet: true}
}

func (v NullableForumCategoryModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumCategoryModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
