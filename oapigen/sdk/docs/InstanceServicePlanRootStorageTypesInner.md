# InstanceServicePlanRootStorageTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to **string** |  | [optional] 
**MinStorage** | Pointer to **NullableString** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DefaultType** | Pointer to **bool** |  | [optional] 
**CreateDatastore** | Pointer to **NullableString** |  | [optional] 
**Resizable** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **NullableString** |  | [optional] 
**AllowSearch** | Pointer to **bool** |  | [optional] 
**VolumeOptionSource** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**MinIOPS** | Pointer to **NullableString** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**CustomSize** | Pointer to **bool** |  | [optional] 
**AutoDelete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**CustomLabel** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**VolumeCategory** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewInstanceServicePlanRootStorageTypesInner

`func NewInstanceServicePlanRootStorageTypesInner() *InstanceServicePlanRootStorageTypesInner`

NewInstanceServicePlanRootStorageTypesInner instantiates a new InstanceServicePlanRootStorageTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceServicePlanRootStorageTypesInnerWithDefaults

`func NewInstanceServicePlanRootStorageTypesInnerWithDefaults() *InstanceServicePlanRootStorageTypesInner`

NewInstanceServicePlanRootStorageTypesInnerWithDefaults instantiates a new InstanceServicePlanRootStorageTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceServicePlanRootStorageTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceServicePlanRootStorageTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceServicePlanRootStorageTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEditable

`func (o *InstanceServicePlanRootStorageTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *InstanceServicePlanRootStorageTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *InstanceServicePlanRootStorageTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceServicePlanRootStorageTypesInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceServicePlanRootStorageTypesInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceServicePlanRootStorageTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetDisplayOrder

`func (o *InstanceServicePlanRootStorageTypesInner) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *InstanceServicePlanRootStorageTypesInner) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *InstanceServicePlanRootStorageTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetCode

`func (o *InstanceServicePlanRootStorageTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceServicePlanRootStorageTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceServicePlanRootStorageTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVolumeType

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *InstanceServicePlanRootStorageTypesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *InstanceServicePlanRootStorageTypesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetMinStorage

`func (o *InstanceServicePlanRootStorageTypesInner) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *InstanceServicePlanRootStorageTypesInner) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *InstanceServicePlanRootStorageTypesInner) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### SetMinStorageNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetMinStorageNil(b bool)`

 SetMinStorageNil sets the value for MinStorage to be an explicit nil

### UnsetMinStorage
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetMinStorage()`

UnsetMinStorage ensures that no value is present for MinStorage, not even an explicit nil
### GetDeletable

`func (o *InstanceServicePlanRootStorageTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *InstanceServicePlanRootStorageTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *InstanceServicePlanRootStorageTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefaultType

`func (o *InstanceServicePlanRootStorageTypesInner) GetDefaultType() bool`

GetDefaultType returns the DefaultType field if non-nil, zero value otherwise.

### GetDefaultTypeOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetDefaultTypeOk() (*bool, bool)`

GetDefaultTypeOk returns a tuple with the DefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultType

`func (o *InstanceServicePlanRootStorageTypesInner) SetDefaultType(v bool)`

SetDefaultType sets DefaultType field to given value.

### HasDefaultType

`func (o *InstanceServicePlanRootStorageTypesInner) HasDefaultType() bool`

HasDefaultType returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) GetCreateDatastore() string`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetCreateDatastoreOk() (*string, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) SetCreateDatastore(v string)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### SetCreateDatastoreNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetCreateDatastoreNil(b bool)`

 SetCreateDatastoreNil sets the value for CreateDatastore to be an explicit nil

### UnsetCreateDatastore
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetCreateDatastore()`

UnsetCreateDatastore ensures that no value is present for CreateDatastore, not even an explicit nil
### GetResizable

`func (o *InstanceServicePlanRootStorageTypesInner) GetResizable() bool`

GetResizable returns the Resizable field if non-nil, zero value otherwise.

### GetResizableOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetResizableOk() (*bool, bool)`

GetResizableOk returns a tuple with the Resizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizable

`func (o *InstanceServicePlanRootStorageTypesInner) SetResizable(v bool)`

SetResizable sets Resizable field to given value.

### HasResizable

`func (o *InstanceServicePlanRootStorageTypesInner) HasResizable() bool`

HasResizable returns a boolean if a field has been set.

### GetStorageType

`func (o *InstanceServicePlanRootStorageTypesInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *InstanceServicePlanRootStorageTypesInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *InstanceServicePlanRootStorageTypesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### SetStorageTypeNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetAllowSearch

`func (o *InstanceServicePlanRootStorageTypesInner) GetAllowSearch() bool`

GetAllowSearch returns the AllowSearch field if non-nil, zero value otherwise.

### GetAllowSearchOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetAllowSearchOk() (*bool, bool)`

GetAllowSearchOk returns a tuple with the AllowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearch

`func (o *InstanceServicePlanRootStorageTypesInner) SetAllowSearch(v bool)`

