# GetVdi200ResponseDesktopAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PoolId** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Instance** | Pointer to [**ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance**](ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerOwner**](ListVDIPools200ResponseAllOfVdiPoolsInnerOwner.md) |  | [optional] 
**LocalUserCreated** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Recyclable** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastReserved** | Pointer to **NullableTime** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetVdi200ResponseDesktopAllocation

`func NewGetVdi200ResponseDesktopAllocation() *GetVdi200ResponseDesktopAllocation`

NewGetVdi200ResponseDesktopAllocation instantiates a new GetVdi200ResponseDesktopAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVdi200ResponseDesktopAllocationWithDefaults

`func NewGetVdi200ResponseDesktopAllocationWithDefaults() *GetVdi200ResponseDesktopAllocation`

NewGetVdi200ResponseDesktopAllocationWithDefaults instantiates a new GetVdi200ResponseDesktopAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVdi200ResponseDesktopAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVdi200ResponseDesktopAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVdi200ResponseDesktopAllocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVdi200ResponseDesktopAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolId

`func (o *GetVdi200ResponseDesktopAllocation) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *GetVdi200ResponseDesktopAllocation) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *GetVdi200ResponseDesktopAllocation) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *GetVdi200ResponseDesktopAllocation) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPool

`func (o *GetVdi200ResponseDesktopAllocation) GetPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetVdi200ResponseDesktopAllocation) GetPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetVdi200ResponseDesktopAllocation) SetPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetVdi200ResponseDesktopAllocation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *GetVdi200ResponseDesktopAllocation) GetInstance() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetVdi200ResponseDesktopAllocation) GetInstanceOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetVdi200ResponseDesktopAllocation) SetInstance(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetVdi200ResponseDesktopAllocation) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *GetVdi200ResponseDesktopAllocation) GetUser() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetVdi200ResponseDesktopAllocation) GetUserOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetVdi200ResponseDesktopAllocation) SetUser(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetUser sets User field to given value.

### HasUser

`func (o *GetVdi200ResponseDesktopAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *GetVdi200ResponseDesktopAllocation) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *GetVdi200ResponseDesktopAllocation) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *GetVdi200ResponseDesktopAllocation) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *GetVdi200ResponseDesktopAllocation) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *GetVdi200ResponseDesktopAllocation) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *GetVdi200ResponseDesktopAllocation) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *GetVdi200ResponseDesktopAllocation) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *GetVdi200ResponseDesktopAllocation) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetRecyclable

`func (o *GetVdi200ResponseDesktopAllocation) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *GetVdi200ResponseDesktopAllocation) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *GetVdi200ResponseDesktopAllocation) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *GetVdi200ResponseDesktopAllocation) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetStatus

`func (o *GetVdi200ResponseDesktopAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVdi200ResponseDesktopAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVdi200ResponseDesktopAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVdi200ResponseDesktopAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetVdi200ResponseDesktopAllocation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetVdi200ResponseDesktopAllocation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetVdi200ResponseDesktopAllocation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetVdi200ResponseDesktopAllocation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetVdi200ResponseDesktopAllocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetVdi200ResponseDesktopAllocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetVdi200ResponseDesktopAllocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetVdi200ResponseDesktopAllocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastReserved

`func (o *GetVdi200ResponseDesktopAllocation) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *GetVdi200ResponseDesktopAllocation) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *GetVdi200ResponseDesktopAllocation) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *GetVdi200ResponseDesktopAllocation) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### SetLastReservedNil

`func (o *GetVdi200ResponseDesktopAllocation) SetLastReservedNil(b bool)`

 SetLastReservedNil sets the value for LastReserved to be an explicit nil

### UnsetLastReserved
`func (o *GetVdi200ResponseDesktopAllocation) UnsetLastReserved()`

UnsetLastReserved ensures that no value is present for LastReserved, not even an explicit nil
### GetReleaseDate

`func (o *GetVdi200ResponseDesktopAllocation) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetVdi200ResponseDesktopAllocation) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetVdi200ResponseDesktopAllocation) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetVdi200ResponseDesktopAllocation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


