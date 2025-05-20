# ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**LoadBalancer** | Pointer to [**ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer**](ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer.md) |  | [optional] 
**Instance** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**Sticky** | Pointer to **bool** |  | [optional] 
**SslEnabled** | Pointer to **string** |  | [optional] 
**ExternalAddress** | Pointer to **bool** |  | [optional] 
**BackendPort** | Pointer to **string** |  | [optional] 
**VipType** | Pointer to **string** |  | [optional] 
**VipAddress** | Pointer to **string** |  | [optional] 
**VipHostname** | Pointer to **string** |  | [optional] 
**VipProtocol** | Pointer to **string** |  | [optional] 
**VipScheme** | Pointer to **string** |  | [optional] 
**VipMode** | Pointer to **string** |  | [optional] 
**VipName** | Pointer to **string** |  | [optional] 
**VipPort** | Pointer to **int64** |  | [optional] 
**VipSticky** | Pointer to **string** |  | [optional] 
**VipBalance** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **string** |  | [optional] 
**SourceAddress** | Pointer to **string** |  | [optional] 
**SslCert** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**SslMode** | Pointer to **string** |  | [optional] 
**SslRedirectMode** | Pointer to **string** |  | [optional] 
**VipShared** | Pointer to **bool** |  | [optional] 
**VipDirectAddress** | Pointer to **string** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**PoolName** | Pointer to **string** |  | [optional] 
**Removing** | Pointer to **bool** |  | [optional] 
**VipSource** | Pointer to **string** |  | [optional] 
**ExtraConfig** | Pointer to **string** |  | [optional] 
**ServiceAccess** | Pointer to **string** |  | [optional] 
**NetworkId** | Pointer to **string** |  | [optional] 
**SubnetId** | Pointer to **string** |  | [optional] 
**ExternalPortId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**VipStatus** | Pointer to **string** |  | [optional] 

## Methods

### NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner

`func NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner() *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner`

NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner instantiates a new ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerWithDefaults

`func NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerWithDefaults() *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner`

NewListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInnerWithDefaults instantiates a new ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoadBalancer

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetLoadBalancer() ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer`

GetLoadBalancer returns the LoadBalancer field if non-nil, zero value otherwise.

### GetLoadBalancerOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetLoadBalancerOk() (*ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer, bool)`

GetLoadBalancerOk returns a tuple with the LoadBalancer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancer

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetLoadBalancer(v ListLoadBalancerMonitors200ResponseAllOfLoadBalancerMonitorsInnerLoadBalancer)`

SetLoadBalancer sets LoadBalancer field to given value.

### HasLoadBalancer

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasLoadBalancer() bool`

HasLoadBalancer returns a boolean if a field has been set.

### GetInstance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetInstance(v string)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetDescription

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActive

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSticky() bool`

GetSticky returns the Sticky field if non-nil, zero value otherwise.

### GetStickyOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetStickyOk() (*bool, bool)`

GetStickyOk returns a tuple with the Sticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSticky(v bool)`

SetSticky sets Sticky field to given value.

### HasSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSticky() bool`

HasSticky returns a boolean if a field has been set.

### GetSslEnabled

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslEnabled() string`

GetSslEnabled returns the SslEnabled field if non-nil, zero value otherwise.

### GetSslEnabledOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslEnabledOk() (*string, bool)`

GetSslEnabledOk returns a tuple with the SslEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslEnabled

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSslEnabled(v string)`

SetSslEnabled sets SslEnabled field to given value.

### HasSslEnabled

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSslEnabled() bool`

HasSslEnabled returns a boolean if a field has been set.

### GetExternalAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalAddress() bool`

GetExternalAddress returns the ExternalAddress field if non-nil, zero value otherwise.

### GetExternalAddressOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalAddressOk() (*bool, bool)`

GetExternalAddressOk returns a tuple with the ExternalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetExternalAddress(v bool)`

SetExternalAddress sets ExternalAddress field to given value.

### HasExternalAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasExternalAddress() bool`

HasExternalAddress returns a boolean if a field has been set.

### GetBackendPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetBackendPort() string`

GetBackendPort returns the BackendPort field if non-nil, zero value otherwise.

### GetBackendPortOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetBackendPortOk() (*string, bool)`

GetBackendPortOk returns a tuple with the BackendPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackendPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetBackendPort(v string)`

SetBackendPort sets BackendPort field to given value.

### HasBackendPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasBackendPort() bool`

HasBackendPort returns a boolean if a field has been set.

### GetVipType

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipType() string`

GetVipType returns the VipType field if non-nil, zero value otherwise.

### GetVipTypeOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipTypeOk() (*string, bool)`

GetVipTypeOk returns a tuple with the VipType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipType

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipType(v string)`

SetVipType sets VipType field to given value.

### HasVipType

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipType() bool`

HasVipType returns a boolean if a field has been set.

### GetVipAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipHostname

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### GetVipProtocol

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipScheme

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipScheme() string`

GetVipScheme returns the VipScheme field if non-nil, zero value otherwise.

### GetVipSchemeOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipSchemeOk() (*string, bool)`

GetVipSchemeOk returns a tuple with the VipScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipScheme

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipScheme(v string)`

SetVipScheme sets VipScheme field to given value.

### HasVipScheme

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipScheme() bool`

HasVipScheme returns a boolean if a field has been set.

### GetVipMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipMode() string`

GetVipMode returns the VipMode field if non-nil, zero value otherwise.

### GetVipModeOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipModeOk() (*string, bool)`

GetVipModeOk returns a tuple with the VipMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipMode(v string)`

SetVipMode sets VipMode field to given value.

### HasVipMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipMode() bool`

HasVipMode returns a boolean if a field has been set.

