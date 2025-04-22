# ListCypherKeys200ResponseAllOfCyphersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**ItemKey** | Pointer to **string** |  | [optional] 
**LeaseTimeout** | Pointer to **int64** |  | [optional] 
**ExpireDate** | Pointer to **time.Time** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastAccessed** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewListCypherKeys200ResponseAllOfCyphersInner

`func NewListCypherKeys200ResponseAllOfCyphersInner() *ListCypherKeys200ResponseAllOfCyphersInner`

NewListCypherKeys200ResponseAllOfCyphersInner instantiates a new ListCypherKeys200ResponseAllOfCyphersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCypherKeys200ResponseAllOfCyphersInnerWithDefaults

`func NewListCypherKeys200ResponseAllOfCyphersInnerWithDefaults() *ListCypherKeys200ResponseAllOfCyphersInner`

NewListCypherKeys200ResponseAllOfCyphersInnerWithDefaults instantiates a new ListCypherKeys200ResponseAllOfCyphersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetItemKey

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetItemKey() string`

GetItemKey returns the ItemKey field if non-nil, zero value otherwise.

### GetItemKeyOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetItemKeyOk() (*string, bool)`

GetItemKeyOk returns a tuple with the ItemKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemKey

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetItemKey(v string)`

SetItemKey sets ItemKey field to given value.

### HasItemKey

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasItemKey() bool`

HasItemKey returns a boolean if a field has been set.

### GetLeaseTimeout

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLeaseTimeout() int64`

GetLeaseTimeout returns the LeaseTimeout field if non-nil, zero value otherwise.

### GetLeaseTimeoutOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLeaseTimeoutOk() (*int64, bool)`

GetLeaseTimeoutOk returns a tuple with the LeaseTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTimeout

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetLeaseTimeout(v int64)`

SetLeaseTimeout sets LeaseTimeout field to given value.

### HasLeaseTimeout

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasLeaseTimeout() bool`

HasLeaseTimeout returns a boolean if a field has been set.

### GetExpireDate

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastAccessed

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLastAccessed() time.Time`

GetLastAccessed returns the LastAccessed field if non-nil, zero value otherwise.

### GetLastAccessedOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetLastAccessedOk() (*time.Time, bool)`

GetLastAccessedOk returns a tuple with the LastAccessed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessed

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetLastAccessed(v time.Time)`

SetLastAccessed sets LastAccessed field to given value.

### HasLastAccessed

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasLastAccessed() bool`

HasLastAccessed returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListCypherKeys200ResponseAllOfCyphersInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


