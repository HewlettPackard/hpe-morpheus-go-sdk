# CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceLoadBalancer**](CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceLoadBalancer.md) |  | [optional] 
**Instance** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by type. | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Sticky** | Pointer to **bool** |  | [optional] 
**SslEnabled** | Pointer to **NullableString** |  | [optional] 
**ExternalAddress** | Pointer to **bool** |  | [optional] 
**BackendPort** | Pointer to **NullableString** |  | [optional] 
**VipType** | Pointer to **NullableString** |  | [optional] 
**VipAddress** | Pointer to **string** |  | [optional] 
**VipHostname** | Pointer to **NullableString** |  | [optional] 
**VipProtocol** | Pointer to **string** |  | [optional] 
**VipScheme** | Pointer to **NullableString** |  | [optional] 
**VipMode** | Pointer to **NullableString** |  | [optional] 
**VipName** | Pointer to **string** |  | [optional] 
**VipPort** | Pointer to **int64** |  | [optional] 
**VipPool** | Pointer to **map[string]interface{}** | Network Pool used for VIP address allocation, if set. | [optional] 
**VipSticky** | Pointer to **NullableString** |  | [optional] 
**VipBalance** | Pointer to **NullableString** |  | [optional] 
**ServicePort** | Pointer to **NullableString** |  | [optional] 
**SourceAddress** | Pointer to **NullableString** |  | [optional] 
**SslCert** | Pointer to [**CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslCert**](CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslCert.md) |  | [optional] 
**SslServerCert** | Pointer to [**CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslServerCert**](CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslServerCert.md) |  | [optional] 
**SslMode** | Pointer to **NullableString** |  | [optional] 
**SslRedirectMode** | Pointer to **NullableString** |  | [optional] 
**VipShared** | Pointer to **bool** |  | [optional] 
**VipDirectAddress** | Pointer to **NullableString** |  | [optional] 
**ServerName** | Pointer to **NullableString** |  | [optional] 
**PoolName** | Pointer to **NullableString** |  | [optional] 
**Removing** | Pointer to **bool** |  | [optional] 
**Pool** | Pointer to [**CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstancePool**](CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstancePool.md) |  | [optional] 
**VipSource** | Pointer to **string** |  | [optional] 
**ExtraConfig** | Pointer to **NullableString** |  | [optional] 
**ServiceAccess** | Pointer to **NullableString** |  | [optional] 
**NetworkId** | Pointer to **NullableString** |  | [optional] 
**SubnetId** | Pointer to **NullableString** |  | [optional] 
**ExternalPortId** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VipStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance

`func NewCreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance() *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance`

NewCreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance instantiates a new CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetLoadBalancer() CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetLoadBalancerOk() (*CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetLoadBalancer(v CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetInstance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### SetInstanceNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetInstanceNil(b bool)`

 SetInstanceNil sets the value for Instance to be an explicit nil

### UnsetInstance
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetInstance()`

UnsetInstance ensures that no value is present for Instance, not even an explicit nil
### GetConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetInternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDateCreated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSticky() bool`

GetSticky returns the Sticky field if non-nil, zero value otherwise.

### GetStickyOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetStickyOk() (*bool, bool)`

GetStickyOk returns a tuple with the Sticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSticky(v bool)`

SetSticky sets Sticky field to given value.

### HasSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSticky() bool`

HasSticky returns a boolean if a field has been set.

### GetSslEnabled

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### SetSslEnabledNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslEnabledNil(b bool)`

 SetSslEnabledNil sets the value for SslEnabled to be an explicit nil

### UnsetSslEnabled
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetSslEnabled()`

UnsetSslEnabled ensures that no value is present for SslEnabled, not even an explicit nil
### GetExternalAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalAddress() bool`

GetExternalAddress returns the ExternalAddress field if non-nil, zero value otherwise.

### GetExternalAddressOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalAddressOk() (*bool, bool)`

GetExternalAddressOk returns a tuple with the ExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExternalAddress(v bool)`

SetExternalAddress sets ExternalAddress field to given value.

### HasExternalAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasExternalAddress() bool`

HasExternalAddress returns a boolean if a field has been set.

### GetBackendPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetBackendPort() string`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetBackendPortOk() (*string, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetBackendPort(v string)`

SetBackendPort sets BackendPort field to given value.

### HasBackendPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### SetBackendPortNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetBackendPortNil(b bool)`

 SetBackendPortNil sets the value for BackendPort to be an explicit nil

### UnsetBackendPort
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetBackendPort()`

UnsetBackendPort ensures that no value is present for BackendPort, not even an explicit nil
### GetVipType

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipType() string`

GetVipType returns the VipType field if non-nil, zero value otherwise.

### GetVipTypeOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipTypeOk() (*string, bool)`

GetVipTypeOk returns a tuple with the VipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipType

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipType(v string)`

SetVipType sets VipType field to given value.

### HasVipType

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipType() bool`

HasVipType returns a boolean if a field has been set.

### SetVipTypeNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipTypeNil(b bool)`

 SetVipTypeNil sets the value for VipType to be an explicit nil

### UnsetVipType
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipType()`

UnsetVipType ensures that no value is present for VipType, not even an explicit nil
### GetVipAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipHostname

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### SetVipHostnameNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipHostnameNil(b bool)`

 SetVipHostnameNil sets the value for VipHostname to be an explicit nil

### UnsetVipHostname
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipHostname()`

UnsetVipHostname ensures that no value is present for VipHostname, not even an explicit nil
### GetVipProtocol

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipScheme

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipScheme() string`

GetVipScheme returns the VipScheme field if non-nil, zero value otherwise.

### GetVipSchemeOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipSchemeOk() (*string, bool)`

GetVipSchemeOk returns a tuple with the VipScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipScheme

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipScheme(v string)`

SetVipScheme sets VipScheme field to given value.

### HasVipScheme

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipScheme() bool`

HasVipScheme returns a boolean if a field has been set.

### SetVipSchemeNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipSchemeNil(b bool)`

 SetVipSchemeNil sets the value for VipScheme to be an explicit nil

### UnsetVipScheme
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipScheme()`

UnsetVipScheme ensures that no value is present for VipScheme, not even an explicit nil
### GetVipMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipMode() string`

GetVipMode returns the VipMode field if non-nil, zero value otherwise.

### GetVipModeOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipModeOk() (*string, bool)`

GetVipModeOk returns a tuple with the VipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipMode(v string)`

SetVipMode sets VipMode field to given value.

### HasVipMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipMode() bool`

HasVipMode returns a boolean if a field has been set.

### SetVipModeNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipModeNil(b bool)`

 SetVipModeNil sets the value for VipMode to be an explicit nil

### UnsetVipMode
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipMode()`

UnsetVipMode ensures that no value is present for VipMode, not even an explicit nil
### GetVipName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetVipPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipPort() int64`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipPortOk() (*int64, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipPort(v int64)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipPool() map[string]interface{}`

GetVipPool returns the VipPool field if non-nil, zero value otherwise.

### GetVipPoolOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipPoolOk() (*map[string]interface{}, bool)`

GetVipPoolOk returns a tuple with the VipPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipPool(v map[string]interface{})`

SetVipPool sets VipPool field to given value.

### HasVipPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipPool() bool`

HasVipPool returns a boolean if a field has been set.

### SetVipPoolNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipPoolNil(b bool)`

 SetVipPoolNil sets the value for VipPool to be an explicit nil

### UnsetVipPool
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipPool()`

UnsetVipPool ensures that no value is present for VipPool, not even an explicit nil
### GetVipSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### SetVipStickyNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipStickyNil(b bool)`

 SetVipStickyNil sets the value for VipSticky to be an explicit nil

### UnsetVipSticky
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipSticky()`

UnsetVipSticky ensures that no value is present for VipSticky, not even an explicit nil
### GetVipBalance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### SetVipBalanceNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipBalanceNil(b bool)`

 SetVipBalanceNil sets the value for VipBalance to be an explicit nil

### UnsetVipBalance
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipBalance()`

UnsetVipBalance ensures that no value is present for VipBalance, not even an explicit nil
### GetServicePort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServicePort() string`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServicePortOk() (*string, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServicePort(v string)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetSourceAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### SetSourceAddressNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSourceAddressNil(b bool)`

 SetSourceAddressNil sets the value for SourceAddress to be an explicit nil

### UnsetSourceAddress
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetSourceAddress()`

UnsetSourceAddress ensures that no value is present for SourceAddress, not even an explicit nil
### GetSslCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslCert() CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslCert`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslCertOk() (*CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslCert, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslCert(v CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslCert)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslServerCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslServerCert() CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslServerCert`

GetSslServerCert returns the SslServerCert field if non-nil, zero value otherwise.

### GetSslServerCertOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslServerCertOk() (*CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslServerCert, bool)`

GetSslServerCertOk returns a tuple with the SslServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslServerCert(v CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstanceSslServerCert)`

SetSslServerCert sets SslServerCert field to given value.

### HasSslServerCert

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSslServerCert() bool`

HasSslServerCert returns a boolean if a field has been set.

### GetSslMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### SetSslModeNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslModeNil(b bool)`

 SetSslModeNil sets the value for SslMode to be an explicit nil

### UnsetSslMode
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetSslMode()`

UnsetSslMode ensures that no value is present for SslMode, not even an explicit nil
### GetSslRedirectMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslRedirectMode() string`

GetSslRedirectMode returns the SslRedirectMode field if non-nil, zero value otherwise.

### GetSslRedirectModeOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSslRedirectModeOk() (*string, bool)`

GetSslRedirectModeOk returns a tuple with the SslRedirectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRedirectMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslRedirectMode(v string)`

SetSslRedirectMode sets SslRedirectMode field to given value.

### HasSslRedirectMode

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSslRedirectMode() bool`

HasSslRedirectMode returns a boolean if a field has been set.

### SetSslRedirectModeNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSslRedirectModeNil(b bool)`

 SetSslRedirectModeNil sets the value for SslRedirectMode to be an explicit nil

### UnsetSslRedirectMode
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetSslRedirectMode()`

UnsetSslRedirectMode ensures that no value is present for SslRedirectMode, not even an explicit nil
### GetVipShared

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipShared() bool`

GetVipShared returns the VipShared field if non-nil, zero value otherwise.

### GetVipSharedOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipSharedOk() (*bool, bool)`

GetVipSharedOk returns a tuple with the VipShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipShared

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipShared(v bool)`

SetVipShared sets VipShared field to given value.

### HasVipShared

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipShared() bool`

HasVipShared returns a boolean if a field has been set.

### GetVipDirectAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipDirectAddress() string`

GetVipDirectAddress returns the VipDirectAddress field if non-nil, zero value otherwise.

### GetVipDirectAddressOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipDirectAddressOk() (*string, bool)`

GetVipDirectAddressOk returns a tuple with the VipDirectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipDirectAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipDirectAddress(v string)`

SetVipDirectAddress sets VipDirectAddress field to given value.

### HasVipDirectAddress

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipDirectAddress() bool`

HasVipDirectAddress returns a boolean if a field has been set.

### SetVipDirectAddressNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipDirectAddressNil(b bool)`

 SetVipDirectAddressNil sets the value for VipDirectAddress to be an explicit nil

### UnsetVipDirectAddress
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetVipDirectAddress()`

UnsetVipDirectAddress ensures that no value is present for VipDirectAddress, not even an explicit nil
### GetServerName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetPoolName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### SetPoolNameNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetPoolNameNil(b bool)`

 SetPoolNameNil sets the value for PoolName to be an explicit nil

### UnsetPoolName
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetPoolName()`

UnsetPoolName ensures that no value is present for PoolName, not even an explicit nil
### GetRemoving

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetRemoving() bool`

GetRemoving returns the Removing field if non-nil, zero value otherwise.

### GetRemovingOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetRemovingOk() (*bool, bool)`

GetRemovingOk returns a tuple with the Removing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoving

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetRemoving(v bool)`

SetRemoving sets Removing field to given value.

### HasRemoving

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasRemoving() bool`

HasRemoving returns a boolean if a field has been set.

### GetPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetPool() CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstancePool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetPoolOk() (*CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstancePool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetPool(v CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstancePool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetVipSource

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipSource() string`

GetVipSource returns the VipSource field if non-nil, zero value otherwise.

### GetVipSourceOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipSourceOk() (*string, bool)`

GetVipSourceOk returns a tuple with the VipSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSource

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipSource(v string)`

SetVipSource sets VipSource field to given value.

### HasVipSource

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipSource() bool`

HasVipSource returns a boolean if a field has been set.

### GetExtraConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetServiceAccess

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### SetServiceAccessNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetServiceAccessNil(b bool)`

 SetServiceAccessNil sets the value for ServiceAccess to be an explicit nil

### UnsetServiceAccess
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetServiceAccess()`

UnsetServiceAccess ensures that no value is present for ServiceAccess, not even an explicit nil
### GetNetworkId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### SetNetworkIdNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetNetworkIdNil(b bool)`

 SetNetworkIdNil sets the value for NetworkId to be an explicit nil

### UnsetNetworkId
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetNetworkId()`

UnsetNetworkId ensures that no value is present for NetworkId, not even an explicit nil
### GetSubnetId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### SetSubnetIdNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetSubnetIdNil(b bool)`

 SetSubnetIdNil sets the value for SubnetId to be an explicit nil

### UnsetSubnetId
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetSubnetId()`

UnsetSubnetId ensures that no value is present for SubnetId, not even an explicit nil
### GetExternalPortId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### SetExternalPortIdNil

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetExternalPortIdNil(b bool)`

 SetExternalPortIdNil sets the value for ExternalPortId to be an explicit nil

### UnsetExternalPortId
`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) UnsetExternalPortId()`

UnsetExternalPortId ensures that no value is present for ExternalPortId, not even an explicit nil
### GetStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVipStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipStatus() string`

GetVipStatus returns the VipStatus field if non-nil, zero value otherwise.

### GetVipStatusOk

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) GetVipStatusOk() (*string, bool)`

GetVipStatusOk returns a tuple with the VipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) SetVipStatus(v string)`

SetVipStatus sets VipStatus field to given value.

### HasVipStatus

`func (o *CreateLoadBalancerVirtualServer200ResponseLoadBalancerInstance) HasVipStatus() bool`

HasVipStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


