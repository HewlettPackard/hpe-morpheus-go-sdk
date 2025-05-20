# AddCloudsRequestZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the cloud | 
**Description** | Pointer to **string** | Optional description field if you want to put more info there | [optional] 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Location** | Pointer to **string** | Optional location for your cloud | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**ZoneType** | [**AddCloudsRequestZoneZoneType**](AddCloudsRequestZoneZoneType.md) |  | 
**GroupId** | **int64** | Specifies which Server group this cloud should be assigned to | 
**AccountId** | Pointer to **int64** | Specifies which Tenant this cloud should be assigned to | [optional] 
**Enabled** | Pointer to **bool** | Can be used to disable the cloud | [optional] [default to true]
**AutoRecoverPowerState** | Pointer to **bool** | Automatically Power on VMs | [optional] [default to false]
**ScalePriority** | Pointer to **int64** | Scale Priority | [optional] [default to 1]
**DefaultDatastoreSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new datastores. | [optional] 
**DefaultNetworkSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new networks. | [optional] 
**DefaultFolderSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new folders. | [optional] 
**DefaultSecurityGroupSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new security groups. | [optional] 
**DefaultPoolSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new resource pools. | [optional] 
**DefaultPlanSyncActive** | Pointer to **bool** | Sets the default active state during discovery of new plans. | [optional] 
**LinkedAccountId** | Pointer to **int64** | Linked Account ID (enter commercial ID to get costing for AWS Govcloud) | [optional] 
**Config** | Pointer to **map[string]interface{}** | Map containing zone configuration settings. See the section on specific zone types for details. | [optional] 
**SecurityMode** | Pointer to **string** | host firewall. &#x60;off&#x60; or &#x60;internal&#x60;. a.k.a. \&quot;local firewall\&quot; | [optional] [default to "off"]
**Credential** | Pointer to [**AddCloudsRequestZoneCredential**](AddCloudsRequestZoneCredential.md) |  | [optional] 

## Methods

### NewAddCloudsRequestZone

`func NewAddCloudsRequestZone(name string, zoneType AddCloudsRequestZoneZoneType, groupId int64, ) *AddCloudsRequestZone`

NewAddCloudsRequestZone instantiates a new AddCloudsRequestZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneWithDefaults

`func NewAddCloudsRequestZoneWithDefaults() *AddCloudsRequestZone`

NewAddCloudsRequestZoneWithDefaults instantiates a new AddCloudsRequestZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddCloudsRequestZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCloudsRequestZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCloudsRequestZone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddCloudsRequestZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCloudsRequestZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCloudsRequestZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCloudsRequestZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AddCloudsRequestZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddCloudsRequestZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddCloudsRequestZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddCloudsRequestZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *AddCloudsRequestZone) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddCloudsRequestZone) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddCloudsRequestZone) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddCloudsRequestZone) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *AddCloudsRequestZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddCloudsRequestZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddCloudsRequestZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddCloudsRequestZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVisibility

`func (o *AddCloudsRequestZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddCloudsRequestZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddCloudsRequestZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddCloudsRequestZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneType

`func (o *AddCloudsRequestZone) GetZoneType() AddCloudsRequestZoneZoneType`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *AddCloudsRequestZone) GetZoneTypeOk() (*AddCloudsRequestZoneZoneType, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *AddCloudsRequestZone) SetZoneType(v AddCloudsRequestZoneZoneType)`

SetZoneType sets ZoneType field to given value.


### GetGroupId

