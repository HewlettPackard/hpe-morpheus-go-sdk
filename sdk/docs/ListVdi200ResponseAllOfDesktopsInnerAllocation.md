# ListVdi200ResponseAllOfDesktopsInnerAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PoolId** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Instance** | Pointer to [**ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance**](ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerOwner**](ListVDIPools200ResponseAllOfVdiPoolsInnerOwner.md) |  | [optional] 
**LocalUserCreated** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Recyclable** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastReserved** | Pointer to **time.Time** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListVdi200ResponseAllOfDesktopsInnerAllocation

`func NewListVdi200ResponseAllOfDesktopsInnerAllocation() *ListVdi200ResponseAllOfDesktopsInnerAllocation`

NewListVdi200ResponseAllOfDesktopsInnerAllocation instantiates a new ListVdi200ResponseAllOfDesktopsInnerAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVdi200ResponseAllOfDesktopsInnerAllocationWithDefaults

`func NewListVdi200ResponseAllOfDesktopsInnerAllocationWithDefaults() *ListVdi200ResponseAllOfDesktopsInnerAllocation`

NewListVdi200ResponseAllOfDesktopsInnerAllocationWithDefaults instantiates a new ListVdi200ResponseAllOfDesktopsInnerAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPoolId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPoolId() int64`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPoolIdOk() (*int64, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetPoolId(v int64)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPool

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetPool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetPool sets Pool field to given value.

### HasPool

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetInstance() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetInstanceOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetInstance(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetUser() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetUserOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetUser(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetUser sets User field to given value.

### HasUser

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetRecyclable

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetStatus

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastReserved

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### GetReleaseDate

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *ListVdi200ResponseAllOfDesktopsInnerAllocation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


