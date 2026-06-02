# GetOptionForm200ResponseOptionTypesInnerOptionsInner

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
**OptionList** | Pointer to [**GetOptionForm200ResponseOptionTypesInnerOptionsInnerOptionList**](GetOptionForm200ResponseOptionTypesInnerOptionsInnerOptionList.md) |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** | Display Value On Details | [optional] [default to false]
**IsLocked** | Pointer to **bool** | Locked | [optional] [default to false]
**IsHidden** | Pointer to **bool** | Hidden | [optional] [default to false]
**ExcludeFromSearch** | Pointer to **bool** | Exclude From Search | [optional] [default to false]
**DependsOnCode** | Pointer to **NullableString** | A fieldName that will trigger reloading this input | [optional] 
**VisibleOnCode** | Pointer to **NullableString** | A fieldName that will trigger visibility of this input | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**RequireOnCode** | Pointer to **NullableString** | A fieldName that will trigger required attribute of this input | [optional] 

## Methods

### NewGetOptionForm200ResponseOptionTypesInnerOptionsInner

`func NewGetOptionForm200ResponseOptionTypesInnerOptionsInner(fieldName string, fieldLabel string, ) *GetOptionForm200ResponseOptionTypesInnerOptionsInner`

NewGetOptionForm200ResponseOptionTypesInnerOptionsInner instantiates a new GetOptionForm200ResponseOptionTypesInnerOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetFieldName

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetType

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.


### GetFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### SetFieldCodeNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetFieldCodeNil(b bool)`

 SetFieldCodeNil sets the value for FieldCode to be an explicit nil

### UnsetFieldCode
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetFieldCode()`

UnsetFieldCode ensures that no value is present for FieldCode, not even an explicit nil
### GetPlaceHolder

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetHelpBlock

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### SetHelpBlockNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetHelpBlockNil(b bool)`

 SetHelpBlockNil sets the value for HelpBlock to be an explicit nil

### UnsetHelpBlock
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetHelpBlock()`

UnsetHelpBlock ensures that no value is present for HelpBlock, not even an explicit nil
### GetHelpBlockFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### SetHelpBlockFieldCodeNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetHelpBlockFieldCodeNil(b bool)`

 SetHelpBlockFieldCodeNil sets the value for HelpBlockFieldCode to be an explicit nil

### UnsetHelpBlockFieldCode
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetHelpBlockFieldCode()`

UnsetHelpBlockFieldCode ensures that no value is present for HelpBlockFieldCode, not even an explicit nil
### GetDefaultValue

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetOptionList() GetOptionForm200ResponseOptionTypesInnerOptionsInnerOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetOptionListOk() (*GetOptionForm200ResponseOptionTypesInnerOptionsInnerOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetOptionList(v GetOptionForm200ResponseOptionTypesInnerOptionsInnerOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetIsLocked

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsHidden

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### SetDependsOnCodeNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetDependsOnCodeNil(b bool)`

 SetDependsOnCodeNil sets the value for DependsOnCode to be an explicit nil

### UnsetDependsOnCode
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetDependsOnCode()`

UnsetDependsOnCode ensures that no value is present for DependsOnCode, not even an explicit nil
### GetVisibleOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetVerifyPattern

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.

### SetRequireOnCodeNil

`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) SetRequireOnCodeNil(b bool)`

 SetRequireOnCodeNil sets the value for RequireOnCode to be an explicit nil

### UnsetRequireOnCode
`func (o *GetOptionForm200ResponseOptionTypesInnerOptionsInner) UnsetRequireOnCode()`

UnsetRequireOnCode ensures that no value is present for RequireOnCode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


