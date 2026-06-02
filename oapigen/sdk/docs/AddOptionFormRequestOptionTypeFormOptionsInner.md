# AddOptionFormRequestOptionTypeFormOptionsInner

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
**OptionList** | Pointer to [**AddOptionFormRequestOptionTypeFormOptionsInnerOptionList**](AddOptionFormRequestOptionTypeFormOptionsInnerOptionList.md) |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** | Display Value On Details | [optional] [default to false]
**IsLocked** | Pointer to **bool** | Locked | [optional] [default to false]
**IsHidden** | Pointer to **bool** | Hidden | [optional] [default to false]
**ExcludeFromSearch** | Pointer to **bool** | Exclude From Search | [optional] [default to false]
**DependsOnCode** | Pointer to **NullableString** | A fieldName that will trigger reloading this input | [optional] 
**VisibleOnCode** | Pointer to **NullableString** | A fieldName that will trigger visibility of this input | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**RequireOnCode** | Pointer to **NullableString** | A fieldName that will trigger required attribute of this input | [optional] 

## Methods

### NewAddOptionFormRequestOptionTypeFormOptionsInner

`func NewAddOptionFormRequestOptionTypeFormOptionsInner(fieldName string, fieldLabel string, ) *AddOptionFormRequestOptionTypeFormOptionsInner`

NewAddOptionFormRequestOptionTypeFormOptionsInner instantiates a new AddOptionFormRequestOptionTypeFormOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldName

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetType

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.


### GetFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetPlaceHolder

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetHelpBlock

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetOptionList() AddOptionFormRequestOptionTypeFormOptionsInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetOptionListOk() (*AddOptionFormRequestOptionTypeFormOptionsInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetOptionList(v AddOptionFormRequestOptionTypeFormOptionsInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetIsLocked

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsHidden

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetVerifyPattern

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *AddOptionFormRequestOptionTypeFormOptionsInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


