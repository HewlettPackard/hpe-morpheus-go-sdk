# GetIntegrationTypeOptionTypes200ResponseOptionTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FieldName** | Pointer to **string** |  | [optional] 
**FieldLabel** | Pointer to **string** |  | [optional] 
**FieldCode** | Pointer to **string** |  | [optional] 
**FieldContext** | Pointer to **string** |  | [optional] 
**FieldGroup** | Pointer to **string** |  | [optional] 
**FieldClass** | Pointer to **string** |  | [optional] 
**FieldAddOn** | Pointer to **string** |  | [optional] 
**FieldComponent** | Pointer to **string** |  | [optional] 
**FieldInput** | Pointer to **string** |  | [optional] 
**PlaceHolder** | Pointer to **string** |  | [optional] 
**VerifyPattern** | Pointer to **string** |  | [optional] 
**HelpBlock** | Pointer to **string** |  | [optional] 
**HelpBlockFieldCode** | Pointer to **string** |  | [optional] 
**DefaultValue** | Pointer to **string** |  | [optional] 
**OptionSource** | Pointer to **string** |  | [optional] 
**OptionSourceType** | Pointer to **string** |  | [optional] 
**OptionList** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Advanced** | Pointer to **bool** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**ExportMeta** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**DisplayOrder** | Pointer to **int64** |  | [optional] 
**WrapperClass** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**NoBlank** | Pointer to **bool** |  | [optional] 
**DependsOnCode** | Pointer to **bool** |  | [optional] 
**VisibleOnCode** | Pointer to **bool** |  | [optional] 
**RequireOnCode** | Pointer to **bool** |  | [optional] 
**ContextualDefault** | Pointer to **bool** |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** |  | [optional] 
**ShowOnCreate** | Pointer to **bool** |  | [optional] 
**ShowOnEdit** | Pointer to **bool** |  | [optional] 
**LocalCredential** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInner

`func NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInner() *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner`

NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInner instantiates a new GetIntegrationTypeOptionTypes200ResponseOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInnerWithDefaults

`func NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInnerWithDefaults() *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner`

NewGetIntegrationTypeOptionTypes200ResponseOptionTypesInnerWithDefaults instantiates a new GetIntegrationTypeOptionTypes200ResponseOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetLabels() []map[string]interface{}`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetLabelsOk() (*[]map[string]interface{}, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetLabels(v []map[string]interface{})`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFieldName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetFieldLabel

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### GetFieldContext

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldContext() string`

GetFieldContext returns the FieldContext field if non-nil, zero value otherwise.

### GetFieldContextOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldContextOk() (*string, bool)`

GetFieldContextOk returns a tuple with the FieldContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldContext

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldContext(v string)`

SetFieldContext sets FieldContext field to given value.

### HasFieldContext

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldContext() bool`

HasFieldContext returns a boolean if a field has been set.

### GetFieldGroup

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldGroup() string`

GetFieldGroup returns the FieldGroup field if non-nil, zero value otherwise.

### GetFieldGroupOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldGroupOk() (*string, bool)`

GetFieldGroupOk returns a tuple with the FieldGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroup

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldGroup(v string)`

SetFieldGroup sets FieldGroup field to given value.

### HasFieldGroup

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldGroup() bool`

HasFieldGroup returns a boolean if a field has been set.

### GetFieldClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldClass() string`

GetFieldClass returns the FieldClass field if non-nil, zero value otherwise.

### GetFieldClassOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldClassOk() (*string, bool)`

GetFieldClassOk returns a tuple with the FieldClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldClass(v string)`

SetFieldClass sets FieldClass field to given value.

### HasFieldClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldClass() bool`

HasFieldClass returns a boolean if a field has been set.

### GetFieldAddOn

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldAddOn() string`

GetFieldAddOn returns the FieldAddOn field if non-nil, zero value otherwise.

### GetFieldAddOnOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldAddOnOk() (*string, bool)`

GetFieldAddOnOk returns a tuple with the FieldAddOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldAddOn

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldAddOn(v string)`

SetFieldAddOn sets FieldAddOn field to given value.

### HasFieldAddOn

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldAddOn() bool`

HasFieldAddOn returns a boolean if a field has been set.

### GetFieldComponent

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldComponent() string`

GetFieldComponent returns the FieldComponent field if non-nil, zero value otherwise.

### GetFieldComponentOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldComponentOk() (*string, bool)`

GetFieldComponentOk returns a tuple with the FieldComponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldComponent

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldComponent(v string)`

SetFieldComponent sets FieldComponent field to given value.

### HasFieldComponent

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldComponent() bool`

HasFieldComponent returns a boolean if a field has been set.

### GetFieldInput

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldInput() string`

GetFieldInput returns the FieldInput field if non-nil, zero value otherwise.

### GetFieldInputOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetFieldInputOk() (*string, bool)`

GetFieldInputOk returns a tuple with the FieldInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldInput

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetFieldInput(v string)`

SetFieldInput sets FieldInput field to given value.

### HasFieldInput

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasFieldInput() bool`

HasFieldInput returns a boolean if a field has been set.

