# UpdateNetworkPoolServer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**NetworkPoolServer** | Pointer to [**UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer.md) |  | [optional] 

## Methods

### NewUpdateNetworkPoolServer200Response

`func NewUpdateNetworkPoolServer200Response() *UpdateNetworkPoolServer200Response`

NewUpdateNetworkPoolServer200Response instantiates a new UpdateNetworkPoolServer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkPoolServer200ResponseWithDefaults

`func NewUpdateNetworkPoolServer200ResponseWithDefaults() *UpdateNetworkPoolServer200Response`

NewUpdateNetworkPoolServer200ResponseWithDefaults instantiates a new UpdateNetworkPoolServer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateNetworkPoolServer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateNetworkPoolServer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateNetworkPoolServer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateNetworkPoolServer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetNetworkPoolServer

`func (o *UpdateNetworkPoolServer200Response) GetNetworkPoolServer() UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer`

GetNetworkPoolServer returns the NetworkPoolServer field if non-nil, zero value otherwise.

### GetNetworkPoolServerOk

`func (o *UpdateNetworkPoolServer200Response) GetNetworkPoolServerOk() (*UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer, bool)`

GetNetworkPoolServerOk returns a tuple with the NetworkPoolServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolServer

`func (o *UpdateNetworkPoolServer200Response) SetNetworkPoolServer(v UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer)`

SetNetworkPoolServer sets NetworkPoolServer field to given value.

### HasNetworkPoolServer

`func (o *UpdateNetworkPoolServer200Response) HasNetworkPoolServer() bool`

HasNetworkPoolServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