### GetVipName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetVipPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipPort() int64`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipPortOk() (*int64, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipPort(v int64)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipSticky() string`

GetVipSticky returns the VipSticky field if non-nil, zero value otherwise.

### GetVipStickyOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipStickyOk() (*string, bool)`

GetVipStickyOk returns a tuple with the VipSticky field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipSticky(v string)`

SetVipSticky sets VipSticky field to given value.

### HasVipSticky

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipSticky() bool`

HasVipSticky returns a boolean if a field has been set.

### GetVipBalance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipBalance() string`

GetVipBalance returns the VipBalance field if non-nil, zero value otherwise.

### GetVipBalanceOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipBalanceOk() (*string, bool)`

GetVipBalanceOk returns a tuple with the VipBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipBalance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipBalance(v string)`

SetVipBalance sets VipBalance field to given value.

### HasVipBalance

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipBalance() bool`

HasVipBalance returns a boolean if a field has been set.

### GetServicePort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServicePort() string`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServicePortOk() (*string, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetServicePort(v string)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetSourceAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSourceAddress() string`

GetSourceAddress returns the SourceAddress field if non-nil, zero value otherwise.

### GetSourceAddressOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSourceAddressOk() (*string, bool)`

GetSourceAddressOk returns a tuple with the SourceAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSourceAddress(v string)`

SetSourceAddress sets SourceAddress field to given value.

### HasSourceAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSourceAddress() bool`

HasSourceAddress returns a boolean if a field has been set.

### GetSslCert

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslCert() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslCertOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSslCert(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslMode() string`

GetSslMode returns the SslMode field if non-nil, zero value otherwise.

### GetSslModeOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslModeOk() (*string, bool)`

GetSslModeOk returns a tuple with the SslMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSslMode(v string)`

SetSslMode sets SslMode field to given value.

### HasSslMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSslMode() bool`

HasSslMode returns a boolean if a field has been set.

### GetSslRedirectMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslRedirectMode() string`

GetSslRedirectMode returns the SslRedirectMode field if non-nil, zero value otherwise.

### GetSslRedirectModeOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSslRedirectModeOk() (*string, bool)`

GetSslRedirectModeOk returns a tuple with the SslRedirectMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslRedirectMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSslRedirectMode(v string)`

SetSslRedirectMode sets SslRedirectMode field to given value.

### HasSslRedirectMode

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSslRedirectMode() bool`

HasSslRedirectMode returns a boolean if a field has been set.

### GetVipShared

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipShared() bool`

GetVipShared returns the VipShared field if non-nil, zero value otherwise.

### GetVipSharedOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipSharedOk() (*bool, bool)`

GetVipSharedOk returns a tuple with the VipShared field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipShared

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipShared(v bool)`

SetVipShared sets VipShared field to given value.

### HasVipShared

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipShared() bool`

HasVipShared returns a boolean if a field has been set.

### GetVipDirectAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipDirectAddress() string`

GetVipDirectAddress returns the VipDirectAddress field if non-nil, zero value otherwise.

### GetVipDirectAddressOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipDirectAddressOk() (*string, bool)`

GetVipDirectAddressOk returns a tuple with the VipDirectAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipDirectAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipDirectAddress(v string)`

SetVipDirectAddress sets VipDirectAddress field to given value.

### HasVipDirectAddress

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipDirectAddress() bool`

HasVipDirectAddress returns a boolean if a field has been set.

### GetServerName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPoolName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetRemoving

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetRemoving() bool`

GetRemoving returns the Removing field if non-nil, zero value otherwise.

### GetRemovingOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetRemovingOk() (*bool, bool)`

GetRemovingOk returns a tuple with the Removing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoving

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetRemoving(v bool)`

SetRemoving sets Removing field to given value.

### HasRemoving

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasRemoving() bool`

HasRemoving returns a boolean if a field has been set.

### GetVipSource

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipSource() string`

GetVipSource returns the VipSource field if non-nil, zero value otherwise.

### GetVipSourceOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipSourceOk() (*string, bool)`

GetVipSourceOk returns a tuple with the VipSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipSource

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipSource(v string)`

SetVipSource sets VipSource field to given value.

### HasVipSource

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipSource() bool`

HasVipSource returns a boolean if a field has been set.

### GetExtraConfig

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### GetServiceAccess

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetNetworkId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetSubnetId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSubnetId() string`

GetSubnetId returns the SubnetId field if non-nil, zero value otherwise.

### GetSubnetIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetSubnetIdOk() (*string, bool)`

GetSubnetIdOk returns a tuple with the SubnetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnetId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetSubnetId(v string)`

SetSubnetId sets SubnetId field to given value.

### HasSubnetId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasSubnetId() bool`

HasSubnetId returns a boolean if a field has been set.

### GetExternalPortId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalPortId() string`

GetExternalPortId returns the ExternalPortId field if non-nil, zero value otherwise.

### GetExternalPortIdOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetExternalPortIdOk() (*string, bool)`

GetExternalPortIdOk returns a tuple with the ExternalPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPortId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetExternalPortId(v string)`

SetExternalPortId sets ExternalPortId field to given value.

### HasExternalPortId

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasExternalPortId() bool`

HasExternalPortId returns a boolean if a field has been set.

### GetStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVipStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipStatus() string`

GetVipStatus returns the VipStatus field if non-nil, zero value otherwise.

### GetVipStatusOk

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) GetVipStatusOk() (*string, bool)`

GetVipStatusOk returns a tuple with the VipStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) SetVipStatus(v string)`

SetVipStatus sets VipStatus field to given value.

### HasVipStatus

`func (o *ListLoadBalancerVirtualServers200ResponseAllOfLoadBalancerInstancesInner) HasVipStatus() bool`

HasVipStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


