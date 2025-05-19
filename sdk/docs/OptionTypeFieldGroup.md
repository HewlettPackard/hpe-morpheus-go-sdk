# OptionTypeFieldGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LocalizedName** | Pointer to **string** |  | [optional] 
**Collapsible** | Pointer to **bool** |  | [optional] 
**DefaultCollapsed** | Pointer to **bool** |  | [optional] 
**VisibleOnCode** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner**](ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner.md) |  | [optional] 

## Methods

### NewOptionTypeFieldGroup

`func NewOptionTypeFieldGroup() *OptionTypeFieldGroup`

NewOptionTypeFieldGroup instantiates a new OptionTypeFieldGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeFieldGroupWithDefaults

`func NewOptionTypeFieldGroupWithDefaults() *OptionTypeFieldGroup`

NewOptionTypeFieldGroupWithDefaults instantiates a new OptionTypeFieldGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionTypeFieldGroup) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeFieldGroup) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeFieldGroup) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeFieldGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeFieldGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeFieldGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeFieldGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeFieldGroup) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeFieldGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeFieldGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeFieldGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeFieldGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeFieldGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeFieldGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeFieldGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeFieldGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLocalizedName

`func (o *OptionTypeFieldGroup) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *OptionTypeFieldGroup) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *OptionTypeFieldGroup) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *OptionTypeFieldGroup) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### GetCollapsible

`func (o *OptionTypeFieldGroup) GetCollapsible() bool`

GetCollapsible returns the Collapsible field if non-nil, zero value otherwise.

### GetCollapsibleOk

`func (o *OptionTypeFieldGroup) GetCollapsibleOk() (*bool, bool)`

GetCollapsibleOk returns a tuple with the Collapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsible

`func (o *OptionTypeFieldGroup) SetCollapsible(v bool)`

SetCollapsible sets Collapsible field to given value.

### HasCollapsible

`func (o *OptionTypeFieldGroup) HasCollapsible() bool`

HasCollapsible returns a boolean if a field has been set.

### GetDefaultCollapsed

`func (o *OptionTypeFieldGroup) GetDefaultCollapsed() bool`

GetDefaultCollapsed returns the DefaultCollapsed field if non-nil, zero value otherwise.

### GetDefaultCollapsedOk

`func (o *OptionTypeFieldGroup) GetDefaultCollapsedOk() (*bool, bool)`

GetDefaultCollapsedOk returns a tuple with the DefaultCollapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollapsed

`func (o *OptionTypeFieldGroup) SetDefaultCollapsed(v bool)`

SetDefaultCollapsed sets DefaultCollapsed field to given value.

### HasDefaultCollapsed

`func (o *OptionTypeFieldGroup) HasDefaultCollapsed() bool`

HasDefaultCollapsed returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *OptionTypeFieldGroup) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *OptionTypeFieldGroup) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *OptionTypeFieldGroup) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *OptionTypeFieldGroup) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### GetOptions

`func (o *OptionTypeFieldGroup) GetOptions() []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OptionTypeFieldGroup) GetOptionsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OptionTypeFieldGroup) SetOptions(v []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OptionTypeFieldGroup) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


