# ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner

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

### NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner

`func NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner() *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner`

NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner instantiates a new ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInnerWithDefaults

`func NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInnerWithDefaults() *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner`

NewListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInnerWithDefaults instantiates a new ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetDefaultType

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDefaultType() bool`

GetDefaultType returns the DefaultType field if non-nil, zero value otherwise.

### GetDefaultTypeOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetDefaultTypeOk() (*bool, bool)`

GetDefaultTypeOk returns a tuple with the DefaultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultType

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetDefaultType(v bool)`

SetDefaultType sets DefaultType field to given value.

### HasDefaultType

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasDefaultType() bool`

HasDefaultType returns a boolean if a field has been set.

### GetCustomLabel

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomLabel() bool`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomLabelOk() (*bool, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetCustomLabel(v bool)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### GetCustomSize

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomSize() bool`

GetCustomSize returns the CustomSize field if non-nil, zero value otherwise.

### GetCustomSizeOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomSizeOk() (*bool, bool)`

GetCustomSizeOk returns a tuple with the CustomSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSize

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetCustomSize(v bool)`

SetCustomSize sets CustomSize field to given value.

### HasCustomSize

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasCustomSize() bool`

HasCustomSize returns a boolean if a field has been set.

### GetCustomSizeOptions

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomSizeOptions() string`

GetCustomSizeOptions returns the CustomSizeOptions field if non-nil, zero value otherwise.

### GetCustomSizeOptionsOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCustomSizeOptionsOk() (*string, bool)`

GetCustomSizeOptionsOk returns a tuple with the CustomSizeOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSizeOptions

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetCustomSizeOptions(v string)`

SetCustomSizeOptions sets CustomSizeOptions field to given value.

### HasCustomSizeOptions

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasCustomSizeOptions() bool`

HasCustomSizeOptions returns a boolean if a field has been set.

### GetConfigurableIOPS

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetConfigurableIOPS() bool`

GetConfigurableIOPS returns the ConfigurableIOPS field if non-nil, zero value otherwise.

### GetConfigurableIOPSOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetConfigurableIOPSOk() (*bool, bool)`

GetConfigurableIOPSOk returns a tuple with the ConfigurableIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurableIOPS

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetConfigurableIOPS(v bool)`

SetConfigurableIOPS sets ConfigurableIOPS field to given value.

### HasConfigurableIOPS

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasConfigurableIOPS() bool`

HasConfigurableIOPS returns a boolean if a field has been set.

### GetHasDatastore

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetHasDatastore() bool`

GetHasDatastore returns the HasDatastore field if non-nil, zero value otherwise.

### GetHasDatastoreOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetHasDatastoreOk() (*bool, bool)`

GetHasDatastoreOk returns a tuple with the HasDatastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDatastore

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetHasDatastore(v bool)`

SetHasDatastore sets HasDatastore field to given value.

### HasHasDatastore

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasHasDatastore() bool`

HasHasDatastore returns a boolean if a field has been set.

### GetCategory

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetEnabled

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetOptionTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) GetOptionTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) SetOptionTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


