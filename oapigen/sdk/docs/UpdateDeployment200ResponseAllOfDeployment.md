# UpdateDeployment200ResponseAllOfDeployment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**VersionCount** | Pointer to **int64** |  | [optional] 
**Versions** | Pointer to [**[]UpdateDeployment200ResponseAllOfDeploymentVersionsInner**](UpdateDeployment200ResponseAllOfDeploymentVersionsInner.md) |  | [optional] 

## Methods

### NewUpdateDeployment200ResponseAllOfDeployment

`func NewUpdateDeployment200ResponseAllOfDeployment() *UpdateDeployment200ResponseAllOfDeployment`

NewUpdateDeployment200ResponseAllOfDeployment instantiates a new UpdateDeployment200ResponseAllOfDeployment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *UpdateDeployment200ResponseAllOfDeployment) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetDateCreated

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVersionCount

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetVersionCount() int64`

GetVersionCount returns the VersionCount field if non-nil, zero value otherwise.

### GetVersionCountOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetVersionCountOk() (*int64, bool)`

GetVersionCountOk returns a tuple with the VersionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionCount

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetVersionCount(v int64)`

SetVersionCount sets VersionCount field to given value.

### HasVersionCount

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasVersionCount() bool`

HasVersionCount returns a boolean if a field has been set.

### GetVersions

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetVersions() []UpdateDeployment200ResponseAllOfDeploymentVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *UpdateDeployment200ResponseAllOfDeployment) GetVersionsOk() (*[]UpdateDeployment200ResponseAllOfDeploymentVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *UpdateDeployment200ResponseAllOfDeployment) SetVersions(v []UpdateDeployment200ResponseAllOfDeploymentVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *UpdateDeployment200ResponseAllOfDeployment) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


