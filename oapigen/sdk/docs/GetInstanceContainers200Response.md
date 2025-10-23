# GetInstanceContainers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Containers** | Pointer to [**[]GetContainer200ResponseContainer**](GetContainer200ResponseContainer.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetInstanceContainers200Response

`func NewGetInstanceContainers200Response() *GetInstanceContainers200Response`

NewGetInstanceContainers200Response instantiates a new GetInstanceContainers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceContainers200ResponseWithDefaults

`func NewGetInstanceContainers200ResponseWithDefaults() *GetInstanceContainers200Response`

NewGetInstanceContainers200ResponseWithDefaults instantiates a new GetInstanceContainers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainers

`func (o *GetInstanceContainers200Response) GetContainers() []GetContainer200ResponseContainer`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetInstanceContainers200Response) GetContainersOk() (*[]GetContainer200ResponseContainer, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetInstanceContainers200Response) SetContainers(v []GetContainer200ResponseContainer)`

SetContainers sets Containers field to given value.

### HasContainers

`func (o *GetInstanceContainers200Response) HasContainers() bool`

HasContainers returns a boolean if a field has been set.

### GetMeta

`func (o *GetInstanceContainers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInstanceContainers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInstanceContainers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInstanceContainers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