### GetPlaceHolder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetVerifyPattern

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetHelpBlock

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### GetHelpBlockFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### GetDefaultValue

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetOptionSource

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionSource() string`

GetOptionSource returns the OptionSource field if non-nil, zero value otherwise.

### GetOptionSourceOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionSourceOk() (*string, bool)`

GetOptionSourceOk returns a tuple with the OptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSource

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetOptionSource(v string)`

SetOptionSource sets OptionSource field to given value.

### HasOptionSource

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasOptionSource() bool`

HasOptionSource returns a boolean if a field has been set.

### GetOptionSourceType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionSourceType() string`

GetOptionSourceType returns the OptionSourceType field if non-nil, zero value otherwise.

### GetOptionSourceTypeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionSourceTypeOk() (*string, bool)`

GetOptionSourceTypeOk returns a tuple with the OptionSourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionSourceType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetOptionSourceType(v string)`

SetOptionSourceType sets OptionSourceType field to given value.

### HasOptionSourceType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasOptionSourceType() bool`

HasOptionSourceType returns a boolean if a field has been set.

### GetOptionList

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionList() string`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetOptionListOk() (*string, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetOptionList(v string)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAdvanced

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetAdvanced() bool`

GetAdvanced returns the Advanced field if non-nil, zero value otherwise.

### GetAdvancedOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetAdvancedOk() (*bool, bool)`

GetAdvancedOk returns a tuple with the Advanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanced

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetAdvanced(v bool)`

SetAdvanced sets Advanced field to given value.

### HasAdvanced

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasAdvanced() bool`

HasAdvanced returns a boolean if a field has been set.

### GetRequired

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetCreatable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDisplayOrder() int64`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDisplayOrderOk() (*int64, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetDisplayOrder(v int64)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetWrapperClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetWrapperClass() string`

GetWrapperClass returns the WrapperClass field if non-nil, zero value otherwise.

### GetWrapperClassOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetWrapperClassOk() (*string, bool)`

GetWrapperClassOk returns a tuple with the WrapperClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapperClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetWrapperClass(v string)`

SetWrapperClass sets WrapperClass field to given value.

### HasWrapperClass

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasWrapperClass() bool`

HasWrapperClass returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetNoBlank

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetNoBlank() bool`

GetNoBlank returns the NoBlank field if non-nil, zero value otherwise.

### GetNoBlankOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetNoBlankOk() (*bool, bool)`

GetNoBlankOk returns a tuple with the NoBlank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoBlank

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetNoBlank(v bool)`

SetNoBlank sets NoBlank field to given value.

### HasNoBlank

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasNoBlank() bool`

HasNoBlank returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDependsOnCode() bool`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDependsOnCodeOk() (*bool, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetDependsOnCode(v bool)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetVisibleOnCode() bool`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetVisibleOnCodeOk() (*bool, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetVisibleOnCode(v bool)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetRequireOnCode() bool`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetRequireOnCodeOk() (*bool, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetRequireOnCode(v bool)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### GetContextualDefault

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetContextualDefault() bool`

GetContextualDefault returns the ContextualDefault field if non-nil, zero value otherwise.

### GetContextualDefaultOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetContextualDefaultOk() (*bool, bool)`

GetContextualDefaultOk returns a tuple with the ContextualDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextualDefault

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetContextualDefault(v bool)`

SetContextualDefault sets ContextualDefault field to given value.

### HasContextualDefault

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasContextualDefault() bool`

HasContextualDefault returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetShowOnCreate

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetShowOnCreate() bool`

GetShowOnCreate returns the ShowOnCreate field if non-nil, zero value otherwise.

### GetShowOnCreateOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetShowOnCreateOk() (*bool, bool)`

GetShowOnCreateOk returns a tuple with the ShowOnCreate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnCreate

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetShowOnCreate(v bool)`

SetShowOnCreate sets ShowOnCreate field to given value.

### HasShowOnCreate

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasShowOnCreate() bool`

HasShowOnCreate returns a boolean if a field has been set.

### GetShowOnEdit

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetShowOnEdit() bool`

GetShowOnEdit returns the ShowOnEdit field if non-nil, zero value otherwise.

### GetShowOnEditOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetShowOnEditOk() (*bool, bool)`

GetShowOnEditOk returns a tuple with the ShowOnEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowOnEdit

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetShowOnEdit(v bool)`

SetShowOnEdit sets ShowOnEdit field to given value.

### HasShowOnEdit

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasShowOnEdit() bool`

HasShowOnEdit returns a boolean if a field has been set.

### GetLocalCredential

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetLocalCredential() bool`

GetLocalCredential returns the LocalCredential field if non-nil, zero value otherwise.

### GetLocalCredentialOk

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) GetLocalCredentialOk() (*bool, bool)`

GetLocalCredentialOk returns a tuple with the LocalCredential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalCredential

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) SetLocalCredential(v bool)`

SetLocalCredential sets LocalCredential field to given value.

### HasLocalCredential

`func (o *GetIntegrationTypeOptionTypes200ResponseOptionTypesInner) HasLocalCredential() bool`

HasLocalCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


