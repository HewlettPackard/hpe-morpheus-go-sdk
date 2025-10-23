# AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer

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
**Volumes** | Pointer to [**[]AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerVolumesInner**](AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerVolumesInner.md) |  | [optional] 

## Methods

### NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer

`func NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer() *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer`

NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer instantiates a new AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerWithDefaults

`func NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerWithDefaults() *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer`

NewAddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerWithDefaults instantiates a new AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAccount() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAccountOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetAccount(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetOwner() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetOwnerOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetOwner(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetZone() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetZoneOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetZone(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerInstance)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPlan() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPlanOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetPlan(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetComputeServerType() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetComputeServerTypeOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetComputeServerType(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetSshHost

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSshPort() int32`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSshPortOk() (*int32, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetSshPort(v int32)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetPlatform

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusPercent() int32`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusPercentOk() (*int32, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusPercent(v int32)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPowerState() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetPowerStateOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetPowerState(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerPowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxCores() int32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxCoresOk() (*int32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetMaxCores(v int32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetSourceImage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSourceImage() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetSourceImageOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetSourceImage(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerContainerType)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetServerOs() ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetServerOsOk() (*ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetServerOs(v ListInstances200ResponseAllOfInstancesInnerContainerDetailsInnerServerServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetVolumes() []AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) GetVolumesOk() (*[]AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) SetVolumes(v []AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddInstance200ResponseAllOfOneOfInstanceContainerDetailsInnerServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


