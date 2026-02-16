# CreateNetworks200ResponseAllOfNetworkConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VlanIDs** | Pointer to **NullableString** |  | [optional] 
**ConnectedGateway** | Pointer to **string** |  | [optional] 
**SubnetIpManagementType** | Pointer to **string** |  | [optional] 
**SubnetIpServerId** | Pointer to **string** |  | [optional] 
**DhcpRange** | Pointer to **string** |  | [optional] 
**SubnetDhcpServerAddress** | Pointer to **string** |  | [optional] 
**SubnetDhcpLeaseTime** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateNetworks200ResponseAllOfNetworkConfig

`func NewCreateNetworks200ResponseAllOfNetworkConfig() *CreateNetworks200ResponseAllOfNetworkConfig`

NewCreateNetworks200ResponseAllOfNetworkConfig instantiates a new CreateNetworks200ResponseAllOfNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNetworks200ResponseAllOfNetworkConfigWithDefaults

`func NewCreateNetworks200ResponseAllOfNetworkConfigWithDefaults() *CreateNetworks200ResponseAllOfNetworkConfig`

NewCreateNetworks200ResponseAllOfNetworkConfigWithDefaults instantiates a new CreateNetworks200ResponseAllOfNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanIDs

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetVlanIDs() string`

GetVlanIDs returns the VlanIDs field if non-nil, zero value otherwise.

### GetVlanIDsOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetVlanIDsOk() (*string, bool)`

GetVlanIDsOk returns a tuple with the VlanIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIDs

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetVlanIDs(v string)`

SetVlanIDs sets VlanIDs field to given value.

### HasVlanIDs

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasVlanIDs() bool`

HasVlanIDs returns a boolean if a field has been set.

### SetVlanIDsNil

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetVlanIDsNil(b bool)`

 SetVlanIDsNil sets the value for VlanIDs to be an explicit nil

### UnsetVlanIDs
`func (o *CreateNetworks200ResponseAllOfNetworkConfig) UnsetVlanIDs()`

UnsetVlanIDs ensures that no value is present for VlanIDs, not even an explicit nil
### GetConnectedGateway

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetConnectedGateway() string`

GetConnectedGateway returns the ConnectedGateway field if non-nil, zero value otherwise.

### GetConnectedGatewayOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetConnectedGatewayOk() (*string, bool)`

GetConnectedGatewayOk returns a tuple with the ConnectedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGateway

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetConnectedGateway(v string)`

SetConnectedGateway sets ConnectedGateway field to given value.

### HasConnectedGateway

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasConnectedGateway() bool`

HasConnectedGateway returns a boolean if a field has been set.

### GetSubnetIpManagementType

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetIpManagementType() string`

GetSubnetIpManagementType returns the SubnetIpManagementType field if non-nil, zero value otherwise.

### GetSubnetIpManagementTypeOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetIpManagementTypeOk() (*string, bool)`

GetSubnetIpManagementTypeOk returns a tuple with the SubnetIpManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpManagementType

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetSubnetIpManagementType(v string)`

SetSubnetIpManagementType sets SubnetIpManagementType field to given value.

### HasSubnetIpManagementType

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasSubnetIpManagementType() bool`

HasSubnetIpManagementType returns a boolean if a field has been set.

### GetSubnetIpServerId

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetIpServerId() string`

GetSubnetIpServerId returns the SubnetIpServerId field if non-nil, zero value otherwise.

### GetSubnetIpServerIdOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetIpServerIdOk() (*string, bool)`

GetSubnetIpServerIdOk returns a tuple with the SubnetIpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpServerId

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetSubnetIpServerId(v string)`

SetSubnetIpServerId sets SubnetIpServerId field to given value.

### HasSubnetIpServerId

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasSubnetIpServerId() bool`

HasSubnetIpServerId returns a boolean if a field has been set.

### GetDhcpRange

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### GetSubnetDhcpServerAddress

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetDhcpServerAddress() string`

GetSubnetDhcpServerAddress returns the SubnetDhcpServerAddress field if non-nil, zero value otherwise.

### GetSubnetDhcpServerAddressOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetDhcpServerAddressOk() (*string, bool)`

GetSubnetDhcpServerAddressOk returns a tuple with the SubnetDhcpServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpServerAddress

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetSubnetDhcpServerAddress(v string)`

SetSubnetDhcpServerAddress sets SubnetDhcpServerAddress field to given value.

### HasSubnetDhcpServerAddress

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasSubnetDhcpServerAddress() bool`

HasSubnetDhcpServerAddress returns a boolean if a field has been set.

### GetSubnetDhcpLeaseTime

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetDhcpLeaseTime() string`

GetSubnetDhcpLeaseTime returns the SubnetDhcpLeaseTime field if non-nil, zero value otherwise.

### GetSubnetDhcpLeaseTimeOk

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) GetSubnetDhcpLeaseTimeOk() (*string, bool)`

GetSubnetDhcpLeaseTimeOk returns a tuple with the SubnetDhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpLeaseTime

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) SetSubnetDhcpLeaseTime(v string)`

SetSubnetDhcpLeaseTime sets SubnetDhcpLeaseTime field to given value.

### HasSubnetDhcpLeaseTime

`func (o *CreateNetworks200ResponseAllOfNetworkConfig) HasSubnetDhcpLeaseTime() bool`

HasSubnetDhcpLeaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


