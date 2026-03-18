# ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner

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
**OptionList** | Pointer to [**ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList**](ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList.md) |  | [optional] 
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

### NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner

`func NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner() *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner`

NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner instantiates a new ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInnerWithDefaults

`func NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInnerWithDefaults() *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner`

NewListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInnerWithDefaults instantiates a new ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionList() ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetOptionListOk() (*ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetOptionList(v ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInnerOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


