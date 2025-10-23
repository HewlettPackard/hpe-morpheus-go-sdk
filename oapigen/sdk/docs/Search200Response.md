# Search200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]Search200ResponseHitsInner**](Search200ResponseHitsInner.md) |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Took** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**Search200ResponseMeta**](Search200ResponseMeta.md) |  | [optional] 

## Methods

### NewSearch200Response

`func NewSearch200Response() *Search200Response`

NewSearch200Response instantiates a new Search200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearch200ResponseWithDefaults

`func NewSearch200ResponseWithDefaults() *Search200Response`

NewSearch200ResponseWithDefaults instantiates a new Search200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *Search200Response) GetHits() []Search200ResponseHitsInner`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *Search200Response) GetHitsOk() (*[]Search200ResponseHitsInner, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *Search200Response) SetHits(v []Search200ResponseHitsInner)`

SetHits sets Hits field to given value.

### HasHits

`func (o *Search200Response) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetQuery

`func (o *Search200Response) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *Search200Response) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *Search200Response) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *Search200Response) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetTook

`func (o *Search200Response) GetTook() int64`

GetTook returns the Took field if non-nil, zero value otherwise.

### GetTookOk

`func (o *Search200Response) GetTookOk() (*int64, bool)`

GetTookOk returns a tuple with the Took field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTook

`func (o *Search200Response) SetTook(v int64)`

SetTook sets Took field to given value.

### HasTook

`func (o *Search200Response) HasTook() bool`

HasTook returns a boolean if a field has been set.

### GetMeta

`func (o *Search200Response) GetMeta() Search200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *Search200Response) GetMetaOk() (*Search200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *Search200Response) SetMeta(v Search200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *Search200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


