# CreateNetworkServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkServer** | Pointer to [**NSXNetworkServer**](NSXNetworkServer.md) |  | [optional] 

## Methods

### NewCreateNetworkServerRequest

`func NewCreateNetworkServerRequest() *CreateNetworkServerRequest`

NewCreateNetworkServerRequest instantiates a new CreateNetworkServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworkServerRequestWithDefaults

`func NewCreateNetworkServerRequestWithDefaults() *CreateNetworkServerRequest`

NewCreateNetworkServerRequestWithDefaults instantiates a new CreateNetworkServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkServer

`func (o *CreateNetworkServerRequest) GetNetworkServer() NSXNetworkServer`

GetNetworkServer returns the NetworkServer field if non-nil, zero value otherwise.

### GetNetworkServerOk

`func (o *CreateNetworkServerRequest) GetNetworkServerOk() (*NSXNetworkServer, bool)`

GetNetworkServerOk returns a tuple with the NetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServer

`func (o *CreateNetworkServerRequest) SetNetworkServer(v NSXNetworkServer)`

SetNetworkServer sets NetworkServer field to given value.

### HasNetworkServer

`func (o *CreateNetworkServerRequest) HasNetworkServer() bool`

HasNetworkServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


