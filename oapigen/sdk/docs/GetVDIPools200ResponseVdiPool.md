# GetVDIPools200ResponseVdiPool

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
**Gateway** | Pointer to [**GetVDIPools200ResponseVdiPoolGateway**](GetVDIPools200ResponseVdiPoolGateway.md) |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]GetVDIPools200ResponseVdiPoolAppsInner**](GetVDIPools200ResponseVdiPoolAppsInner.md) |  | [optional] 
**Owner** | Pointer to [**GetVDIPools200ResponseVdiPoolOwner**](GetVDIPools200ResponseVdiPoolOwner.md) |  | [optional] 
**Config** | Pointer to [**GetVDIPools200ResponseVdiPoolConfig**](GetVDIPools200ResponseVdiPoolConfig.md) |  | [optional] 
**Group** | Pointer to [**GetVDIPools200ResponseVdiPoolGroup**](GetVDIPools200ResponseVdiPoolGroup.md) |  | [optional] 
**Cloud** | Pointer to [**GetVDIPools200ResponseVdiPoolCloud**](GetVDIPools200ResponseVdiPoolCloud.md) |  | [optional] 
**UsedCount** | Pointer to **int64** |  | [optional] 
**ReservedCount** | Pointer to **int64** |  | [optional] 
**PreparingCount** | Pointer to **int64** |  | [optional] 
**IdleCount** | Pointer to **int64** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetVDIPools200ResponseVdiPool

`func NewGetVDIPools200ResponseVdiPool() *GetVDIPools200ResponseVdiPool`

NewGetVDIPools200ResponseVdiPool instantiates a new GetVDIPools200ResponseVdiPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetVDIPools200ResponseVdiPool) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVDIPools200ResponseVdiPool) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVDIPools200ResponseVdiPool) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetVDIPools200ResponseVdiPool) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetVDIPools200ResponseVdiPool) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetVDIPools200ResponseVdiPool) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetVDIPools200ResponseVdiPool) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetVDIPools200ResponseVdiPool) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetVDIPools200ResponseVdiPool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVDIPools200ResponseVdiPool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVDIPools200ResponseVdiPool) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVDIPools200ResponseVdiPool) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetVDIPools200ResponseVdiPool) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetVDIPools200ResponseVdiPool) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMinIdle

`func (o *GetVDIPools200ResponseVdiPool) GetMinIdle() int64`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *GetVDIPools200ResponseVdiPool) GetMinIdleOk() (*int64, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *GetVDIPools200ResponseVdiPool) SetMinIdle(v int64)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *GetVDIPools200ResponseVdiPool) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetMaxIdle

`func (o *GetVDIPools200ResponseVdiPool) GetMaxIdle() int64`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *GetVDIPools200ResponseVdiPool) GetMaxIdleOk() (*int64, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *GetVDIPools200ResponseVdiPool) SetMaxIdle(v int64)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *GetVDIPools200ResponseVdiPool) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *GetVDIPools200ResponseVdiPool) GetInitialPoolSize() int64`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *GetVDIPools200ResponseVdiPool) GetInitialPoolSizeOk() (*int64, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *GetVDIPools200ResponseVdiPool) SetInitialPoolSize(v int64)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *GetVDIPools200ResponseVdiPool) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *GetVDIPools200ResponseVdiPool) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *GetVDIPools200ResponseVdiPool) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *GetVDIPools200ResponseVdiPool) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *GetVDIPools200ResponseVdiPool) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *GetVDIPools200ResponseVdiPool) GetAllocationTimeoutMinutes() int64`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *GetVDIPools200ResponseVdiPool) GetAllocationTimeoutMinutesOk() (*int64, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *GetVDIPools200ResponseVdiPool) SetAllocationTimeoutMinutes(v int64)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *GetVDIPools200ResponseVdiPool) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *GetVDIPools200ResponseVdiPool) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *GetVDIPools200ResponseVdiPool) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *GetVDIPools200ResponseVdiPool) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *GetVDIPools200ResponseVdiPool) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### SetPersistentUserNil

`func (o *GetVDIPools200ResponseVdiPool) SetPersistentUserNil(b bool)`

 SetPersistentUserNil sets the value for PersistentUser to be an explicit nil

### UnsetPersistentUser
`func (o *GetVDIPools200ResponseVdiPool) UnsetPersistentUser()`

