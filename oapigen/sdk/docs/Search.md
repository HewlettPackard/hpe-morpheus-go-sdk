# Search

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]SearchHitsInner**](SearchHitsInner.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Took** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**SearchMeta**](SearchMeta.md) |  | [optional] 

## Methods

### NewSearch

`func NewSearch() *Search`

NewSearch instantiates a new Search object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchWithDefaults

`func NewSearchWithDefaults() *Search`

NewSearchWithDefaults instantiates a new Search object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *Search) GetHits() []SearchHitsInner`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *Search) GetHitsOk() (*[]SearchHitsInner, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *Search) SetHits(v []SearchHitsInner)`

SetHits sets Hits field to given value.

### HasHits

`func (o *Search) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetQuery

`func (o *Search) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Search) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Search) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Search) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTook

`func (o *Search) GetTook() int64`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *Search) GetTookOk() (*int64, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *Search) SetTook(v int64)`

SetTook sets Took field to given value.

### HasTook

`func (o *Search) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetMeta

`func (o *Search) GetMeta() SearchMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Search) GetMetaOk() (*SearchMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Search) SetMeta(v SearchMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Search) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


