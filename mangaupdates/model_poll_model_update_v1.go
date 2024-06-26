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

// checks if the PollModelUpdateV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PollModelUpdateV1{}

// PollModelUpdateV1 struct for PollModelUpdateV1
type PollModelUpdateV1 struct {
	Question *string `json:"question,omitempty"`
	Answers []PollModelUpdateV1Answers `json:"answers,omitempty"`
}

// NewPollModelUpdateV1 instantiates a new PollModelUpdateV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPollModelUpdateV1() *PollModelUpdateV1 {
	this := PollModelUpdateV1{}
	return &this
}

// NewPollModelUpdateV1WithDefaults instantiates a new PollModelUpdateV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPollModelUpdateV1WithDefaults() *PollModelUpdateV1 {
	this := PollModelUpdateV1{}
	return &this
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *PollModelUpdateV1) GetQuestion() string {
	if o == nil || IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollModelUpdateV1) GetQuestionOk() (*string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *PollModelUpdateV1) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *PollModelUpdateV1) SetQuestion(v string) {
	o.Question = &v
}

// GetAnswers returns the Answers field value if set, zero value otherwise.
func (o *PollModelUpdateV1) GetAnswers() []PollModelUpdateV1Answers {
	if o == nil || IsNil(o.Answers) {
		var ret []PollModelUpdateV1Answers
		return ret
	}
	return o.Answers
}

// GetAnswersOk returns a tuple with the Answers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PollModelUpdateV1) GetAnswersOk() ([]PollModelUpdateV1Answers, bool) {
	if o == nil || IsNil(o.Answers) {
		return nil, false
	}
	return o.Answers, true
}

// HasAnswers returns a boolean if a field has been set.
func (o *PollModelUpdateV1) HasAnswers() bool {
	if o != nil && !IsNil(o.Answers) {
		return true
	}

	return false
}

// SetAnswers gets a reference to the given []PollModelUpdateV1Answers and assigns it to the Answers field.
func (o *PollModelUpdateV1) SetAnswers(v []PollModelUpdateV1Answers) {
	o.Answers = v
}

func (o PollModelUpdateV1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PollModelUpdateV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.Answers) {
		toSerialize["answers"] = o.Answers
	}
	return toSerialize, nil
}

type NullablePollModelUpdateV1 struct {
	value *PollModelUpdateV1
	isSet bool
}

func (v NullablePollModelUpdateV1) Get() *PollModelUpdateV1 {
	return v.value
}

func (v *NullablePollModelUpdateV1) Set(val *PollModelUpdateV1) {
	v.value = val
	v.isSet = true
}

func (v NullablePollModelUpdateV1) IsSet() bool {
	return v.isSet
}

func (v *NullablePollModelUpdateV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePollModelUpdateV1(val *PollModelUpdateV1) *NullablePollModelUpdateV1 {
	return &NullablePollModelUpdateV1{value: val, isSet: true}
}

func (v NullablePollModelUpdateV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePollModelUpdateV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


