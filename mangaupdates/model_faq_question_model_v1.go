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

// checks if the FaqQuestionModelV1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaqQuestionModelV1{}

// FaqQuestionModelV1 struct for FaqQuestionModelV1
type FaqQuestionModelV1 struct {
	QuestionId int64   `json:"question_id"`
	Question   *string `json:"question,omitempty"`
	Answer     *string `json:"answer,omitempty"`
	Position   *int64  `json:"position,omitempty"`
}

// NewFaqQuestionModelV1 instantiates a new FaqQuestionModelV1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaqQuestionModelV1(questionId int64) *FaqQuestionModelV1 {
	this := FaqQuestionModelV1{}
	this.QuestionId = questionId
	return &this
}

// NewFaqQuestionModelV1WithDefaults instantiates a new FaqQuestionModelV1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaqQuestionModelV1WithDefaults() *FaqQuestionModelV1 {
	this := FaqQuestionModelV1{}
	return &this
}

// GetQuestionId returns the QuestionId field value
func (o *FaqQuestionModelV1) GetQuestionId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.QuestionId
}

// GetQuestionIdOk returns a tuple with the QuestionId field value
// and a boolean to check if the value has been set.
func (o *FaqQuestionModelV1) GetQuestionIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QuestionId, true
}

// SetQuestionId sets field value
func (o *FaqQuestionModelV1) SetQuestionId(v int64) {
	o.QuestionId = v
}

// GetQuestion returns the Question field value if set, zero value otherwise.
func (o *FaqQuestionModelV1) GetQuestion() string {
	if o == nil || IsNil(o.Question) {
		var ret string
		return ret
	}
	return *o.Question
}

// GetQuestionOk returns a tuple with the Question field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaqQuestionModelV1) GetQuestionOk() (*string, bool) {
	if o == nil || IsNil(o.Question) {
		return nil, false
	}
	return o.Question, true
}

// HasQuestion returns a boolean if a field has been set.
func (o *FaqQuestionModelV1) HasQuestion() bool {
	if o != nil && !IsNil(o.Question) {
		return true
	}

	return false
}

// SetQuestion gets a reference to the given string and assigns it to the Question field.
func (o *FaqQuestionModelV1) SetQuestion(v string) {
	o.Question = &v
}

// GetAnswer returns the Answer field value if set, zero value otherwise.
func (o *FaqQuestionModelV1) GetAnswer() string {
	if o == nil || IsNil(o.Answer) {
		var ret string
		return ret
	}
	return *o.Answer
}

// GetAnswerOk returns a tuple with the Answer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaqQuestionModelV1) GetAnswerOk() (*string, bool) {
	if o == nil || IsNil(o.Answer) {
		return nil, false
	}
	return o.Answer, true
}

// HasAnswer returns a boolean if a field has been set.
func (o *FaqQuestionModelV1) HasAnswer() bool {
	if o != nil && !IsNil(o.Answer) {
		return true
	}

	return false
}

// SetAnswer gets a reference to the given string and assigns it to the Answer field.
func (o *FaqQuestionModelV1) SetAnswer(v string) {
	o.Answer = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *FaqQuestionModelV1) GetPosition() int64 {
	if o == nil || IsNil(o.Position) {
		var ret int64
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaqQuestionModelV1) GetPositionOk() (*int64, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *FaqQuestionModelV1) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given int64 and assigns it to the Position field.
func (o *FaqQuestionModelV1) SetPosition(v int64) {
	o.Position = &v
}

func (o FaqQuestionModelV1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaqQuestionModelV1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["question_id"] = o.QuestionId
	if !IsNil(o.Question) {
		toSerialize["question"] = o.Question
	}
	if !IsNil(o.Answer) {
		toSerialize["answer"] = o.Answer
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

type NullableFaqQuestionModelV1 struct {
	value *FaqQuestionModelV1
	isSet bool
}

func (v NullableFaqQuestionModelV1) Get() *FaqQuestionModelV1 {
	return v.value
}

func (v *NullableFaqQuestionModelV1) Set(val *FaqQuestionModelV1) {
	v.value = val
	v.isSet = true
}

func (v NullableFaqQuestionModelV1) IsSet() bool {
	return v.isSet
}

func (v *NullableFaqQuestionModelV1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaqQuestionModelV1(val *FaqQuestionModelV1) *NullableFaqQuestionModelV1 {
	return &NullableFaqQuestionModelV1{value: val, isSet: true}
}

func (v NullableFaqQuestionModelV1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaqQuestionModelV1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
