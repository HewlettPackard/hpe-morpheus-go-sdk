# InstanceUpdateInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the instance. | [optional] 
**Description** | Pointer to **string** | Optional description field. | [optional] 
**InstanceContext** | Pointer to **string** | Environment | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]InstanceUpdateInstanceTagsInner**](InstanceUpdateInstanceTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**AddTags** | Pointer to [**[]InstanceUpdateInstanceAddTagsInner**](InstanceUpdateInstanceAddTagsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**RemoveTags** | Pointer to [**[]InstanceUpdateInstanceRemoveTagsInner**](InstanceUpdateInstanceRemoveTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**PowerScheduleType** | Pointer to **int64** | Power schedule ID. | [optional] 
**Site** | Pointer to [**InstanceUpdateInstanceSite**](InstanceUpdateInstanceSite.md) |  | [optional] 
**OwnerId** | Pointer to **int64** | User ID, can be used to change instance owner. | [optional] 
**DisplayName** | Pointer to **string** | Name used in the UI for display | [optional] 

## Methods

### NewInstanceUpdateInstance

`func NewInstanceUpdateInstance() *InstanceUpdateInstance`

NewInstanceUpdateInstance instantiates a new InstanceUpdateInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpdateInstanceWithDefaults

`func NewInstanceUpdateInstanceWithDefaults() *InstanceUpdateInstance`

NewInstanceUpdateInstanceWithDefaults instantiates a new InstanceUpdateInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceUpdateInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceUpdateInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceUpdateInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceUpdateInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceUpdateInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceUpdateInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceUpdateInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceUpdateInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstanceContext

`func (o *InstanceUpdateInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *InstanceUpdateInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *InstanceUpdateInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *InstanceUpdateInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceUpdateInstance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceUpdateInstance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceUpdateInstance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceUpdateInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceUpdateInstance) GetTags() []InstanceUpdateInstanceTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceUpdateInstance) GetTagsOk() (*[]InstanceUpdateInstanceTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceUpdateInstance) SetTags(v []InstanceUpdateInstanceTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceUpdateInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddTags

`func (o *InstanceUpdateInstance) GetAddTags() []InstanceUpdateInstanceAddTagsInner`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *InstanceUpdateInstance) GetAddTagsOk() (*[]InstanceUpdateInstanceAddTagsInner, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *InstanceUpdateInstance) SetAddTags(v []InstanceUpdateInstanceAddTagsInner)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *InstanceUpdateInstance) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *InstanceUpdateInstance) GetRemoveTags() []InstanceUpdateInstanceRemoveTagsInner`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *InstanceUpdateInstance) GetRemoveTagsOk() (*[]InstanceUpdateInstanceRemoveTagsInner, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *InstanceUpdateInstance) SetRemoveTags(v []InstanceUpdateInstanceRemoveTagsInner)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *InstanceUpdateInstance) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *InstanceUpdateInstance) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *InstanceUpdateInstance) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *InstanceUpdateInstance) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *InstanceUpdateInstance) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### GetSite

`func (o *InstanceUpdateInstance) GetSite() InstanceUpdateInstanceSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *InstanceUpdateInstance) GetSiteOk() (*InstanceUpdateInstanceSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *InstanceUpdateInstance) SetSite(v InstanceUpdateInstanceSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *InstanceUpdateInstance) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetOwnerId

`func (o *InstanceUpdateInstance) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *InstanceUpdateInstance) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *InstanceUpdateInstance) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *InstanceUpdateInstance) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceUpdateInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceUpdateInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceUpdateInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceUpdateInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


