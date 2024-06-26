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

// checks if the ConvoSearchResponseV1ResultsMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConvoSearchResponseV1ResultsMetadata{}

// ConvoSearchResponseV1ResultsMetadata struct for ConvoSearchResponseV1ResultsMetadata
type ConvoSearchResponseV1ResultsMetadata struct {
	Message *ConvoMessageModelV1 `json:"message,omitempty"`
	Participant *ConvoParticipantModelV1 `json:"participant,omitempty"`
}

// NewConvoSearchResponseV1ResultsMetadata instantiates a new ConvoSearchResponseV1ResultsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConvoSearchResponseV1ResultsMetadata() *ConvoSearchResponseV1ResultsMetadata {
	this := ConvoSearchResponseV1ResultsMetadata{}
	return &this
}

// NewConvoSearchResponseV1ResultsMetadataWithDefaults instantiates a new ConvoSearchResponseV1ResultsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConvoSearchResponseV1ResultsMetadataWithDefaults() *ConvoSearchResponseV1ResultsMetadata {
	this := ConvoSearchResponseV1ResultsMetadata{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ConvoSearchResponseV1ResultsMetadata) GetMessage() ConvoMessageModelV1 {
	if o == nil || IsNil(o.Message) {
		var ret ConvoMessageModelV1
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoSearchResponseV1ResultsMetadata) GetMessageOk() (*ConvoMessageModelV1, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ConvoSearchResponseV1ResultsMetadata) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ConvoMessageModelV1 and assigns it to the Message field.
func (o *ConvoSearchResponseV1ResultsMetadata) SetMessage(v ConvoMessageModelV1) {
	o.Message = &v
}

// GetParticipant returns the Participant field value if set, zero value otherwise.
func (o *ConvoSearchResponseV1ResultsMetadata) GetParticipant() ConvoParticipantModelV1 {
	if o == nil || IsNil(o.Participant) {
		var ret ConvoParticipantModelV1
		return ret
	}
	return *o.Participant
}

// GetParticipantOk returns a tuple with the Participant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConvoSearchResponseV1ResultsMetadata) GetParticipantOk() (*ConvoParticipantModelV1, bool) {
	if o == nil || IsNil(o.Participant) {
		return nil, false
	}
	return o.Participant, true
}

// HasParticipant returns a boolean if a field has been set.
func (o *ConvoSearchResponseV1ResultsMetadata) HasParticipant() bool {
	if o != nil && !IsNil(o.Participant) {
		return true
	}

	return false
}

// SetParticipant gets a reference to the given ConvoParticipantModelV1 and assigns it to the Participant field.
func (o *ConvoSearchResponseV1ResultsMetadata) SetParticipant(v ConvoParticipantModelV1) {
	o.Participant = &v
}

func (o ConvoSearchResponseV1ResultsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConvoSearchResponseV1ResultsMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Participant) {
		toSerialize["participant"] = o.Participant
	}
	return toSerialize, nil
}

type NullableConvoSearchResponseV1ResultsMetadata struct {
	value *ConvoSearchResponseV1ResultsMetadata
	isSet bool
}

func (v NullableConvoSearchResponseV1ResultsMetadata) Get() *ConvoSearchResponseV1ResultsMetadata {
	return v.value
}

func (v *NullableConvoSearchResponseV1ResultsMetadata) Set(val *ConvoSearchResponseV1ResultsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableConvoSearchResponseV1ResultsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableConvoSearchResponseV1ResultsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConvoSearchResponseV1ResultsMetadata(val *ConvoSearchResponseV1ResultsMetadata) *NullableConvoSearchResponseV1ResultsMetadata {
	return &NullableConvoSearchResponseV1ResultsMetadata{value: val, isSet: true}
}

func (v NullableConvoSearchResponseV1ResultsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConvoSearchResponseV1ResultsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


