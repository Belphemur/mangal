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

// checks if the GenreModelStatsV1Stats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenreModelStatsV1Stats{}

// GenreModelStatsV1Stats struct for GenreModelStatsV1Stats
type GenreModelStatsV1Stats struct {
	Series     *int64 `json:"series,omitempty"`
	Authors    *int64 `json:"authors,omitempty"`
	Filters    *int64 `json:"filters,omitempty"`
	Highlights *int64 `json:"highlights,omitempty"`
}

// NewGenreModelStatsV1Stats instantiates a new GenreModelStatsV1Stats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenreModelStatsV1Stats() *GenreModelStatsV1Stats {
	this := GenreModelStatsV1Stats{}
	return &this
}

// NewGenreModelStatsV1StatsWithDefaults instantiates a new GenreModelStatsV1Stats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenreModelStatsV1StatsWithDefaults() *GenreModelStatsV1Stats {
	this := GenreModelStatsV1Stats{}
	return &this
}

// GetSeries returns the Series field value if set, zero value otherwise.
func (o *GenreModelStatsV1Stats) GetSeries() int64 {
	if o == nil || IsNil(o.Series) {
		var ret int64
		return ret
	}
	return *o.Series
}

// GetSeriesOk returns a tuple with the Series field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreModelStatsV1Stats) GetSeriesOk() (*int64, bool) {
	if o == nil || IsNil(o.Series) {
		return nil, false
	}
	return o.Series, true
}

// HasSeries returns a boolean if a field has been set.
func (o *GenreModelStatsV1Stats) HasSeries() bool {
	if o != nil && !IsNil(o.Series) {
		return true
	}

	return false
}

// SetSeries gets a reference to the given int64 and assigns it to the Series field.
func (o *GenreModelStatsV1Stats) SetSeries(v int64) {
	o.Series = &v
}

// GetAuthors returns the Authors field value if set, zero value otherwise.
func (o *GenreModelStatsV1Stats) GetAuthors() int64 {
	if o == nil || IsNil(o.Authors) {
		var ret int64
		return ret
	}
	return *o.Authors
}

// GetAuthorsOk returns a tuple with the Authors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreModelStatsV1Stats) GetAuthorsOk() (*int64, bool) {
	if o == nil || IsNil(o.Authors) {
		return nil, false
	}
	return o.Authors, true
}

// HasAuthors returns a boolean if a field has been set.
func (o *GenreModelStatsV1Stats) HasAuthors() bool {
	if o != nil && !IsNil(o.Authors) {
		return true
	}

	return false
}

// SetAuthors gets a reference to the given int64 and assigns it to the Authors field.
func (o *GenreModelStatsV1Stats) SetAuthors(v int64) {
	o.Authors = &v
}

// GetFilters returns the Filters field value if set, zero value otherwise.
func (o *GenreModelStatsV1Stats) GetFilters() int64 {
	if o == nil || IsNil(o.Filters) {
		var ret int64
		return ret
	}
	return *o.Filters
}

// GetFiltersOk returns a tuple with the Filters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreModelStatsV1Stats) GetFiltersOk() (*int64, bool) {
	if o == nil || IsNil(o.Filters) {
		return nil, false
	}
	return o.Filters, true
}

// HasFilters returns a boolean if a field has been set.
func (o *GenreModelStatsV1Stats) HasFilters() bool {
	if o != nil && !IsNil(o.Filters) {
		return true
	}

	return false
}

// SetFilters gets a reference to the given int64 and assigns it to the Filters field.
func (o *GenreModelStatsV1Stats) SetFilters(v int64) {
	o.Filters = &v
}

// GetHighlights returns the Highlights field value if set, zero value otherwise.
func (o *GenreModelStatsV1Stats) GetHighlights() int64 {
	if o == nil || IsNil(o.Highlights) {
		var ret int64
		return ret
	}
	return *o.Highlights
}

// GetHighlightsOk returns a tuple with the Highlights field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenreModelStatsV1Stats) GetHighlightsOk() (*int64, bool) {
	if o == nil || IsNil(o.Highlights) {
		return nil, false
	}
	return o.Highlights, true
}

// HasHighlights returns a boolean if a field has been set.
func (o *GenreModelStatsV1Stats) HasHighlights() bool {
	if o != nil && !IsNil(o.Highlights) {
		return true
	}

	return false
}

// SetHighlights gets a reference to the given int64 and assigns it to the Highlights field.
func (o *GenreModelStatsV1Stats) SetHighlights(v int64) {
	o.Highlights = &v
}

func (o GenreModelStatsV1Stats) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenreModelStatsV1Stats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Series) {
		toSerialize["series"] = o.Series
	}
	if !IsNil(o.Authors) {
		toSerialize["authors"] = o.Authors
	}
	if !IsNil(o.Filters) {
		toSerialize["filters"] = o.Filters
	}
	if !IsNil(o.Highlights) {
		toSerialize["highlights"] = o.Highlights
	}
	return toSerialize, nil
}

type NullableGenreModelStatsV1Stats struct {
	value *GenreModelStatsV1Stats
	isSet bool
}

func (v NullableGenreModelStatsV1Stats) Get() *GenreModelStatsV1Stats {
	return v.value
}

func (v *NullableGenreModelStatsV1Stats) Set(val *GenreModelStatsV1Stats) {
	v.value = val
	v.isSet = true
}

func (v NullableGenreModelStatsV1Stats) IsSet() bool {
	return v.isSet
}

func (v *NullableGenreModelStatsV1Stats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenreModelStatsV1Stats(val *GenreModelStatsV1Stats) *NullableGenreModelStatsV1Stats {
	return &NullableGenreModelStatsV1Stats{value: val, isSet: true}
}

func (v NullableGenreModelStatsV1Stats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenreModelStatsV1Stats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}