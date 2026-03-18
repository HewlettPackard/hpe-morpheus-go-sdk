# OptionTypeFormCreateOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the input, include this to use an existing input or to update an existing form input record instead of creating a new one. | [optional] 
**Code** | Pointer to **string** | The code of the option type as a globally unique identifier. By default a UUID will be used. | [optional] 
**Name** | Pointer to **string** | The name of the option type for handy reference. By default a UUID will be used. | [optional] 
**Description** | Pointer to **NullableString** | Short description of the option type | [optional] 
**FieldName** | **string** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | 
**Type** | Pointer to **string** | Type, the type of input. eg. text, checkbox, select, etc. | [optional] [default to "text"]
**FieldLabel** | **string** | Field Label, the label for user input. | 
**FieldCode** | Pointer to **NullableString** | Localized Label, i18n code for the label | [optional] 
**PlaceHolder** | Pointer to **string** | Any placeholder text when nothing is yet entered | [optional] 
**HelpBlock** | Pointer to **NullableString** | This is the explaination of the input that shows typically underneath the option | [optional] 
**HelpBlockFieldCode** | Pointer to **NullableString** | Localized Help Block, i18n code for the help block | [optional] 
**DefaultValue** | Pointer to **string** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional] 
**Required** | Pointer to **bool** | Is this field entry required for the request | [optional] [default to false]
**ExportMeta** | Pointer to **bool** | Export as Tag | [optional] [default to false]
**Editable** | Pointer to **bool** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional] [default to false]
**OptionList** | Pointer to [**OptionTypeFormCreateOptionsInnerOptionList**](OptionTypeFormCreateOptionsInnerOptionList.md) |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** | Display Value On Details | [optional] [default to false]
**IsLocked** | Pointer to **bool** | Locked | [optional] [default to false]
**IsHidden** | Pointer to **bool** | Hidden | [optional] [default to false]
**ExcludeFromSearch** | Pointer to **bool** | Exclude From Search | [optional] [default to false]
**DependsOnCode** | Pointer to **NullableString** | A fieldName that will trigger reloading this input | [optional] 
**VisibleOnCode** | Pointer to **NullableString** | A fieldName that will trigger visibility of this input | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**RequireOnCode** | Pointer to **NullableString** | A fieldName that will trigger required attribute of this input | [optional] 

## Methods

### NewOptionTypeFormCreateOptionsInner

`func NewOptionTypeFormCreateOptionsInner(fieldName string, fieldLabel string, ) *OptionTypeFormCreateOptionsInner`

NewOptionTypeFormCreateOptionsInner instantiates a new OptionTypeFormCreateOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeFormCreateOptionsInnerWithDefaults

`func NewOptionTypeFormCreateOptionsInnerWithDefaults() *OptionTypeFormCreateOptionsInner`

NewOptionTypeFormCreateOptionsInnerWithDefaults instantiates a new OptionTypeFormCreateOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionTypeFormCreateOptionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeFormCreateOptionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeFormCreateOptionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeFormCreateOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeFormCreateOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeFormCreateOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeFormCreateOptionsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeFormCreateOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeFormCreateOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeFormCreateOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeFormCreateOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeFormCreateOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeFormCreateOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeFormCreateOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeFormCreateOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptionTypeFormCreateOptionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptionTypeFormCreateOptionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldName

`func (o *OptionTypeFormCreateOptionsInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *OptionTypeFormCreateOptionsInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *OptionTypeFormCreateOptionsInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetType

`func (o *OptionTypeFormCreateOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionTypeFormCreateOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionTypeFormCreateOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionTypeFormCreateOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *OptionTypeFormCreateOptionsInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *OptionTypeFormCreateOptionsInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *OptionTypeFormCreateOptionsInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.


### GetFieldCode

`func (o *OptionTypeFormCreateOptionsInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *OptionTypeFormCreateOptionsInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *OptionTypeFormCreateOptionsInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *OptionTypeFormCreateOptionsInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *OptionTypeFormCreateOptionsInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetPlaceHolder

`func (o *OptionTypeFormCreateOptionsInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *OptionTypeFormCreateOptionsInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *OptionTypeFormCreateOptionsInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *OptionTypeFormCreateOptionsInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetHelpBlock

`func (o *OptionTypeFormCreateOptionsInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *OptionTypeFormCreateOptionsInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *OptionTypeFormCreateOptionsInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *OptionTypeFormCreateOptionsInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *OptionTypeFormCreateOptionsInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *OptionTypeFormCreateOptionsInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *OptionTypeFormCreateOptionsInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *OptionTypeFormCreateOptionsInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *OptionTypeFormCreateOptionsInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *OptionTypeFormCreateOptionsInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *OptionTypeFormCreateOptionsInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *OptionTypeFormCreateOptionsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *OptionTypeFormCreateOptionsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *OptionTypeFormCreateOptionsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *OptionTypeFormCreateOptionsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *OptionTypeFormCreateOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OptionTypeFormCreateOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OptionTypeFormCreateOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *OptionTypeFormCreateOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *OptionTypeFormCreateOptionsInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *OptionTypeFormCreateOptionsInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *OptionTypeFormCreateOptionsInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *OptionTypeFormCreateOptionsInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *OptionTypeFormCreateOptionsInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *OptionTypeFormCreateOptionsInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *OptionTypeFormCreateOptionsInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *OptionTypeFormCreateOptionsInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *OptionTypeFormCreateOptionsInner) GetOptionList() OptionTypeFormCreateOptionsInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *OptionTypeFormCreateOptionsInner) GetOptionListOk() (*OptionTypeFormCreateOptionsInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *OptionTypeFormCreateOptionsInner) SetOptionList(v OptionTypeFormCreateOptionsInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *OptionTypeFormCreateOptionsInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *OptionTypeFormCreateOptionsInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *OptionTypeFormCreateOptionsInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *OptionTypeFormCreateOptionsInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *OptionTypeFormCreateOptionsInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetIsLocked

`func (o *OptionTypeFormCreateOptionsInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *OptionTypeFormCreateOptionsInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *OptionTypeFormCreateOptionsInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *OptionTypeFormCreateOptionsInner) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsHidden

`func (o *OptionTypeFormCreateOptionsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *OptionTypeFormCreateOptionsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *OptionTypeFormCreateOptionsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *OptionTypeFormCreateOptionsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *OptionTypeFormCreateOptionsInner) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *OptionTypeFormCreateOptionsInner) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *OptionTypeFormCreateOptionsInner) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *OptionTypeFormCreateOptionsInner) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *OptionTypeFormCreateOptionsInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *OptionTypeFormCreateOptionsInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *OptionTypeFormCreateOptionsInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *OptionTypeFormCreateOptionsInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *OptionTypeFormCreateOptionsInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *OptionTypeFormCreateOptionsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *OptionTypeFormCreateOptionsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *OptionTypeFormCreateOptionsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *OptionTypeFormCreateOptionsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *OptionTypeFormCreateOptionsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetVerifyPattern

`func (o *OptionTypeFormCreateOptionsInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *OptionTypeFormCreateOptionsInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *OptionTypeFormCreateOptionsInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *OptionTypeFormCreateOptionsInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *OptionTypeFormCreateOptionsInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *OptionTypeFormCreateOptionsInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *OptionTypeFormCreateOptionsInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *OptionTypeFormCreateOptionsInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *OptionTypeFormCreateOptionsInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *OptionTypeFormCreateOptionsInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


