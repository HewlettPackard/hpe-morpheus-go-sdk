# GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner

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

### NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner

`func NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner() *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner`

NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner instantiates a new GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInnerWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInnerWithDefaults() *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner`

NewGetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInnerWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionList() ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetOptionListOk() (*ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetOptionList(v ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *GetNetworkRouter200ResponseNetworkRouterTypeFirewallGroupOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


