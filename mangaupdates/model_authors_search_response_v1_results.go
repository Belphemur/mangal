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

// checks if the AuthorsSearchResponseV1Results type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorsSearchResponseV1Results{}

// AuthorsSearchResponseV1Results struct for AuthorsSearchResponseV1Results
type AuthorsSearchResponseV1Results struct {
	Record *AuthorsModelSearchV1 `json:"record,omitempty"`
	HitName *string `json:"hit_name,omitempty"`
	HitGenre []string `json:"hit_genre,omitempty"`
}

// NewAuthorsSearchResponseV1Results instantiates a new AuthorsSearchResponseV1Results object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorsSearchResponseV1Results() *AuthorsSearchResponseV1Results {
	this := AuthorsSearchResponseV1Results{}
	return &this
}

// NewAuthorsSearchResponseV1ResultsWithDefaults instantiates a new AuthorsSearchResponseV1Results object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorsSearchResponseV1ResultsWithDefaults() *AuthorsSearchResponseV1Results {
	this := AuthorsSearchResponseV1Results{}
	return &this
}

// GetRecord returns the Record field value if set, zero value otherwise.
func (o *AuthorsSearchResponseV1Results) GetRecord() AuthorsModelSearchV1 {
	if o == nil || IsNil(o.Record) {
		var ret AuthorsModelSearchV1
		return ret
	}
	return *o.Record
}

// GetRecordOk returns a tuple with the Record field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSearchResponseV1Results) GetRecordOk() (*AuthorsModelSearchV1, bool) {
	if o == nil || IsNil(o.Record) {
		return nil, false
	}
	return o.Record, true
}

// HasRecord returns a boolean if a field has been set.
func (o *AuthorsSearchResponseV1Results) HasRecord() bool {
	if o != nil && !IsNil(o.Record) {
		return true
	}

	return false
}

// SetRecord gets a reference to the given AuthorsModelSearchV1 and assigns it to the Record field.
func (o *AuthorsSearchResponseV1Results) SetRecord(v AuthorsModelSearchV1) {
	o.Record = &v
}

// GetHitName returns the HitName field value if set, zero value otherwise.
func (o *AuthorsSearchResponseV1Results) GetHitName() string {
	if o == nil || IsNil(o.HitName) {
		var ret string
		return ret
	}
	return *o.HitName
}

// GetHitNameOk returns a tuple with the HitName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSearchResponseV1Results) GetHitNameOk() (*string, bool) {
	if o == nil || IsNil(o.HitName) {
		return nil, false
	}
	return o.HitName, true
}

// HasHitName returns a boolean if a field has been set.
func (o *AuthorsSearchResponseV1Results) HasHitName() bool {
	if o != nil && !IsNil(o.HitName) {
		return true
	}

	return false
}

// SetHitName gets a reference to the given string and assigns it to the HitName field.
func (o *AuthorsSearchResponseV1Results) SetHitName(v string) {
	o.HitName = &v
}

// GetHitGenre returns the HitGenre field value if set, zero value otherwise.
func (o *AuthorsSearchResponseV1Results) GetHitGenre() []string {
	if o == nil || IsNil(o.HitGenre) {
		var ret []string
		return ret
	}
	return o.HitGenre
}

// GetHitGenreOk returns a tuple with the HitGenre field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorsSearchResponseV1Results) GetHitGenreOk() ([]string, bool) {
	if o == nil || IsNil(o.HitGenre) {
		return nil, false
	}
	return o.HitGenre, true
}

// HasHitGenre returns a boolean if a field has been set.
func (o *AuthorsSearchResponseV1Results) HasHitGenre() bool {
	if o != nil && !IsNil(o.HitGenre) {
		return true
	}

	return false
}

// SetHitGenre gets a reference to the given []string and assigns it to the HitGenre field.
func (o *AuthorsSearchResponseV1Results) SetHitGenre(v []string) {
	o.HitGenre = v
}

func (o AuthorsSearchResponseV1Results) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorsSearchResponseV1Results) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Record) {
		toSerialize["record"] = o.Record
	}
	if !IsNil(o.HitName) {
		toSerialize["hit_name"] = o.HitName
	}
	if !IsNil(o.HitGenre) {
		toSerialize["hit_genre"] = o.HitGenre
	}
	return toSerialize, nil
}

type NullableAuthorsSearchResponseV1Results struct {
	value *AuthorsSearchResponseV1Results
	isSet bool
}

func (v NullableAuthorsSearchResponseV1Results) Get() *AuthorsSearchResponseV1Results {
	return v.value
}

func (v *NullableAuthorsSearchResponseV1Results) Set(val *AuthorsSearchResponseV1Results) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorsSearchResponseV1Results) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorsSearchResponseV1Results) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorsSearchResponseV1Results(val *AuthorsSearchResponseV1Results) *NullableAuthorsSearchResponseV1Results {
	return &NullableAuthorsSearchResponseV1Results{value: val, isSet: true}
}

func (v NullableAuthorsSearchResponseV1Results) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorsSearchResponseV1Results) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

