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

// checks if the ReleaseModelV1Admin type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseModelV1Admin{}

// ReleaseModelV1Admin struct for ReleaseModelV1Admin
type ReleaseModelV1Admin struct {
	Approved *bool `json:"approved,omitempty"`
	Archived *bool `json:"archived,omitempty"`
	AddedBy *UserModelSearchV1 `json:"added_by,omitempty"`
}

// NewReleaseModelV1Admin instantiates a new ReleaseModelV1Admin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseModelV1Admin() *ReleaseModelV1Admin {
	this := ReleaseModelV1Admin{}
	return &this
}

// NewReleaseModelV1AdminWithDefaults instantiates a new ReleaseModelV1Admin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseModelV1AdminWithDefaults() *ReleaseModelV1Admin {
	this := ReleaseModelV1Admin{}
	return &this
}

// GetApproved returns the Approved field value if set, zero value otherwise.
func (o *ReleaseModelV1Admin) GetApproved() bool {
	if o == nil || IsNil(o.Approved) {
		var ret bool
		return ret
	}
	return *o.Approved
}

// GetApprovedOk returns a tuple with the Approved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelV1Admin) GetApprovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Approved) {
		return nil, false
	}
	return o.Approved, true
}

// HasApproved returns a boolean if a field has been set.
func (o *ReleaseModelV1Admin) HasApproved() bool {
	if o != nil && !IsNil(o.Approved) {
		return true
	}

	return false
}

// SetApproved gets a reference to the given bool and assigns it to the Approved field.
func (o *ReleaseModelV1Admin) SetApproved(v bool) {
	o.Approved = &v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *ReleaseModelV1Admin) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelV1Admin) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *ReleaseModelV1Admin) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *ReleaseModelV1Admin) SetArchived(v bool) {
	o.Archived = &v
}

// GetAddedBy returns the AddedBy field value if set, zero value otherwise.
func (o *ReleaseModelV1Admin) GetAddedBy() UserModelSearchV1 {
	if o == nil || IsNil(o.AddedBy) {
		var ret UserModelSearchV1
		return ret
	}
	return *o.AddedBy
}

// GetAddedByOk returns a tuple with the AddedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelV1Admin) GetAddedByOk() (*UserModelSearchV1, bool) {
	if o == nil || IsNil(o.AddedBy) {
		return nil, false
	}
	return o.AddedBy, true
}

// HasAddedBy returns a boolean if a field has been set.
func (o *ReleaseModelV1Admin) HasAddedBy() bool {
	if o != nil && !IsNil(o.AddedBy) {
		return true
	}

	return false
}

// SetAddedBy gets a reference to the given UserModelSearchV1 and assigns it to the AddedBy field.
func (o *ReleaseModelV1Admin) SetAddedBy(v UserModelSearchV1) {
	o.AddedBy = &v
}

func (o ReleaseModelV1Admin) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseModelV1Admin) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Approved) {
		toSerialize["approved"] = o.Approved
	}
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.AddedBy) {
		toSerialize["added_by"] = o.AddedBy
	}
	return toSerialize, nil
}

type NullableReleaseModelV1Admin struct {
	value *ReleaseModelV1Admin
	isSet bool
}

func (v NullableReleaseModelV1Admin) Get() *ReleaseModelV1Admin {
	return v.value
}

func (v *NullableReleaseModelV1Admin) Set(val *ReleaseModelV1Admin) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseModelV1Admin) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseModelV1Admin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseModelV1Admin(val *ReleaseModelV1Admin) *NullableReleaseModelV1Admin {
	return &NullableReleaseModelV1Admin{value: val, isSet: true}
}

func (v NullableReleaseModelV1Admin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseModelV1Admin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


