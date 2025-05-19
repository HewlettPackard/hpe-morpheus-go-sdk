# GetReportTypes200ResponseReportTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**MasterOnly** | Pointer to **bool** |  | [optional] 
**OwnerOnly** | Pointer to **bool** |  | [optional] 
**SupportsAllZoneTypes** | Pointer to **bool** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**OptionTypes** | Pointer to [**[]GetReportTypes200ResponseReportTypesInnerOptionTypesInner**](GetReportTypes200ResponseReportTypesInnerOptionTypesInner.md) |  | [optional] 
**SupportedZoneTypes** | Pointer to [**[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 

## Methods

### NewGetReportTypes200ResponseReportTypesInner

`func NewGetReportTypes200ResponseReportTypesInner() *GetReportTypes200ResponseReportTypesInner`

NewGetReportTypes200ResponseReportTypesInner instantiates a new GetReportTypes200ResponseReportTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReportTypes200ResponseReportTypesInnerWithDefaults

`func NewGetReportTypes200ResponseReportTypesInnerWithDefaults() *GetReportTypes200ResponseReportTypesInner`

NewGetReportTypes200ResponseReportTypesInnerWithDefaults instantiates a new GetReportTypes200ResponseReportTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReportTypes200ResponseReportTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReportTypes200ResponseReportTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetReportTypes200ResponseReportTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetReportTypes200ResponseReportTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetReportTypes200ResponseReportTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetReportTypes200ResponseReportTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetReportTypes200ResponseReportTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetReportTypes200ResponseReportTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetReportTypes200ResponseReportTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetReportTypes200ResponseReportTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetReportTypes200ResponseReportTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetReportTypes200ResponseReportTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *GetReportTypes200ResponseReportTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetReportTypes200ResponseReportTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetReportTypes200ResponseReportTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisible

`func (o *GetReportTypes200ResponseReportTypesInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *GetReportTypes200ResponseReportTypesInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *GetReportTypes200ResponseReportTypesInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetMasterOnly

`func (o *GetReportTypes200ResponseReportTypesInner) GetMasterOnly() bool`

GetMasterOnly returns the MasterOnly field if non-nil, zero value otherwise.

### GetMasterOnlyOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetMasterOnlyOk() (*bool, bool)`

GetMasterOnlyOk returns a tuple with the MasterOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOnly

`func (o *GetReportTypes200ResponseReportTypesInner) SetMasterOnly(v bool)`

SetMasterOnly sets MasterOnly field to given value.

### HasMasterOnly

`func (o *GetReportTypes200ResponseReportTypesInner) HasMasterOnly() bool`

HasMasterOnly returns a boolean if a field has been set.

### GetOwnerOnly

`func (o *GetReportTypes200ResponseReportTypesInner) GetOwnerOnly() bool`

GetOwnerOnly returns the OwnerOnly field if non-nil, zero value otherwise.

### GetOwnerOnlyOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetOwnerOnlyOk() (*bool, bool)`

GetOwnerOnlyOk returns a tuple with the OwnerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOnly

`func (o *GetReportTypes200ResponseReportTypesInner) SetOwnerOnly(v bool)`

SetOwnerOnly sets OwnerOnly field to given value.

### HasOwnerOnly

`func (o *GetReportTypes200ResponseReportTypesInner) HasOwnerOnly() bool`

HasOwnerOnly returns a boolean if a field has been set.

### GetSupportsAllZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) GetSupportsAllZoneTypes() bool`

GetSupportsAllZoneTypes returns the SupportsAllZoneTypes field if non-nil, zero value otherwise.

### GetSupportsAllZoneTypesOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetSupportsAllZoneTypesOk() (*bool, bool)`

GetSupportsAllZoneTypesOk returns a tuple with the SupportsAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAllZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) SetSupportsAllZoneTypes(v bool)`

SetSupportsAllZoneTypes sets SupportsAllZoneTypes field to given value.

### HasSupportsAllZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) HasSupportsAllZoneTypes() bool`

HasSupportsAllZoneTypes returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetReportTypes200ResponseReportTypesInner) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetReportTypes200ResponseReportTypesInner) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetReportTypes200ResponseReportTypesInner) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetReportTypes200ResponseReportTypesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetReportTypes200ResponseReportTypesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetReportTypes200ResponseReportTypesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetReportTypes200ResponseReportTypesInner) GetOptionTypes() []GetReportTypes200ResponseReportTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetOptionTypesOk() (*[]GetReportTypes200ResponseReportTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetReportTypes200ResponseReportTypesInner) SetOptionTypes(v []GetReportTypes200ResponseReportTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetReportTypes200ResponseReportTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetSupportedZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) GetSupportedZoneTypes() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetSupportedZoneTypes returns the SupportedZoneTypes field if non-nil, zero value otherwise.

### GetSupportedZoneTypesOk

`func (o *GetReportTypes200ResponseReportTypesInner) GetSupportedZoneTypesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetSupportedZoneTypesOk returns a tuple with the SupportedZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) SetSupportedZoneTypes(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetSupportedZoneTypes sets SupportedZoneTypes field to given value.

### HasSupportedZoneTypes

`func (o *GetReportTypes200ResponseReportTypesInner) HasSupportedZoneTypes() bool`

HasSupportedZoneTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


