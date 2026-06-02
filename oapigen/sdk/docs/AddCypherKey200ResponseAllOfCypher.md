# AddCypherKey200ResponseAllOfCypher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ItemKey** | Pointer to **string** |  | [optional] 
**LeaseTimeout** | Pointer to **int64** |  | [optional] 
**ExpireDate** | Pointer to **NullableTime** |  | [optional] 
**DateCreated** | Pointer to **NullableTime** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastAccessed** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewAddCypherKey200ResponseAllOfCypher

`func NewAddCypherKey200ResponseAllOfCypher() *AddCypherKey200ResponseAllOfCypher`

NewAddCypherKey200ResponseAllOfCypher instantiates a new AddCypherKey200ResponseAllOfCypher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddCypherKey200ResponseAllOfCypher) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCypherKey200ResponseAllOfCypher) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AddCypherKey200ResponseAllOfCypher) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemKey

`func (o *AddCypherKey200ResponseAllOfCypher) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *AddCypherKey200ResponseAllOfCypher) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *AddCypherKey200ResponseAllOfCypher) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetLeaseTimeout

`func (o *AddCypherKey200ResponseAllOfCypher) GetLeaseTimeout() int64`

GetLeaseTimeout returns the LeaseTimeout field if non-nil, zero value otherwise.

### GetLeaseTimeoutOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetLeaseTimeoutOk() (*int64, bool)`

GetLeaseTimeoutOk returns a tuple with the LeaseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeout

`func (o *AddCypherKey200ResponseAllOfCypher) SetLeaseTimeout(v int64)`

SetLeaseTimeout sets LeaseTimeout field to given value.

### HasLeaseTimeout

`func (o *AddCypherKey200ResponseAllOfCypher) HasLeaseTimeout() bool`

HasLeaseTimeout returns a boolean if a field has been set.

### GetExpireDate

`func (o *AddCypherKey200ResponseAllOfCypher) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *AddCypherKey200ResponseAllOfCypher) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *AddCypherKey200ResponseAllOfCypher) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### SetExpireDateNil

`func (o *AddCypherKey200ResponseAllOfCypher) SetExpireDateNil(b bool)`

 SetExpireDateNil sets the value for ExpireDate to be an explicit nil

### UnsetExpireDate
`func (o *AddCypherKey200ResponseAllOfCypher) UnsetExpireDate()`

UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
### GetDateCreated

`func (o *AddCypherKey200ResponseAllOfCypher) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCypherKey200ResponseAllOfCypher) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCypherKey200ResponseAllOfCypher) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *AddCypherKey200ResponseAllOfCypher) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *AddCypherKey200ResponseAllOfCypher) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetLastUpdated

`func (o *AddCypherKey200ResponseAllOfCypher) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCypherKey200ResponseAllOfCypher) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCypherKey200ResponseAllOfCypher) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastAccessed

`func (o *AddCypherKey200ResponseAllOfCypher) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *AddCypherKey200ResponseAllOfCypher) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *AddCypherKey200ResponseAllOfCypher) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddCypherKey200ResponseAllOfCypher) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddCypherKey200ResponseAllOfCypher) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddCypherKey200ResponseAllOfCypher) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddCypherKey200ResponseAllOfCypher) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


