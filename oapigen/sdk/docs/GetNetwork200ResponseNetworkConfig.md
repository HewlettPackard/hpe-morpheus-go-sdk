# GetNetwork200ResponseNetworkConfig

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

### NewGetNetwork200ResponseNetworkConfig

`func NewGetNetwork200ResponseNetworkConfig() *GetNetwork200ResponseNetworkConfig`

NewGetNetwork200ResponseNetworkConfig instantiates a new GetNetwork200ResponseNetworkConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetwork200ResponseNetworkConfigWithDefaults

`func NewGetNetwork200ResponseNetworkConfigWithDefaults() *GetNetwork200ResponseNetworkConfig`

NewGetNetwork200ResponseNetworkConfigWithDefaults instantiates a new GetNetwork200ResponseNetworkConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVlanIDs

`func (o *GetNetwork200ResponseNetworkConfig) GetVlanIDs() string`

GetVlanIDs returns the VlanIDs field if non-nil, zero value otherwise.

### GetVlanIDsOk

`func (o *GetNetwork200ResponseNetworkConfig) GetVlanIDsOk() (*string, bool)`

GetVlanIDsOk returns a tuple with the VlanIDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanIDs

`func (o *GetNetwork200ResponseNetworkConfig) SetVlanIDs(v string)`

SetVlanIDs sets VlanIDs field to given value.

### HasVlanIDs

`func (o *GetNetwork200ResponseNetworkConfig) HasVlanIDs() bool`

HasVlanIDs returns a boolean if a field has been set.

### SetVlanIDsNil

`func (o *GetNetwork200ResponseNetworkConfig) SetVlanIDsNil(b bool)`

 SetVlanIDsNil sets the value for VlanIDs to be an explicit nil

### UnsetVlanIDs
`func (o *GetNetwork200ResponseNetworkConfig) UnsetVlanIDs()`

UnsetVlanIDs ensures that no value is present for VlanIDs, not even an explicit nil
### GetConnectedGateway

`func (o *GetNetwork200ResponseNetworkConfig) GetConnectedGateway() string`

GetConnectedGateway returns the ConnectedGateway field if non-nil, zero value otherwise.

### GetConnectedGatewayOk

`func (o *GetNetwork200ResponseNetworkConfig) GetConnectedGatewayOk() (*string, bool)`

GetConnectedGatewayOk returns a tuple with the ConnectedGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedGateway

`func (o *GetNetwork200ResponseNetworkConfig) SetConnectedGateway(v string)`

SetConnectedGateway sets ConnectedGateway field to given value.

### HasConnectedGateway

`func (o *GetNetwork200ResponseNetworkConfig) HasConnectedGateway() bool`

HasConnectedGateway returns a boolean if a field has been set.

### GetSubnetIpManagementType

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetIpManagementType() string`

GetSubnetIpManagementType returns the SubnetIpManagementType field if non-nil, zero value otherwise.

### GetSubnetIpManagementTypeOk

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetIpManagementTypeOk() (*string, bool)`

GetSubnetIpManagementTypeOk returns a tuple with the SubnetIpManagementType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpManagementType

`func (o *GetNetwork200ResponseNetworkConfig) SetSubnetIpManagementType(v string)`

SetSubnetIpManagementType sets SubnetIpManagementType field to given value.

### HasSubnetIpManagementType

`func (o *GetNetwork200ResponseNetworkConfig) HasSubnetIpManagementType() bool`

HasSubnetIpManagementType returns a boolean if a field has been set.

### GetSubnetIpServerId

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetIpServerId() string`

GetSubnetIpServerId returns the SubnetIpServerId field if non-nil, zero value otherwise.

### GetSubnetIpServerIdOk

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetIpServerIdOk() (*string, bool)`

GetSubnetIpServerIdOk returns a tuple with the SubnetIpServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetIpServerId

`func (o *GetNetwork200ResponseNetworkConfig) SetSubnetIpServerId(v string)`

SetSubnetIpServerId sets SubnetIpServerId field to given value.

### HasSubnetIpServerId

`func (o *GetNetwork200ResponseNetworkConfig) HasSubnetIpServerId() bool`

HasSubnetIpServerId returns a boolean if a field has been set.

### GetDhcpRange

`func (o *GetNetwork200ResponseNetworkConfig) GetDhcpRange() string`

GetDhcpRange returns the DhcpRange field if non-nil, zero value otherwise.

### GetDhcpRangeOk

`func (o *GetNetwork200ResponseNetworkConfig) GetDhcpRangeOk() (*string, bool)`

GetDhcpRangeOk returns a tuple with the DhcpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpRange

`func (o *GetNetwork200ResponseNetworkConfig) SetDhcpRange(v string)`

SetDhcpRange sets DhcpRange field to given value.

### HasDhcpRange

`func (o *GetNetwork200ResponseNetworkConfig) HasDhcpRange() bool`

HasDhcpRange returns a boolean if a field has been set.

### GetSubnetDhcpServerAddress

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetDhcpServerAddress() string`

GetSubnetDhcpServerAddress returns the SubnetDhcpServerAddress field if non-nil, zero value otherwise.

### GetSubnetDhcpServerAddressOk

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetDhcpServerAddressOk() (*string, bool)`

GetSubnetDhcpServerAddressOk returns a tuple with the SubnetDhcpServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpServerAddress

`func (o *GetNetwork200ResponseNetworkConfig) SetSubnetDhcpServerAddress(v string)`

SetSubnetDhcpServerAddress sets SubnetDhcpServerAddress field to given value.

### HasSubnetDhcpServerAddress

`func (o *GetNetwork200ResponseNetworkConfig) HasSubnetDhcpServerAddress() bool`

HasSubnetDhcpServerAddress returns a boolean if a field has been set.

### GetSubnetDhcpLeaseTime

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetDhcpLeaseTime() string`

GetSubnetDhcpLeaseTime returns the SubnetDhcpLeaseTime field if non-nil, zero value otherwise.

### GetSubnetDhcpLeaseTimeOk

`func (o *GetNetwork200ResponseNetworkConfig) GetSubnetDhcpLeaseTimeOk() (*string, bool)`

GetSubnetDhcpLeaseTimeOk returns a tuple with the SubnetDhcpLeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetDhcpLeaseTime

`func (o *GetNetwork200ResponseNetworkConfig) SetSubnetDhcpLeaseTime(v string)`

SetSubnetDhcpLeaseTime sets SubnetDhcpLeaseTime field to given value.

### HasSubnetDhcpLeaseTime

`func (o *GetNetwork200ResponseNetworkConfig) HasSubnetDhcpLeaseTime() bool`

HasSubnetDhcpLeaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


