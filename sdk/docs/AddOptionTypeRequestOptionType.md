# AddOptionTypeRequestOptionType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the option type for handy reference | 
**Description** | Pointer to **string** | Short description of the option type | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**FieldName** | Pointer to **string** | Field Name, the name for user input. This along with fieldContext determines the configuration property name.  The property key for when posting this option type to a JSON POST request | [optional] 
**Type** | Pointer to **string** | Type, the type of input. eg. text, checkbox, select, etc. | [optional] [default to "text"]
**FieldLabel** | Pointer to **string** | Field Label, the label for user input. | [optional] 
**PlaceHolder** | Pointer to **string** | Any placeholder text when nothing is yet entered | [optional] 
**VerifyPattern** | Pointer to **string** | Verify Pattern, A regexp string that validates the input, use (?i) to make the matcher case insensitive | [optional] 
**DefaultValue** | Pointer to **string** | The default value if no user entry is specified. This value should be passed to the desired JSON Map if nothing else is entered | [optional] 
**Required** | Pointer to **bool** | Is this field entry required for the request | [optional] [default to false]
**ExportMeta** | Pointer to **bool** | Export as Tag | [optional] [default to false]
**Editable** | Pointer to **bool** | Used primarily on tasks and workflows. Basically wether or not the field can be overridden optionally when the object is run | [optional] [default to false]
**OptionList** | Pointer to [**AddOptionTypeRequestOptionTypeOptionList**](AddOptionTypeRequestOptionTypeOptionList.md) |  | [optional] 

## Methods

### NewAddOptionTypeRequestOptionType

`func NewAddOptionTypeRequestOptionType(name string, ) *AddOptionTypeRequestOptionType`

NewAddOptionTypeRequestOptionType instantiates a new AddOptionTypeRequestOptionType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOptionTypeRequestOptionTypeWithDefaults

`func NewAddOptionTypeRequestOptionTypeWithDefaults() *AddOptionTypeRequestOptionType`

NewAddOptionTypeRequestOptionTypeWithDefaults instantiates a new AddOptionTypeRequestOptionType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddOptionTypeRequestOptionType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOptionTypeRequestOptionType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOptionTypeRequestOptionType) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddOptionTypeRequestOptionType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOptionTypeRequestOptionType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOptionTypeRequestOptionType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOptionTypeRequestOptionType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddOptionTypeRequestOptionType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddOptionTypeRequestOptionType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddOptionTypeRequestOptionType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddOptionTypeRequestOptionType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetFieldName

`func (o *AddOptionTypeRequestOptionType) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *AddOptionTypeRequestOptionType) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *AddOptionTypeRequestOptionType) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *AddOptionTypeRequestOptionType) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetType

`func (o *AddOptionTypeRequestOptionType) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddOptionTypeRequestOptionType) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddOptionTypeRequestOptionType) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddOptionTypeRequestOptionType) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFieldLabel

`func (o *AddOptionTypeRequestOptionType) GetFieldLabel() string`

GetFieldLabel returns the FieldLabel field if non-nil, zero value otherwise.

### GetFieldLabelOk

`func (o *AddOptionTypeRequestOptionType) GetFieldLabelOk() (*string, bool)`

GetFieldLabelOk returns a tuple with the FieldLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldLabel

`func (o *AddOptionTypeRequestOptionType) SetFieldLabel(v string)`

SetFieldLabel sets FieldLabel field to given value.

### HasFieldLabel

`func (o *AddOptionTypeRequestOptionType) HasFieldLabel() bool`

HasFieldLabel returns a boolean if a field has been set.

### GetPlaceHolder

`func (o *AddOptionTypeRequestOptionType) GetPlaceHolder() string`

GetPlaceHolder returns the PlaceHolder field if non-nil, zero value otherwise.

### GetPlaceHolderOk

`func (o *AddOptionTypeRequestOptionType) GetPlaceHolderOk() (*string, bool)`

GetPlaceHolderOk returns a tuple with the PlaceHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceHolder

`func (o *AddOptionTypeRequestOptionType) SetPlaceHolder(v string)`

SetPlaceHolder sets PlaceHolder field to given value.

### HasPlaceHolder

`func (o *AddOptionTypeRequestOptionType) HasPlaceHolder() bool`

HasPlaceHolder returns a boolean if a field has been set.

### GetVerifyPattern

`func (o *AddOptionTypeRequestOptionType) GetVerifyPattern() string`

GetVerifyPattern returns the VerifyPattern field if non-nil, zero value otherwise.

### GetVerifyPatternOk

`func (o *AddOptionTypeRequestOptionType) GetVerifyPatternOk() (*string, bool)`

GetVerifyPatternOk returns a tuple with the VerifyPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyPattern

`func (o *AddOptionTypeRequestOptionType) SetVerifyPattern(v string)`

SetVerifyPattern sets VerifyPattern field to given value.

### HasVerifyPattern

`func (o *AddOptionTypeRequestOptionType) HasVerifyPattern() bool`

HasVerifyPattern returns a boolean if a field has been set.

### GetDefaultValue

`func (o *AddOptionTypeRequestOptionType) GetDefaultValue() string`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *AddOptionTypeRequestOptionType) GetDefaultValueOk() (*string, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *AddOptionTypeRequestOptionType) SetDefaultValue(v string)`

SetDefaultValue sets DefaultValue field to given value.

### HasDefaultValue

`func (o *AddOptionTypeRequestOptionType) HasDefaultValue() bool`

HasDefaultValue returns a boolean if a field has been set.

### GetRequired

`func (o *AddOptionTypeRequestOptionType) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *AddOptionTypeRequestOptionType) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *AddOptionTypeRequestOptionType) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *AddOptionTypeRequestOptionType) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetExportMeta

`func (o *AddOptionTypeRequestOptionType) GetExportMeta() bool`

GetExportMeta returns the ExportMeta field if non-nil, zero value otherwise.

### GetExportMetaOk

`func (o *AddOptionTypeRequestOptionType) GetExportMetaOk() (*bool, bool)`

GetExportMetaOk returns a tuple with the ExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportMeta

`func (o *AddOptionTypeRequestOptionType) SetExportMeta(v bool)`

SetExportMeta sets ExportMeta field to given value.

### HasExportMeta

`func (o *AddOptionTypeRequestOptionType) HasExportMeta() bool`

HasExportMeta returns a boolean if a field has been set.

### GetEditable

`func (o *AddOptionTypeRequestOptionType) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddOptionTypeRequestOptionType) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddOptionTypeRequestOptionType) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddOptionTypeRequestOptionType) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionList

`func (o *AddOptionTypeRequestOptionType) GetOptionList() AddOptionTypeRequestOptionTypeOptionList`

GetOptionList returns the OptionList field if non-nil, zero value otherwise.

### GetOptionListOk

`func (o *AddOptionTypeRequestOptionType) GetOptionListOk() (*AddOptionTypeRequestOptionTypeOptionList, bool)`

GetOptionListOk returns a tuple with the OptionList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionList

`func (o *AddOptionTypeRequestOptionType) SetOptionList(v AddOptionTypeRequestOptionTypeOptionList)`

SetOptionList sets OptionList field to given value.

### HasOptionList

`func (o *AddOptionTypeRequestOptionType) HasOptionList() bool`

HasOptionList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


