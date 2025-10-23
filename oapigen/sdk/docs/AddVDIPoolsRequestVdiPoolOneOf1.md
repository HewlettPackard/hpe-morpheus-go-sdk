# AddVDIPoolsRequestVdiPoolOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Virtual Desktop name | 
**Description** | Pointer to **string** | Virtual Desktop description | [optional] 
**Owner** | Pointer to **int64** | Owner (User) ID | [optional] 
**MinIdle** | Pointer to **float32** | Min Idle - Sets the minimum number of idle instances on standby in the pool. The pool will always try to maintain this number of available instances on standby.  | [optional] 
**InitialPoolSize** | Pointer to **float32** | The initial size of the pool to be allocated on creation | [optional] 
**MaxIdle** | Pointer to **float32** | Sets the maximum number of idle instances on standby in the pool. If the number of idle instances supersedes this, the pool will start removing instances  | [optional] 
**MaxPoolSize** | **float32** | Max limit on number of allocations and instances within the pool.  | 
**AllocationTimeoutMinutes** | Pointer to **float32** | Time (in minutes) after a user disconnects before an allocation is recycled or shutdown depending on persistence.  | [optional] 
**PersistentUser** | Pointer to **bool** | Persistent Desktop Pool | [optional] [default to false]
**Recyclable** | Pointer to **bool** | Recyclable VDI Pools only work with cloud types that support snapshot management (i.e. Vmware, Nutanix, VCD) | [optional] [default to false]
**AllowCopy** | Pointer to **bool** | Allow copy from desktop | [optional] [default to false]
**AllowPrinter** | Pointer to **bool** | Allow local printers from Desktop | [optional] [default to false]
**AllowFileshare** | Pointer to **bool** | Allow File Share | [optional] [default to false]
**AllowHypervisorConsole** | Pointer to **bool** | Allow hypervisor console | [optional] [default to false]
**AutoCreateLocalUserOnReservation** | Pointer to **bool** | Auto create local user upon reservation | [optional] [default to false]
**Enabled** | Pointer to **bool** | Can be used to enable or disable the VDI pool | [optional] [default to true]
**IconPath** | Pointer to **string** | The relative location of an icon image | [optional] 
**Apps** | Pointer to **[]int64** | Array of VDI App IDs | [optional] 
**Gateway** | Pointer to **int64** | VDI Gateway ID | [optional] 
**InstanceConfig** | Pointer to **string** | Instance Config JSON. Passing as a string will preserve property order.  See &#x60;config&#x60; object for required values. | [optional] 
**Config** | [**AddVDIPoolsRequestVdiPoolOneOfConfig**](AddVDIPoolsRequestVdiPoolOneOfConfig.md) |  | 
**GuestConsoleJumpHost** | Pointer to **string** | Guest Console Jump Host | [optional] 
**GuestConsoleJumpPort** | Pointer to **int64** | Guest Console Jump Port | [optional] 
**GuestConsoleJumpUsername** | Pointer to **string** | Guest Console Jump Username | [optional] 
**GuestConsoleJumpPassword** | Pointer to **string** | Guest Console Jump Password | [optional] 
**GuestConsoleJumpKeypair** | Pointer to **int64** | Guest Console Jump Key Pair. see &#x60;Key Pair&#x60; | [optional] 

## Methods

### NewAddVDIPoolsRequestVdiPoolOneOf1

`func NewAddVDIPoolsRequestVdiPoolOneOf1(name string, maxPoolSize float32, config AddVDIPoolsRequestVdiPoolOneOfConfig, ) *AddVDIPoolsRequestVdiPoolOneOf1`

NewAddVDIPoolsRequestVdiPoolOneOf1 instantiates a new AddVDIPoolsRequestVdiPoolOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIPoolsRequestVdiPoolOneOf1WithDefaults

`func NewAddVDIPoolsRequestVdiPoolOneOf1WithDefaults() *AddVDIPoolsRequestVdiPoolOneOf1`

NewAddVDIPoolsRequestVdiPoolOneOf1WithDefaults instantiates a new AddVDIPoolsRequestVdiPoolOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOwner

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetOwner() int64`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetOwnerOk() (*int64, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetOwner(v int64)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetMinIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMinIdle() float32`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMinIdleOk() (*float32, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetMinIdle(v float32)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetInitialPoolSize() float32`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetInitialPoolSizeOk() (*float32, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetInitialPoolSize(v float32)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMaxIdle() float32`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMaxIdleOk() (*float32, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetMaxIdle(v float32)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMaxPoolSize() float32`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetMaxPoolSizeOk() (*float32, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetMaxPoolSize(v float32)`

SetMaxPoolSize sets MaxPoolSize field to given value.


### GetAllocationTimeoutMinutes

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllocationTimeoutMinutes() float32`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllocationTimeoutMinutesOk() (*float32, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAllocationTimeoutMinutes(v float32)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### GetRecyclable

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetAllowCopy

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### GetAllowPrinter

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### GetAllowFileshare

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetEnabled

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconPath

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetApps

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetApps() []int64`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetAppsOk() (*[]int64, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetApps(v []int64)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetGateway

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGateway() int64`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGatewayOk() (*int64, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGateway(v int64)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInstanceConfig

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetInstanceConfig() string`

GetInstanceConfig returns the InstanceConfig field if non-nil, zero value otherwise.

### GetInstanceConfigOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetInstanceConfigOk() (*string, bool)`

GetInstanceConfigOk returns a tuple with the InstanceConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceConfig

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetInstanceConfig(v string)`

SetInstanceConfig sets InstanceConfig field to given value.

### HasInstanceConfig

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasInstanceConfig() bool`

HasInstanceConfig returns a boolean if a field has been set.

### GetConfig

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetConfig() AddVDIPoolsRequestVdiPoolOneOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetConfigOk() (*AddVDIPoolsRequestVdiPoolOneOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetConfig(v AddVDIPoolsRequestVdiPoolOneOfConfig)`

SetConfig sets Config field to given value.


### GetGuestConsoleJumpHost

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### GetGuestConsoleJumpPort

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpPort() int64`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpPortOk() (*int64, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGuestConsoleJumpPort(v int64)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### GetGuestConsoleJumpUsername

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### GetGuestConsoleJumpPassword

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### GetGuestConsoleJumpKeypair

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpKeypair() int64`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) GetGuestConsoleJumpKeypairOk() (*int64, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) SetGuestConsoleJumpKeypair(v int64)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *AddVDIPoolsRequestVdiPoolOneOf1) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


