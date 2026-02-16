# InstanceContainerServer1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**InternalId** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**InstanceContainerServer1Account**](InstanceContainerServer1Account.md) |  | [optional] 
**Owner** | Pointer to [**InstanceContainerServer1Owner**](InstanceContainerServer1Owner.md) |  | [optional] 
**Zone** | Pointer to [**InstanceContainerServer1Zone**](InstanceContainerServer1Zone.md) |  | [optional] 
**Plan** | Pointer to [**InstanceContainerServer1Plan**](InstanceContainerServer1Plan.md) |  | [optional] 
**ComputeServerType** | Pointer to [**InstanceContainerServer1ComputeServerType**](InstanceContainerServer1ComputeServerType.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**ResourcePoolId** | Pointer to **int64** |  | [optional] 
**SshHost** | Pointer to **NullableString** |  | [optional] 
**SshPort** | Pointer to **int64** |  | [optional] 
**ExternalIp** | Pointer to **NullableString** |  | [optional] 
**InternalIp** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusPercent** | Pointer to **NullableInt64** |  | [optional] 
**StatusEta** | Pointer to **NullableInt64** |  | [optional] 
**PowerState** | Pointer to [**InstanceContainerServer1PowerState**](InstanceContainerServer1PowerState.md) |  | [optional] 
**AgentInstalled** | Pointer to **bool** |  | [optional] 
**LastAgentUpdate** | Pointer to **NullableTime** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**SourceImage** | Pointer to [**InstanceContainerServer1SourceImage**](InstanceContainerServer1SourceImage.md) |  | [optional] 
**ServerOs** | Pointer to [**InstanceContainerServer1ServerOs**](InstanceContainerServer1ServerOs.md) |  | [optional] 
**Volumes** | Pointer to [**[]InstanceContainerServerVolume1**](InstanceContainerServerVolume1.md) |  | [optional] 
**Interfaces** | Pointer to [**[]InstanceContainerServerInterfacesInner1**](InstanceContainerServerInterfacesInner1.md) |  | [optional] 

## Methods

### NewInstanceContainerServer1

`func NewInstanceContainerServer1() *InstanceContainerServer1`

NewInstanceContainerServer1 instantiates a new InstanceContainerServer1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceContainerServer1WithDefaults

`func NewInstanceContainerServer1WithDefaults() *InstanceContainerServer1`

NewInstanceContainerServer1WithDefaults instantiates a new InstanceContainerServer1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceContainerServer1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceContainerServer1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceContainerServer1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceContainerServer1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *InstanceContainerServer1) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *InstanceContainerServer1) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *InstanceContainerServer1) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *InstanceContainerServer1) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExternalId

`func (o *InstanceContainerServer1) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceContainerServer1) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceContainerServer1) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InstanceContainerServer1) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InstanceContainerServer1) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InstanceContainerServer1) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetInternalId

`func (o *InstanceContainerServer1) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *InstanceContainerServer1) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *InstanceContainerServer1) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *InstanceContainerServer1) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### SetInternalIdNil

`func (o *InstanceContainerServer1) SetInternalIdNil(b bool)`

 SetInternalIdNil sets the value for InternalId to be an explicit nil

### UnsetInternalId
`func (o *InstanceContainerServer1) UnsetInternalId()`

UnsetInternalId ensures that no value is present for InternalId, not even an explicit nil
### GetHostName

