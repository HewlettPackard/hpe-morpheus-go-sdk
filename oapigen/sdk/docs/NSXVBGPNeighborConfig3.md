# NSXVBGPNeighborConfig3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterId** | Pointer to **string** | The router identifier (auto-populated for edge routers) | [optional] 
**Interface** | Pointer to **string** | The interface name for the BGP session (distributed routers only) | [optional] 

## Methods

### NewNSXVBGPNeighborConfig3

`func NewNSXVBGPNeighborConfig3() *NSXVBGPNeighborConfig3`

NewNSXVBGPNeighborConfig3 instantiates a new NSXVBGPNeighborConfig3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXVBGPNeighborConfig3WithDefaults

`func NewNSXVBGPNeighborConfig3WithDefaults() *NSXVBGPNeighborConfig3`

NewNSXVBGPNeighborConfig3WithDefaults instantiates a new NSXVBGPNeighborConfig3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouterId

`func (o *NSXVBGPNeighborConfig3) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *NSXVBGPNeighborConfig3) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *NSXVBGPNeighborConfig3) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *NSXVBGPNeighborConfig3) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetInterface

`func (o *NSXVBGPNeighborConfig3) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NSXVBGPNeighborConfig3) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NSXVBGPNeighborConfig3) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NSXVBGPNeighborConfig3) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


