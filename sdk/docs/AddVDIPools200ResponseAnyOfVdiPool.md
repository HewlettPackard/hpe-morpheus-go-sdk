# AddVDIPools200ResponseAnyOfVdiPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MinIdle** | Pointer to **int64** |  | [optional] 
**MaxIdle** | Pointer to **int64** |  | [optional] 
**InitialPoolSize** | Pointer to **int64** |  | [optional] 
**MaxPoolSize** | Pointer to **int64** |  | [optional] 
**AllocationTimeoutMinutes** | Pointer to **int64** |  | [optional] 
**PersistentUser** | Pointer to **NullableBool** |  | [optional] 
**Recyclable** | Pointer to **NullableBool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AutoCreateLocalUserOnReservation** | Pointer to **bool** |  | [optional] 
**AllowHypervisorConsole** | Pointer to **NullableBool** |  | [optional] 
**AllowCopy** | Pointer to **NullableBool** |  | [optional] 
**AllowPrinter** | Pointer to **NullableBool** |  | [optional] 
**AllowFileshare** | Pointer to **NullableBool** |  | [optional] 
**GuestConsoleJumpHost** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpPort** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpUsername** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpPassword** | Pointer to **NullableString** |  | [optional] 
**GuestConsoleJumpKeypair** | Pointer to **NullableString** |  | [optional] 
**Gateway** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerOwner**](ListVDIPools200ResponseAllOfVdiPoolsInnerOwner.md) |  | [optional] 
**Config** | Pointer to [**AddVDIPools200ResponseAnyOfVdiPoolConfig**](AddVDIPools200ResponseAnyOfVdiPoolConfig.md) |  | [optional] 
**Group** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Cloud** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**UsedCount** | Pointer to **int64** |  | [optional] 
**ReservedCount** | Pointer to **int64** |  | [optional] 
**PreparingCount** | Pointer to **int64** |  | [optional] 
**IdleCount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddVDIPools200ResponseAnyOfVdiPool

`func NewAddVDIPools200ResponseAnyOfVdiPool() *AddVDIPools200ResponseAnyOfVdiPool`

NewAddVDIPools200ResponseAnyOfVdiPool instantiates a new AddVDIPools200ResponseAnyOfVdiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIPools200ResponseAnyOfVdiPoolWithDefaults

`func NewAddVDIPools200ResponseAnyOfVdiPoolWithDefaults() *AddVDIPools200ResponseAnyOfVdiPool`