`func (o *InstanceContainerServer1) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceContainerServer1) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceContainerServer1) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceContainerServer1) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetAccountId

`func (o *InstanceContainerServer1) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstanceContainerServer1) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstanceContainerServer1) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstanceContainerServer1) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceContainerServer1) GetAccount() InstanceContainerServer1Account`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceContainerServer1) GetAccountOk() (*InstanceContainerServer1Account, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceContainerServer1) SetAccount(v InstanceContainerServer1Account)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceContainerServer1) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetOwner

`func (o *InstanceContainerServer1) GetOwner() InstanceContainerServer1Owner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *InstanceContainerServer1) GetOwnerOk() (*InstanceContainerServer1Owner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *InstanceContainerServer1) SetOwner(v InstanceContainerServer1Owner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *InstanceContainerServer1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetZone

`func (o *InstanceContainerServer1) GetZone() InstanceContainerServer1Zone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *InstanceContainerServer1) GetZoneOk() (*InstanceContainerServer1Zone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *InstanceContainerServer1) SetZone(v InstanceContainerServer1Zone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *InstanceContainerServer1) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceContainerServer1) GetPlan() InstanceContainerServer1Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceContainerServer1) GetPlanOk() (*InstanceContainerServer1Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceContainerServer1) SetPlan(v InstanceContainerServer1Plan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceContainerServer1) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetComputeServerType

`func (o *InstanceContainerServer1) GetComputeServerType() InstanceContainerServer1ComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *InstanceContainerServer1) GetComputeServerTypeOk() (*InstanceContainerServer1ComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *InstanceContainerServer1) SetComputeServerType(v InstanceContainerServer1ComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *InstanceContainerServer1) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetVisibility

`func (o *InstanceContainerServer1) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InstanceContainerServer1) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InstanceContainerServer1) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InstanceContainerServer1) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceContainerServer1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceContainerServer1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceContainerServer1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceContainerServer1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceContainerServer1) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceContainerServer1) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetZoneId

`func (o *InstanceContainerServer1) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *InstanceContainerServer1) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *InstanceContainerServer1) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *InstanceContainerServer1) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetSiteId

`func (o *InstanceContainerServer1) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *InstanceContainerServer1) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *InstanceContainerServer1) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *InstanceContainerServer1) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *InstanceContainerServer1) GetResourcePoolId() int64`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstanceContainerServer1) GetResourcePoolIdOk() (*int64, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstanceContainerServer1) SetResourcePoolId(v int64)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstanceContainerServer1) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetSshHost

`func (o *InstanceContainerServer1) GetSshHost() string`

GetSshHost returns the SshHost field if non-nil, zero value otherwise.

### GetSshHostOk

`func (o *InstanceContainerServer1) GetSshHostOk() (*string, bool)`

GetSshHostOk returns a tuple with the SshHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHost

`func (o *InstanceContainerServer1) SetSshHost(v string)`

SetSshHost sets SshHost field to given value.

### HasSshHost

`func (o *InstanceContainerServer1) HasSshHost() bool`

HasSshHost returns a boolean if a field has been set.

### SetSshHostNil

`func (o *InstanceContainerServer1) SetSshHostNil(b bool)`

 SetSshHostNil sets the value for SshHost to be an explicit nil

### UnsetSshHost
`func (o *InstanceContainerServer1) UnsetSshHost()`

UnsetSshHost ensures that no value is present for SshHost, not even an explicit nil
### GetSshPort

`func (o *InstanceContainerServer1) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *InstanceContainerServer1) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *InstanceContainerServer1) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *InstanceContainerServer1) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetExternalIp

`func (o *InstanceContainerServer1) GetExternalIp() string`

GetExternalIp returns the ExternalIp field if non-nil, zero value otherwise.

### GetExternalIpOk

`func (o *InstanceContainerServer1) GetExternalIpOk() (*string, bool)`

GetExternalIpOk returns a tuple with the ExternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIp

`func (o *InstanceContainerServer1) SetExternalIp(v string)`

SetExternalIp sets ExternalIp field to given value.

### HasExternalIp

`func (o *InstanceContainerServer1) HasExternalIp() bool`

HasExternalIp returns a boolean if a field has been set.

### SetExternalIpNil

`func (o *InstanceContainerServer1) SetExternalIpNil(b bool)`

 SetExternalIpNil sets the value for ExternalIp to be an explicit nil

### UnsetExternalIp
`func (o *InstanceContainerServer1) UnsetExternalIp()`

UnsetExternalIp ensures that no value is present for ExternalIp, not even an explicit nil
### GetInternalIp

`func (o *InstanceContainerServer1) GetInternalIp() string`

GetInternalIp returns the InternalIp field if non-nil, zero value otherwise.

### GetInternalIpOk

`func (o *InstanceContainerServer1) GetInternalIpOk() (*string, bool)`

GetInternalIpOk returns a tuple with the InternalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalIp

`func (o *InstanceContainerServer1) SetInternalIp(v string)`

SetInternalIp sets InternalIp field to given value.

### HasInternalIp

`func (o *InstanceContainerServer1) HasInternalIp() bool`

HasInternalIp returns a boolean if a field has been set.

### SetInternalIpNil

`func (o *InstanceContainerServer1) SetInternalIpNil(b bool)`

 SetInternalIpNil sets the value for InternalIp to be an explicit nil

### UnsetInternalIp
`func (o *InstanceContainerServer1) UnsetInternalIp()`

UnsetInternalIp ensures that no value is present for InternalIp, not even an explicit nil
### GetPlatform

`func (o *InstanceContainerServer1) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InstanceContainerServer1) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InstanceContainerServer1) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InstanceContainerServer1) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *InstanceContainerServer1) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstanceContainerServer1) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstanceContainerServer1) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstanceContainerServer1) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *InstanceContainerServer1) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *InstanceContainerServer1) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *InstanceContainerServer1) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *InstanceContainerServer1) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *InstanceContainerServer1) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *InstanceContainerServer1) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetErrorMessage

