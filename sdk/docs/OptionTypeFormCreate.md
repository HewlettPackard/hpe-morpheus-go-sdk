# OptionTypeFormCreate

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

### NewOptionTypeFormCreate

`func NewOptionTypeFormCreate() *OptionTypeFormCreate`

NewOptionTypeFormCreate instantiates a new OptionTypeFormCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeFormCreateWithDefaults

`func NewOptionTypeFormCreateWithDefaults() *OptionTypeFormCreate`

NewOptionTypeFormCreateWithDefaults instantiates a new OptionTypeFormCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OptionTypeFormCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeFormCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeFormCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeFormCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeFormCreate) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeFormCreate) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeFormCreate) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeFormCreate) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeFormCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeFormCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeFormCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeFormCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *OptionTypeFormCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OptionTypeFormCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OptionTypeFormCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OptionTypeFormCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOptions

`func (o *OptionTypeFormCreate) GetOptions() []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OptionTypeFormCreate) GetOptionsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OptionTypeFormCreate) SetOptions(v []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OptionTypeFormCreate) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFieldGroups

`func (o *OptionTypeFormCreate) GetFieldGroups() []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *OptionTypeFormCreate) GetFieldGroupsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *OptionTypeFormCreate) SetFieldGroups(v []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *OptionTypeFormCreate) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