UnsetPersistentUser ensures that no value is present for PersistentUser, not even an explicit nil
### GetRecyclable

`func (o *GetVDIPools200ResponseVdiPool) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *GetVDIPools200ResponseVdiPool) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *GetVDIPools200ResponseVdiPool) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *GetVDIPools200ResponseVdiPool) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### SetRecyclableNil

`func (o *GetVDIPools200ResponseVdiPool) SetRecyclableNil(b bool)`

 SetRecyclableNil sets the value for Recyclable to be an explicit nil

### UnsetRecyclable
`func (o *GetVDIPools200ResponseVdiPool) UnsetRecyclable()`

UnsetRecyclable ensures that no value is present for Recyclable, not even an explicit nil
### GetEnabled

`func (o *GetVDIPools200ResponseVdiPool) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetVDIPools200ResponseVdiPool) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetVDIPools200ResponseVdiPool) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetVDIPools200ResponseVdiPool) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *GetVDIPools200ResponseVdiPool) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *GetVDIPools200ResponseVdiPool) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *GetVDIPools200ResponseVdiPool) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *GetVDIPools200ResponseVdiPool) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *GetVDIPools200ResponseVdiPool) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *GetVDIPools200ResponseVdiPool) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *GetVDIPools200ResponseVdiPool) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *GetVDIPools200ResponseVdiPool) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### SetAllowHypervisorConsoleNil

`func (o *GetVDIPools200ResponseVdiPool) SetAllowHypervisorConsoleNil(b bool)`

 SetAllowHypervisorConsoleNil sets the value for AllowHypervisorConsole to be an explicit nil

### UnsetAllowHypervisorConsole
`func (o *GetVDIPools200ResponseVdiPool) UnsetAllowHypervisorConsole()`

UnsetAllowHypervisorConsole ensures that no value is present for AllowHypervisorConsole, not even an explicit nil
### GetAllowCopy

`func (o *GetVDIPools200ResponseVdiPool) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *GetVDIPools200ResponseVdiPool) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *GetVDIPools200ResponseVdiPool) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *GetVDIPools200ResponseVdiPool) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### SetAllowCopyNil

`func (o *GetVDIPools200ResponseVdiPool) SetAllowCopyNil(b bool)`

 SetAllowCopyNil sets the value for AllowCopy to be an explicit nil

### UnsetAllowCopy
`func (o *GetVDIPools200ResponseVdiPool) UnsetAllowCopy()`

UnsetAllowCopy ensures that no value is present for AllowCopy, not even an explicit nil
### GetAllowPrinter

`func (o *GetVDIPools200ResponseVdiPool) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *GetVDIPools200ResponseVdiPool) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *GetVDIPools200ResponseVdiPool) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *GetVDIPools200ResponseVdiPool) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### SetAllowPrinterNil

`func (o *GetVDIPools200ResponseVdiPool) SetAllowPrinterNil(b bool)`

 SetAllowPrinterNil sets the value for AllowPrinter to be an explicit nil

### UnsetAllowPrinter
`func (o *GetVDIPools200ResponseVdiPool) UnsetAllowPrinter()`

UnsetAllowPrinter ensures that no value is present for AllowPrinter, not even an explicit nil
### GetAllowFileshare

`func (o *GetVDIPools200ResponseVdiPool) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *GetVDIPools200ResponseVdiPool) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *GetVDIPools200ResponseVdiPool) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *GetVDIPools200ResponseVdiPool) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### SetAllowFileshareNil

`func (o *GetVDIPools200ResponseVdiPool) SetAllowFileshareNil(b bool)`

 SetAllowFileshareNil sets the value for AllowFileshare to be an explicit nil

### UnsetAllowFileshare
`func (o *GetVDIPools200ResponseVdiPool) UnsetAllowFileshare()`

UnsetAllowFileshare ensures that no value is present for AllowFileshare, not even an explicit nil
### GetGuestConsoleJumpHost

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *GetVDIPools200ResponseVdiPool) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### SetGuestConsoleJumpHostNil

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpHostNil(b bool)`

 SetGuestConsoleJumpHostNil sets the value for GuestConsoleJumpHost to be an explicit nil

### UnsetGuestConsoleJumpHost
`func (o *GetVDIPools200ResponseVdiPool) UnsetGuestConsoleJumpHost()`

