# UpdateVDIPools200ResponseAnyOfVdiPool

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
**Gateway** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolGateway**](UpdateVDIPools200ResponseAnyOfVdiPoolGateway.md) |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]UpdateVDIPools200ResponseAnyOfVdiPoolAppsInner**](UpdateVDIPools200ResponseAnyOfVdiPoolAppsInner.md) |  | [optional] 
**Owner** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolOwner**](UpdateVDIPools200ResponseAnyOfVdiPoolOwner.md) |  | [optional] 
**Config** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolConfig**](UpdateVDIPools200ResponseAnyOfVdiPoolConfig.md) |  | [optional] 
**Group** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolGroup**](UpdateVDIPools200ResponseAnyOfVdiPoolGroup.md) |  | [optional] 
**Cloud** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPoolCloud**](UpdateVDIPools200ResponseAnyOfVdiPoolCloud.md) |  | [optional] 
**UsedCount** | Pointer to **int64** |  | [optional] 
**ReservedCount** | Pointer to **int64** |  | [optional] 
**PreparingCount** | Pointer to **int64** |  | [optional] 
**IdleCount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateVDIPools200ResponseAnyOfVdiPool

`func NewUpdateVDIPools200ResponseAnyOfVdiPool() *UpdateVDIPools200ResponseAnyOfVdiPool`

