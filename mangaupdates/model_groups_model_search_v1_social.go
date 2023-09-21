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

// checks if the GroupsModelSearchV1Social type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupsModelSearchV1Social{}

// GroupsModelSearchV1Social struct for GroupsModelSearchV1Social
type GroupsModelSearchV1Social struct {
	Site *string `json:"site,omitempty"`
	Facebook *string `json:"facebook,omitempty"`
	Twitter *string `json:"twitter,omitempty"`
	Irc *GroupsModelSearchV1SocialIrc `json:"irc,omitempty"`
	Forum *string `json:"forum,omitempty"`
	Discord *string `json:"discord,omitempty"`
}

// NewGroupsModelSearchV1Social instantiates a new GroupsModelSearchV1Social object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupsModelSearchV1Social() *GroupsModelSearchV1Social {
	this := GroupsModelSearchV1Social{}
	return &this
}

// NewGroupsModelSearchV1SocialWithDefaults instantiates a new GroupsModelSearchV1Social object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupsModelSearchV1SocialWithDefaults() *GroupsModelSearchV1Social {
	this := GroupsModelSearchV1Social{}
	return &this
}

// GetSite returns the Site field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetSite() string {
	if o == nil || IsNil(o.Site) {
		var ret string
		return ret
	}
	return *o.Site
}

// GetSiteOk returns a tuple with the Site field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetSiteOk() (*string, bool) {
	if o == nil || IsNil(o.Site) {
		return nil, false
	}
	return o.Site, true
}

// HasSite returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasSite() bool {
	if o != nil && !IsNil(o.Site) {
		return true
	}

	return false
}

// SetSite gets a reference to the given string and assigns it to the Site field.
func (o *GroupsModelSearchV1Social) SetSite(v string) {
	o.Site = &v
}

// GetFacebook returns the Facebook field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetFacebook() string {
	if o == nil || IsNil(o.Facebook) {
		var ret string
		return ret
	}
	return *o.Facebook
}

// GetFacebookOk returns a tuple with the Facebook field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetFacebookOk() (*string, bool) {
	if o == nil || IsNil(o.Facebook) {
		return nil, false
	}
	return o.Facebook, true
}

// HasFacebook returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasFacebook() bool {
	if o != nil && !IsNil(o.Facebook) {
		return true
	}

	return false
}

// SetFacebook gets a reference to the given string and assigns it to the Facebook field.
func (o *GroupsModelSearchV1Social) SetFacebook(v string) {
	o.Facebook = &v
}

// GetTwitter returns the Twitter field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetTwitter() string {
	if o == nil || IsNil(o.Twitter) {
		var ret string
		return ret
	}
	return *o.Twitter
}

// GetTwitterOk returns a tuple with the Twitter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetTwitterOk() (*string, bool) {
	if o == nil || IsNil(o.Twitter) {
		return nil, false
	}
	return o.Twitter, true
}

// HasTwitter returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasTwitter() bool {
	if o != nil && !IsNil(o.Twitter) {
		return true
	}

	return false
}

// SetTwitter gets a reference to the given string and assigns it to the Twitter field.
func (o *GroupsModelSearchV1Social) SetTwitter(v string) {
	o.Twitter = &v
}

// GetIrc returns the Irc field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetIrc() GroupsModelSearchV1SocialIrc {
	if o == nil || IsNil(o.Irc) {
		var ret GroupsModelSearchV1SocialIrc
		return ret
	}
	return *o.Irc
}

// GetIrcOk returns a tuple with the Irc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetIrcOk() (*GroupsModelSearchV1SocialIrc, bool) {
	if o == nil || IsNil(o.Irc) {
		return nil, false
	}
	return o.Irc, true
}

// HasIrc returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasIrc() bool {
	if o != nil && !IsNil(o.Irc) {
		return true
	}

	return false
}

// SetIrc gets a reference to the given GroupsModelSearchV1SocialIrc and assigns it to the Irc field.
func (o *GroupsModelSearchV1Social) SetIrc(v GroupsModelSearchV1SocialIrc) {
	o.Irc = &v
}

// GetForum returns the Forum field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetForum() string {
	if o == nil || IsNil(o.Forum) {
		var ret string
		return ret
	}
	return *o.Forum
}

// GetForumOk returns a tuple with the Forum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetForumOk() (*string, bool) {
	if o == nil || IsNil(o.Forum) {
		return nil, false
	}
	return o.Forum, true
}

// HasForum returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasForum() bool {
	if o != nil && !IsNil(o.Forum) {
		return true
	}

	return false
}

// SetForum gets a reference to the given string and assigns it to the Forum field.
func (o *GroupsModelSearchV1Social) SetForum(v string) {
	o.Forum = &v
}

// GetDiscord returns the Discord field value if set, zero value otherwise.
func (o *GroupsModelSearchV1Social) GetDiscord() string {
	if o == nil || IsNil(o.Discord) {
		var ret string
		return ret
	}
	return *o.Discord
}

// GetDiscordOk returns a tuple with the Discord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupsModelSearchV1Social) GetDiscordOk() (*string, bool) {
	if o == nil || IsNil(o.Discord) {
		return nil, false
	}
	return o.Discord, true
}

// HasDiscord returns a boolean if a field has been set.
func (o *GroupsModelSearchV1Social) HasDiscord() bool {
	if o != nil && !IsNil(o.Discord) {
		return true
	}

	return false
}

// SetDiscord gets a reference to the given string and assigns it to the Discord field.
func (o *GroupsModelSearchV1Social) SetDiscord(v string) {
	o.Discord = &v
}

func (o GroupsModelSearchV1Social) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupsModelSearchV1Social) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Site) {
		toSerialize["site"] = o.Site
	}
	if !IsNil(o.Facebook) {
		toSerialize["facebook"] = o.Facebook
	}
	if !IsNil(o.Twitter) {
		toSerialize["twitter"] = o.Twitter
	}
	if !IsNil(o.Irc) {
		toSerialize["irc"] = o.Irc
	}
	if !IsNil(o.Forum) {
		toSerialize["forum"] = o.Forum
	}
	if !IsNil(o.Discord) {
		toSerialize["discord"] = o.Discord
	}
	return toSerialize, nil
}

type NullableGroupsModelSearchV1Social struct {
	value *GroupsModelSearchV1Social
	isSet bool
}

func (v NullableGroupsModelSearchV1Social) Get() *GroupsModelSearchV1Social {
	return v.value
}

func (v *NullableGroupsModelSearchV1Social) Set(val *GroupsModelSearchV1Social) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupsModelSearchV1Social) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupsModelSearchV1Social) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupsModelSearchV1Social(val *GroupsModelSearchV1Social) *NullableGroupsModelSearchV1Social {
	return &NullableGroupsModelSearchV1Social{value: val, isSet: true}
}

func (v NullableGroupsModelSearchV1Social) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupsModelSearchV1Social) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

