# ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner

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

### NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner

`func NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner() *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner`

NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner instantiates a new ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInnerWithDefaults

`func NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInnerWithDefaults() *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner`

NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInnerWithDefaults instantiates a new ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetDescription

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetFieldContext

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### SetFieldGroupNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldGroupNil(b bool)`

 SetFieldGroupNil sets the value for FieldGroup to be an explicit nil

### UnsetFieldGroup
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldGroup()`

UnsetFieldGroup ensures that no value is present for FieldGroup, not even an explicit nil
### GetFieldClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### SetFieldClassNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldClassNil(b bool)`

 SetFieldClassNil sets the value for FieldClass to be an explicit nil

### UnsetFieldClass
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldClass()`

UnsetFieldClass ensures that no value is present for FieldClass, not even an explicit nil
### GetFieldAddOn

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### SetFieldAddOnNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldAddOnNil(b bool)`

 SetFieldAddOnNil sets the value for FieldAddOn to be an explicit nil

### UnsetFieldAddOn
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldAddOn()`

UnsetFieldAddOn ensures that no value is present for FieldAddOn, not even an explicit nil
### GetFieldComponent

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### SetFieldComponentNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldComponentNil(b bool)`

 SetFieldComponentNil sets the value for FieldComponent to be an explicit nil

### UnsetFieldComponent
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldComponent()`

UnsetFieldComponent ensures that no value is present for FieldComponent, not even an explicit nil
### GetFieldInput

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### SetFieldInputNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetFieldInputNil(b bool)`

 SetFieldInputNil sets the value for FieldInput to be an explicit nil

### UnsetFieldInput
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetFieldInput()`

UnsetFieldInput ensures that no value is present for FieldInput, not even an explicit nil
### GetPlaceHolder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### SetPlaceHolderNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetPlaceHolderNil(b bool)`

 SetPlaceHolderNil sets the value for PlaceHolder to be an explicit nil

### UnsetPlaceHolder
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetPlaceHolder()`

UnsetPlaceHolder ensures that no value is present for PlaceHolder, not even an explicit nil
### GetVerifyPattern

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### SetVerifyPatternNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetVerifyPatternNil(b bool)`

 SetVerifyPatternNil sets the value for VerifyPattern to be an explicit nil

### UnsetVerifyPattern
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetVerifyPattern()`

UnsetVerifyPattern ensures that no value is present for VerifyPattern, not even an explicit nil
### GetHelpBlock

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### SetDefaultValueNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDefaultValueNil(b bool)`

 SetDefaultValueNil sets the value for DefaultValue to be an explicit nil

### UnsetDefaultValue
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetDefaultValue()`

UnsetDefaultValue ensures that no value is present for DefaultValue, not even an explicit nil
### GetOptionSource

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### SetOptionSourceNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetOptionSourceNil(b bool)`

 SetOptionSourceNil sets the value for OptionSource to be an explicit nil

### UnsetOptionSource
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetOptionSource()`

UnsetOptionSource ensures that no value is present for OptionSource, not even an explicit nil
### GetOptionSourceType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### SetOptionSourceTypeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetOptionSourceTypeNil(b bool)`

 SetOptionSourceTypeNil sets the value for OptionSourceType to be an explicit nil

### UnsetOptionSourceType
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetOptionSourceType()`

UnsetOptionSourceType ensures that no value is present for OptionSourceType, not even an explicit nil
### GetOptionList

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionList() ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetOptionListOk() (*ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetOptionList(v ListClusterTypes200ResponseAllOfClusterTypesInnerOptionTypesInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetDisplayOrder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### SetWrapperClassNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetWrapperClassNil(b bool)`

 SetWrapperClassNil sets the value for WrapperClass to be an explicit nil

### UnsetWrapperClass
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetWrapperClass()`

UnsetWrapperClass ensures that no value is present for WrapperClass, not even an explicit nil
### GetEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetRequireOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil
### GetContextualDefault

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### SetContextualDefaultNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetContextualDefaultNil(b bool)`

 SetContextualDefaultNil sets the value for ContextualDefault to be an explicit nil

### UnsetContextualDefault
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetContextualDefault()`

UnsetContextualDefault ensures that no value is present for ContextualDefault, not even an explicit nil
### GetDisplayValueOnDetails

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### SetDisplayValueOnDetailsNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetDisplayValueOnDetailsNil(b bool)`

 SetDisplayValueOnDetailsNil sets the value for DisplayValueOnDetails to be an explicit nil

### UnsetDisplayValueOnDetails
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetDisplayValueOnDetails()`

UnsetDisplayValueOnDetails ensures that no value is present for DisplayValueOnDetails, not even an explicit nil
### GetShowOnCreate

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### SetShowOnCreateNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetShowOnCreateNil(b bool)`

 SetShowOnCreateNil sets the value for ShowOnCreate to be an explicit nil

### UnsetShowOnCreate
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetShowOnCreate()`

UnsetShowOnCreate ensures that no value is present for ShowOnCreate, not even an explicit nil
### GetShowOnEdit

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### SetShowOnEditNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetShowOnEditNil(b bool)`

 SetShowOnEditNil sets the value for ShowOnEdit to be an explicit nil

### UnsetShowOnEdit
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetShowOnEdit()`

UnsetShowOnEdit ensures that no value is present for ShowOnEdit, not even an explicit nil
### GetLocalCredential

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.

### SetLocalCredentialNil

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) SetLocalCredentialNil(b bool)`

 SetLocalCredentialNil sets the value for LocalCredential to be an explicit nil

### UnsetLocalCredential
`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerVipOptionTypesInner) UnsetLocalCredential()`

UnsetLocalCredential ensures that no value is present for LocalCredential, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