NewAddVDIPools200ResponseAnyOfVdiPoolWithDefaults instantiates a new AddVDIPools200ResponseAnyOfVdiPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMinIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMinIdle() int64`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMinIdleOk() (*int64, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetMinIdle(v int64)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetMaxIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMaxIdle() int64`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMaxIdleOk() (*int64, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetMaxIdle(v int64)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetInitialPoolSize() int64`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetInitialPoolSizeOk() (*int64, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetInitialPoolSize(v int64)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllocationTimeoutMinutes() int64`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllocationTimeoutMinutesOk() (*int64, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllocationTimeoutMinutes(v int64)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### SetPersistentUserNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetPersistentUserNil(b bool)`

 SetPersistentUserNil sets the value for PersistentUser to be an explicit nil

### UnsetPersistentUser
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetPersistentUser()`

UnsetPersistentUser ensures that no value is present for PersistentUser, not even an explicit nil
### GetRecyclable

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### SetRecyclableNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetRecyclableNil(b bool)`

 SetRecyclableNil sets the value for Recyclable to be an explicit nil

### UnsetRecyclable
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetRecyclable()`

UnsetRecyclable ensures that no value is present for Recyclable, not even an explicit nil
### GetEnabled

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### SetAllowHypervisorConsoleNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowHypervisorConsoleNil(b bool)`

 SetAllowHypervisorConsoleNil sets the value for AllowHypervisorConsole to be an explicit nil

### UnsetAllowHypervisorConsole
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetAllowHypervisorConsole()`

UnsetAllowHypervisorConsole ensures that no value is present for AllowHypervisorConsole, not even an explicit nil
### GetAllowCopy

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### SetAllowCopyNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowCopyNil(b bool)`

 SetAllowCopyNil sets the value for AllowCopy to be an explicit nil

### UnsetAllowCopy
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetAllowCopy()`

UnsetAllowCopy ensures that no value is present for AllowCopy, not even an explicit nil
### GetAllowPrinter

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### SetAllowPrinterNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowPrinterNil(b bool)`

 SetAllowPrinterNil sets the value for AllowPrinter to be an explicit nil

### UnsetAllowPrinter
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetAllowPrinter()`

UnsetAllowPrinter ensures that no value is present for AllowPrinter, not even an explicit nil
### GetAllowFileshare

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### SetAllowFileshareNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetAllowFileshareNil(b bool)`

 SetAllowFileshareNil sets the value for AllowFileshare to be an explicit nil

### UnsetAllowFileshare
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetAllowFileshare()`

UnsetAllowFileshare ensures that no value is present for AllowFileshare, not even an explicit nil
### GetGuestConsoleJumpHost

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### SetGuestConsoleJumpHostNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpHostNil(b bool)`

 SetGuestConsoleJumpHostNil sets the value for GuestConsoleJumpHost to be an explicit nil

### UnsetGuestConsoleJumpHost
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpHost()`

UnsetGuestConsoleJumpHost ensures that no value is present for GuestConsoleJumpHost, not even an explicit nil
### GetGuestConsoleJumpPort

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPort() string`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPortOk() (*string, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPort(v string)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### SetGuestConsoleJumpPortNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPortNil(b bool)`

 SetGuestConsoleJumpPortNil sets the value for GuestConsoleJumpPort to be an explicit nil

### UnsetGuestConsoleJumpPort
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpPort()`

UnsetGuestConsoleJumpPort ensures that no value is present for GuestConsoleJumpPort, not even an explicit nil
### GetGuestConsoleJumpUsername

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### SetGuestConsoleJumpUsernameNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpUsernameNil(b bool)`

 SetGuestConsoleJumpUsernameNil sets the value for GuestConsoleJumpUsername to be an explicit nil

### UnsetGuestConsoleJumpUsername
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpUsername()`

UnsetGuestConsoleJumpUsername ensures that no value is present for GuestConsoleJumpUsername, not even an explicit nil
### GetGuestConsoleJumpPassword

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### SetGuestConsoleJumpPasswordNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPasswordNil(b bool)`

 SetGuestConsoleJumpPasswordNil sets the value for GuestConsoleJumpPassword to be an explicit nil

### UnsetGuestConsoleJumpPassword
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpPassword()`

UnsetGuestConsoleJumpPassword ensures that no value is present for GuestConsoleJumpPassword, not even an explicit nil
### GetGuestConsoleJumpKeypair

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpKeypair() string`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpKeypairOk() (*string, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpKeypair(v string)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.

### SetGuestConsoleJumpKeypairNil

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpKeypairNil(b bool)`

 SetGuestConsoleJumpKeypairNil sets the value for GuestConsoleJumpKeypair to be an explicit nil

### UnsetGuestConsoleJumpKeypair
`func (o *AddVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpKeypair()`

UnsetGuestConsoleJumpKeypair ensures that no value is present for GuestConsoleJumpKeypair, not even an explicit nil
### GetGateway

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGateway() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGatewayOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGateway(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIconPath

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLogo

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetApps

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetApps() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetAppsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetApps(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetApps sets Apps field to given value.

### HasApps

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetOwner

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetOwner() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetOwnerOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetOwner(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetConfig() AddVDIPools200ResponseAnyOfVdiPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetConfigOk() (*AddVDIPools200ResponseAnyOfVdiPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetConfig(v AddVDIPools200ResponseAnyOfVdiPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetCloud() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetCloudOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetCloud(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUsedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetReservedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetReservedCount() int64`

GetReservedCount returns the ReservedCount field if non-nil, zero value otherwise.

### GetReservedCountOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetReservedCountOk() (*int64, bool)`

GetReservedCountOk returns a tuple with the ReservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetReservedCount(v int64)`

SetReservedCount sets ReservedCount field to given value.

### HasReservedCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasReservedCount() bool`

HasReservedCount returns a boolean if a field has been set.

### GetPreparingCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetPreparingCount() int64`

GetPreparingCount returns the PreparingCount field if non-nil, zero value otherwise.

### GetPreparingCountOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetPreparingCountOk() (*int64, bool)`

GetPreparingCountOk returns a tuple with the PreparingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetPreparingCount(v int64)`

SetPreparingCount sets PreparingCount field to given value.

### HasPreparingCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasPreparingCount() bool`

HasPreparingCount returns a boolean if a field has been set.

### GetIdleCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetIdleCount() int64`

GetIdleCount returns the IdleCount field if non-nil, zero value otherwise.

### GetIdleCountOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetIdleCountOk() (*int64, bool)`

GetIdleCountOk returns a tuple with the IdleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetIdleCount(v int64)`

SetIdleCount sets IdleCount field to given value.

### HasIdleCount

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasIdleCount() bool`

HasIdleCount returns a boolean if a field has been set.

### GetStatus

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddVDIPools200ResponseAnyOfVdiPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddVDIPools200ResponseAnyOfVdiPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


