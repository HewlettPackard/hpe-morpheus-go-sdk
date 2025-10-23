# ListHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | Pointer to [**ListHealth200ResponseAllOfHealth**](ListHealth200ResponseAllOfHealth.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListHealth200Response

`func NewListHealth200Response() *ListHealth200Response`

NewListHealth200Response instantiates a new ListHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseWithDefaults

`func NewListHealth200ResponseWithDefaults() *ListHealth200Response`

NewListHealth200ResponseWithDefaults instantiates a new ListHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *ListHealth200Response) GetHealth() ListHealth200ResponseAllOfHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ListHealth200Response) GetHealthOk() (*ListHealth200ResponseAllOfHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ListHealth200Response) SetHealth(v ListHealth200ResponseAllOfHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ListHealth200Response) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetMeta

`func (o *ListHealth200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListHealth200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListHealth200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListHealth200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