`func (o *InstanceContainerServer1) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *InstanceContainerServer1) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *InstanceContainerServer1) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *InstanceContainerServer1) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *InstanceContainerServer1) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *InstanceContainerServer1) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetStatusDate

`func (o *InstanceContainerServer1) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *InstanceContainerServer1) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *InstanceContainerServer1) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *InstanceContainerServer1) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusPercent

`func (o *InstanceContainerServer1) GetStatusPercent() int64`

GetStatusPercent returns the StatusPercent field if non-nil, zero value otherwise.

### GetStatusPercentOk

`func (o *InstanceContainerServer1) GetStatusPercentOk() (*int64, bool)`

GetStatusPercentOk returns a tuple with the StatusPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPercent

`func (o *InstanceContainerServer1) SetStatusPercent(v int64)`

SetStatusPercent sets StatusPercent field to given value.

### HasStatusPercent

`func (o *InstanceContainerServer1) HasStatusPercent() bool`

HasStatusPercent returns a boolean if a field has been set.

### SetStatusPercentNil

`func (o *InstanceContainerServer1) SetStatusPercentNil(b bool)`

 SetStatusPercentNil sets the value for StatusPercent to be an explicit nil

### UnsetStatusPercent
`func (o *InstanceContainerServer1) UnsetStatusPercent()`

UnsetStatusPercent ensures that no value is present for StatusPercent, not even an explicit nil
### GetStatusEta

