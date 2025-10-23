# UpdateInstanceRequestInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Unique name scoped to your account for the instance. | [optional] 
**Description** | Pointer to **string** | Optional description field. | [optional] 
**InstanceContext** | Pointer to **string** | Environment | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**AddTags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Add or update value of Metadata tags, Array of objects having a name and value. | [optional] 
**RemoveTags** | Pointer to [**[]ListInstances200ResponseAllOfInstancesInnerTagsInner**](ListInstances200ResponseAllOfInstancesInnerTagsInner.md) | Remove Metadata tags, Array of objects having a name and an optional value. If value is passed, it must match to be removed. | [optional] 
**PowerScheduleType** | Pointer to **int64** | Power schedule ID. | [optional] 
**Site** | Pointer to [**UpdateInstanceRequestInstanceSite**](UpdateInstanceRequestInstanceSite.md) |  | [optional] 
**OwnerId** | Pointer to **int64** | User ID, can be used to change instance owner. | [optional] 
**DisplayName** | Pointer to **string** | Name used in the UI for display | [optional] 

## Methods

### NewUpdateInstanceRequestInstance

`func NewUpdateInstanceRequestInstance() *UpdateInstanceRequestInstance`

NewUpdateInstanceRequestInstance instantiates a new UpdateInstanceRequestInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceRequestInstanceWithDefaults

`func NewUpdateInstanceRequestInstanceWithDefaults() *UpdateInstanceRequestInstance`

NewUpdateInstanceRequestInstanceWithDefaults instantiates a new UpdateInstanceRequestInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateInstanceRequestInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceRequestInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceRequestInstance) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateInstanceRequestInstance) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateInstanceRequestInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateInstanceRequestInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateInstanceRequestInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateInstanceRequestInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstanceContext

`func (o *UpdateInstanceRequestInstance) GetInstanceContext() string`

GetInstanceContext returns the InstanceContext field if non-nil, zero value otherwise.

### GetInstanceContextOk

`func (o *UpdateInstanceRequestInstance) GetInstanceContextOk() (*string, bool)`

GetInstanceContextOk returns a tuple with the InstanceContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceContext

`func (o *UpdateInstanceRequestInstance) SetInstanceContext(v string)`

SetInstanceContext sets InstanceContext field to given value.

### HasInstanceContext

`func (o *UpdateInstanceRequestInstance) HasInstanceContext() bool`

HasInstanceContext returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateInstanceRequestInstance) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateInstanceRequestInstance) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateInstanceRequestInstance) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateInstanceRequestInstance) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *UpdateInstanceRequestInstance) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateInstanceRequestInstance) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateInstanceRequestInstance) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateInstanceRequestInstance) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetAddTags

`func (o *UpdateInstanceRequestInstance) GetAddTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetAddTags returns the AddTags field if non-nil, zero value otherwise.

### GetAddTagsOk

`func (o *UpdateInstanceRequestInstance) GetAddTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetAddTagsOk returns a tuple with the AddTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddTags

`func (o *UpdateInstanceRequestInstance) SetAddTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetAddTags sets AddTags field to given value.

### HasAddTags

`func (o *UpdateInstanceRequestInstance) HasAddTags() bool`

HasAddTags returns a boolean if a field has been set.

### GetRemoveTags

`func (o *UpdateInstanceRequestInstance) GetRemoveTags() []ListInstances200ResponseAllOfInstancesInnerTagsInner`

GetRemoveTags returns the RemoveTags field if non-nil, zero value otherwise.

### GetRemoveTagsOk

`func (o *UpdateInstanceRequestInstance) GetRemoveTagsOk() (*[]ListInstances200ResponseAllOfInstancesInnerTagsInner, bool)`

GetRemoveTagsOk returns a tuple with the RemoveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveTags

`func (o *UpdateInstanceRequestInstance) SetRemoveTags(v []ListInstances200ResponseAllOfInstancesInnerTagsInner)`

SetRemoveTags sets RemoveTags field to given value.

### HasRemoveTags

`func (o *UpdateInstanceRequestInstance) HasRemoveTags() bool`

HasRemoveTags returns a boolean if a field has been set.

### GetPowerScheduleType

`func (o *UpdateInstanceRequestInstance) GetPowerScheduleType() int64`

GetPowerScheduleType returns the PowerScheduleType field if non-nil, zero value otherwise.

### GetPowerScheduleTypeOk

`func (o *UpdateInstanceRequestInstance) GetPowerScheduleTypeOk() (*int64, bool)`

GetPowerScheduleTypeOk returns a tuple with the PowerScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerScheduleType

`func (o *UpdateInstanceRequestInstance) SetPowerScheduleType(v int64)`

SetPowerScheduleType sets PowerScheduleType field to given value.

### HasPowerScheduleType

`func (o *UpdateInstanceRequestInstance) HasPowerScheduleType() bool`

HasPowerScheduleType returns a boolean if a field has been set.

### GetSite

`func (o *UpdateInstanceRequestInstance) GetSite() UpdateInstanceRequestInstanceSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdateInstanceRequestInstance) GetSiteOk() (*UpdateInstanceRequestInstanceSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdateInstanceRequestInstance) SetSite(v UpdateInstanceRequestInstanceSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdateInstanceRequestInstance) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetOwnerId

`func (o *UpdateInstanceRequestInstance) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *UpdateInstanceRequestInstance) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *UpdateInstanceRequestInstance) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *UpdateInstanceRequestInstance) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateInstanceRequestInstance) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateInstanceRequestInstance) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateInstanceRequestInstance) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateInstanceRequestInstance) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


