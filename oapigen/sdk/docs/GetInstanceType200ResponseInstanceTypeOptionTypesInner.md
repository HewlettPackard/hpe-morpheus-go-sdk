# GetInstanceType200ResponseInstanceTypeOptionTypesInner

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
**OptionList** | Pointer to [**GetInstanceType200ResponseInstanceTypeOptionTypesInnerOptionList**](GetInstanceType200ResponseInstanceTypeOptionTypesInnerOptionList.md) |  | [optional] 
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

### NewGetInstanceType200ResponseInstanceTypeOptionTypesInner

`func NewGetInstanceType200ResponseInstanceTypeOptionTypesInner() *GetInstanceType200ResponseInstanceTypeOptionTypesInner`

NewGetInstanceType200ResponseInstanceTypeOptionTypesInner instantiates a new GetInstanceType200ResponseInstanceTypeOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeOptionTypesInnerWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeOptionTypesInnerWithDefaults() *GetInstanceType200ResponseInstanceTypeOptionTypesInner`

NewGetInstanceType200ResponseInstanceTypeOptionTypesInnerWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionList() GetInstanceType200ResponseInstanceTypeOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetOptionListOk() (*GetInstanceType200ResponseInstanceTypeOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetOptionList(v GetInstanceType200ResponseInstanceTypeOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *GetInstanceType200ResponseInstanceTypeOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