UnsetGuestConsoleJumpHost ensures that no value is present for GuestConsoleJumpHost, not even an explicit nil
### GetGuestConsoleJumpPort

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpPort() string`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpPortOk() (*string, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpPort(v string)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *GetVDIPools200ResponseVdiPool) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### SetGuestConsoleJumpPortNil

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpPortNil(b bool)`

 SetGuestConsoleJumpPortNil sets the value for GuestConsoleJumpPort to be an explicit nil

### UnsetGuestConsoleJumpPort
`func (o *GetVDIPools200ResponseVdiPool) UnsetGuestConsoleJumpPort()`

UnsetGuestConsoleJumpPort ensures that no value is present for GuestConsoleJumpPort, not even an explicit nil
### GetGuestConsoleJumpUsername

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *GetVDIPools200ResponseVdiPool) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### SetGuestConsoleJumpUsernameNil

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpUsernameNil(b bool)`

 SetGuestConsoleJumpUsernameNil sets the value for GuestConsoleJumpUsername to be an explicit nil

### UnsetGuestConsoleJumpUsername
`func (o *GetVDIPools200ResponseVdiPool) UnsetGuestConsoleJumpUsername()`

UnsetGuestConsoleJumpUsername ensures that no value is present for GuestConsoleJumpUsername, not even an explicit nil
### GetGuestConsoleJumpPassword

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *GetVDIPools200ResponseVdiPool) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### SetGuestConsoleJumpPasswordNil

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpPasswordNil(b bool)`

 SetGuestConsoleJumpPasswordNil sets the value for GuestConsoleJumpPassword to be an explicit nil

### UnsetGuestConsoleJumpPassword
`func (o *GetVDIPools200ResponseVdiPool) UnsetGuestConsoleJumpPassword()`

UnsetGuestConsoleJumpPassword ensures that no value is present for GuestConsoleJumpPassword, not even an explicit nil
### GetGuestConsoleJumpKeypair

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpKeypair() string`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *GetVDIPools200ResponseVdiPool) GetGuestConsoleJumpKeypairOk() (*string, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpKeypair(v string)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *GetVDIPools200ResponseVdiPool) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.

### SetGuestConsoleJumpKeypairNil

`func (o *GetVDIPools200ResponseVdiPool) SetGuestConsoleJumpKeypairNil(b bool)`

 SetGuestConsoleJumpKeypairNil sets the value for GuestConsoleJumpKeypair to be an explicit nil

### UnsetGuestConsoleJumpKeypair
`func (o *GetVDIPools200ResponseVdiPool) UnsetGuestConsoleJumpKeypair()`

UnsetGuestConsoleJumpKeypair ensures that no value is present for GuestConsoleJumpKeypair, not even an explicit nil
### GetGateway

`func (o *GetVDIPools200ResponseVdiPool) GetGateway() GetVDIPools200ResponseVdiPoolGateway`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *GetVDIPools200ResponseVdiPool) GetGatewayOk() (*GetVDIPools200ResponseVdiPoolGateway, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *GetVDIPools200ResponseVdiPool) SetGateway(v GetVDIPools200ResponseVdiPoolGateway)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *GetVDIPools200ResponseVdiPool) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIconPath

`func (o *GetVDIPools200ResponseVdiPool) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *GetVDIPools200ResponseVdiPool) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *GetVDIPools200ResponseVdiPool) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *GetVDIPools200ResponseVdiPool) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLogo

`func (o *GetVDIPools200ResponseVdiPool) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *GetVDIPools200ResponseVdiPool) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *GetVDIPools200ResponseVdiPool) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *GetVDIPools200ResponseVdiPool) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetApps

`func (o *GetVDIPools200ResponseVdiPool) GetApps() []GetVDIPools200ResponseVdiPoolAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *GetVDIPools200ResponseVdiPool) GetAppsOk() (*[]GetVDIPools200ResponseVdiPoolAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *GetVDIPools200ResponseVdiPool) SetApps(v []GetVDIPools200ResponseVdiPoolAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *GetVDIPools200ResponseVdiPool) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetOwner

