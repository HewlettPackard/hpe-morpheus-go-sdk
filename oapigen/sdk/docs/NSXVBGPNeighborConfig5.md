# NSXVBGPNeighborConfig5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RouterId** | Pointer to **string** | The router identifier (auto-populated for edge routers) | [optional] 
**Interface** | Pointer to **string** | The interface name for the BGP session (distributed routers only) | [optional] 

## Methods

### NewNSXVBGPNeighborConfig5

`func NewNSXVBGPNeighborConfig5() *NSXVBGPNeighborConfig5`

NewNSXVBGPNeighborConfig5 instantiates a new NSXVBGPNeighborConfig5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXVBGPNeighborConfig5WithDefaults

`func NewNSXVBGPNeighborConfig5WithDefaults() *NSXVBGPNeighborConfig5`

NewNSXVBGPNeighborConfig5WithDefaults instantiates a new NSXVBGPNeighborConfig5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRouterId

`func (o *NSXVBGPNeighborConfig5) GetRouterId() string`

GetRouterId returns the RouterId field if non-nil, zero value otherwise.

### GetRouterIdOk

`func (o *NSXVBGPNeighborConfig5) GetRouterIdOk() (*string, bool)`

GetRouterIdOk returns a tuple with the RouterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterId

`func (o *NSXVBGPNeighborConfig5) SetRouterId(v string)`

SetRouterId sets RouterId field to given value.

### HasRouterId

`func (o *NSXVBGPNeighborConfig5) HasRouterId() bool`

HasRouterId returns a boolean if a field has been set.

### GetInterface

`func (o *NSXVBGPNeighborConfig5) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *NSXVBGPNeighborConfig5) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *NSXVBGPNeighborConfig5) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *NSXVBGPNeighborConfig5) HasInterface() bool`

HasInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


