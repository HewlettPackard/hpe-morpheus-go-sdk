# GetVDIAllocations200ResponseVdiAllocation

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

### NewGetVDIAllocations200ResponseVdiAllocation

`func NewGetVDIAllocations200ResponseVdiAllocation() *GetVDIAllocations200ResponseVdiAllocation`

NewGetVDIAllocations200ResponseVdiAllocation instantiates a new GetVDIAllocations200ResponseVdiAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVDIAllocations200ResponseVdiAllocationWithDefaults

`func NewGetVDIAllocations200ResponseVdiAllocationWithDefaults() *GetVDIAllocations200ResponseVdiAllocation`

NewGetVDIAllocations200ResponseVdiAllocationWithDefaults instantiates a new GetVDIAllocations200ResponseVdiAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetPool() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetPoolOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetPool(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetPool sets Pool field to given value.

### HasPool

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetInstance() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetInstanceOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetInstance(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetUser() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetUserOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetUser(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetUser sets User field to given value.

### HasUser

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetStatus

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *GetVDIAllocations200ResponseVdiAllocation) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
### GetLastReserved

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### SetLastReservedNil

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetLastReservedNil(b bool)`

 SetLastReservedNil sets the value for LastReserved to be an explicit nil

### UnsetLastReserved
`func (o *GetVDIAllocations200ResponseVdiAllocation) UnsetLastReserved()`

UnsetLastReserved ensures that no value is present for LastReserved, not even an explicit nil
### GetReleaseDate

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetVDIAllocations200ResponseVdiAllocation) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetVDIAllocations200ResponseVdiAllocation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *GetVDIAllocations200ResponseVdiAllocation) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *GetVDIAllocations200ResponseVdiAllocation) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


