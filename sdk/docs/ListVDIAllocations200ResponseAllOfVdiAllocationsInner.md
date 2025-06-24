# ListVDIAllocations200ResponseAllOfVdiAllocationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Instance** | Pointer to [**ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance**](ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerOwner**](ListVDIPools200ResponseAllOfVdiPoolsInnerOwner.md) |  | [optional] 
**LocalUserCreated** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **NullableTime** |  | [optional] 
**LastReserved** | Pointer to **NullableTime** |  | [optional] 
**ReleaseDate** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListVDIAllocations200ResponseAllOfVdiAllocationsInner

`func NewListVDIAllocations200ResponseAllOfVdiAllocationsInner() *ListVDIAllocations200ResponseAllOfVdiAllocationsInner`

NewListVDIAllocations200ResponseAllOfVdiAllocationsInner instantiates a new ListVDIAllocations200ResponseAllOfVdiAllocationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIAllocations200ResponseAllOfVdiAllocationsInnerWithDefaults

`func NewListVDIAllocations200ResponseAllOfVdiAllocationsInnerWithDefaults() *ListVDIAllocations200ResponseAllOfVdiAllocationsInner`

NewListVDIAllocations200ResponseAllOfVdiAllocationsInnerWithDefaults instantiates a new ListVDIAllocations200ResponseAllOfVdiAllocationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetInstance() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetInstanceOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetInstance(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetUser() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetUserOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetUser(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetUser sets User field to given value.

### HasUser

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetStatus

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastReserved

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### SetLastReservedNil

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetLastReservedNil(b bool)`

 SetLastReservedNil sets the value for LastReserved to be an explicit nil

### UnsetLastReserved
`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) UnsetLastReserved()`

UnsetLastReserved ensures that no value is present for LastReserved, not even an explicit nil
### GetReleaseDate

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *ListVDIAllocations200ResponseAllOfVdiAllocationsInner) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


