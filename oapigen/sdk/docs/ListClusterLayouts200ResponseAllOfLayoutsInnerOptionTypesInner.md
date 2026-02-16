# ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**FieldLabel** | Pointer to **string** |  | [optional] 
**FieldCode** | Pointer to **NullableString** |  | [optional] 
**FieldContext** | Pointer to **string** |  | [optional] 
**FieldGroup** | Pointer to **NullableString** |  | [optional] 
**FieldClass** | Pointer to **NullableString** |  | [optional] 
**FieldAddOn** | Pointer to **NullableString** |  | [optional] 
**FieldComponent** | Pointer to **NullableString** |  | [optional] 
**FieldInput** | Pointer to **NullableString** |  | [optional] 
**PlaceHolder** | Pointer to **NullableString** |  | [optional] 
**VerifyPattern** | Pointer to **NullableString** |  | [optional] 
**HelpBlock** | Pointer to **NullableString** |  | [optional] 
**HelpBlockFieldCode** | Pointer to **NullableString** |  | [optional] 
**DefaultValue** | Pointer to **NullableString** |  | [optional] 
**OptionSource** | Pointer to **NullableString** |  | [optional] 
**OptionSourceType** | Pointer to **NullableString** |  | [optional] 
**OptionList** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerOptionList**](ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerOptionList.md) |  | [optional] 
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
**ContextualDefault** | Pointer to **NullableBool** |  | [optional] 
**DisplayValueOnDetails** | Pointer to **NullableBool** |  | [optional] 
**ShowOnCreate** | Pointer to **NullableBool** |  | [optional] 
**ShowOnEdit** | Pointer to **NullableBool** |  | [optional] 
**LocalCredential** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner() *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerWithDefaults

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerWithDefaults() *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerWithDefaults instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionList() ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetOptionListOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetOptionList(v ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


