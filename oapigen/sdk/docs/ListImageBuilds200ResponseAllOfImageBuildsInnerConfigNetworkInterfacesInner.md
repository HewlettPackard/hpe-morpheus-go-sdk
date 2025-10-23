# ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpMode** | Pointer to **string** |  | [optional] 
**PrimaryInterface** | Pointer to **bool** |  | [optional] 
**ShowNetworkPoolLabel** | Pointer to **bool** |  | [optional] 
**ShowNetworkDhcpLabel** | Pointer to **bool** |  | [optional] 
**Network** | Pointer to [**ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork**](ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork.md) |  | [optional] 
**NetworkInterfaceTypeId** | Pointer to **int64** |  | [optional] 
**NetworkInterfaceTypeIdName** | Pointer to **string** |  | [optional] 

## Methods

### NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner

`func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner`

NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerWithDefaults

`func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner`

NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpMode

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetIpMode() string`

GetIpMode returns the IpMode field if non-nil, zero value otherwise.

### GetIpModeOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetIpModeOk() (*string, bool)`

GetIpModeOk returns a tuple with the IpMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpMode

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetIpMode(v string)`

SetIpMode sets IpMode field to given value.

### HasIpMode

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasIpMode() bool`

HasIpMode returns a boolean if a field has been set.

### GetPrimaryInterface

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetPrimaryInterface() bool`

GetPrimaryInterface returns the PrimaryInterface field if non-nil, zero value otherwise.

### GetPrimaryInterfaceOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool)`

GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryInterface

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetPrimaryInterface(v bool)`

SetPrimaryInterface sets PrimaryInterface field to given value.

### HasPrimaryInterface

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasPrimaryInterface() bool`

HasPrimaryInterface returns a boolean if a field has been set.

### GetShowNetworkPoolLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkPoolLabel() bool`

GetShowNetworkPoolLabel returns the ShowNetworkPoolLabel field if non-nil, zero value otherwise.

### GetShowNetworkPoolLabelOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkPoolLabelOk() (*bool, bool)`

GetShowNetworkPoolLabelOk returns a tuple with the ShowNetworkPoolLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkPoolLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetShowNetworkPoolLabel(v bool)`

SetShowNetworkPoolLabel sets ShowNetworkPoolLabel field to given value.

### HasShowNetworkPoolLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasShowNetworkPoolLabel() bool`

HasShowNetworkPoolLabel returns a boolean if a field has been set.

### GetShowNetworkDhcpLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkDhcpLabel() bool`

GetShowNetworkDhcpLabel returns the ShowNetworkDhcpLabel field if non-nil, zero value otherwise.

### GetShowNetworkDhcpLabelOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkDhcpLabelOk() (*bool, bool)`

GetShowNetworkDhcpLabelOk returns a tuple with the ShowNetworkDhcpLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowNetworkDhcpLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetShowNetworkDhcpLabel(v bool)`

SetShowNetworkDhcpLabel sets ShowNetworkDhcpLabel field to given value.

### HasShowNetworkDhcpLabel

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasShowNetworkDhcpLabel() bool`

HasShowNetworkDhcpLabel returns a boolean if a field has been set.

### GetNetwork

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetwork() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetwork(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkInterfaceTypeId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64`

GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool)`

GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64)`

SetNetworkInterfaceTypeId sets NetworkInterfaceTypeId field to given value.

### HasNetworkInterfaceTypeId

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasNetworkInterfaceTypeId() bool`

HasNetworkInterfaceTypeId returns a boolean if a field has been set.

### GetNetworkInterfaceTypeIdName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdName() string`

GetNetworkInterfaceTypeIdName returns the NetworkInterfaceTypeIdName field if non-nil, zero value otherwise.

### GetNetworkInterfaceTypeIdNameOk

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdNameOk() (*string, bool)`

GetNetworkInterfaceTypeIdNameOk returns a tuple with the NetworkInterfaceTypeIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaceTypeIdName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetworkInterfaceTypeIdName(v string)`

SetNetworkInterfaceTypeIdName sets NetworkInterfaceTypeIdName field to given value.

### HasNetworkInterfaceTypeIdName

`func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) HasNetworkInterfaceTypeIdName() bool`

HasNetworkInterfaceTypeIdName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


