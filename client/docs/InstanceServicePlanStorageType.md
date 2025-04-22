# InstanceServicePlanStorageType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**MinStorage** | Pointer to **string** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DefaultType** | Pointer to **bool** |  | [optional] 
**CreateDatastore** | Pointer to **string** |  | [optional] 
**Resizable** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **string** |  | [optional] 
**AllowSearch** | Pointer to **bool** |  | [optional] 
**VolumeOptionSource** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**MinIOPS** | Pointer to **string** |  | [optional] 
**MaxIOPS** | Pointer to **string** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**CustomSize** | Pointer to **bool** |  | [optional] 
**AutoDelete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**CustomLabel** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**VolumeCategory** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **string** |  | [optional] 

## Methods

### NewInstanceServicePlanStorageType

`func NewInstanceServicePlanStorageType() *InstanceServicePlanStorageType`

NewInstanceServicePlanStorageType instantiates a new InstanceServicePlanStorageType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceServicePlanStorageTypeWithDefaults

`func NewInstanceServicePlanStorageTypeWithDefaults() *InstanceServicePlanStorageType`

NewInstanceServicePlanStorageTypeWithDefaults instantiates a new InstanceServicePlanStorageType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceServicePlanStorageType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceServicePlanStorageType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceServicePlanStorageType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceServicePlanStorageType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEditable

`func (o *InstanceServicePlanStorageType) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *InstanceServicePlanStorageType) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *InstanceServicePlanStorageType) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *InstanceServicePlanStorageType) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceServicePlanStorageType) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceServicePlanStorageType) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceServicePlanStorageType) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceServicePlanStorageType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *InstanceServicePlanStorageType) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *InstanceServicePlanStorageType) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *InstanceServicePlanStorageType) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *InstanceServicePlanStorageType) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetCode

`func (o *InstanceServicePlanStorageType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceServicePlanStorageType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceServicePlanStorageType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceServicePlanStorageType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVolumeType

`func (o *InstanceServicePlanStorageType) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *InstanceServicePlanStorageType) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *InstanceServicePlanStorageType) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *InstanceServicePlanStorageType) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetMinStorage

`func (o *InstanceServicePlanStorageType) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *InstanceServicePlanStorageType) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *InstanceServicePlanStorageType) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *InstanceServicePlanStorageType) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### GetDeletable

`func (o *InstanceServicePlanStorageType) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *InstanceServicePlanStorageType) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *InstanceServicePlanStorageType) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *InstanceServicePlanStorageType) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefaultType

`func (o *InstanceServicePlanStorageType) GetDefaultType() bool`

GetDefaultType returns the DefaultType field if non-nil, zero value otherwise.

### GetDefaultTypeOk

`func (o *InstanceServicePlanStorageType) GetDefaultTypeOk() (*bool, bool)`

GetDefaultTypeOk returns a tuple with the DefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultType

`func (o *InstanceServicePlanStorageType) SetDefaultType(v bool)`

SetDefaultType sets DefaultType field to given value.

### HasDefaultType

`func (o *InstanceServicePlanStorageType) HasDefaultType() bool`

HasDefaultType returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *InstanceServicePlanStorageType) GetCreateDatastore() string`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *InstanceServicePlanStorageType) GetCreateDatastoreOk() (*string, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *InstanceServicePlanStorageType) SetCreateDatastore(v string)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *InstanceServicePlanStorageType) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### GetResizable

`func (o *InstanceServicePlanStorageType) GetResizable() bool`

GetResizable returns the Resizable field if non-nil, zero value otherwise.

### GetResizableOk

`func (o *InstanceServicePlanStorageType) GetResizableOk() (*bool, bool)`

GetResizableOk returns a tuple with the Resizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizable

`func (o *InstanceServicePlanStorageType) SetResizable(v bool)`

SetResizable sets Resizable field to given value.

### HasResizable

`func (o *InstanceServicePlanStorageType) HasResizable() bool`

HasResizable returns a boolean if a field has been set.

### GetStorageType

