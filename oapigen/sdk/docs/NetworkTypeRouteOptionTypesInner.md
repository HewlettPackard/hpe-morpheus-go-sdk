# NetworkTypeRouteOptionTypesInner

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
**PlaceHolder** | Pointer to **string** |  | [optional] 
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
**DisplayValueOnDetails** | Pointer to **bool** |  | [optional] 
**ShowOnCreate** | Pointer to **bool** |  | [optional] 
**ShowOnEdit** | Pointer to **bool** |  | [optional] 

## Methods

### NewNetworkTypeRouteOptionTypesInner

`func NewNetworkTypeRouteOptionTypesInner() *NetworkTypeRouteOptionTypesInner`

NewNetworkTypeRouteOptionTypesInner instantiates a new NetworkTypeRouteOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTypeRouteOptionTypesInnerWithDefaults

`func NewNetworkTypeRouteOptionTypesInnerWithDefaults() *NetworkTypeRouteOptionTypesInner`

NewNetworkTypeRouteOptionTypesInnerWithDefaults instantiates a new NetworkTypeRouteOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkTypeRouteOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkTypeRouteOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkTypeRouteOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkTypeRouteOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkTypeRouteOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTypeRouteOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTypeRouteOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkTypeRouteOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkTypeRouteOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkTypeRouteOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkTypeRouteOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkTypeRouteOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *NetworkTypeRouteOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *NetworkTypeRouteOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *NetworkTypeRouteOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NetworkTypeRouteOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *NetworkTypeRouteOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### GetFieldContext

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *NetworkTypeRouteOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *NetworkTypeRouteOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *NetworkTypeRouteOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *NetworkTypeRouteOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *NetworkTypeRouteOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *NetworkTypeRouteOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *NetworkTypeRouteOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *NetworkTypeRouteOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *NetworkTypeRouteOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *NetworkTypeRouteOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *NetworkTypeRouteOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *NetworkTypeRouteOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetVerifyPattern

`func (o *NetworkTypeRouteOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *NetworkTypeRouteOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *NetworkTypeRouteOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *NetworkTypeRouteOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *NetworkTypeRouteOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *NetworkTypeRouteOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *NetworkTypeRouteOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *NetworkTypeRouteOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *NetworkTypeRouteOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *NetworkTypeRouteOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### GetHelpBlockFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *NetworkTypeRouteOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *NetworkTypeRouteOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *NetworkTypeRouteOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *NetworkTypeRouteOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *NetworkTypeRouteOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *NetworkTypeRouteOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *NetworkTypeRouteOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *NetworkTypeRouteOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *NetworkTypeRouteOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *NetworkTypeRouteOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *NetworkTypeRouteOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *NetworkTypeRouteOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *NetworkTypeRouteOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *NetworkTypeRouteOptionTypesInner) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *NetworkTypeRouteOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### SetOptionListNil

`func (o *NetworkTypeRouteOptionTypesInner) SetOptionListNil(b bool)`

 SetOptionListNil sets the value for OptionList to be an explicit nil

### UnsetOptionList
`func (o *NetworkTypeRouteOptionTypesInner) UnsetOptionList()`

UnsetOptionList ensures that no value is present for OptionList, not even an explicit nil
### GetType

`func (o *NetworkTypeRouteOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkTypeRouteOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkTypeRouteOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *NetworkTypeRouteOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *NetworkTypeRouteOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *NetworkTypeRouteOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *NetworkTypeRouteOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *NetworkTypeRouteOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NetworkTypeRouteOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NetworkTypeRouteOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NetworkTypeRouteOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *NetworkTypeRouteOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *NetworkTypeRouteOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *NetworkTypeRouteOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *NetworkTypeRouteOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *NetworkTypeRouteOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *NetworkTypeRouteOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *NetworkTypeRouteOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *NetworkTypeRouteOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *NetworkTypeRouteOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *NetworkTypeRouteOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *NetworkTypeRouteOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *NetworkTypeRouteOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkTypeRouteOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkTypeRouteOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkTypeRouteOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkTypeRouteOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *NetworkTypeRouteOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *NetworkTypeRouteOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *NetworkTypeRouteOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *NetworkTypeRouteOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *NetworkTypeRouteOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *NetworkTypeRouteOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *NetworkTypeRouteOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *NetworkTypeRouteOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *NetworkTypeRouteOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *NetworkTypeRouteOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *NetworkTypeRouteOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *NetworkTypeRouteOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *NetworkTypeRouteOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkTypeRouteOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkTypeRouteOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkTypeRouteOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *NetworkTypeRouteOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *NetworkTypeRouteOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *NetworkTypeRouteOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *NetworkTypeRouteOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *NetworkTypeRouteOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *NetworkTypeRouteOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *NetworkTypeRouteOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *NetworkTypeRouteOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *NetworkTypeRouteOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *NetworkTypeRouteOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *NetworkTypeRouteOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *NetworkTypeRouteOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *NetworkTypeRouteOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *NetworkTypeRouteOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *NetworkTypeRouteOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *NetworkTypeRouteOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *NetworkTypeRouteOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *NetworkTypeRouteOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *NetworkTypeRouteOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *NetworkTypeRouteOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *NetworkTypeRouteOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *NetworkTypeRouteOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *NetworkTypeRouteOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *NetworkTypeRouteOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *NetworkTypeRouteOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *NetworkTypeRouteOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *NetworkTypeRouteOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *NetworkTypeRouteOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetShowOnCreate

`func (o *NetworkTypeRouteOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *NetworkTypeRouteOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *NetworkTypeRouteOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *NetworkTypeRouteOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### GetShowOnEdit

`func (o *NetworkTypeRouteOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *NetworkTypeRouteOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *NetworkTypeRouteOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *NetworkTypeRouteOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


