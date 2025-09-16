# AddEnvironments200ResponseAllOfEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddEnvironments200ResponseAllOfEnvironment

`func NewAddEnvironments200ResponseAllOfEnvironment() *AddEnvironments200ResponseAllOfEnvironment`

NewAddEnvironments200ResponseAllOfEnvironment instantiates a new AddEnvironments200ResponseAllOfEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEnvironments200ResponseAllOfEnvironmentWithDefaults

`func NewAddEnvironments200ResponseAllOfEnvironmentWithDefaults() *AddEnvironments200ResponseAllOfEnvironment`

NewAddEnvironments200ResponseAllOfEnvironmentWithDefaults instantiates a new AddEnvironments200ResponseAllOfEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddEnvironments200ResponseAllOfEnvironment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddEnvironments200ResponseAllOfEnvironment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddEnvironments200ResponseAllOfEnvironment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