`func (o *InstanceServicePlanStorageType) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InstanceServicePlanStorageType) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InstanceServicePlanStorageType) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *InstanceServicePlanStorageType) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetAllowSearch

`func (o *InstanceServicePlanStorageType) GetAllowSearch() bool`

GetAllowSearch returns the AllowSearch field if non-nil, zero value otherwise.

### GetAllowSearchOk

`func (o *InstanceServicePlanStorageType) GetAllowSearchOk() (*bool, bool)`

GetAllowSearchOk returns a tuple with the AllowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearch

`func (o *InstanceServicePlanStorageType) SetAllowSearch(v bool)`

SetAllowSearch sets AllowSearch field to given value.

### HasAllowSearch

`func (o *InstanceServicePlanStorageType) HasAllowSearch() bool`

HasAllowSearch returns a boolean if a field has been set.

### GetVolumeOptionSource

`func (o *InstanceServicePlanStorageType) GetVolumeOptionSource() string`

GetVolumeOptionSource returns the VolumeOptionSource field if non-nil, zero value otherwise.

### GetVolumeOptionSourceOk

`func (o *InstanceServicePlanStorageType) GetVolumeOptionSourceOk() (*string, bool)`

GetVolumeOptionSourceOk returns a tuple with the VolumeOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeOptionSource

`func (o *InstanceServicePlanStorageType) SetVolumeOptionSource(v string)`

SetVolumeOptionSource sets VolumeOptionSource field to given value.

### HasVolumeOptionSource

`func (o *InstanceServicePlanStorageType) HasVolumeOptionSource() bool`

HasVolumeOptionSource returns a boolean if a field has been set.

### GetDisplayName

`func (o *InstanceServicePlanStorageType) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceServicePlanStorageType) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceServicePlanStorageType) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceServicePlanStorageType) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMinIOPS

`func (o *InstanceServicePlanStorageType) GetMinIOPS() string`

GetMinIOPS returns the MinIOPS field if non-nil, zero value otherwise.

### GetMinIOPSOk

`func (o *InstanceServicePlanStorageType) GetMinIOPSOk() (*string, bool)`

GetMinIOPSOk returns a tuple with the MinIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIOPS

`func (o *InstanceServicePlanStorageType) SetMinIOPS(v string)`

SetMinIOPS sets MinIOPS field to given value.

### HasMinIOPS

`func (o *InstanceServicePlanStorageType) HasMinIOPS() bool`

HasMinIOPS returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *InstanceServicePlanStorageType) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *InstanceServicePlanStorageType) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *InstanceServicePlanStorageType) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *InstanceServicePlanStorageType) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### GetHasDatastore

`func (o *InstanceServicePlanStorageType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *InstanceServicePlanStorageType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *InstanceServicePlanStorageType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *InstanceServicePlanStorageType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetCustomSize

`func (o *InstanceServicePlanStorageType) GetCustomSize() bool`

GetCustomSize returns the CustomSize field if non-nil, zero value otherwise.

### GetCustomSizeOk

`func (o *InstanceServicePlanStorageType) GetCustomSizeOk() (*bool, bool)`

GetCustomSizeOk returns a tuple with the CustomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSize

`func (o *InstanceServicePlanStorageType) SetCustomSize(v bool)`

SetCustomSize sets CustomSize field to given value.

### HasCustomSize

`func (o *InstanceServicePlanStorageType) HasCustomSize() bool`

HasCustomSize returns a boolean if a field has been set.

### GetAutoDelete

`func (o *InstanceServicePlanStorageType) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *InstanceServicePlanStorageType) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *InstanceServicePlanStorageType) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *InstanceServicePlanStorageType) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetName

`func (o *InstanceServicePlanStorageType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceServicePlanStorageType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceServicePlanStorageType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceServicePlanStorageType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *InstanceServicePlanStorageType) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *InstanceServicePlanStorageType) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *InstanceServicePlanStorageType) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *InstanceServicePlanStorageType) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetCustomLabel

`func (o *InstanceServicePlanStorageType) GetCustomLabel() bool`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *InstanceServicePlanStorageType) GetCustomLabelOk() (*bool, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *InstanceServicePlanStorageType) SetCustomLabel(v bool)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *InstanceServicePlanStorageType) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetEnabled

`func (o *InstanceServicePlanStorageType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InstanceServicePlanStorageType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InstanceServicePlanStorageType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InstanceServicePlanStorageType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceServicePlanStorageType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceServicePlanStorageType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceServicePlanStorageType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceServicePlanStorageType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVolumeCategory

`func (o *InstanceServicePlanStorageType) GetVolumeCategory() string`

GetVolumeCategory returns the VolumeCategory field if non-nil, zero value otherwise.

### GetVolumeCategoryOk

`func (o *InstanceServicePlanStorageType) GetVolumeCategoryOk() (*string, bool)`

GetVolumeCategoryOk returns a tuple with the VolumeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCategory

`func (o *InstanceServicePlanStorageType) SetVolumeCategory(v string)`

SetVolumeCategory sets VolumeCategory field to given value.

### HasVolumeCategory

`func (o *InstanceServicePlanStorageType) HasVolumeCategory() bool`

HasVolumeCategory returns a boolean if a field has been set.

### GetExternalId

`func (o *InstanceServicePlanStorageType) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceServicePlanStorageType) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceServicePlanStorageType) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InstanceServicePlanStorageType) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetMaxStorage

`func (o *InstanceServicePlanStorageType) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceServicePlanStorageType) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceServicePlanStorageType) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceServicePlanStorageType) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


