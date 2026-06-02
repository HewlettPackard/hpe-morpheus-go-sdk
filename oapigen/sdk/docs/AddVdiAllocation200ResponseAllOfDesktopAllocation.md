# AddVdiAllocation200ResponseAllOfDesktopAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PoolId** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**AddVdiAllocation200ResponseAllOfDesktopAllocationPool**](AddVdiAllocation200ResponseAllOfDesktopAllocationPool.md) |  | [optional] 
**Instance** | Pointer to [**AddVdiAllocation200ResponseAllOfDesktopAllocationInstance**](AddVdiAllocation200ResponseAllOfDesktopAllocationInstance.md) |  | [optional] 
**User** | Pointer to [**AddVdiAllocation200ResponseAllOfDesktopAllocationUser**](AddVdiAllocation200ResponseAllOfDesktopAllocationUser.md) |  | [optional] 
**LocalUserCreated** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Recyclable** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastReserved** | Pointer to **NullableTime** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddVdiAllocation200ResponseAllOfDesktopAllocation

`func NewAddVdiAllocation200ResponseAllOfDesktopAllocation() *AddVdiAllocation200ResponseAllOfDesktopAllocation`

NewAddVdiAllocation200ResponseAllOfDesktopAllocation instantiates a new AddVdiAllocation200ResponseAllOfDesktopAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPool

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPool() AddVdiAllocation200ResponseAllOfDesktopAllocationPool`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPoolOk() (*AddVdiAllocation200ResponseAllOfDesktopAllocationPool, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetPool(v AddVdiAllocation200ResponseAllOfDesktopAllocationPool)`

SetPool sets Pool field to given value.

### HasPool

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetInstance() AddVdiAllocation200ResponseAllOfDesktopAllocationInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetInstanceOk() (*AddVdiAllocation200ResponseAllOfDesktopAllocationInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetInstance(v AddVdiAllocation200ResponseAllOfDesktopAllocationInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetUser() AddVdiAllocation200ResponseAllOfDesktopAllocationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetUserOk() (*AddVdiAllocation200ResponseAllOfDesktopAllocationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetUser(v AddVdiAllocation200ResponseAllOfDesktopAllocationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetRecyclable

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetStatus

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastReserved

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### SetLastReservedNil

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetLastReservedNil(b bool)`

 SetLastReservedNil sets the value for LastReserved to be an explicit nil

### UnsetLastReserved
`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) UnsetLastReserved()`

UnsetLastReserved ensures that no value is present for LastReserved, not even an explicit nil
### GetReleaseDate

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *AddVdiAllocation200ResponseAllOfDesktopAllocation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