NewUpdateVDIPools200ResponseAnyOfVdiPool instantiates a new UpdateVDIPools200ResponseAnyOfVdiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMinIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMinIdle() int64`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMinIdleOk() (*int64, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetMinIdle(v int64)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetMaxIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMaxIdle() int64`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMaxIdleOk() (*int64, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetMaxIdle(v int64)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetInitialPoolSize() int64`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetInitialPoolSizeOk() (*int64, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetInitialPoolSize(v int64)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllocationTimeoutMinutes() int64`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllocationTimeoutMinutesOk() (*int64, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllocationTimeoutMinutes(v int64)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### SetPersistentUserNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetPersistentUserNil(b bool)`

 SetPersistentUserNil sets the value for PersistentUser to be an explicit nil

### UnsetPersistentUser
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetPersistentUser()`

UnsetPersistentUser ensures that no value is present for PersistentUser, not even an explicit nil
### GetRecyclable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### SetRecyclableNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetRecyclableNil(b bool)`

 SetRecyclableNil sets the value for Recyclable to be an explicit nil

### UnsetRecyclable
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetRecyclable()`

UnsetRecyclable ensures that no value is present for Recyclable, not even an explicit nil
### GetEnabled

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### SetAllowHypervisorConsoleNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowHypervisorConsoleNil(b bool)`

 SetAllowHypervisorConsoleNil sets the value for AllowHypervisorConsole to be an explicit nil

### UnsetAllowHypervisorConsole
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetAllowHypervisorConsole()`

UnsetAllowHypervisorConsole ensures that no value is present for AllowHypervisorConsole, not even an explicit nil
### GetAllowCopy

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### SetAllowCopyNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowCopyNil(b bool)`

 SetAllowCopyNil sets the value for AllowCopy to be an explicit nil

### UnsetAllowCopy
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetAllowCopy()`

UnsetAllowCopy ensures that no value is present for AllowCopy, not even an explicit nil
### GetAllowPrinter

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### SetAllowPrinterNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowPrinterNil(b bool)`

 SetAllowPrinterNil sets the value for AllowPrinter to be an explicit nil

### UnsetAllowPrinter
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetAllowPrinter()`

UnsetAllowPrinter ensures that no value is present for AllowPrinter, not even an explicit nil
### GetAllowFileshare

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### SetAllowFileshareNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetAllowFileshareNil(b bool)`

 SetAllowFileshareNil sets the value for AllowFileshare to be an explicit nil

### UnsetAllowFileshare
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetAllowFileshare()`

UnsetAllowFileshare ensures that no value is present for AllowFileshare, not even an explicit nil
### GetGuestConsoleJumpHost

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### SetGuestConsoleJumpHostNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpHostNil(b bool)`

 SetGuestConsoleJumpHostNil sets the value for GuestConsoleJumpHost to be an explicit nil

### UnsetGuestConsoleJumpHost
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpHost()`

UnsetGuestConsoleJumpHost ensures that no value is present for GuestConsoleJumpHost, not even an explicit nil
### GetGuestConsoleJumpPort

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPort() string`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPortOk() (*string, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPort(v string)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### SetGuestConsoleJumpPortNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPortNil(b bool)`

 SetGuestConsoleJumpPortNil sets the value for GuestConsoleJumpPort to be an explicit nil

### UnsetGuestConsoleJumpPort
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpPort()`

UnsetGuestConsoleJumpPort ensures that no value is present for GuestConsoleJumpPort, not even an explicit nil
### GetGuestConsoleJumpUsername

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### SetGuestConsoleJumpUsernameNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpUsernameNil(b bool)`

 SetGuestConsoleJumpUsernameNil sets the value for GuestConsoleJumpUsername to be an explicit nil

### UnsetGuestConsoleJumpUsername
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpUsername()`

UnsetGuestConsoleJumpUsername ensures that no value is present for GuestConsoleJumpUsername, not even an explicit nil
### GetGuestConsoleJumpPassword

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### SetGuestConsoleJumpPasswordNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpPasswordNil(b bool)`

 SetGuestConsoleJumpPasswordNil sets the value for GuestConsoleJumpPassword to be an explicit nil

### UnsetGuestConsoleJumpPassword
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpPassword()`

UnsetGuestConsoleJumpPassword ensures that no value is present for GuestConsoleJumpPassword, not even an explicit nil
### GetGuestConsoleJumpKeypair

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpKeypair() string`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGuestConsoleJumpKeypairOk() (*string, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpKeypair(v string)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.

### SetGuestConsoleJumpKeypairNil

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGuestConsoleJumpKeypairNil(b bool)`

 SetGuestConsoleJumpKeypairNil sets the value for GuestConsoleJumpKeypair to be an explicit nil

### UnsetGuestConsoleJumpKeypair
`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) UnsetGuestConsoleJumpKeypair()`

UnsetGuestConsoleJumpKeypair ensures that no value is present for GuestConsoleJumpKeypair, not even an explicit nil
### GetGateway

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGateway() UpdateVDIPools200ResponseAnyOfVdiPoolGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGatewayOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGateway(v UpdateVDIPools200ResponseAnyOfVdiPoolGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIconPath

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLogo

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetApps

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetApps() []UpdateVDIPools200ResponseAnyOfVdiPoolAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetAppsOk() (*[]UpdateVDIPools200ResponseAnyOfVdiPoolAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetApps(v []UpdateVDIPools200ResponseAnyOfVdiPoolAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetOwner() UpdateVDIPools200ResponseAnyOfVdiPoolOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetOwnerOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetOwner(v UpdateVDIPools200ResponseAnyOfVdiPoolOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetConfig() UpdateVDIPools200ResponseAnyOfVdiPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetConfigOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetConfig(v UpdateVDIPools200ResponseAnyOfVdiPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGroup() UpdateVDIPools200ResponseAnyOfVdiPoolGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetGroupOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetGroup(v UpdateVDIPools200ResponseAnyOfVdiPoolGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetCloud() UpdateVDIPools200ResponseAnyOfVdiPoolCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetCloudOk() (*UpdateVDIPools200ResponseAnyOfVdiPoolCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetCloud(v UpdateVDIPools200ResponseAnyOfVdiPoolCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUsedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetReservedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetReservedCount() int64`

GetReservedCount returns the ReservedCount field if non-nil, zero value otherwise.

### GetReservedCountOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetReservedCountOk() (*int64, bool)`

GetReservedCountOk returns a tuple with the ReservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetReservedCount(v int64)`

SetReservedCount sets ReservedCount field to given value.

### HasReservedCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasReservedCount() bool`

HasReservedCount returns a boolean if a field has been set.

### GetPreparingCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetPreparingCount() int64`

GetPreparingCount returns the PreparingCount field if non-nil, zero value otherwise.

### GetPreparingCountOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetPreparingCountOk() (*int64, bool)`

GetPreparingCountOk returns a tuple with the PreparingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetPreparingCount(v int64)`

SetPreparingCount sets PreparingCount field to given value.

### HasPreparingCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasPreparingCount() bool`

HasPreparingCount returns a boolean if a field has been set.

### GetIdleCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetIdleCount() int64`

GetIdleCount returns the IdleCount field if non-nil, zero value otherwise.

### GetIdleCountOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetIdleCountOk() (*int64, bool)`

GetIdleCountOk returns a tuple with the IdleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetIdleCount(v int64)`

SetIdleCount sets IdleCount field to given value.

### HasIdleCount

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasIdleCount() bool`

HasIdleCount returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateVDIPools200ResponseAnyOfVdiPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


