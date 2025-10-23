# ListIntegrationTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IntegrationTypes** | Pointer to [**[]ListIntegrationTypes200ResponseAllOfIntegrationTypesInner**](ListIntegrationTypes200ResponseAllOfIntegrationTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListIntegrationTypes200Response

`func NewListIntegrationTypes200Response() *ListIntegrationTypes200Response`

NewListIntegrationTypes200Response instantiates a new ListIntegrationTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrationTypes200ResponseWithDefaults

`func NewListIntegrationTypes200ResponseWithDefaults() *ListIntegrationTypes200Response`

NewListIntegrationTypes200ResponseWithDefaults instantiates a new ListIntegrationTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIntegrationTypes

`func (o *ListIntegrationTypes200Response) GetIntegrationTypes() []ListIntegrationTypes200ResponseAllOfIntegrationTypesInner`

GetIntegrationTypes returns the IntegrationTypes field if non-nil, zero value otherwise.

### GetIntegrationTypesOk

`func (o *ListIntegrationTypes200Response) GetIntegrationTypesOk() (*[]ListIntegrationTypes200ResponseAllOfIntegrationTypesInner, bool)`

GetIntegrationTypesOk returns a tuple with the IntegrationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationTypes

`func (o *ListIntegrationTypes200Response) SetIntegrationTypes(v []ListIntegrationTypes200ResponseAllOfIntegrationTypesInner)`

SetIntegrationTypes sets IntegrationTypes field to given value.

### HasIntegrationTypes

`func (o *ListIntegrationTypes200Response) HasIntegrationTypes() bool`

HasIntegrationTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListIntegrationTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIntegrationTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIntegrationTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIntegrationTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