`func (o *GetVDIPools200ResponseVdiPool) GetOwner() GetVDIPools200ResponseVdiPoolOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetVDIPools200ResponseVdiPool) GetOwnerOk() (*GetVDIPools200ResponseVdiPoolOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetVDIPools200ResponseVdiPool) SetOwner(v GetVDIPools200ResponseVdiPoolOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetVDIPools200ResponseVdiPool) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *GetVDIPools200ResponseVdiPool) GetConfig() GetVDIPools200ResponseVdiPoolConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetVDIPools200ResponseVdiPool) GetConfigOk() (*GetVDIPools200ResponseVdiPoolConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetVDIPools200ResponseVdiPool) SetConfig(v GetVDIPools200ResponseVdiPoolConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetVDIPools200ResponseVdiPool) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGroup

`func (o *GetVDIPools200ResponseVdiPool) GetGroup() GetVDIPools200ResponseVdiPoolGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *GetVDIPools200ResponseVdiPool) GetGroupOk() (*GetVDIPools200ResponseVdiPoolGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *GetVDIPools200ResponseVdiPool) SetGroup(v GetVDIPools200ResponseVdiPoolGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *GetVDIPools200ResponseVdiPool) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *GetVDIPools200ResponseVdiPool) GetCloud() GetVDIPools200ResponseVdiPoolCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *GetVDIPools200ResponseVdiPool) GetCloudOk() (*GetVDIPools200ResponseVdiPoolCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *GetVDIPools200ResponseVdiPool) SetCloud(v GetVDIPools200ResponseVdiPoolCloud)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *GetVDIPools200ResponseVdiPool) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUsedCount

`func (o *GetVDIPools200ResponseVdiPool) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *GetVDIPools200ResponseVdiPool) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *GetVDIPools200ResponseVdiPool) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *GetVDIPools200ResponseVdiPool) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetReservedCount

`func (o *GetVDIPools200ResponseVdiPool) GetReservedCount() int64`

GetReservedCount returns the ReservedCount field if non-nil, zero value otherwise.

### GetReservedCountOk

`func (o *GetVDIPools200ResponseVdiPool) GetReservedCountOk() (*int64, bool)`

GetReservedCountOk returns a tuple with the ReservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCount

`func (o *GetVDIPools200ResponseVdiPool) SetReservedCount(v int64)`

SetReservedCount sets ReservedCount field to given value.

### HasReservedCount

`func (o *GetVDIPools200ResponseVdiPool) HasReservedCount() bool`

HasReservedCount returns a boolean if a field has been set.

### GetPreparingCount

`func (o *GetVDIPools200ResponseVdiPool) GetPreparingCount() int64`

GetPreparingCount returns the PreparingCount field if non-nil, zero value otherwise.

### GetPreparingCountOk

`func (o *GetVDIPools200ResponseVdiPool) GetPreparingCountOk() (*int64, bool)`

GetPreparingCountOk returns a tuple with the PreparingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingCount

`func (o *GetVDIPools200ResponseVdiPool) SetPreparingCount(v int64)`

SetPreparingCount sets PreparingCount field to given value.

### HasPreparingCount

`func (o *GetVDIPools200ResponseVdiPool) HasPreparingCount() bool`

HasPreparingCount returns a boolean if a field has been set.

### GetIdleCount

`func (o *GetVDIPools200ResponseVdiPool) GetIdleCount() int64`

GetIdleCount returns the IdleCount field if non-nil, zero value otherwise.

### GetIdleCountOk

`func (o *GetVDIPools200ResponseVdiPool) GetIdleCountOk() (*int64, bool)`

GetIdleCountOk returns a tuple with the IdleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCount

`func (o *GetVDIPools200ResponseVdiPool) SetIdleCount(v int64)`

SetIdleCount sets IdleCount field to given value.

### HasIdleCount

`func (o *GetVDIPools200ResponseVdiPool) HasIdleCount() bool`

HasIdleCount returns a boolean if a field has been set.

### GetStatus

`func (o *GetVDIPools200ResponseVdiPool) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVDIPools200ResponseVdiPool) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVDIPools200ResponseVdiPool) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVDIPools200ResponseVdiPool) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetVDIPools200ResponseVdiPool) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetVDIPools200ResponseVdiPool) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetVDIPools200ResponseVdiPool) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetVDIPools200ResponseVdiPool) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetVDIPools200ResponseVdiPool) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetVDIPools200ResponseVdiPool) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetVDIPools200ResponseVdiPool) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetVDIPools200ResponseVdiPool) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


