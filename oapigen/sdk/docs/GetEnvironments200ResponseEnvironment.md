# GetEnvironments200ResponseEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetEnvironments200ResponseEnvironmentAccount**](GetEnvironments200ResponseEnvironmentAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetEnvironments200ResponseEnvironment

`func NewGetEnvironments200ResponseEnvironment() *GetEnvironments200ResponseEnvironment`

NewGetEnvironments200ResponseEnvironment instantiates a new GetEnvironments200ResponseEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetEnvironments200ResponseEnvironmentWithDefaults

`func NewGetEnvironments200ResponseEnvironmentWithDefaults() *GetEnvironments200ResponseEnvironment`

NewGetEnvironments200ResponseEnvironmentWithDefaults instantiates a new GetEnvironments200ResponseEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetEnvironments200ResponseEnvironment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetEnvironments200ResponseEnvironment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetEnvironments200ResponseEnvironment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetEnvironments200ResponseEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetEnvironments200ResponseEnvironment) GetAccount() GetEnvironments200ResponseEnvironmentAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetEnvironments200ResponseEnvironment) GetAccountOk() (*GetEnvironments200ResponseEnvironmentAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetEnvironments200ResponseEnvironment) SetAccount(v GetEnvironments200ResponseEnvironmentAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetEnvironments200ResponseEnvironment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *GetEnvironments200ResponseEnvironment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetEnvironments200ResponseEnvironment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetEnvironments200ResponseEnvironment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetEnvironments200ResponseEnvironment) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetEnvironments200ResponseEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetEnvironments200ResponseEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetEnvironments200ResponseEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetEnvironments200ResponseEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetEnvironments200ResponseEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetEnvironments200ResponseEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetEnvironments200ResponseEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetEnvironments200ResponseEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *GetEnvironments200ResponseEnvironment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetEnvironments200ResponseEnvironment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetEnvironments200ResponseEnvironment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetEnvironments200ResponseEnvironment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *GetEnvironments200ResponseEnvironment) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetEnvironments200ResponseEnvironment) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetEnvironments200ResponseEnvironment) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetEnvironments200ResponseEnvironment) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *GetEnvironments200ResponseEnvironment) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetEnvironments200ResponseEnvironment) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetEnvironments200ResponseEnvironment) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetEnvironments200ResponseEnvironment) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetEnvironments200ResponseEnvironment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetEnvironments200ResponseEnvironment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetEnvironments200ResponseEnvironment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetEnvironments200ResponseEnvironment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetEnvironments200ResponseEnvironment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetEnvironments200ResponseEnvironment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetEnvironments200ResponseEnvironment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetEnvironments200ResponseEnvironment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


