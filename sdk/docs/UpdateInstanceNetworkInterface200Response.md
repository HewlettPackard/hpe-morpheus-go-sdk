# UpdateInstanceNetworkInterface200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkInterface** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface.md) |  | [optional] 
**InterfaceType** | Pointer to **string** |  | [optional] 
**NetId** | Pointer to **int64** |  | [optional] 
**Server** | Pointer to [**UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer**](UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewUpdateInstanceNetworkInterface200Response

`func NewUpdateInstanceNetworkInterface200Response() *UpdateInstanceNetworkInterface200Response`

NewUpdateInstanceNetworkInterface200Response instantiates a new UpdateInstanceNetworkInterface200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceNetworkInterface200ResponseWithDefaults

`func NewUpdateInstanceNetworkInterface200ResponseWithDefaults() *UpdateInstanceNetworkInterface200Response`

NewUpdateInstanceNetworkInterface200ResponseWithDefaults instantiates a new UpdateInstanceNetworkInterface200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkInterface

`func (o *UpdateInstanceNetworkInterface200Response) GetNetworkInterface() UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface`

GetNetworkInterface returns the NetworkInterface field if non-nil, zero value otherwise.

### GetNetworkInterfaceOk

`func (o *UpdateInstanceNetworkInterface200Response) GetNetworkInterfaceOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface, bool)`

GetNetworkInterfaceOk returns a tuple with the NetworkInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterface

`func (o *UpdateInstanceNetworkInterface200Response) SetNetworkInterface(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfNetworkInterface)`

SetNetworkInterface sets NetworkInterface field to given value.

### HasNetworkInterface

`func (o *UpdateInstanceNetworkInterface200Response) HasNetworkInterface() bool`

HasNetworkInterface returns a boolean if a field has been set.

### GetInterfaceType

`func (o *UpdateInstanceNetworkInterface200Response) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *UpdateInstanceNetworkInterface200Response) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *UpdateInstanceNetworkInterface200Response) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *UpdateInstanceNetworkInterface200Response) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetNetId

`func (o *UpdateInstanceNetworkInterface200Response) GetNetId() int64`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *UpdateInstanceNetworkInterface200Response) GetNetIdOk() (*int64, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *UpdateInstanceNetworkInterface200Response) SetNetId(v int64)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *UpdateInstanceNetworkInterface200Response) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetServer

`func (o *UpdateInstanceNetworkInterface200Response) GetServer() UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateInstanceNetworkInterface200Response) GetServerOk() (*UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateInstanceNetworkInterface200Response) SetServer(v UpdateInstanceNetworkInterface200ResponseAllOfOneOfServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateInstanceNetworkInterface200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateInstanceNetworkInterface200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateInstanceNetworkInterface200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateInstanceNetworkInterface200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateInstanceNetworkInterface200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *UpdateInstanceNetworkInterface200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UpdateInstanceNetworkInterface200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UpdateInstanceNetworkInterface200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UpdateInstanceNetworkInterface200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