`func (o *InstanceContainerServer1) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *InstanceContainerServer1) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *InstanceContainerServer1) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *InstanceContainerServer1) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### SetStatusEtaNil

`func (o *InstanceContainerServer1) SetStatusEtaNil(b bool)`

 SetStatusEtaNil sets the value for StatusEta to be an explicit nil

### UnsetStatusEta
`func (o *InstanceContainerServer1) UnsetStatusEta()`

UnsetStatusEta ensures that no value is present for StatusEta, not even an explicit nil
### GetPowerState

`func (o *InstanceContainerServer1) GetPowerState() InstanceContainerServer1PowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *InstanceContainerServer1) GetPowerStateOk() (*InstanceContainerServer1PowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *InstanceContainerServer1) SetPowerState(v InstanceContainerServer1PowerState)`

SetPowerState sets PowerState field to given value.

### HasPowerState

`func (o *InstanceContainerServer1) HasPowerState() bool`

HasPowerState returns a boolean if a field has been set.

### GetAgentInstalled

`func (o *InstanceContainerServer1) GetAgentInstalled() bool`

GetAgentInstalled returns the AgentInstalled field if non-nil, zero value otherwise.

### GetAgentInstalledOk

`func (o *InstanceContainerServer1) GetAgentInstalledOk() (*bool, bool)`

GetAgentInstalledOk returns a tuple with the AgentInstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentInstalled

`func (o *InstanceContainerServer1) SetAgentInstalled(v bool)`

SetAgentInstalled sets AgentInstalled field to given value.

### HasAgentInstalled

`func (o *InstanceContainerServer1) HasAgentInstalled() bool`

HasAgentInstalled returns a boolean if a field has been set.

### GetLastAgentUpdate

`func (o *InstanceContainerServer1) GetLastAgentUpdate() time.Time`

GetLastAgentUpdate returns the LastAgentUpdate field if non-nil, zero value otherwise.

### GetLastAgentUpdateOk

`func (o *InstanceContainerServer1) GetLastAgentUpdateOk() (*time.Time, bool)`

GetLastAgentUpdateOk returns a tuple with the LastAgentUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAgentUpdate

`func (o *InstanceContainerServer1) SetLastAgentUpdate(v time.Time)`

SetLastAgentUpdate sets LastAgentUpdate field to given value.

### HasLastAgentUpdate

`func (o *InstanceContainerServer1) HasLastAgentUpdate() bool`

HasLastAgentUpdate returns a boolean if a field has been set.

### SetLastAgentUpdateNil

`func (o *InstanceContainerServer1) SetLastAgentUpdateNil(b bool)`

 SetLastAgentUpdateNil sets the value for LastAgentUpdate to be an explicit nil

### UnsetLastAgentUpdate
`func (o *InstanceContainerServer1) UnsetLastAgentUpdate()`

UnsetLastAgentUpdate ensures that no value is present for LastAgentUpdate, not even an explicit nil
### GetMaxCores

`func (o *InstanceContainerServer1) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *InstanceContainerServer1) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *InstanceContainerServer1) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *InstanceContainerServer1) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxMemory

`func (o *InstanceContainerServer1) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *InstanceContainerServer1) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *InstanceContainerServer1) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *InstanceContainerServer1) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxStorage

`func (o *InstanceContainerServer1) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceContainerServer1) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceContainerServer1) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceContainerServer1) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetSourceImage

`func (o *InstanceContainerServer1) GetSourceImage() InstanceContainerServer1SourceImage`

GetSourceImage returns the SourceImage field if non-nil, zero value otherwise.

### GetSourceImageOk

`func (o *InstanceContainerServer1) GetSourceImageOk() (*InstanceContainerServer1SourceImage, bool)`

GetSourceImageOk returns a tuple with the SourceImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceImage

`func (o *InstanceContainerServer1) SetSourceImage(v InstanceContainerServer1SourceImage)`

SetSourceImage sets SourceImage field to given value.

### HasSourceImage

`func (o *InstanceContainerServer1) HasSourceImage() bool`

HasSourceImage returns a boolean if a field has been set.

### GetServerOs

`func (o *InstanceContainerServer1) GetServerOs() InstanceContainerServer1ServerOs`

GetServerOs returns the ServerOs field if non-nil, zero value otherwise.

### GetServerOsOk

`func (o *InstanceContainerServer1) GetServerOsOk() (*InstanceContainerServer1ServerOs, bool)`

GetServerOsOk returns a tuple with the ServerOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerOs

`func (o *InstanceContainerServer1) SetServerOs(v InstanceContainerServer1ServerOs)`

SetServerOs sets ServerOs field to given value.

### HasServerOs

`func (o *InstanceContainerServer1) HasServerOs() bool`

HasServerOs returns a boolean if a field has been set.

### GetVolumes

`func (o *InstanceContainerServer1) GetVolumes() []InstanceContainerServerVolume1`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceContainerServer1) GetVolumesOk() (*[]InstanceContainerServerVolume1, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceContainerServer1) SetVolumes(v []InstanceContainerServerVolume1)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *InstanceContainerServer1) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetInterfaces

`func (o *InstanceContainerServer1) GetInterfaces() []InstanceContainerServerInterfacesInner1`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *InstanceContainerServer1) GetInterfacesOk() (*[]InstanceContainerServerInterfacesInner1, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *InstanceContainerServer1) SetInterfaces(v []InstanceContainerServerInterfacesInner1)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *InstanceContainerServer1) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


