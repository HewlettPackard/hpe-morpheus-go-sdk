# StorageVolumeType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**DefaultType** | Pointer to **bool** |  | [optional] 
**CustomLabel** | Pointer to **bool** |  | [optional] 
**CustomSize** | Pointer to **bool** |  | [optional] 
**CustomSizeOptions** | Pointer to **string** |  | [optional] 
**ConfigurableIOPS** | Pointer to **bool** |  | [optional] 
**HasDatastore** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner**](ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewStorageVolumeType

`func NewStorageVolumeType() *StorageVolumeType`

NewStorageVolumeType instantiates a new StorageVolumeType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageVolumeTypeWithDefaults

`func NewStorageVolumeTypeWithDefaults() *StorageVolumeType`

NewStorageVolumeTypeWithDefaults instantiates a new StorageVolumeType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageVolumeType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageVolumeType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageVolumeType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *StorageVolumeType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *StorageVolumeType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *StorageVolumeType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *StorageVolumeType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *StorageVolumeType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *StorageVolumeType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StorageVolumeType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StorageVolumeType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StorageVolumeType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *StorageVolumeType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StorageVolumeType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StorageVolumeType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StorageVolumeType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *StorageVolumeType) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *StorageVolumeType) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *StorageVolumeType) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *StorageVolumeType) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDefaultType

`func (o *StorageVolumeType) GetDefaultType() bool`

GetDefaultType returns the DefaultType field if non-nil, zero value otherwise.

### GetDefaultTypeOk

`func (o *StorageVolumeType) GetDefaultTypeOk() (*bool, bool)`

GetDefaultTypeOk returns a tuple with the DefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultType

`func (o *StorageVolumeType) SetDefaultType(v bool)`

SetDefaultType sets DefaultType field to given value.

### HasDefaultType

`func (o *StorageVolumeType) HasDefaultType() bool`

HasDefaultType returns a boolean if a field has been set.

### GetCustomLabel

`func (o *StorageVolumeType) GetCustomLabel() bool`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *StorageVolumeType) GetCustomLabelOk() (*bool, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *StorageVolumeType) SetCustomLabel(v bool)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *StorageVolumeType) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetCustomSize

`func (o *StorageVolumeType) GetCustomSize() bool`

GetCustomSize returns the CustomSize field if non-nil, zero value otherwise.

### GetCustomSizeOk

`func (o *StorageVolumeType) GetCustomSizeOk() (*bool, bool)`

GetCustomSizeOk returns a tuple with the CustomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSize

`func (o *StorageVolumeType) SetCustomSize(v bool)`

SetCustomSize sets CustomSize field to given value.

### HasCustomSize

`func (o *StorageVolumeType) HasCustomSize() bool`

HasCustomSize returns a boolean if a field has been set.

### GetCustomSizeOptions

`func (o *StorageVolumeType) GetCustomSizeOptions() string`

GetCustomSizeOptions returns the CustomSizeOptions field if non-nil, zero value otherwise.

### GetCustomSizeOptionsOk

`func (o *StorageVolumeType) GetCustomSizeOptionsOk() (*string, bool)`

GetCustomSizeOptionsOk returns a tuple with the CustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSizeOptions

`func (o *StorageVolumeType) SetCustomSizeOptions(v string)`

SetCustomSizeOptions sets CustomSizeOptions field to given value.

### HasCustomSizeOptions

`func (o *StorageVolumeType) HasCustomSizeOptions() bool`

HasCustomSizeOptions returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *StorageVolumeType) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *StorageVolumeType) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *StorageVolumeType) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *StorageVolumeType) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetHasDatastore

`func (o *StorageVolumeType) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *StorageVolumeType) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *StorageVolumeType) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *StorageVolumeType) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetCategory

`func (o *StorageVolumeType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StorageVolumeType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StorageVolumeType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StorageVolumeType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *StorageVolumeType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StorageVolumeType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StorageVolumeType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StorageVolumeType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOptionTypes

`func (o *StorageVolumeType) GetOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *StorageVolumeType) GetOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *StorageVolumeType) SetOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *StorageVolumeType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


