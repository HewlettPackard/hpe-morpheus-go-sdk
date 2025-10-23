# UpdateCloudsRequestZone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the cloud | 
**Description** | Pointer to **string** | Optional description field if you want to put more info there | [optional] 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Location** | Pointer to **string** | Optional location for your cloud | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**ZoneType** | **string** | Map containing code or id of the cloud type | [default to "standard"]
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
**DefaultCloudLogos** | Pointer to **bool** | Can be used to clear any custom logo and darkLogo, reverting to the defaults for the cloud type | [optional] 
**Credential** | **map[string]interface{}** | Map containing Credential ID. &#x60;local&#x60; means use the values set in the local cloud config instead of associating a credential. | [default to {"type":"local"}]

## Methods

### NewUpdateCloudsRequestZone

`func NewUpdateCloudsRequestZone(name string, zoneType string, groupId int64, credential map[string]interface{}, ) *UpdateCloudsRequestZone`

NewUpdateCloudsRequestZone instantiates a new UpdateCloudsRequestZone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudsRequestZoneWithDefaults

`func NewUpdateCloudsRequestZoneWithDefaults() *UpdateCloudsRequestZone`

NewUpdateCloudsRequestZoneWithDefaults instantiates a new UpdateCloudsRequestZone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCloudsRequestZone) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCloudsRequestZone) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCloudsRequestZone) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateCloudsRequestZone) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCloudsRequestZone) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCloudsRequestZone) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCloudsRequestZone) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateCloudsRequestZone) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateCloudsRequestZone) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateCloudsRequestZone) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateCloudsRequestZone) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateCloudsRequestZone) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateCloudsRequestZone) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateCloudsRequestZone) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateCloudsRequestZone) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *UpdateCloudsRequestZone) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *UpdateCloudsRequestZone) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *UpdateCloudsRequestZone) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *UpdateCloudsRequestZone) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateCloudsRequestZone) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateCloudsRequestZone) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateCloudsRequestZone) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateCloudsRequestZone) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneType

`func (o *UpdateCloudsRequestZone) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *UpdateCloudsRequestZone) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *UpdateCloudsRequestZone) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.


### GetGroupId

`func (o *UpdateCloudsRequestZone) GetGroupId() int64`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *UpdateCloudsRequestZone) GetGroupIdOk() (*int64, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *UpdateCloudsRequestZone) SetGroupId(v int64)`

SetGroupId sets GroupId field to given value.


### GetAccountId

`func (o *UpdateCloudsRequestZone) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateCloudsRequestZone) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateCloudsRequestZone) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateCloudsRequestZone) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateCloudsRequestZone) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateCloudsRequestZone) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateCloudsRequestZone) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateCloudsRequestZone) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoRecoverPowerState

`func (o *UpdateCloudsRequestZone) GetAutoRecoverPowerState() bool`

GetAutoRecoverPowerState returns the AutoRecoverPowerState field if non-nil, zero value otherwise.

### GetAutoRecoverPowerStateOk

`func (o *UpdateCloudsRequestZone) GetAutoRecoverPowerStateOk() (*bool, bool)`

GetAutoRecoverPowerStateOk returns a tuple with the AutoRecoverPowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRecoverPowerState

`func (o *UpdateCloudsRequestZone) SetAutoRecoverPowerState(v bool)`

SetAutoRecoverPowerState sets AutoRecoverPowerState field to given value.

### HasAutoRecoverPowerState

`func (o *UpdateCloudsRequestZone) HasAutoRecoverPowerState() bool`

HasAutoRecoverPowerState returns a boolean if a field has been set.

### GetScalePriority

`func (o *UpdateCloudsRequestZone) GetScalePriority() int64`

GetScalePriority returns the ScalePriority field if non-nil, zero value otherwise.

### GetScalePriorityOk

`func (o *UpdateCloudsRequestZone) GetScalePriorityOk() (*int64, bool)`

GetScalePriorityOk returns a tuple with the ScalePriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScalePriority

`func (o *UpdateCloudsRequestZone) SetScalePriority(v int64)`

SetScalePriority sets ScalePriority field to given value.

### HasScalePriority

`func (o *UpdateCloudsRequestZone) HasScalePriority() bool`

HasScalePriority returns a boolean if a field has been set.

### GetDefaultDatastoreSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultDatastoreSyncActive() bool`

GetDefaultDatastoreSyncActive returns the DefaultDatastoreSyncActive field if non-nil, zero value otherwise.

### GetDefaultDatastoreSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultDatastoreSyncActiveOk() (*bool, bool)`

GetDefaultDatastoreSyncActiveOk returns a tuple with the DefaultDatastoreSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDatastoreSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultDatastoreSyncActive(v bool)`

SetDefaultDatastoreSyncActive sets DefaultDatastoreSyncActive field to given value.

### HasDefaultDatastoreSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultDatastoreSyncActive() bool`

HasDefaultDatastoreSyncActive returns a boolean if a field has been set.

### GetDefaultNetworkSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultNetworkSyncActive() bool`

GetDefaultNetworkSyncActive returns the DefaultNetworkSyncActive field if non-nil, zero value otherwise.

### GetDefaultNetworkSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultNetworkSyncActiveOk() (*bool, bool)`

GetDefaultNetworkSyncActiveOk returns a tuple with the DefaultNetworkSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultNetworkSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultNetworkSyncActive(v bool)`

SetDefaultNetworkSyncActive sets DefaultNetworkSyncActive field to given value.

### HasDefaultNetworkSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultNetworkSyncActive() bool`

