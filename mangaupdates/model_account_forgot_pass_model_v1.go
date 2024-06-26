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

// checks if the AccountForgotPassModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountForgotPassModelV1{}

// AccountForgotPassModelV1 struct for AccountForgotPassModelV1
type AccountForgotPassModelV1 struct {
	Email *string `json:"email,omitempty"`
}

// NewAccountForgotPassModelV1 instantiates a new AccountForgotPassModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountForgotPassModelV1() *AccountForgotPassModelV1 {
	this := AccountForgotPassModelV1{}
	return &this
}

// NewAccountForgotPassModelV1WithDefaults instantiates a new AccountForgotPassModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountForgotPassModelV1WithDefaults() *AccountForgotPassModelV1 {
	this := AccountForgotPassModelV1{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *AccountForgotPassModelV1) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountForgotPassModelV1) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *AccountForgotPassModelV1) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *AccountForgotPassModelV1) SetEmail(v string) {
	o.Email = &v
}

func (o AccountForgotPassModelV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountForgotPassModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

type NullableAccountForgotPassModelV1 struct {
	value *AccountForgotPassModelV1
	isSet bool
}

func (v NullableAccountForgotPassModelV1) Get() *AccountForgotPassModelV1 {
	return v.value
}

func (v *NullableAccountForgotPassModelV1) Set(val *AccountForgotPassModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountForgotPassModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountForgotPassModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountForgotPassModelV1(val *AccountForgotPassModelV1) *NullableAccountForgotPassModelV1 {
	return &NullableAccountForgotPassModelV1{value: val, isSet: true}
}

func (v NullableAccountForgotPassModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountForgotPassModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