SetAllowSearch sets AllowSearch field to given value.

### HasAllowSearch

`func (o *InstanceServicePlanRootStorageTypesInner) HasAllowSearch() bool`

HasAllowSearch returns a boolean if a field has been set.

### GetVolumeOptionSource

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeOptionSource() string`

GetVolumeOptionSource returns the VolumeOptionSource field if non-nil, zero value otherwise.

### GetVolumeOptionSourceOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeOptionSourceOk() (*string, bool)`

GetVolumeOptionSourceOk returns a tuple with the VolumeOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeOptionSource

`func (o *InstanceServicePlanRootStorageTypesInner) SetVolumeOptionSource(v string)`

SetVolumeOptionSource sets VolumeOptionSource field to given value.

### HasVolumeOptionSource

`func (o *InstanceServicePlanRootStorageTypesInner) HasVolumeOptionSource() bool`

HasVolumeOptionSource returns a boolean if a field has been set.

### SetVolumeOptionSourceNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetVolumeOptionSourceNil(b bool)`

 SetVolumeOptionSourceNil sets the value for VolumeOptionSource to be an explicit nil

### UnsetVolumeOptionSource
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetVolumeOptionSource()`

UnsetVolumeOptionSource ensures that no value is present for VolumeOptionSource, not even an explicit nil
### GetDisplayName

`func (o *InstanceServicePlanRootStorageTypesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InstanceServicePlanRootStorageTypesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InstanceServicePlanRootStorageTypesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMinIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) GetMinIOPS() string`

GetMinIOPS returns the MinIOPS field if non-nil, zero value otherwise.

### GetMinIOPSOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetMinIOPSOk() (*string, bool)`

GetMinIOPSOk returns a tuple with the MinIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) SetMinIOPS(v string)`

SetMinIOPS sets MinIOPS field to given value.

### HasMinIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) HasMinIOPS() bool`

HasMinIOPS returns a boolean if a field has been set.

### SetMinIOPSNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetMinIOPSNil(b bool)`

 SetMinIOPSNil sets the value for MinIOPS to be an explicit nil

### UnsetMinIOPS
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetMinIOPS()`

UnsetMinIOPS ensures that no value is present for MinIOPS, not even an explicit nil
### GetMaxIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetHasDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *InstanceServicePlanRootStorageTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetCustomSize

`func (o *InstanceServicePlanRootStorageTypesInner) GetCustomSize() bool`

GetCustomSize returns the CustomSize field if non-nil, zero value otherwise.

### GetCustomSizeOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetCustomSizeOk() (*bool, bool)`

GetCustomSizeOk returns a tuple with the CustomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSize

`func (o *InstanceServicePlanRootStorageTypesInner) SetCustomSize(v bool)`

SetCustomSize sets CustomSize field to given value.

### HasCustomSize

`func (o *InstanceServicePlanRootStorageTypesInner) HasCustomSize() bool`

HasCustomSize returns a boolean if a field has been set.

### GetAutoDelete

`func (o *InstanceServicePlanRootStorageTypesInner) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *InstanceServicePlanRootStorageTypesInner) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *InstanceServicePlanRootStorageTypesInner) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetName

`func (o *InstanceServicePlanRootStorageTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceServicePlanRootStorageTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceServicePlanRootStorageTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *InstanceServicePlanRootStorageTypesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetCustomLabel

`func (o *InstanceServicePlanRootStorageTypesInner) GetCustomLabel() bool`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetCustomLabelOk() (*bool, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *InstanceServicePlanRootStorageTypesInner) SetCustomLabel(v bool)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *InstanceServicePlanRootStorageTypesInner) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetEnabled

`func (o *InstanceServicePlanRootStorageTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InstanceServicePlanRootStorageTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InstanceServicePlanRootStorageTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceServicePlanRootStorageTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceServicePlanRootStorageTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceServicePlanRootStorageTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVolumeCategory

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeCategory() string`

GetVolumeCategory returns the VolumeCategory field if non-nil, zero value otherwise.

### GetVolumeCategoryOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetVolumeCategoryOk() (*string, bool)`

GetVolumeCategoryOk returns a tuple with the VolumeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCategory

`func (o *InstanceServicePlanRootStorageTypesInner) SetVolumeCategory(v string)`

SetVolumeCategory sets VolumeCategory field to given value.

### HasVolumeCategory

`func (o *InstanceServicePlanRootStorageTypesInner) HasVolumeCategory() bool`

HasVolumeCategory returns a boolean if a field has been set.

### GetExternalId

`func (o *InstanceServicePlanRootStorageTypesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *InstanceServicePlanRootStorageTypesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *InstanceServicePlanRootStorageTypesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMaxStorage

`func (o *InstanceServicePlanRootStorageTypesInner) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *InstanceServicePlanRootStorageTypesInner) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *InstanceServicePlanRootStorageTypesInner) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *InstanceServicePlanRootStorageTypesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *InstanceServicePlanRootStorageTypesInner) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *InstanceServicePlanRootStorageTypesInner) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