`func (o *AddCloudsRequestZone) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AddCloudsRequestZone) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AddCloudsRequestZone) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetAccountId

`func (o *AddCloudsRequestZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddCloudsRequestZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddCloudsRequestZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddCloudsRequestZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *AddCloudsRequestZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCloudsRequestZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCloudsRequestZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCloudsRequestZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *AddCloudsRequestZone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *AddCloudsRequestZone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *AddCloudsRequestZone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *AddCloudsRequestZone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *AddCloudsRequestZone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *AddCloudsRequestZone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *AddCloudsRequestZone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *AddCloudsRequestZone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetDefaultDatastoreSyncActive

`func (o *AddCloudsRequestZone) GetDefaultDatastoreSyncActive() bool`

GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field if non-nil, zero value otherwise.

### GetDefaultDatastoreSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultDatastoreSyncActiveOk() (*bool, bool)`

GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreSyncActive

`func (o *AddCloudsRequestZone) SetDefaultDatastoreSyncActive(v bool)`

SetDefaultDatastoreSyncActive sets DefaultDatastoreSyncActive field to given value.

### HasDefaultDatastoreSyncActive

`func (o *AddCloudsRequestZone) HasDefaultDatastoreSyncActive() bool`

HasDefaultDatastoreSyncActive returns a boolean if a field has been set.

### GetDefaultNetworkSyncActive

`func (o *AddCloudsRequestZone) GetDefaultNetworkSyncActive() bool`

GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field if non-nil, zero value otherwise.

### GetDefaultNetworkSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultNetworkSyncActiveOk() (*bool, bool)`

GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkSyncActive

`func (o *AddCloudsRequestZone) SetDefaultNetworkSyncActive(v bool)`

SetDefaultNetworkSyncActive sets DefaultNetworkSyncActive field to given value.

### HasDefaultNetworkSyncActive

`func (o *AddCloudsRequestZone) HasDefaultNetworkSyncActive() bool`

HasDefaultNetworkSyncActive returns a boolean if a field has been set.

### GetDefaultFolderSyncActive

`func (o *AddCloudsRequestZone) GetDefaultFolderSyncActive() bool`

GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field if non-nil, zero value otherwise.

### GetDefaultFolderSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultFolderSyncActiveOk() (*bool, bool)`

GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderSyncActive

`func (o *AddCloudsRequestZone) SetDefaultFolderSyncActive(v bool)`

SetDefaultFolderSyncActive sets DefaultFolderSyncActive field to given value.

### HasDefaultFolderSyncActive

`func (o *AddCloudsRequestZone) HasDefaultFolderSyncActive() bool`

HasDefaultFolderSyncActive returns a boolean if a field has been set.

### GetDefaultSecurityGroupSyncActive

`func (o *AddCloudsRequestZone) GetDefaultSecurityGroupSyncActive() bool`

GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool)`

GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupSyncActive

`func (o *AddCloudsRequestZone) SetDefaultSecurityGroupSyncActive(v bool)`

SetDefaultSecurityGroupSyncActive sets DefaultSecurityGroupSyncActive field to given value.

### HasDefaultSecurityGroupSyncActive

`func (o *AddCloudsRequestZone) HasDefaultSecurityGroupSyncActive() bool`

HasDefaultSecurityGroupSyncActive returns a boolean if a field has been set.

### GetDefaultPoolSyncActive

`func (o *AddCloudsRequestZone) GetDefaultPoolSyncActive() bool`

GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field if non-nil, zero value otherwise.

### GetDefaultPoolSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultPoolSyncActiveOk() (*bool, bool)`

GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolSyncActive

`func (o *AddCloudsRequestZone) SetDefaultPoolSyncActive(v bool)`

SetDefaultPoolSyncActive sets DefaultPoolSyncActive field to given value.

### HasDefaultPoolSyncActive

`func (o *AddCloudsRequestZone) HasDefaultPoolSyncActive() bool`

HasDefaultPoolSyncActive returns a boolean if a field has been set.

### GetDefaultPlanSyncActive

`func (o *AddCloudsRequestZone) GetDefaultPlanSyncActive() bool`

GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field if non-nil, zero value otherwise.

### GetDefaultPlanSyncActiveOk

`func (o *AddCloudsRequestZone) GetDefaultPlanSyncActiveOk() (*bool, bool)`

GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanSyncActive

`func (o *AddCloudsRequestZone) SetDefaultPlanSyncActive(v bool)`

SetDefaultPlanSyncActive sets DefaultPlanSyncActive field to given value.

### HasDefaultPlanSyncActive

`func (o *AddCloudsRequestZone) HasDefaultPlanSyncActive() bool`

HasDefaultPlanSyncActive returns a boolean if a field has been set.

### GetLinkedAccountId

`func (o *AddCloudsRequestZone) GetLinkedAccountId() int64`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *AddCloudsRequestZone) GetLinkedAccountIdOk() (*int64, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *AddCloudsRequestZone) SetLinkedAccountId(v int64)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *AddCloudsRequestZone) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### GetConfig

`func (o *AddCloudsRequestZone) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCloudsRequestZone) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCloudsRequestZone) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddCloudsRequestZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecurityMode

`func (o *AddCloudsRequestZone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *AddCloudsRequestZone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *AddCloudsRequestZone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *AddCloudsRequestZone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetCredential

`func (o *AddCloudsRequestZone) GetCredential() AddCloudsRequestZoneCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddCloudsRequestZone) GetCredentialOk() (*AddCloudsRequestZoneCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddCloudsRequestZone) SetCredential(v AddCloudsRequestZoneCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddCloudsRequestZone) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


