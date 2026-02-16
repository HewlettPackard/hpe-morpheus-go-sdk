# ClusterLayoutOptionTypesInner

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
**OptionList** | Pointer to [**ClusterLayoutOptionTypesInnerOptionList**](ClusterLayoutOptionTypesInnerOptionList.md) |  | [optional] 
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

### NewClusterLayoutOptionTypesInner

`func NewClusterLayoutOptionTypesInner() *ClusterLayoutOptionTypesInner`

NewClusterLayoutOptionTypesInner instantiates a new ClusterLayoutOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutOptionTypesInnerWithDefaults

`func NewClusterLayoutOptionTypesInnerWithDefaults() *ClusterLayoutOptionTypesInner`

NewClusterLayoutOptionTypesInnerWithDefaults instantiates a new ClusterLayoutOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterLayoutOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterLayoutOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterLayoutOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterLayoutOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ClusterLayoutOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterLayoutOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterLayoutOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterLayoutOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterLayoutOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterLayoutOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ClusterLayoutOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ClusterLayoutOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *ClusterLayoutOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterLayoutOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterLayoutOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterLayoutOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterLayoutOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterLayoutOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ClusterLayoutOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ClusterLayoutOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ClusterLayoutOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *ClusterLayoutOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ClusterLayoutOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ClusterLayoutOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ClusterLayoutOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ClusterLayoutOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *ClusterLayoutOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *ClusterLayoutOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ClusterLayoutOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ClusterLayoutOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *ClusterLayoutOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *ClusterLayoutOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *ClusterLayoutOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *ClusterLayoutOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *ClusterLayoutOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *ClusterLayoutOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *ClusterLayoutOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *ClusterLayoutOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *ClusterLayoutOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *ClusterLayoutOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *ClusterLayoutOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *ClusterLayoutOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *ClusterLayoutOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *ClusterLayoutOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *ClusterLayoutOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *ClusterLayoutOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *ClusterLayoutOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *ClusterLayoutOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *ClusterLayoutOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *ClusterLayoutOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *ClusterLayoutOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *ClusterLayoutOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ClusterLayoutOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ClusterLayoutOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ClusterLayoutOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *ClusterLayoutOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *ClusterLayoutOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *ClusterLayoutOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ClusterLayoutOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ClusterLayoutOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ClusterLayoutOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *ClusterLayoutOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *ClusterLayoutOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *ClusterLayoutOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ClusterLayoutOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ClusterLayoutOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ClusterLayoutOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *ClusterLayoutOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *ClusterLayoutOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *ClusterLayoutOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ClusterLayoutOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ClusterLayoutOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *ClusterLayoutOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *ClusterLayoutOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *ClusterLayoutOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ClusterLayoutOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ClusterLayoutOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ClusterLayoutOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ClusterLayoutOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ClusterLayoutOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *ClusterLayoutOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *ClusterLayoutOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *ClusterLayoutOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *ClusterLayoutOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *ClusterLayoutOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *ClusterLayoutOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *ClusterLayoutOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *ClusterLayoutOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *ClusterLayoutOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *ClusterLayoutOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *ClusterLayoutOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *ClusterLayoutOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *ClusterLayoutOptionTypesInner) GetOptionList() ClusterLayoutOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ClusterLayoutOptionTypesInner) GetOptionListOk() (*ClusterLayoutOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ClusterLayoutOptionTypesInner) SetOptionList(v ClusterLayoutOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ClusterLayoutOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *ClusterLayoutOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClusterLayoutOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClusterLayoutOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ClusterLayoutOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *ClusterLayoutOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ClusterLayoutOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ClusterLayoutOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *ClusterLayoutOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *ClusterLayoutOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ClusterLayoutOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ClusterLayoutOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ClusterLayoutOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ClusterLayoutOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ClusterLayoutOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ClusterLayoutOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ClusterLayoutOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ClusterLayoutOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ClusterLayoutOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ClusterLayoutOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ClusterLayoutOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *ClusterLayoutOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ClusterLayoutOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ClusterLayoutOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ClusterLayoutOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *ClusterLayoutOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterLayoutOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterLayoutOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ClusterLayoutOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ClusterLayoutOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ClusterLayoutOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *ClusterLayoutOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ClusterLayoutOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ClusterLayoutOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ClusterLayoutOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *ClusterLayoutOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *ClusterLayoutOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *ClusterLayoutOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *ClusterLayoutOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *ClusterLayoutOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *ClusterLayoutOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *ClusterLayoutOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ClusterLayoutOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ClusterLayoutOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ClusterLayoutOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *ClusterLayoutOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *ClusterLayoutOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *ClusterLayoutOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *ClusterLayoutOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ClusterLayoutOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ClusterLayoutOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ClusterLayoutOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *ClusterLayoutOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *ClusterLayoutOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *ClusterLayoutOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ClusterLayoutOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ClusterLayoutOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *ClusterLayoutOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *ClusterLayoutOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *ClusterLayoutOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ClusterLayoutOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ClusterLayoutOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ClusterLayoutOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *ClusterLayoutOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *ClusterLayoutOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *ClusterLayoutOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *ClusterLayoutOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *ClusterLayoutOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *ClusterLayoutOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *ClusterLayoutOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *ClusterLayoutOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *ClusterLayoutOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ClusterLayoutOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ClusterLayoutOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ClusterLayoutOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *ClusterLayoutOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *ClusterLayoutOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *ClusterLayoutOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *ClusterLayoutOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *ClusterLayoutOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *ClusterLayoutOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *ClusterLayoutOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *ClusterLayoutOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *ClusterLayoutOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *ClusterLayoutOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *ClusterLayoutOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *ClusterLayoutOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *ClusterLayoutOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *ClusterLayoutOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *ClusterLayoutOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *ClusterLayoutOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *ClusterLayoutOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *ClusterLayoutOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *ClusterLayoutOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *ClusterLayoutOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


