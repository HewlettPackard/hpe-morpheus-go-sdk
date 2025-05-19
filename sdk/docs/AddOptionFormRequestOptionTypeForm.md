# AddOptionFormRequestOptionTypeForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Form name | [optional] 
**Code** | Pointer to **string** | Unique form code | [optional] 
**Description** | Pointer to **string** | A short description of the form | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to [**[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner**](ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner.md) | Inputs | [optional] 
**FieldGroups** | Pointer to [**[]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner**](ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner.md) | Field Groups | [optional] 

## Methods

### NewAddOptionFormRequestOptionTypeForm

`func NewAddOptionFormRequestOptionTypeForm() *AddOptionFormRequestOptionTypeForm`

NewAddOptionFormRequestOptionTypeForm instantiates a new AddOptionFormRequestOptionTypeForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOptionFormRequestOptionTypeFormWithDefaults

`func NewAddOptionFormRequestOptionTypeFormWithDefaults() *AddOptionFormRequestOptionTypeForm`

NewAddOptionFormRequestOptionTypeFormWithDefaults instantiates a new AddOptionFormRequestOptionTypeForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddOptionFormRequestOptionTypeForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOptionFormRequestOptionTypeForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOptionFormRequestOptionTypeForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddOptionFormRequestOptionTypeForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddOptionFormRequestOptionTypeForm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddOptionFormRequestOptionTypeForm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddOptionFormRequestOptionTypeForm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddOptionFormRequestOptionTypeForm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *AddOptionFormRequestOptionTypeForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOptionFormRequestOptionTypeForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOptionFormRequestOptionTypeForm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOptionFormRequestOptionTypeForm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *AddOptionFormRequestOptionTypeForm) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddOptionFormRequestOptionTypeForm) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddOptionFormRequestOptionTypeForm) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddOptionFormRequestOptionTypeForm) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOptions

`func (o *AddOptionFormRequestOptionTypeForm) GetOptions() []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AddOptionFormRequestOptionTypeForm) GetOptionsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AddOptionFormRequestOptionTypeForm) SetOptions(v []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AddOptionFormRequestOptionTypeForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFieldGroups

`func (o *AddOptionFormRequestOptionTypeForm) GetFieldGroups() []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *AddOptionFormRequestOptionTypeForm) GetFieldGroupsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *AddOptionFormRequestOptionTypeForm) SetFieldGroups(v []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *AddOptionFormRequestOptionTypeForm) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


