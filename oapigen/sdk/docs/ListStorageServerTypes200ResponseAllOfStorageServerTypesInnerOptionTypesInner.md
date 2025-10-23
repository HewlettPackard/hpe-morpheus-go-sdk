# ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**FieldLabel** | Pointer to **string** |  | [optional] 
**FieldCode** | Pointer to **string** |  | [optional] 
**FieldContext** | Pointer to **string** |  | [optional] 
**FieldGroup** | Pointer to **NullableString** |  | [optional] 
**FieldClass** | Pointer to **NullableString** |  | [optional] 
**FieldAddOn** | Pointer to **NullableString** |  | [optional] 
**FieldComponent** | Pointer to **NullableString** |  | [optional] 
**FieldInput** | Pointer to **NullableString** |  | [optional] 
**PlaceHolder** | Pointer to **NullableString** |  | [optional] 
**VerifyPattern** | Pointer to **NullableString** |  | [optional] 
**HelpBlock** | Pointer to **string** |  | [optional] 
**HelpBlockFieldCode** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**OptionSource** | Pointer to **NullableString** |  | [optional] 
**OptionSourceType** | Pointer to **NullableString** |  | [optional] 
**OptionList** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Advanced** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**ExportMeta** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**WrapperClass** | Pointer to **NullableString** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NoBlank** | Pointer to **bool** |  | [optional] 
**DependsOnCode** | Pointer to **NullableString** |  | [optional] 
**VisibleOnCode** | Pointer to **NullableString** |  | [optional] 
**RequireOnCode** | Pointer to **NullableString** |  | [optional] 
**ContextualDefault** | Pointer to **bool** |  | [optional] 
**DisplayValueOnDetails** | Pointer to **NullableBool** |  | [optional] 
**ShowOnCreate** | Pointer to **NullableBool** |  | [optional] 
**ShowOnEdit** | Pointer to **NullableBool** |  | [optional] 
**LocalCredential** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner

`func NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner() *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner instantiates a new ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInnerWithDefaults

`func NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInnerWithDefaults() *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner`

NewListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInnerWithDefaults instantiates a new ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### GetFieldContext

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### GetHelpBlockFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *ListStorageServerTypes200ResponseAllOfStorageServerTypesInnerOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


