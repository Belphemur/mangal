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

// checks if the ReleaseModelSearchV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReleaseModelSearchV1{}

// ReleaseModelSearchV1 struct for ReleaseModelSearchV1
type ReleaseModelSearchV1 struct {
	Id          *int64                       `json:"id,omitempty"`
	Title       *string                      `json:"title,omitempty"`
	Volume      *string                      `json:"volume,omitempty"`
	Chapter     *string                      `json:"chapter,omitempty"`
	Groups      []ReleaseModelSearchV1Groups `json:"groups,omitempty"`
	ReleaseDate *string                      `json:"release_date,omitempty"`
	TimeAdded   *TimeV1                      `json:"time_added,omitempty"`
}

// NewReleaseModelSearchV1 instantiates a new ReleaseModelSearchV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReleaseModelSearchV1() *ReleaseModelSearchV1 {
	this := ReleaseModelSearchV1{}
	return &this
}

// NewReleaseModelSearchV1WithDefaults instantiates a new ReleaseModelSearchV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReleaseModelSearchV1WithDefaults() *ReleaseModelSearchV1 {
	this := ReleaseModelSearchV1{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ReleaseModelSearchV1) SetId(v int64) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *ReleaseModelSearchV1) SetTitle(v string) {
	o.Title = &v
}

// GetVolume returns the Volume field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetVolume() string {
	if o == nil || IsNil(o.Volume) {
		var ret string
		return ret
	}
	return *o.Volume
}

// GetVolumeOk returns a tuple with the Volume field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetVolumeOk() (*string, bool) {
	if o == nil || IsNil(o.Volume) {
		return nil, false
	}
	return o.Volume, true
}

// HasVolume returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasVolume() bool {
	if o != nil && !IsNil(o.Volume) {
		return true
	}

	return false
}

// SetVolume gets a reference to the given string and assigns it to the Volume field.
func (o *ReleaseModelSearchV1) SetVolume(v string) {
	o.Volume = &v
}

// GetChapter returns the Chapter field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetChapter() string {
	if o == nil || IsNil(o.Chapter) {
		var ret string
		return ret
	}
	return *o.Chapter
}

// GetChapterOk returns a tuple with the Chapter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetChapterOk() (*string, bool) {
	if o == nil || IsNil(o.Chapter) {
		return nil, false
	}
	return o.Chapter, true
}

// HasChapter returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasChapter() bool {
	if o != nil && !IsNil(o.Chapter) {
		return true
	}

	return false
}

// SetChapter gets a reference to the given string and assigns it to the Chapter field.
func (o *ReleaseModelSearchV1) SetChapter(v string) {
	o.Chapter = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetGroups() []ReleaseModelSearchV1Groups {
	if o == nil || IsNil(o.Groups) {
		var ret []ReleaseModelSearchV1Groups
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetGroupsOk() ([]ReleaseModelSearchV1Groups, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []ReleaseModelSearchV1Groups and assigns it to the Groups field.
func (o *ReleaseModelSearchV1) SetGroups(v []ReleaseModelSearchV1Groups) {
	o.Groups = v
}

// GetReleaseDate returns the ReleaseDate field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetReleaseDate() string {
	if o == nil || IsNil(o.ReleaseDate) {
		var ret string
		return ret
	}
	return *o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetReleaseDateOk() (*string, bool) {
	if o == nil || IsNil(o.ReleaseDate) {
		return nil, false
	}
	return o.ReleaseDate, true
}

// HasReleaseDate returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasReleaseDate() bool {
	if o != nil && !IsNil(o.ReleaseDate) {
		return true
	}

	return false
}

// SetReleaseDate gets a reference to the given string and assigns it to the ReleaseDate field.
func (o *ReleaseModelSearchV1) SetReleaseDate(v string) {
	o.ReleaseDate = &v
}

// GetTimeAdded returns the TimeAdded field value if set, zero value otherwise.
func (o *ReleaseModelSearchV1) GetTimeAdded() TimeV1 {
	if o == nil || IsNil(o.TimeAdded) {
		var ret TimeV1
		return ret
	}
	return *o.TimeAdded
}

// GetTimeAddedOk returns a tuple with the TimeAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReleaseModelSearchV1) GetTimeAddedOk() (*TimeV1, bool) {
	if o == nil || IsNil(o.TimeAdded) {
		return nil, false
	}
	return o.TimeAdded, true
}

// HasTimeAdded returns a boolean if a field has been set.
func (o *ReleaseModelSearchV1) HasTimeAdded() bool {
	if o != nil && !IsNil(o.TimeAdded) {
		return true
	}

	return false
}

// SetTimeAdded gets a reference to the given TimeV1 and assigns it to the TimeAdded field.
func (o *ReleaseModelSearchV1) SetTimeAdded(v TimeV1) {
	o.TimeAdded = &v
}

func (o ReleaseModelSearchV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReleaseModelSearchV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Volume) {
		toSerialize["volume"] = o.Volume
	}
	if !IsNil(o.Chapter) {
		toSerialize["chapter"] = o.Chapter
	}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !IsNil(o.ReleaseDate) {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if !IsNil(o.TimeAdded) {
		toSerialize["time_added"] = o.TimeAdded
	}
	return toSerialize, nil
}

type NullableReleaseModelSearchV1 struct {
	value *ReleaseModelSearchV1
	isSet bool
}

func (v NullableReleaseModelSearchV1) Get() *ReleaseModelSearchV1 {
	return v.value
}

func (v *NullableReleaseModelSearchV1) Set(val *ReleaseModelSearchV1) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseModelSearchV1) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseModelSearchV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseModelSearchV1(val *ReleaseModelSearchV1) *NullableReleaseModelSearchV1 {
	return &NullableReleaseModelSearchV1{value: val, isSet: true}
}

func (v NullableReleaseModelSearchV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseModelSearchV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
