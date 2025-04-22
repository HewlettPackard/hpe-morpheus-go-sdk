# CreateNetworkServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**NetworkServer** | Pointer to [**ListNetworkServers200ResponseAllOfNetworkServersInner**](ListNetworkServers200ResponseAllOfNetworkServersInner.md) |  | [optional] 

## Methods

### NewCreateNetworkServer200Response

`func NewCreateNetworkServer200Response() *CreateNetworkServer200Response`

NewCreateNetworkServer200Response instantiates a new CreateNetworkServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkServer200ResponseWithDefaults

`func NewCreateNetworkServer200ResponseWithDefaults() *CreateNetworkServer200Response`

NewCreateNetworkServer200ResponseWithDefaults instantiates a new CreateNetworkServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateNetworkServer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateNetworkServer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateNetworkServer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateNetworkServer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetNetworkServer

`func (o *CreateNetworkServer200Response) GetNetworkServer() ListNetworkServers200ResponseAllOfNetworkServersInner`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *CreateNetworkServer200Response) GetNetworkServerOk() (*ListNetworkServers200ResponseAllOfNetworkServersInner, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *CreateNetworkServer200Response) SetNetworkServer(v ListNetworkServers200ResponseAllOfNetworkServersInner)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *CreateNetworkServer200Response) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


