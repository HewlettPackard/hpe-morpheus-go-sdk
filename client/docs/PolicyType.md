# PolicyType

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

### NewPolicyType

`func NewPolicyType() *PolicyType`

NewPolicyType instantiates a new PolicyType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyTypeWithDefaults

`func NewPolicyTypeWithDefaults() *PolicyType`

NewPolicyTypeWithDefaults instantiates a new PolicyType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *PolicyType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PolicyType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PolicyType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *PolicyType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *PolicyType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PolicyType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *PolicyType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PolicyType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PolicyType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PolicyType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *PolicyType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PolicyType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PolicyType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PolicyType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetLoadMethod

`func (o *PolicyType) GetLoadMethod() string`

GetLoadMethod returns the LoadMethod field if non-nil, zero value otherwise.

### GetLoadMethodOk

`func (o *PolicyType) GetLoadMethodOk() (*string, bool)`

GetLoadMethodOk returns a tuple with the LoadMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadMethod

`func (o *PolicyType) SetLoadMethod(v string)`

SetLoadMethod sets LoadMethod field to given value.

### HasLoadMethod

`func (o *PolicyType) HasLoadMethod() bool`

HasLoadMethod returns a boolean if a field has been set.

### GetEnforceMethod

`func (o *PolicyType) GetEnforceMethod() string`

GetEnforceMethod returns the EnforceMethod field if non-nil, zero value otherwise.

### GetEnforceMethodOk

`func (o *PolicyType) GetEnforceMethodOk() (*string, bool)`

GetEnforceMethodOk returns a tuple with the EnforceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceMethod

`func (o *PolicyType) SetEnforceMethod(v string)`

SetEnforceMethod sets EnforceMethod field to given value.

### HasEnforceMethod

`func (o *PolicyType) HasEnforceMethod() bool`

HasEnforceMethod returns a boolean if a field has been set.

### GetPrepareMethod

`func (o *PolicyType) GetPrepareMethod() string`

GetPrepareMethod returns the PrepareMethod field if non-nil, zero value otherwise.

### GetPrepareMethodOk

`func (o *PolicyType) GetPrepareMethodOk() (*string, bool)`

GetPrepareMethodOk returns a tuple with the PrepareMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrepareMethod

`func (o *PolicyType) SetPrepareMethod(v string)`

SetPrepareMethod sets PrepareMethod field to given value.

### HasPrepareMethod

`func (o *PolicyType) HasPrepareMethod() bool`

HasPrepareMethod returns a boolean if a field has been set.

### GetValidateMethod

`func (o *PolicyType) GetValidateMethod() string`

GetValidateMethod returns the ValidateMethod field if non-nil, zero value otherwise.

### GetValidateMethodOk

`func (o *PolicyType) GetValidateMethodOk() (*string, bool)`

GetValidateMethodOk returns a tuple with the ValidateMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateMethod

`func (o *PolicyType) SetValidateMethod(v string)`

SetValidateMethod sets ValidateMethod field to given value.

### HasValidateMethod

`func (o *PolicyType) HasValidateMethod() bool`

HasValidateMethod returns a boolean if a field has been set.

### GetEnforceOnProvision

`func (o *PolicyType) GetEnforceOnProvision() bool`

GetEnforceOnProvision returns the EnforceOnProvision field if non-nil, zero value otherwise.

### GetEnforceOnProvisionOk

`func (o *PolicyType) GetEnforceOnProvisionOk() (*bool, bool)`

GetEnforceOnProvisionOk returns a tuple with the EnforceOnProvision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceOnProvision

`func (o *PolicyType) SetEnforceOnProvision(v bool)`

SetEnforceOnProvision sets EnforceOnProvision field to given value.

### HasEnforceOnProvision

`func (o *PolicyType) HasEnforceOnProvision() bool`

HasEnforceOnProvision returns a boolean if a field has been set.

### GetEnforceOnManaged

`func (o *PolicyType) GetEnforceOnManaged() bool`

GetEnforceOnManaged returns the EnforceOnManaged field if non-nil, zero value otherwise.

### GetEnforceOnManagedOk

`func (o *PolicyType) GetEnforceOnManagedOk() (*bool, bool)`

GetEnforceOnManagedOk returns a tuple with the EnforceOnManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceOnManaged

`func (o *PolicyType) SetEnforceOnManaged(v bool)`

SetEnforceOnManaged sets EnforceOnManaged field to given value.

### HasEnforceOnManaged

`func (o *PolicyType) HasEnforceOnManaged() bool`

HasEnforceOnManaged returns a boolean if a field has been set.

### GetOptionTypes

`func (o *PolicyType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *PolicyType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *PolicyType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *PolicyType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


