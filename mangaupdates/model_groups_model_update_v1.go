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

// checks if the GroupsModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsModelUpdateV1{}

// GroupsModelUpdateV1 struct for GroupsModelUpdateV1
type GroupsModelUpdateV1 struct {
	Name *string `json:"name,omitempty"`
	Associated []GroupsModelUpdateV1Associated `json:"associated,omitempty"`
	Social *GroupsModelUpdateV1Social `json:"social,omitempty"`
	Active *bool `json:"active,omitempty"`
	Notes *string `json:"notes,omitempty"`
	Admin *GroupsModelUpdateV1Admin `json:"admin,omitempty"`
}

// NewGroupsModelUpdateV1 instantiates a new GroupsModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsModelUpdateV1() *GroupsModelUpdateV1 {
	this := GroupsModelUpdateV1{}
	return &this
}

// NewGroupsModelUpdateV1WithDefaults instantiates a new GroupsModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsModelUpdateV1WithDefaults() *GroupsModelUpdateV1 {
	this := GroupsModelUpdateV1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GroupsModelUpdateV1) SetName(v string) {
	o.Name = &v
}

// GetAssociated returns the Associated field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetAssociated() []GroupsModelUpdateV1Associated {
	if o == nil || IsNil(o.Associated) {
		var ret []GroupsModelUpdateV1Associated
		return ret
	}
	return o.Associated
}

// GetAssociatedOk returns a tuple with the Associated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetAssociatedOk() ([]GroupsModelUpdateV1Associated, bool) {
	if o == nil || IsNil(o.Associated) {
		return nil, false
	}
	return o.Associated, true
}

// HasAssociated returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasAssociated() bool {
	if o != nil && !IsNil(o.Associated) {
		return true
	}

	return false
}

// SetAssociated gets a reference to the given []GroupsModelUpdateV1Associated and assigns it to the Associated field.
func (o *GroupsModelUpdateV1) SetAssociated(v []GroupsModelUpdateV1Associated) {
	o.Associated = v
}

// GetSocial returns the Social field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetSocial() GroupsModelUpdateV1Social {
	if o == nil || IsNil(o.Social) {
		var ret GroupsModelUpdateV1Social
		return ret
	}
	return *o.Social
}

// GetSocialOk returns a tuple with the Social field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetSocialOk() (*GroupsModelUpdateV1Social, bool) {
	if o == nil || IsNil(o.Social) {
		return nil, false
	}
	return o.Social, true
}

// HasSocial returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasSocial() bool {
	if o != nil && !IsNil(o.Social) {
		return true
	}

	return false
}

// SetSocial gets a reference to the given GroupsModelUpdateV1Social and assigns it to the Social field.
func (o *GroupsModelUpdateV1) SetSocial(v GroupsModelUpdateV1Social) {
	o.Social = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *GroupsModelUpdateV1) SetActive(v bool) {
	o.Active = &v
}

// GetNotes returns the Notes field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetNotes() string {
	if o == nil || IsNil(o.Notes) {
		var ret string
		return ret
	}
	return *o.Notes
}

// GetNotesOk returns a tuple with the Notes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetNotesOk() (*string, bool) {
	if o == nil || IsNil(o.Notes) {
		return nil, false
	}
	return o.Notes, true
}

// HasNotes returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasNotes() bool {
	if o != nil && !IsNil(o.Notes) {
		return true
	}

	return false
}

// SetNotes gets a reference to the given string and assigns it to the Notes field.
func (o *GroupsModelUpdateV1) SetNotes(v string) {
	o.Notes = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *GroupsModelUpdateV1) GetAdmin() GroupsModelUpdateV1Admin {
	if o == nil || IsNil(o.Admin) {
		var ret GroupsModelUpdateV1Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelUpdateV1) GetAdminOk() (*GroupsModelUpdateV1Admin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *GroupsModelUpdateV1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given GroupsModelUpdateV1Admin and assigns it to the Admin field.
func (o *GroupsModelUpdateV1) SetAdmin(v GroupsModelUpdateV1Admin) {
	o.Admin = &v
}

func (o GroupsModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Associated) {
		toSerialize["associated"] = o.Associated
	}
	if !IsNil(o.Social) {
		toSerialize["social"] = o.Social
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Notes) {
		toSerialize["notes"] = o.Notes
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

type NullableGroupsModelUpdateV1 struct {
	value *GroupsModelUpdateV1
	isSet bool
}

func (v NullableGroupsModelUpdateV1) Get() *GroupsModelUpdateV1 {
	return v.value
}

func (v *NullableGroupsModelUpdateV1) Set(val *GroupsModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsModelUpdateV1(val *GroupsModelUpdateV1) *NullableGroupsModelUpdateV1 {
	return &NullableGroupsModelUpdateV1{value: val, isSet: true}
}

func (v NullableGroupsModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

