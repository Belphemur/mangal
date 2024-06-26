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

// checks if the ForumForumModelV1Category type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForumForumModelV1Category{}

// ForumForumModelV1Category struct for ForumForumModelV1Category
type ForumForumModelV1Category struct {
	CategoryId   *int64  `json:"category_id,omitempty"`
	CategoryName *string `json:"category_name,omitempty"`
}

// NewForumForumModelV1Category instantiates a new ForumForumModelV1Category object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForumForumModelV1Category() *ForumForumModelV1Category {
	this := ForumForumModelV1Category{}
	return &this
}

// NewForumForumModelV1CategoryWithDefaults instantiates a new ForumForumModelV1Category object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForumForumModelV1CategoryWithDefaults() *ForumForumModelV1Category {
	this := ForumForumModelV1Category{}
	return &this
}

// GetCategoryId returns the CategoryId field value if set, zero value otherwise.
func (o *ForumForumModelV1Category) GetCategoryId() int64 {
	if o == nil || IsNil(o.CategoryId) {
		var ret int64
		return ret
	}
	return *o.CategoryId
}

// GetCategoryIdOk returns a tuple with the CategoryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumForumModelV1Category) GetCategoryIdOk() (*int64, bool) {
	if o == nil || IsNil(o.CategoryId) {
		return nil, false
	}
	return o.CategoryId, true
}

// HasCategoryId returns a boolean if a field has been set.
func (o *ForumForumModelV1Category) HasCategoryId() bool {
	if o != nil && !IsNil(o.CategoryId) {
		return true
	}

	return false
}

// SetCategoryId gets a reference to the given int64 and assigns it to the CategoryId field.
func (o *ForumForumModelV1Category) SetCategoryId(v int64) {
	o.CategoryId = &v
}

// GetCategoryName returns the CategoryName field value if set, zero value otherwise.
func (o *ForumForumModelV1Category) GetCategoryName() string {
	if o == nil || IsNil(o.CategoryName) {
		var ret string
		return ret
	}
	return *o.CategoryName
}

// GetCategoryNameOk returns a tuple with the CategoryName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForumForumModelV1Category) GetCategoryNameOk() (*string, bool) {
	if o == nil || IsNil(o.CategoryName) {
		return nil, false
	}
	return o.CategoryName, true
}

// HasCategoryName returns a boolean if a field has been set.
func (o *ForumForumModelV1Category) HasCategoryName() bool {
	if o != nil && !IsNil(o.CategoryName) {
		return true
	}

	return false
}

// SetCategoryName gets a reference to the given string and assigns it to the CategoryName field.
func (o *ForumForumModelV1Category) SetCategoryName(v string) {
	o.CategoryName = &v
}

func (o ForumForumModelV1Category) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForumForumModelV1Category) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CategoryId) {
		toSerialize["category_id"] = o.CategoryId
	}
	if !IsNil(o.CategoryName) {
		toSerialize["category_name"] = o.CategoryName
	}
	return toSerialize, nil
}

type NullableForumForumModelV1Category struct {
	value *ForumForumModelV1Category
	isSet bool
}

func (v NullableForumForumModelV1Category) Get() *ForumForumModelV1Category {
	return v.value
}

func (v *NullableForumForumModelV1Category) Set(val *ForumForumModelV1Category) {
	v.value = val
	v.isSet = true
}

func (v NullableForumForumModelV1Category) IsSet() bool {
	return v.isSet
}

func (v *NullableForumForumModelV1Category) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForumForumModelV1Category(val *ForumForumModelV1Category) *NullableForumForumModelV1Category {
	return &NullableForumForumModelV1Category{value: val, isSet: true}
}

func (v NullableForumForumModelV1Category) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForumForumModelV1Category) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
