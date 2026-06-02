# UpdateSubnetRequestSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name for the subnet | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Type** | Pointer to [**UpdateSubnetRequestSubnetType**](UpdateSubnetRequestSubnetType.md) |  | [optional] 
**NetworkId** | Pointer to **int64** | The ID of the Network this subnet belongs to. Required when not using the nested route &#x60;/api/networks/{networkId}/subnets&#x60;. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**Cidr** | Pointer to **string** | Subnet CIDR. Used directly by subnet types with &#x60;cidrEditable&#x60; and &#x60;cidrRequired&#x60; (e.g. Google). For Azure subnets, this is derived from &#x60;config.subnetCidr&#x60; and does not need to be set explicitly. | [optional] 
**Active** | Pointer to **bool** | Activate (true) or disable (false) the subnet | [optional] 
**DhcpServer** | Pointer to **bool** | DHCP Server enabled subnet | [optional] 
**AllowStaticOverride** | Pointer to **bool** | Allow IP Override | [optional] 
**Pool** | Pointer to [**UpdateSubnetRequestSubnetPool**](UpdateSubnetRequestSubnetPool.md) |  | [optional] 
**Tenants** | Pointer to [**[]UpdateSubnetRequestSubnetTenantsInner**](UpdateSubnetRequestSubnetTenantsInner.md) | Array of tenant account ID objects that are allowed access | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 

## Methods

### NewUpdateSubnetRequestSubnet

`func NewUpdateSubnetRequestSubnet() *UpdateSubnetRequestSubnet`

NewUpdateSubnetRequestSubnet instantiates a new UpdateSubnetRequestSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateSubnetRequestSubnet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSubnetRequestSubnet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSubnetRequestSubnet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSubnetRequestSubnet) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSubnetRequestSubnet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSubnetRequestSubnet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSubnetRequestSubnet) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSubnetRequestSubnet) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UpdateSubnetRequestSubnet) GetType() UpdateSubnetRequestSubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSubnetRequestSubnet) GetTypeOk() (*UpdateSubnetRequestSubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSubnetRequestSubnet) SetType(v UpdateSubnetRequestSubnetType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSubnetRequestSubnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNetworkId

`func (o *UpdateSubnetRequestSubnet) GetNetworkId() int64`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *UpdateSubnetRequestSubnet) GetNetworkIdOk() (*int64, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *UpdateSubnetRequestSubnet) SetNetworkId(v int64)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *UpdateSubnetRequestSubnet) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateSubnetRequestSubnet) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSubnetRequestSubnet) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSubnetRequestSubnet) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSubnetRequestSubnet) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCidr

`func (o *UpdateSubnetRequestSubnet) GetCidr() string`

GetCidr returns the Cidr field if non-nil, zero value otherwise.

### GetCidrOk

`func (o *UpdateSubnetRequestSubnet) GetCidrOk() (*string, bool)`

GetCidrOk returns a tuple with the Cidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidr

`func (o *UpdateSubnetRequestSubnet) SetCidr(v string)`

SetCidr sets Cidr field to given value.

### HasCidr

`func (o *UpdateSubnetRequestSubnet) HasCidr() bool`

HasCidr returns a boolean if a field has been set.

### GetActive

`func (o *UpdateSubnetRequestSubnet) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateSubnetRequestSubnet) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateSubnetRequestSubnet) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateSubnetRequestSubnet) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetDhcpServer

`func (o *UpdateSubnetRequestSubnet) GetDhcpServer() bool`

GetDhcpServer returns the DhcpServer field if non-nil, zero value otherwise.

### GetDhcpServerOk

`func (o *UpdateSubnetRequestSubnet) GetDhcpServerOk() (*bool, bool)`

GetDhcpServerOk returns a tuple with the DhcpServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServer

`func (o *UpdateSubnetRequestSubnet) SetDhcpServer(v bool)`

SetDhcpServer sets DhcpServer field to given value.

### HasDhcpServer

`func (o *UpdateSubnetRequestSubnet) HasDhcpServer() bool`

HasDhcpServer returns a boolean if a field has been set.

### GetAllowStaticOverride

`func (o *UpdateSubnetRequestSubnet) GetAllowStaticOverride() bool`

GetAllowStaticOverride returns the AllowStaticOverride field if non-nil, zero value otherwise.

### GetAllowStaticOverrideOk

`func (o *UpdateSubnetRequestSubnet) GetAllowStaticOverrideOk() (*bool, bool)`

GetAllowStaticOverrideOk returns a tuple with the AllowStaticOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStaticOverride

`func (o *UpdateSubnetRequestSubnet) SetAllowStaticOverride(v bool)`

SetAllowStaticOverride sets AllowStaticOverride field to given value.

### HasAllowStaticOverride

`func (o *UpdateSubnetRequestSubnet) HasAllowStaticOverride() bool`

HasAllowStaticOverride returns a boolean if a field has been set.

### GetPool

`func (o *UpdateSubnetRequestSubnet) GetPool() UpdateSubnetRequestSubnetPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateSubnetRequestSubnet) GetPoolOk() (*UpdateSubnetRequestSubnetPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateSubnetRequestSubnet) SetPool(v UpdateSubnetRequestSubnetPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateSubnetRequestSubnet) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateSubnetRequestSubnet) GetTenants() []UpdateSubnetRequestSubnetTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateSubnetRequestSubnet) GetTenantsOk() (*[]UpdateSubnetRequestSubnetTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateSubnetRequestSubnet) SetTenants(v []UpdateSubnetRequestSubnetTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateSubnetRequestSubnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateSubnetRequestSubnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateSubnetRequestSubnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateSubnetRequestSubnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateSubnetRequestSubnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateSubnetRequestSubnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateSubnetRequestSubnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateSubnetRequestSubnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateSubnetRequestSubnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateSubnetRequestSubnet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateSubnetRequestSubnet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


