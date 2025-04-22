# ListPolicyTypes200ResponseAllOfPolicyTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**LoadMethod** | Pointer to **string** |  | [optional] 
**EnforceMethod** | Pointer to **string** |  | [optional] 
**PrepareMethod** | Pointer to **string** |  | [optional] 
**ValidateMethod** | Pointer to **string** |  | [optional] 
**EnforceOnProvision** | Pointer to **bool** |  | [optional] 
**EnforceOnManaged** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewListPolicyTypes200ResponseAllOfPolicyTypesInner

`func NewListPolicyTypes200ResponseAllOfPolicyTypesInner() *ListPolicyTypes200ResponseAllOfPolicyTypesInner`

NewListPolicyTypes200ResponseAllOfPolicyTypesInner instantiates a new ListPolicyTypes200ResponseAllOfPolicyTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPolicyTypes200ResponseAllOfPolicyTypesInnerWithDefaults

`func NewListPolicyTypes200ResponseAllOfPolicyTypesInnerWithDefaults() *ListPolicyTypes200ResponseAllOfPolicyTypesInner`

NewListPolicyTypes200ResponseAllOfPolicyTypesInnerWithDefaults instantiates a new ListPolicyTypes200ResponseAllOfPolicyTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLoadMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetLoadMethod() string`

GetLoadMethod returns the LoadMethod field if non-nil, zero value otherwise.

### GetLoadMethodOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetLoadMethodOk() (*string, bool)`

GetLoadMethodOk returns a tuple with the LoadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetLoadMethod(v string)`

SetLoadMethod sets LoadMethod field to given value.

### HasLoadMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasLoadMethod() bool`

HasLoadMethod returns a boolean if a field has been set.

### GetEnforceMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceMethod() string`

GetEnforceMethod returns the EnforceMethod field if non-nil, zero value otherwise.

### GetEnforceMethodOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceMethodOk() (*string, bool)`

GetEnforceMethodOk returns a tuple with the EnforceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetEnforceMethod(v string)`

SetEnforceMethod sets EnforceMethod field to given value.

### HasEnforceMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasEnforceMethod() bool`

HasEnforceMethod returns a boolean if a field has been set.

### GetPrepareMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetPrepareMethod() string`

GetPrepareMethod returns the PrepareMethod field if non-nil, zero value otherwise.

### GetPrepareMethodOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetPrepareMethodOk() (*string, bool)`

GetPrepareMethodOk returns a tuple with the PrepareMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetPrepareMethod(v string)`

SetPrepareMethod sets PrepareMethod field to given value.

### HasPrepareMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasPrepareMethod() bool`

HasPrepareMethod returns a boolean if a field has been set.

### GetValidateMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetValidateMethod() string`

GetValidateMethod returns the ValidateMethod field if non-nil, zero value otherwise.

### GetValidateMethodOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetValidateMethodOk() (*string, bool)`

GetValidateMethodOk returns a tuple with the ValidateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetValidateMethod(v string)`

SetValidateMethod sets ValidateMethod field to given value.

### HasValidateMethod

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasValidateMethod() bool`

HasValidateMethod returns a boolean if a field has been set.

### GetEnforceOnProvision

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceOnProvision() bool`

GetEnforceOnProvision returns the EnforceOnProvision field if non-nil, zero value otherwise.

### GetEnforceOnProvisionOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceOnProvisionOk() (*bool, bool)`

GetEnforceOnProvisionOk returns a tuple with the EnforceOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceOnProvision

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetEnforceOnProvision(v bool)`

SetEnforceOnProvision sets EnforceOnProvision field to given value.

### HasEnforceOnProvision

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasEnforceOnProvision() bool`

HasEnforceOnProvision returns a boolean if a field has been set.

### GetEnforceOnManaged

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceOnManaged() bool`

GetEnforceOnManaged returns the EnforceOnManaged field if non-nil, zero value otherwise.

### GetEnforceOnManagedOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetEnforceOnManagedOk() (*bool, bool)`

GetEnforceOnManagedOk returns a tuple with the EnforceOnManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceOnManaged

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetEnforceOnManaged(v bool)`

SetEnforceOnManaged sets EnforceOnManaged field to given value.

### HasEnforceOnManaged

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasEnforceOnManaged() bool`

HasEnforceOnManaged returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListPolicyTypes200ResponseAllOfPolicyTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


