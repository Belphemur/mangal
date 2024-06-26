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

// checks if the PublishersModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishersModelUpdateV1{}

// PublishersModelUpdateV1 struct for PublishersModelUpdateV1
type PublishersModelUpdateV1 struct {
	Name *string `json:"name,omitempty"`
	Associated []PublishersModelUpdateV1Associated `json:"associated,omitempty"`
	Type *string `json:"type,omitempty"`
	Info *string `json:"info,omitempty"`
	Site *string `json:"site,omitempty"`
	Admin *PublishersModelUpdateV1Admin `json:"admin,omitempty"`
}

// NewPublishersModelUpdateV1 instantiates a new PublishersModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishersModelUpdateV1() *PublishersModelUpdateV1 {
	this := PublishersModelUpdateV1{}
	return &this
}

// NewPublishersModelUpdateV1WithDefaults instantiates a new PublishersModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishersModelUpdateV1WithDefaults() *PublishersModelUpdateV1 {
	this := PublishersModelUpdateV1{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublishersModelUpdateV1) SetName(v string) {
	o.Name = &v
}

// GetAssociated returns the Associated field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetAssociated() []PublishersModelUpdateV1Associated {
	if o == nil || IsNil(o.Associated) {
		var ret []PublishersModelUpdateV1Associated
		return ret
	}
	return o.Associated
}

// GetAssociatedOk returns a tuple with the Associated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetAssociatedOk() ([]PublishersModelUpdateV1Associated, bool) {
	if o == nil || IsNil(o.Associated) {
		return nil, false
	}
	return o.Associated, true
}

// HasAssociated returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasAssociated() bool {
	if o != nil && !IsNil(o.Associated) {
		return true
	}

	return false
}

// SetAssociated gets a reference to the given []PublishersModelUpdateV1Associated and assigns it to the Associated field.
func (o *PublishersModelUpdateV1) SetAssociated(v []PublishersModelUpdateV1Associated) {
	o.Associated = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *PublishersModelUpdateV1) SetType(v string) {
	o.Type = &v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetInfo() string {
	if o == nil || IsNil(o.Info) {
		var ret string
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetInfoOk() (*string, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given string and assigns it to the Info field.
func (o *PublishersModelUpdateV1) SetInfo(v string) {
	o.Info = &v
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *PublishersModelUpdateV1) SetSite(v string) {
	o.Site = &v
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *PublishersModelUpdateV1) GetAdmin() PublishersModelUpdateV1Admin {
	if o == nil || IsNil(o.Admin) {
		var ret PublishersModelUpdateV1Admin
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublishersModelUpdateV1) GetAdminOk() (*PublishersModelUpdateV1Admin, bool) {
	if o == nil || IsNil(o.Admin) {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *PublishersModelUpdateV1) HasAdmin() bool {
	if o != nil && !IsNil(o.Admin) {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given PublishersModelUpdateV1Admin and assigns it to the Admin field.
func (o *PublishersModelUpdateV1) SetAdmin(v PublishersModelUpdateV1Admin) {
	o.Admin = &v
}

func (o PublishersModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishersModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Associated) {
		toSerialize["associated"] = o.Associated
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Admin) {
		toSerialize["admin"] = o.Admin
	}
	return toSerialize, nil
}

type NullablePublishersModelUpdateV1 struct {
	value *PublishersModelUpdateV1
	isSet bool
}

func (v NullablePublishersModelUpdateV1) Get() *PublishersModelUpdateV1 {
	return v.value
}

func (v *NullablePublishersModelUpdateV1) Set(val *PublishersModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishersModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishersModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishersModelUpdateV1(val *PublishersModelUpdateV1) *NullablePublishersModelUpdateV1 {
	return &NullablePublishersModelUpdateV1{value: val, isSet: true}
}

func (v NullablePublishersModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishersModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


