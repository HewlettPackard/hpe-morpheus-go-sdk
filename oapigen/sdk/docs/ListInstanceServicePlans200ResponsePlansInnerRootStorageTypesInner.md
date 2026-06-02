# ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner

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

### NewListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner

`func NewListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner() *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner`

NewListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner instantiates a new ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEditable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetDisplayOrder

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetCode

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeType() string`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeTypeOk() (*string, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetVolumeType(v string)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetMinStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### SetMinStorageNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMinStorageNil(b bool)`

 SetMinStorageNil sets the value for MinStorage to be an explicit nil

### UnsetMinStorage
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetMinStorage()`

UnsetMinStorage ensures that no value is present for MinStorage, not even an explicit nil
### GetDeletable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDefaultType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDefaultType() bool`

GetDefaultType returns the DefaultType field if non-nil, zero value otherwise.

### GetDefaultTypeOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDefaultTypeOk() (*bool, bool)`

GetDefaultTypeOk returns a tuple with the DefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDefaultType(v bool)`

SetDefaultType sets DefaultType field to given value.

### HasDefaultType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasDefaultType() bool`

HasDefaultType returns a boolean if a field has been set.

### GetCreateDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCreateDatastore() string`

GetCreateDatastore returns the CreateDatastore field if non-nil, zero value otherwise.

### GetCreateDatastoreOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCreateDatastoreOk() (*string, bool)`

GetCreateDatastoreOk returns a tuple with the CreateDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetCreateDatastore(v string)`

SetCreateDatastore sets CreateDatastore field to given value.

### HasCreateDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasCreateDatastore() bool`

HasCreateDatastore returns a boolean if a field has been set.

### SetCreateDatastoreNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetCreateDatastoreNil(b bool)`

 SetCreateDatastoreNil sets the value for CreateDatastore to be an explicit nil

### UnsetCreateDatastore
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetCreateDatastore()`

UnsetCreateDatastore ensures that no value is present for CreateDatastore, not even an explicit nil
### GetResizable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetResizable() bool`

GetResizable returns the Resizable field if non-nil, zero value otherwise.

### GetResizableOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetResizableOk() (*bool, bool)`

GetResizableOk returns a tuple with the Resizable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResizable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetResizable(v bool)`

SetResizable sets Resizable field to given value.

### HasResizable

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasResizable() bool`

HasResizable returns a boolean if a field has been set.

### GetStorageType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetStorageType() string`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetStorageTypeOk() (*string, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetStorageType(v string)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### SetStorageTypeNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetStorageTypeNil(b bool)`

 SetStorageTypeNil sets the value for StorageType to be an explicit nil

### UnsetStorageType
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetStorageType()`

UnsetStorageType ensures that no value is present for StorageType, not even an explicit nil
### GetAllowSearch

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetAllowSearch() bool`

GetAllowSearch returns the AllowSearch field if non-nil, zero value otherwise.

### GetAllowSearchOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetAllowSearchOk() (*bool, bool)`

GetAllowSearchOk returns a tuple with the AllowSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowSearch

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetAllowSearch(v bool)`

SetAllowSearch sets AllowSearch field to given value.

### HasAllowSearch

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasAllowSearch() bool`

HasAllowSearch returns a boolean if a field has been set.

### GetVolumeOptionSource

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeOptionSource() string`

GetVolumeOptionSource returns the VolumeOptionSource field if non-nil, zero value otherwise.

### GetVolumeOptionSourceOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeOptionSourceOk() (*string, bool)`

GetVolumeOptionSourceOk returns a tuple with the VolumeOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeOptionSource

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetVolumeOptionSource(v string)`

SetVolumeOptionSource sets VolumeOptionSource field to given value.

### HasVolumeOptionSource

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasVolumeOptionSource() bool`

HasVolumeOptionSource returns a boolean if a field has been set.

### SetVolumeOptionSourceNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetVolumeOptionSourceNil(b bool)`

 SetVolumeOptionSourceNil sets the value for VolumeOptionSource to be an explicit nil

### UnsetVolumeOptionSource
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetVolumeOptionSource()`

UnsetVolumeOptionSource ensures that no value is present for VolumeOptionSource, not even an explicit nil
### GetDisplayName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetMinIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMinIOPS() string`

GetMinIOPS returns the MinIOPS field if non-nil, zero value otherwise.

### GetMinIOPSOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMinIOPSOk() (*string, bool)`

GetMinIOPSOk returns a tuple with the MinIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMinIOPS(v string)`

SetMinIOPS sets MinIOPS field to given value.

### HasMinIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasMinIOPS() bool`

HasMinIOPS returns a boolean if a field has been set.

### SetMinIOPSNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMinIOPSNil(b bool)`

 SetMinIOPSNil sets the value for MinIOPS to be an explicit nil

### UnsetMinIOPS
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetMinIOPS()`

UnsetMinIOPS ensures that no value is present for MinIOPS, not even an explicit nil
### GetMaxIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetCustomSize

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCustomSize() bool`

GetCustomSize returns the CustomSize field if non-nil, zero value otherwise.

### GetCustomSizeOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCustomSizeOk() (*bool, bool)`

GetCustomSizeOk returns a tuple with the CustomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSize

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetCustomSize(v bool)`

SetCustomSize sets CustomSize field to given value.

### HasCustomSize

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasCustomSize() bool`

HasCustomSize returns a boolean if a field has been set.

### GetAutoDelete

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetAutoDelete() bool`

GetAutoDelete returns the AutoDelete field if non-nil, zero value otherwise.

### GetAutoDeleteOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetAutoDeleteOk() (*bool, bool)`

GetAutoDeleteOk returns a tuple with the AutoDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDelete

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetAutoDelete(v bool)`

SetAutoDelete sets AutoDelete field to given value.

### HasAutoDelete

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasAutoDelete() bool`

HasAutoDelete returns a boolean if a field has been set.

### GetName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetCustomLabel

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCustomLabel() bool`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetCustomLabelOk() (*bool, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetCustomLabel(v bool)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetEnabled

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDescription

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVolumeCategory

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeCategory() string`

GetVolumeCategory returns the VolumeCategory field if non-nil, zero value otherwise.

### GetVolumeCategoryOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetVolumeCategoryOk() (*string, bool)`

GetVolumeCategoryOk returns a tuple with the VolumeCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeCategory

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetVolumeCategory(v string)`

SetVolumeCategory sets VolumeCategory field to given value.

### HasVolumeCategory

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasVolumeCategory() bool`

HasVolumeCategory returns a boolean if a field has been set.

### GetExternalId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### SetMaxStorageNil

`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) SetMaxStorageNil(b bool)`

 SetMaxStorageNil sets the value for MaxStorage to be an explicit nil

### UnsetMaxStorage
`func (o *ListInstanceServicePlans200ResponsePlansInnerRootStorageTypesInner) UnsetMaxStorage()`

UnsetMaxStorage ensures that no value is present for MaxStorage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


