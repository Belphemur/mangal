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

// checks if the ImageModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageModelV1{}

// ImageModelV1 struct for ImageModelV1
type ImageModelV1 struct {
	Url    *ImageModelV1Url `json:"url,omitempty"`
	Height *int64           `json:"height,omitempty"`
	Width  *int64           `json:"width,omitempty"`
}

// NewImageModelV1 instantiates a new ImageModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageModelV1() *ImageModelV1 {
	this := ImageModelV1{}
	return &this
}

// NewImageModelV1WithDefaults instantiates a new ImageModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageModelV1WithDefaults() *ImageModelV1 {
	this := ImageModelV1{}
	return &this
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *ImageModelV1) GetUrl() ImageModelV1Url {
	if o == nil || IsNil(o.Url) {
		var ret ImageModelV1Url
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageModelV1) GetUrlOk() (*ImageModelV1Url, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *ImageModelV1) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given ImageModelV1Url and assigns it to the Url field.
func (o *ImageModelV1) SetUrl(v ImageModelV1Url) {
	o.Url = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ImageModelV1) GetHeight() int64 {
	if o == nil || IsNil(o.Height) {
		var ret int64
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageModelV1) GetHeightOk() (*int64, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ImageModelV1) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int64 and assigns it to the Height field.
func (o *ImageModelV1) SetHeight(v int64) {
	o.Height = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ImageModelV1) GetWidth() int64 {
	if o == nil || IsNil(o.Width) {
		var ret int64
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageModelV1) GetWidthOk() (*int64, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ImageModelV1) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int64 and assigns it to the Width field.
func (o *ImageModelV1) SetWidth(v int64) {
	o.Width = &v
}

func (o ImageModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableImageModelV1 struct {
	value *ImageModelV1
	isSet bool
}

func (v NullableImageModelV1) Get() *ImageModelV1 {
	return v.value
}

func (v *NullableImageModelV1) Set(val *ImageModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableImageModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableImageModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageModelV1(val *ImageModelV1) *NullableImageModelV1 {
	return &NullableImageModelV1{value: val, isSet: true}
}

func (v NullableImageModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
