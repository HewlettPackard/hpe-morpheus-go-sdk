# ListEnvironments200ResponseAllOfEnvironmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListEnvironments200ResponseAllOfEnvironmentsInner

`func NewListEnvironments200ResponseAllOfEnvironmentsInner() *ListEnvironments200ResponseAllOfEnvironmentsInner`

NewListEnvironments200ResponseAllOfEnvironmentsInner instantiates a new ListEnvironments200ResponseAllOfEnvironmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListEnvironments200ResponseAllOfEnvironmentsInnerWithDefaults

`func NewListEnvironments200ResponseAllOfEnvironmentsInnerWithDefaults() *ListEnvironments200ResponseAllOfEnvironmentsInner`

NewListEnvironments200ResponseAllOfEnvironmentsInnerWithDefaults instantiates a new ListEnvironments200ResponseAllOfEnvironmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetActive

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListEnvironments200ResponseAllOfEnvironmentsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


