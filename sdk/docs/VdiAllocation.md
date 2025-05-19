# VdiAllocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Pool** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Instance** | Pointer to [**ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance**](ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance.md) |  | [optional] 
**User** | Pointer to [**ListVDIAllocations200ResponseAllOfVdiAllocationsInnerUser**](ListVDIAllocations200ResponseAllOfVdiAllocationsInnerUser.md) |  | [optional] 
**LocalUserCreated** | Pointer to **bool** |  | [optional] 
**Persistent** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastReserved** | Pointer to **time.Time** |  | [optional] 
**ReleaseDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVdiAllocation

`func NewVdiAllocation() *VdiAllocation`

NewVdiAllocation instantiates a new VdiAllocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVdiAllocationWithDefaults

`func NewVdiAllocationWithDefaults() *VdiAllocation`

NewVdiAllocationWithDefaults instantiates a new VdiAllocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VdiAllocation) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VdiAllocation) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VdiAllocation) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *VdiAllocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPool

`func (o *VdiAllocation) GetPool() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *VdiAllocation) GetPoolOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *VdiAllocation) SetPool(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetPool sets Pool field to given value.

### HasPool

`func (o *VdiAllocation) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetInstance

`func (o *VdiAllocation) GetInstance() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *VdiAllocation) GetInstanceOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *VdiAllocation) SetInstance(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *VdiAllocation) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetUser

`func (o *VdiAllocation) GetUser() ListVDIAllocations200ResponseAllOfVdiAllocationsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VdiAllocation) GetUserOk() (*ListVDIAllocations200ResponseAllOfVdiAllocationsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VdiAllocation) SetUser(v ListVDIAllocations200ResponseAllOfVdiAllocationsInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *VdiAllocation) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetLocalUserCreated

`func (o *VdiAllocation) GetLocalUserCreated() bool`

GetLocalUserCreated returns the LocalUserCreated field if non-nil, zero value otherwise.

### GetLocalUserCreatedOk

`func (o *VdiAllocation) GetLocalUserCreatedOk() (*bool, bool)`

GetLocalUserCreatedOk returns a tuple with the LocalUserCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserCreated

`func (o *VdiAllocation) SetLocalUserCreated(v bool)`

SetLocalUserCreated sets LocalUserCreated field to given value.

### HasLocalUserCreated

`func (o *VdiAllocation) HasLocalUserCreated() bool`

HasLocalUserCreated returns a boolean if a field has been set.

### GetPersistent

`func (o *VdiAllocation) GetPersistent() bool`

GetPersistent returns the Persistent field if non-nil, zero value otherwise.

### GetPersistentOk

`func (o *VdiAllocation) GetPersistentOk() (*bool, bool)`

GetPersistentOk returns a tuple with the Persistent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistent

`func (o *VdiAllocation) SetPersistent(v bool)`

SetPersistent sets Persistent field to given value.

### HasPersistent

`func (o *VdiAllocation) HasPersistent() bool`

HasPersistent returns a boolean if a field has been set.

### GetStatus

`func (o *VdiAllocation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VdiAllocation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VdiAllocation) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VdiAllocation) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *VdiAllocation) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *VdiAllocation) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *VdiAllocation) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *VdiAllocation) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *VdiAllocation) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *VdiAllocation) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *VdiAllocation) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *VdiAllocation) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastReserved

`func (o *VdiAllocation) GetLastReserved() time.Time`

GetLastReserved returns the LastReserved field if non-nil, zero value otherwise.

### GetLastReservedOk

`func (o *VdiAllocation) GetLastReservedOk() (*time.Time, bool)`

GetLastReservedOk returns a tuple with the LastReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReserved

`func (o *VdiAllocation) SetLastReserved(v time.Time)`

SetLastReserved sets LastReserved field to given value.

### HasLastReserved

`func (o *VdiAllocation) HasLastReserved() bool`

HasLastReserved returns a boolean if a field has been set.

### GetReleaseDate

`func (o *VdiAllocation) GetReleaseDate() time.Time`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *VdiAllocation) GetReleaseDateOk() (*time.Time, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *VdiAllocation) SetReleaseDate(v time.Time)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *VdiAllocation) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