HasDefaultNetworkSyncActive returns a boolean if a field has been set.

### GetDefaultFolderSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultFolderSyncActive() bool`

GetDefaultFolderSyncActive returns the DefaultFolderSyncActive field if non-nil, zero value otherwise.

### GetDefaultFolderSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultFolderSyncActiveOk() (*bool, bool)`

GetDefaultFolderSyncActiveOk returns a tuple with the DefaultFolderSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultFolderSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultFolderSyncActive(v bool)`

SetDefaultFolderSyncActive sets DefaultFolderSyncActive field to given value.

### HasDefaultFolderSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultFolderSyncActive() bool`

HasDefaultFolderSyncActive returns a boolean if a field has been set.

### GetDefaultSecurityGroupSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultSecurityGroupSyncActive() bool`

GetDefaultSecurityGroupSyncActive returns the DefaultSecurityGroupSyncActive field if non-nil, zero value otherwise.

### GetDefaultSecurityGroupSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultSecurityGroupSyncActiveOk() (*bool, bool)`

GetDefaultSecurityGroupSyncActiveOk returns a tuple with the DefaultSecurityGroupSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityGroupSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultSecurityGroupSyncActive(v bool)`

SetDefaultSecurityGroupSyncActive sets DefaultSecurityGroupSyncActive field to given value.

### HasDefaultSecurityGroupSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultSecurityGroupSyncActive() bool`

HasDefaultSecurityGroupSyncActive returns a boolean if a field has been set.

### GetDefaultPoolSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultPoolSyncActive() bool`

GetDefaultPoolSyncActive returns the DefaultPoolSyncActive field if non-nil, zero value otherwise.

### GetDefaultPoolSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultPoolSyncActiveOk() (*bool, bool)`

GetDefaultPoolSyncActiveOk returns a tuple with the DefaultPoolSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPoolSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultPoolSyncActive(v bool)`

SetDefaultPoolSyncActive sets DefaultPoolSyncActive field to given value.

### HasDefaultPoolSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultPoolSyncActive() bool`

HasDefaultPoolSyncActive returns a boolean if a field has been set.

### GetDefaultPlanSyncActive

`func (o *UpdateCloudsRequestZone) GetDefaultPlanSyncActive() bool`

GetDefaultPlanSyncActive returns the DefaultPlanSyncActive field if non-nil, zero value otherwise.

### GetDefaultPlanSyncActiveOk

`func (o *UpdateCloudsRequestZone) GetDefaultPlanSyncActiveOk() (*bool, bool)`

GetDefaultPlanSyncActiveOk returns a tuple with the DefaultPlanSyncActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPlanSyncActive

`func (o *UpdateCloudsRequestZone) SetDefaultPlanSyncActive(v bool)`

SetDefaultPlanSyncActive sets DefaultPlanSyncActive field to given value.

### HasDefaultPlanSyncActive

`func (o *UpdateCloudsRequestZone) HasDefaultPlanSyncActive() bool`

HasDefaultPlanSyncActive returns a boolean if a field has been set.

### GetLinkedAccountId

`func (o *UpdateCloudsRequestZone) GetLinkedAccountId() int64`

GetLinkedAccountId returns the LinkedAccountId field if non-nil, zero value otherwise.

### GetLinkedAccountIdOk

`func (o *UpdateCloudsRequestZone) GetLinkedAccountIdOk() (*int64, bool)`

GetLinkedAccountIdOk returns a tuple with the LinkedAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkedAccountId

`func (o *UpdateCloudsRequestZone) SetLinkedAccountId(v int64)`

SetLinkedAccountId sets LinkedAccountId field to given value.

### HasLinkedAccountId

`func (o *UpdateCloudsRequestZone) HasLinkedAccountId() bool`

HasLinkedAccountId returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateCloudsRequestZone) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateCloudsRequestZone) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateCloudsRequestZone) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateCloudsRequestZone) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetSecurityMode

`func (o *UpdateCloudsRequestZone) GetSecurityMode() string`

GetSecurityMode returns the SecurityMode field if non-nil, zero value otherwise.

### GetSecurityModeOk

`func (o *UpdateCloudsRequestZone) GetSecurityModeOk() (*string, bool)`

GetSecurityModeOk returns a tuple with the SecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityMode

`func (o *UpdateCloudsRequestZone) SetSecurityMode(v string)`

SetSecurityMode sets SecurityMode field to given value.

### HasSecurityMode

`func (o *UpdateCloudsRequestZone) HasSecurityMode() bool`

HasSecurityMode returns a boolean if a field has been set.

### GetDefaultCloudLogos

`func (o *UpdateCloudsRequestZone) GetDefaultCloudLogos() bool`

GetDefaultCloudLogos returns the DefaultCloudLogos field if non-nil, zero value otherwise.

### GetDefaultCloudLogosOk

`func (o *UpdateCloudsRequestZone) GetDefaultCloudLogosOk() (*bool, bool)`

GetDefaultCloudLogosOk returns a tuple with the DefaultCloudLogos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloudLogos

`func (o *UpdateCloudsRequestZone) SetDefaultCloudLogos(v bool)`

SetDefaultCloudLogos sets DefaultCloudLogos field to given value.

### HasDefaultCloudLogos

`func (o *UpdateCloudsRequestZone) HasDefaultCloudLogos() bool`

HasDefaultCloudLogos returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateCloudsRequestZone) GetCredential() map[string]interface{}`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateCloudsRequestZone) GetCredentialOk() (*map[string]interface{}, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateCloudsRequestZone) SetCredential(v map[string]interface{})`

SetCredential sets Credential field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


