# ListClients200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clients** | Pointer to [**[]ListClients200ResponseAllOfClientsInner**](ListClients200ResponseAllOfClientsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClients200Response

`func NewListClients200Response() *ListClients200Response`

NewListClients200Response instantiates a new ListClients200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClients200ResponseWithDefaults

`func NewListClients200ResponseWithDefaults() *ListClients200Response`

NewListClients200ResponseWithDefaults instantiates a new ListClients200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClients

`func (o *ListClients200Response) GetClients() []ListClients200ResponseAllOfClientsInner`

GetClients returns the Clients field if non-nil, zero value otherwise.

### GetClientsOk

`func (o *ListClients200Response) GetClientsOk() (*[]ListClients200ResponseAllOfClientsInner, bool)`

GetClientsOk returns a tuple with the Clients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClients

`func (o *ListClients200Response) SetClients(v []ListClients200ResponseAllOfClientsInner)`

SetClients sets Clients field to given value.

### HasClients

`func (o *ListClients200Response) HasClients() bool`

HasClients returns a boolean if a field has been set.

### GetMeta

`func (o *ListClients200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClients200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClients200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClients200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


