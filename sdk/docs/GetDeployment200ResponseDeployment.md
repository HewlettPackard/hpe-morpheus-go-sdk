# GetDeployment200ResponseDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**VersionCount** | Pointer to **int64** |  | [optional] 
**Versions** | Pointer to [**[]GetDeployment200ResponseDeploymentVersionsInner**](GetDeployment200ResponseDeploymentVersionsInner.md) |  | [optional] 

## Methods

### NewGetDeployment200ResponseDeployment

`func NewGetDeployment200ResponseDeployment() *GetDeployment200ResponseDeployment`

NewGetDeployment200ResponseDeployment instantiates a new GetDeployment200ResponseDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeployment200ResponseDeploymentWithDefaults

`func NewGetDeployment200ResponseDeploymentWithDefaults() *GetDeployment200ResponseDeployment`

NewGetDeployment200ResponseDeploymentWithDefaults instantiates a new GetDeployment200ResponseDeployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeployment200ResponseDeployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeployment200ResponseDeployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeployment200ResponseDeployment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeployment200ResponseDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetDeployment200ResponseDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDeployment200ResponseDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDeployment200ResponseDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDeployment200ResponseDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetDeployment200ResponseDeployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetDeployment200ResponseDeployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetDeployment200ResponseDeployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetDeployment200ResponseDeployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *GetDeployment200ResponseDeployment) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetDeployment200ResponseDeployment) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetDeployment200ResponseDeployment) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetDeployment200ResponseDeployment) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *GetDeployment200ResponseDeployment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetDeployment200ResponseDeployment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetDeployment200ResponseDeployment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetDeployment200ResponseDeployment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetDeployment200ResponseDeployment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetDeployment200ResponseDeployment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetDeployment200ResponseDeployment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetDeployment200ResponseDeployment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetDeployment200ResponseDeployment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetDeployment200ResponseDeployment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetDeployment200ResponseDeployment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetDeployment200ResponseDeployment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVersionCount

`func (o *GetDeployment200ResponseDeployment) GetVersionCount() int64`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *GetDeployment200ResponseDeployment) GetVersionCountOk() (*int64, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *GetDeployment200ResponseDeployment) SetVersionCount(v int64)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *GetDeployment200ResponseDeployment) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### GetVersions

`func (o *GetDeployment200ResponseDeployment) GetVersions() []GetDeployment200ResponseDeploymentVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetDeployment200ResponseDeployment) GetVersionsOk() (*[]GetDeployment200ResponseDeploymentVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetDeployment200ResponseDeployment) SetVersions(v []GetDeployment200ResponseDeploymentVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetDeployment200ResponseDeployment) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


