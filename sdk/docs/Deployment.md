# Deployment

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

### NewDeployment

`func NewDeployment() *Deployment`

NewDeployment instantiates a new Deployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentWithDefaults

`func NewDeploymentWithDefaults() *Deployment`

NewDeploymentWithDefaults instantiates a new Deployment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Deployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Deployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Deployment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Deployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Deployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Deployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Deployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Deployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Deployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Deployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Deployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Deployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *Deployment) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Deployment) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Deployment) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Deployment) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *Deployment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *Deployment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *Deployment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *Deployment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDateCreated

`func (o *Deployment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Deployment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Deployment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Deployment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Deployment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Deployment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Deployment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Deployment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVersionCount

`func (o *Deployment) GetVersionCount() int64`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *Deployment) GetVersionCountOk() (*int64, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *Deployment) SetVersionCount(v int64)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *Deployment) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### GetVersions

`func (o *Deployment) GetVersions() []GetDeployment200ResponseDeploymentVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *Deployment) GetVersionsOk() (*[]GetDeployment200ResponseDeploymentVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *Deployment) SetVersions(v []GetDeployment200ResponseDeploymentVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *Deployment) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


