# ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the input, include this to use an existing input or to update an existing form input record instead of creating a new one. | [optional] 
**Code** | Pointer to **string** | The code of the option type as a globally unique identifier. By default a UUID will be used. | [optional] 
**Name** | Pointer to **string** | The name of the option type for handy reference. By default a UUID will be used. | [optional] 
**Description** | Pointer to **string** | Short description of the option type | [optional] 
**FieldName** | **string** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | 
**Type** | Pointer to **string** | Type, the type of input. eg. text, checkbox, select, etc. | [optional] [default to "text"]
**FieldLabel** | **string** | Field Label, the label for user input. | 
**FieldCode** | Pointer to **string** | Localized Label, i18n code for the label | [optional] 
**PlaceHolder** | Pointer to **string** | Any placeholder text when nothing is yet entered | [optional] 
**HelpBlock** | Pointer to **string** | This is the explaination of the input that shows typically underneath the option | [optional] 
**HelpBlockFieldCode** | Pointer to **string** | Localized Help Block, i18n code for the help block | [optional] 
**DefaultValue** | Pointer to **string** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional] 
**Required** | Pointer to **bool** | Is this field entry required for the request | [optional] [default to false]
**ExportMeta** | Pointer to **bool** | Export as Tag | [optional] [default to false]
**Editable** | Pointer to **bool** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional] [default to false]
**OptionList** | Pointer to [**AddOptionTypeRequestOptionTypeOptionList**](AddOptionTypeRequestOptionTypeOptionList.md) |  | [optional] 
**DisplayValueOnDetails** | Pointer to **bool** | Display Value On Details | [optional] [default to false]
**IsLocked** | Pointer to **bool** | Locked | [optional] [default to false]
**IsHidden** | Pointer to **bool** | Hidden | [optional] [default to false]
**ExcludeFromSearch** | Pointer to **bool** | Exclude From Search | [optional] [default to false]
**DependsOnCode** | Pointer to **string** | A fieldName that will trigger reloading this input | [optional] 
**VisibleOnCode** | Pointer to **string** | A fieldName that will trigger visibility of this input | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**RequireOnCode** | Pointer to **string** | A fieldName that will trigger required attribute of this input | [optional] 

## Methods

### NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner

`func NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner(fieldName string, fieldLabel string, ) *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner instantiates a new ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInnerWithDefaults

`func NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInnerWithDefaults() *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

NewListOptionForms200ResponseAllOfOptionTypesInnerOptionsInnerWithDefaults instantiates a new ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldName

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.


### GetType

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.


### GetFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldCode() string`

GetFieldCode returns the FieldCode field if non-nil, zero value otherwise.

### GetFieldCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetFieldCodeOk() (*string, bool)`

GetFieldCodeOk returns a tuple with the FieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetFieldCode(v string)`

SetFieldCode sets FieldCode field to given value.

### HasFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasFieldCode() bool`

HasFieldCode returns a boolean if a field has been set.

### GetPlaceHolder

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetHelpBlock

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetHelpBlock() string`

GetHelpBlock returns the HelpBlock field if non-nil, zero value otherwise.

### GetHelpBlockOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetHelpBlockOk() (*string, bool)`

GetHelpBlockOk returns a tuple with the HelpBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlock

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetHelpBlock(v string)`

SetHelpBlock sets HelpBlock field to given value.

### HasHelpBlock

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasHelpBlock() bool`

HasHelpBlock returns a boolean if a field has been set.

### GetHelpBlockFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetHelpBlockFieldCode() string`

GetHelpBlockFieldCode returns the HelpBlockFieldCode field if non-nil, zero value otherwise.

### GetHelpBlockFieldCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetHelpBlockFieldCodeOk() (*string, bool)`

GetHelpBlockFieldCodeOk returns a tuple with the HelpBlockFieldCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpBlockFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetHelpBlockFieldCode(v string)`

SetHelpBlockFieldCode sets HelpBlockFieldCode field to given value.

### HasHelpBlockFieldCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasHelpBlockFieldCode() bool`

HasHelpBlockFieldCode returns a boolean if a field has been set.

### GetDefaultValue

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetOptionList() AddOptionTypeRequestOptionTypeOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetOptionListOk() (*AddOptionTypeRequestOptionTypeOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetOptionList(v AddOptionTypeRequestOptionTypeOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.

### GetDisplayValueOnDetails

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDisplayValueOnDetails() bool`

GetDisplayValueOnDetails returns the DisplayValueOnDetails field if non-nil, zero value otherwise.

### GetDisplayValueOnDetailsOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDisplayValueOnDetailsOk() (*bool, bool)`

GetDisplayValueOnDetailsOk returns a tuple with the DisplayValueOnDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayValueOnDetails

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetDisplayValueOnDetails(v bool)`

SetDisplayValueOnDetails sets DisplayValueOnDetails field to given value.

### HasDisplayValueOnDetails

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasDisplayValueOnDetails() bool`

HasDisplayValueOnDetails returns a boolean if a field has been set.

### GetIsLocked

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsHidden

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetExcludeFromSearch

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetExcludeFromSearch() bool`

GetExcludeFromSearch returns the ExcludeFromSearch field if non-nil, zero value otherwise.

### GetExcludeFromSearchOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetExcludeFromSearchOk() (*bool, bool)`

GetExcludeFromSearchOk returns a tuple with the ExcludeFromSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeFromSearch

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetExcludeFromSearch(v bool)`

SetExcludeFromSearch sets ExcludeFromSearch field to given value.

### HasExcludeFromSearch

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasExcludeFromSearch() bool`

HasExcludeFromSearch returns a boolean if a field has been set.

### GetDependsOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDependsOnCode() string`

GetDependsOnCode returns the DependsOnCode field if non-nil, zero value otherwise.

### GetDependsOnCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetDependsOnCodeOk() (*string, bool)`

GetDependsOnCodeOk returns a tuple with the DependsOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependsOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetDependsOnCode(v string)`

SetDependsOnCode sets DependsOnCode field to given value.

### HasDependsOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasDependsOnCode() bool`

HasDependsOnCode returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### GetVerifyPattern

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetRequireOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetRequireOnCode() string`

GetRequireOnCode returns the RequireOnCode field if non-nil, zero value otherwise.

### GetRequireOnCodeOk

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) GetRequireOnCodeOk() (*string, bool)`

GetRequireOnCodeOk returns a tuple with the RequireOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) SetRequireOnCode(v string)`

SetRequireOnCode sets RequireOnCode field to given value.

### HasRequireOnCode

`func (o *ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner) HasRequireOnCode() bool`

HasRequireOnCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


