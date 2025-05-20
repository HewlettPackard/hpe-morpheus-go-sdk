# CreateNetworkPoolServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**NetworkPoolServer** | Pointer to [**ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner.md) |  | [optional] 

## Methods

### NewCreateNetworkPoolServer200Response

`func NewCreateNetworkPoolServer200Response() *CreateNetworkPoolServer200Response`

NewCreateNetworkPoolServer200Response instantiates a new CreateNetworkPoolServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkPoolServer200ResponseWithDefaults

`func NewCreateNetworkPoolServer200ResponseWithDefaults() *CreateNetworkPoolServer200Response`

NewCreateNetworkPoolServer200ResponseWithDefaults instantiates a new CreateNetworkPoolServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateNetworkPoolServer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateNetworkPoolServer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateNetworkPoolServer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateNetworkPoolServer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetNetworkPoolServer

`func (o *CreateNetworkPoolServer200Response) GetNetworkPoolServer() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner`

GetNetworkPoolServer returns the NetworkPoolServer field if non-nil, zero value otherwise.

### GetNetworkPoolServerOk

`func (o *CreateNetworkPoolServer200Response) GetNetworkPoolServerOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner, bool)`

GetNetworkPoolServerOk returns a tuple with the NetworkPoolServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolServer

`func (o *CreateNetworkPoolServer200Response) SetNetworkPoolServer(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner)`

SetNetworkPoolServer sets NetworkPoolServer field to given value.

### HasNetworkPoolServer

`func (o *CreateNetworkPoolServer200Response) HasNetworkPoolServer() bool`

HasNetworkPoolServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


