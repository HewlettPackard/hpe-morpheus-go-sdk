# ReportTypesInner

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
**IsPlugin** | Pointer to **NullableBool** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**OptionTypes** | Pointer to [**[]ReportTypesInnerOptionTypesInner**](ReportTypesInnerOptionTypesInner.md) |  | [optional] 
**SupportedZoneTypes** | Pointer to [**[]ReportTypesInnerSupportedZoneTypesInner**](ReportTypesInnerSupportedZoneTypesInner.md) |  | [optional] 

## Methods

### NewReportTypesInner

`func NewReportTypesInner() *ReportTypesInner`

NewReportTypesInner instantiates a new ReportTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportTypesInnerWithDefaults

`func NewReportTypesInnerWithDefaults() *ReportTypesInner`

NewReportTypesInnerWithDefaults instantiates a new ReportTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ReportTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ReportTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReportTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReportTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReportTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ReportTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ReportTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ReportTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ReportTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ReportTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ReportTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ReportTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ReportTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ReportTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ReportTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ReportTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ReportTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetVisible

`func (o *ReportTypesInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ReportTypesInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ReportTypesInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ReportTypesInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetMasterOnly

`func (o *ReportTypesInner) GetMasterOnly() bool`

GetMasterOnly returns the MasterOnly field if non-nil, zero value otherwise.

### GetMasterOnlyOk

`func (o *ReportTypesInner) GetMasterOnlyOk() (*bool, bool)`

GetMasterOnlyOk returns a tuple with the MasterOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterOnly

`func (o *ReportTypesInner) SetMasterOnly(v bool)`

SetMasterOnly sets MasterOnly field to given value.

### HasMasterOnly

`func (o *ReportTypesInner) HasMasterOnly() bool`

HasMasterOnly returns a boolean if a field has been set.

### GetOwnerOnly

`func (o *ReportTypesInner) GetOwnerOnly() bool`

GetOwnerOnly returns the OwnerOnly field if non-nil, zero value otherwise.

### GetOwnerOnlyOk

`func (o *ReportTypesInner) GetOwnerOnlyOk() (*bool, bool)`

GetOwnerOnlyOk returns a tuple with the OwnerOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerOnly

`func (o *ReportTypesInner) SetOwnerOnly(v bool)`

SetOwnerOnly sets OwnerOnly field to given value.

### HasOwnerOnly

`func (o *ReportTypesInner) HasOwnerOnly() bool`

HasOwnerOnly returns a boolean if a field has been set.

### GetSupportsAllZoneTypes

`func (o *ReportTypesInner) GetSupportsAllZoneTypes() bool`

GetSupportsAllZoneTypes returns the SupportsAllZoneTypes field if non-nil, zero value otherwise.

### GetSupportsAllZoneTypesOk

`func (o *ReportTypesInner) GetSupportsAllZoneTypesOk() (*bool, bool)`

GetSupportsAllZoneTypesOk returns a tuple with the SupportsAllZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsAllZoneTypes

`func (o *ReportTypesInner) SetSupportsAllZoneTypes(v bool)`

SetSupportsAllZoneTypes sets SupportsAllZoneTypes field to given value.

### HasSupportsAllZoneTypes

`func (o *ReportTypesInner) HasSupportsAllZoneTypes() bool`

HasSupportsAllZoneTypes returns a boolean if a field has been set.

### GetIsPlugin

`func (o *ReportTypesInner) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *ReportTypesInner) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *ReportTypesInner) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *ReportTypesInner) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### SetIsPluginNil

`func (o *ReportTypesInner) SetIsPluginNil(b bool)`

 SetIsPluginNil sets the value for IsPlugin to be an explicit nil

### UnsetIsPlugin
`func (o *ReportTypesInner) UnsetIsPlugin()`

UnsetIsPlugin ensures that no value is present for IsPlugin, not even an explicit nil
### GetDateCreated

`func (o *ReportTypesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ReportTypesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ReportTypesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ReportTypesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ReportTypesInner) GetOptionTypes() []ReportTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ReportTypesInner) GetOptionTypesOk() (*[]ReportTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ReportTypesInner) SetOptionTypes(v []ReportTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ReportTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetSupportedZoneTypes

`func (o *ReportTypesInner) GetSupportedZoneTypes() []ReportTypesInnerSupportedZoneTypesInner`

GetSupportedZoneTypes returns the SupportedZoneTypes field if non-nil, zero value otherwise.

### GetSupportedZoneTypesOk

`func (o *ReportTypesInner) GetSupportedZoneTypesOk() (*[]ReportTypesInnerSupportedZoneTypesInner, bool)`

GetSupportedZoneTypesOk returns a tuple with the SupportedZoneTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedZoneTypes

`func (o *ReportTypesInner) SetSupportedZoneTypes(v []ReportTypesInnerSupportedZoneTypesInner)`

SetSupportedZoneTypes sets SupportedZoneTypes field to given value.

### HasSupportedZoneTypes

`func (o *ReportTypesInner) HasSupportedZoneTypes() bool`

HasSupportedZoneTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


