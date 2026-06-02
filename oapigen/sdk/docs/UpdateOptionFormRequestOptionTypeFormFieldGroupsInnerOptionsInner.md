# UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner

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
**OptionList** | Pointer to [**UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInnerOptionList**](UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInnerOptionList.md) |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** | Display Value On Details | [optional] [default to false]
**IsLocked** | Pointer to **bool** | Locked | [optional] [default to false]
**IsHidden** | Pointer to **bool** | Hidden | [optional] [default to false]
**ExcludeFromSearch** | Pointer to **bool** | Exclude From Search | [optional] [default to false]
**DependsOnCode** | Pointer to **NullableString** | A fieldName that will trigger reloading this input | [optional] 
**VisibleOnCode** | Pointer to **NullableString** | A fieldName that will trigger visibility of this input | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**RequireOnCode** | Pointer to **NullableString** | A fieldName that will trigger required attribute of this input | [optional] 

## Methods

### NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner

`func NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner(fieldName string, fieldLabel string, ) *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner`

NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner instantiates a new UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetType

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.


### GetFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetPlaceHolder

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetHelpBlock

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetOptionList() UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetOptionListOk() (*UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetOptionList(v UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetIsLocked

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsHidden

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetVerifyPattern

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


