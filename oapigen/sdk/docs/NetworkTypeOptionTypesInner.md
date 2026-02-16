# NetworkTypeOptionTypesInner

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
**OptionList** | Pointer to [**NetworkTypeOptionTypesInnerOptionList**](NetworkTypeOptionTypesInnerOptionList.md) |  | [optional] 
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

### NewNetworkTypeOptionTypesInner

`func NewNetworkTypeOptionTypesInner() *NetworkTypeOptionTypesInner`

NewNetworkTypeOptionTypesInner instantiates a new NetworkTypeOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTypeOptionTypesInnerWithDefaults

`func NewNetworkTypeOptionTypesInnerWithDefaults() *NetworkTypeOptionTypesInner`

NewNetworkTypeOptionTypesInnerWithDefaults instantiates a new NetworkTypeOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTypeOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTypeOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTypeOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTypeOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTypeOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTypeOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTypeOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTypeOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *NetworkTypeOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *NetworkTypeOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *NetworkTypeOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *NetworkTypeOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *NetworkTypeOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *NetworkTypeOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *NetworkTypeOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkTypeOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkTypeOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkTypeOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkTypeOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkTypeOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *NetworkTypeOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkTypeOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkTypeOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkTypeOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *NetworkTypeOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *NetworkTypeOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *NetworkTypeOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *NetworkTypeOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *NetworkTypeOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *NetworkTypeOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *NetworkTypeOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *NetworkTypeOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *NetworkTypeOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *NetworkTypeOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *NetworkTypeOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *NetworkTypeOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *NetworkTypeOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *NetworkTypeOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *NetworkTypeOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *NetworkTypeOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *NetworkTypeOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *NetworkTypeOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *NetworkTypeOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *NetworkTypeOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *NetworkTypeOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *NetworkTypeOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *NetworkTypeOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *NetworkTypeOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *NetworkTypeOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *NetworkTypeOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *NetworkTypeOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *NetworkTypeOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *NetworkTypeOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *NetworkTypeOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *NetworkTypeOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *NetworkTypeOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *NetworkTypeOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *NetworkTypeOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *NetworkTypeOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *NetworkTypeOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *NetworkTypeOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *NetworkTypeOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *NetworkTypeOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *NetworkTypeOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *NetworkTypeOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *NetworkTypeOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *NetworkTypeOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *NetworkTypeOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *NetworkTypeOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *NetworkTypeOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *NetworkTypeOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *NetworkTypeOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *NetworkTypeOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *NetworkTypeOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *NetworkTypeOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *NetworkTypeOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *NetworkTypeOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *NetworkTypeOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *NetworkTypeOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *NetworkTypeOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *NetworkTypeOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *NetworkTypeOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *NetworkTypeOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *NetworkTypeOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *NetworkTypeOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *NetworkTypeOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *NetworkTypeOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *NetworkTypeOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *NetworkTypeOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *NetworkTypeOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *NetworkTypeOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *NetworkTypeOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *NetworkTypeOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *NetworkTypeOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *NetworkTypeOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *NetworkTypeOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *NetworkTypeOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NetworkTypeOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NetworkTypeOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NetworkTypeOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NetworkTypeOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NetworkTypeOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *NetworkTypeOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *NetworkTypeOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *NetworkTypeOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *NetworkTypeOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *NetworkTypeOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *NetworkTypeOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *NetworkTypeOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *NetworkTypeOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *NetworkTypeOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *NetworkTypeOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *NetworkTypeOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *NetworkTypeOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *NetworkTypeOptionTypesInner) GetOptionList() NetworkTypeOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *NetworkTypeOptionTypesInner) GetOptionListOk() (*NetworkTypeOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *NetworkTypeOptionTypesInner) SetOptionList(v NetworkTypeOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *NetworkTypeOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *NetworkTypeOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTypeOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTypeOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkTypeOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *NetworkTypeOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *NetworkTypeOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *NetworkTypeOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *NetworkTypeOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *NetworkTypeOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NetworkTypeOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NetworkTypeOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NetworkTypeOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *NetworkTypeOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *NetworkTypeOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *NetworkTypeOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *NetworkTypeOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *NetworkTypeOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NetworkTypeOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NetworkTypeOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *NetworkTypeOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *NetworkTypeOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *NetworkTypeOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *NetworkTypeOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *NetworkTypeOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkTypeOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkTypeOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkTypeOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkTypeOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *NetworkTypeOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *NetworkTypeOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *NetworkTypeOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkTypeOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkTypeOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkTypeOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *NetworkTypeOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *NetworkTypeOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *NetworkTypeOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *NetworkTypeOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *NetworkTypeOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *NetworkTypeOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *NetworkTypeOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkTypeOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkTypeOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkTypeOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *NetworkTypeOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *NetworkTypeOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *NetworkTypeOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *NetworkTypeOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *NetworkTypeOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *NetworkTypeOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *NetworkTypeOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *NetworkTypeOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *NetworkTypeOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *NetworkTypeOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *NetworkTypeOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *NetworkTypeOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *NetworkTypeOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *NetworkTypeOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *NetworkTypeOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *NetworkTypeOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *NetworkTypeOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *NetworkTypeOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *NetworkTypeOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *NetworkTypeOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *NetworkTypeOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *NetworkTypeOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *NetworkTypeOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *NetworkTypeOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *NetworkTypeOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *NetworkTypeOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *NetworkTypeOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *NetworkTypeOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *NetworkTypeOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *NetworkTypeOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *NetworkTypeOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *NetworkTypeOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *NetworkTypeOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *NetworkTypeOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *NetworkTypeOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *NetworkTypeOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *NetworkTypeOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *NetworkTypeOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *NetworkTypeOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *NetworkTypeOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *NetworkTypeOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *NetworkTypeOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *NetworkTypeOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *NetworkTypeOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *NetworkTypeOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *NetworkTypeOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *NetworkTypeOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *NetworkTypeOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *NetworkTypeOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *NetworkTypeOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *NetworkTypeOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *NetworkTypeOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


