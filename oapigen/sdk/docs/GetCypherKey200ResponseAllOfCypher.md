# GetCypherKey200ResponseAllOfCypher

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

### NewGetCypherKey200ResponseAllOfCypher

`func NewGetCypherKey200ResponseAllOfCypher() *GetCypherKey200ResponseAllOfCypher`

NewGetCypherKey200ResponseAllOfCypher instantiates a new GetCypherKey200ResponseAllOfCypher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetCypherKey200ResponseAllOfCypher) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCypherKey200ResponseAllOfCypher) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetCypherKey200ResponseAllOfCypher) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemKey

`func (o *GetCypherKey200ResponseAllOfCypher) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *GetCypherKey200ResponseAllOfCypher) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *GetCypherKey200ResponseAllOfCypher) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetLeaseTimeout

`func (o *GetCypherKey200ResponseAllOfCypher) GetLeaseTimeout() int64`

GetLeaseTimeout returns the LeaseTimeout field if non-nil, zero value otherwise.

### GetLeaseTimeoutOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetLeaseTimeoutOk() (*int64, bool)`

GetLeaseTimeoutOk returns a tuple with the LeaseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeout

`func (o *GetCypherKey200ResponseAllOfCypher) SetLeaseTimeout(v int64)`

SetLeaseTimeout sets LeaseTimeout field to given value.

### HasLeaseTimeout

`func (o *GetCypherKey200ResponseAllOfCypher) HasLeaseTimeout() bool`

HasLeaseTimeout returns a boolean if a field has been set.

### GetExpireDate

`func (o *GetCypherKey200ResponseAllOfCypher) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *GetCypherKey200ResponseAllOfCypher) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *GetCypherKey200ResponseAllOfCypher) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### SetExpireDateNil

`func (o *GetCypherKey200ResponseAllOfCypher) SetExpireDateNil(b bool)`

 SetExpireDateNil sets the value for ExpireDate to be an explicit nil

### UnsetExpireDate
`func (o *GetCypherKey200ResponseAllOfCypher) UnsetExpireDate()`

UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
### GetDateCreated

`func (o *GetCypherKey200ResponseAllOfCypher) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetCypherKey200ResponseAllOfCypher) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetCypherKey200ResponseAllOfCypher) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### SetDateCreatedNil

`func (o *GetCypherKey200ResponseAllOfCypher) SetDateCreatedNil(b bool)`

 SetDateCreatedNil sets the value for DateCreated to be an explicit nil

### UnsetDateCreated
`func (o *GetCypherKey200ResponseAllOfCypher) UnsetDateCreated()`

UnsetDateCreated ensures that no value is present for DateCreated, not even an explicit nil
### GetLastUpdated

`func (o *GetCypherKey200ResponseAllOfCypher) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetCypherKey200ResponseAllOfCypher) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetCypherKey200ResponseAllOfCypher) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastAccessed

`func (o *GetCypherKey200ResponseAllOfCypher) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *GetCypherKey200ResponseAllOfCypher) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *GetCypherKey200ResponseAllOfCypher) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetCypherKey200ResponseAllOfCypher) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetCypherKey200ResponseAllOfCypher) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetCypherKey200ResponseAllOfCypher) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetCypherKey200ResponseAllOfCypher) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


