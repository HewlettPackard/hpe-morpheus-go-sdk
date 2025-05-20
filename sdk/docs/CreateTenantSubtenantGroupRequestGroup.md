# CreateTenantSubtenantGroupRequestGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to the subtenant for the group | 
**Description** | Pointer to **string** | Optional description field if you want to put more info there | [optional] 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Location** | Pointer to **string** | location | [optional] 

## Methods

### NewCreateTenantSubtenantGroupRequestGroup

`func NewCreateTenantSubtenantGroupRequestGroup(name string, ) *CreateTenantSubtenantGroupRequestGroup`

NewCreateTenantSubtenantGroupRequestGroup instantiates a new CreateTenantSubtenantGroupRequestGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTenantSubtenantGroupRequestGroupWithDefaults

`func NewCreateTenantSubtenantGroupRequestGroupWithDefaults() *CreateTenantSubtenantGroupRequestGroup`

NewCreateTenantSubtenantGroupRequestGroupWithDefaults instantiates a new CreateTenantSubtenantGroupRequestGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTenantSubtenantGroupRequestGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTenantSubtenantGroupRequestGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTenantSubtenantGroupRequestGroup) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateTenantSubtenantGroupRequestGroup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateTenantSubtenantGroupRequestGroup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateTenantSubtenantGroupRequestGroup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateTenantSubtenantGroupRequestGroup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *CreateTenantSubtenantGroupRequestGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CreateTenantSubtenantGroupRequestGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CreateTenantSubtenantGroupRequestGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CreateTenantSubtenantGroupRequestGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLocation

`func (o *CreateTenantSubtenantGroupRequestGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateTenantSubtenantGroupRequestGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateTenantSubtenantGroupRequestGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateTenantSubtenantGroupRequestGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


