# GetCatalogType200ResponseAllOfCatalogItemTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**AllowQuantity** | Pointer to **bool** | Can users order more than one of this item at a time. | [optional] 
**ImagePath** | Pointer to **string** |  | [optional] 
**DarkImagePath** | Pointer to **string** |  | [optional] 
**FormType** | Pointer to **string** |  | [optional] 
**Form** | Pointer to [**GetCatalogType200ResponseAllOfCatalogItemTypesInnerForm**](GetCatalogType200ResponseAllOfCatalogItemTypesInnerForm.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetCatalogType200ResponseAllOfCatalogItemTypesInner

`func NewGetCatalogType200ResponseAllOfCatalogItemTypesInner() *GetCatalogType200ResponseAllOfCatalogItemTypesInner`

NewGetCatalogType200ResponseAllOfCatalogItemTypesInner instantiates a new GetCatalogType200ResponseAllOfCatalogItemTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogType200ResponseAllOfCatalogItemTypesInnerWithDefaults

`func NewGetCatalogType200ResponseAllOfCatalogItemTypesInnerWithDefaults() *GetCatalogType200ResponseAllOfCatalogItemTypesInner`

NewGetCatalogType200ResponseAllOfCatalogItemTypesInnerWithDefaults instantiates a new GetCatalogType200ResponseAllOfCatalogItemTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetContext

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetFeatured

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetAllowQuantity

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetAllowQuantity() bool`

GetAllowQuantity returns the AllowQuantity field if non-nil, zero value otherwise.

### GetAllowQuantityOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetAllowQuantityOk() (*bool, bool)`

GetAllowQuantityOk returns a tuple with the AllowQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowQuantity

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetAllowQuantity(v bool)`

SetAllowQuantity sets AllowQuantity field to given value.

### HasAllowQuantity

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasAllowQuantity() bool`

HasAllowQuantity returns a boolean if a field has been set.

### GetImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.

### GetFormType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetFormType() string`

GetFormType returns the FormType field if non-nil, zero value otherwise.

### GetFormTypeOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetFormTypeOk() (*string, bool)`

GetFormTypeOk returns a tuple with the FormType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetFormType(v string)`

SetFormType sets FormType field to given value.

### HasFormType

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasFormType() bool`

HasFormType returns a boolean if a field has been set.

### GetForm

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetForm() GetCatalogType200ResponseAllOfCatalogItemTypesInnerForm`

GetForm returns the Form field if non-nil, zero value otherwise.

### GetFormOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetFormOk() (*GetCatalogType200ResponseAllOfCatalogItemTypesInnerForm, bool)`

GetFormOk returns a tuple with the Form field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForm

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetForm(v GetCatalogType200ResponseAllOfCatalogItemTypesInnerForm)`

SetForm sets Form field to given value.

### HasForm

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasForm() bool`

HasForm returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetCatalogType200ResponseAllOfCatalogItemTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


