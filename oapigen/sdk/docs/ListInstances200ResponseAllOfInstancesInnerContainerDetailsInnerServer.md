# ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner.md) |  | [optional] 
**Zone** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance.md) |  | [optional] 
**Plan** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **NullableString** |  | [optional] 
**SshPort** | Pointer to **int32** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **NullableInt32** |  | [optional] 
**StatusEta** | Pointer to **NullableInt64** |  | [optional] 
**PowerState** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState.md) |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **NullableTime** |  | [optional] 
**MaxCores** | Pointer to **int32** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**SourceImage** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType.md) |  | [optional] 
**ServerOs** | Pointer to [**ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerVolumesInner**](ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerVolumesInner.md) |  | [optional] 

## Methods

### NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer

`func NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer() *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer`

NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer instantiates a new ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerWithDefaults

`func NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerWithDefaults() *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer`

NewListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAccount() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAccountOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetAccount(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetOwner() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetOwnerOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetOwner(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetZone() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetZoneOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetZone(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPlan() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPlanOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetPlan(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetComputeServerType() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetComputeServerTypeOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetComputeServerType(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetSshHost

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetPlatform

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusPercent() int32`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusPercentOk() (*int32, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusPercent(v int32)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPowerState() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetPowerStateOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetPowerState(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetSourceImage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSourceImage() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetSourceImageOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetSourceImage(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetServerOs() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetServerOsOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetServerOs(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetVolumes() []ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) GetVolumesOk() (*[]ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) SetVolumes(v []ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


