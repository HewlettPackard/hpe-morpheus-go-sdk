# ListVDIPools200ResponseAllOfVdiPoolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MinIdle** | Pointer to **int64** |  | [optional] 
**MaxIdle** | Pointer to **int64** |  | [optional] 
**InitialPoolSize** | Pointer to **int64** |  | [optional] 
**MaxPoolSize** | Pointer to **int64** |  | [optional] 
**AllocationTimeoutMinutes** | Pointer to **int64** |  | [optional] 
**PersistentUser** | Pointer to **bool** |  | [optional] 
**Recyclable** | Pointer to **bool** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**AutoCreateLocalUserOnReservation** | Pointer to **bool** |  | [optional] 
**AllowHypervisorConsole** | Pointer to **bool** |  | [optional] 
**AllowCopy** | Pointer to **bool** |  | [optional] 
**AllowPrinter** | Pointer to **bool** |  | [optional] 
**AllowFileshare** | Pointer to **bool** |  | [optional] 
**GuestConsoleJumpHost** | Pointer to **string** |  | [optional] 
**GuestConsoleJumpPort** | Pointer to **string** |  | [optional] 
**GuestConsoleJumpUsername** | Pointer to **string** |  | [optional] 
**GuestConsoleJumpPassword** | Pointer to **string** |  | [optional] 
**GuestConsoleJumpKeypair** | Pointer to **string** |  | [optional] 
**Gateway** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**IconPath** | Pointer to **string** |  | [optional] 
**Logo** | Pointer to **string** |  | [optional] 
**Apps** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Owner** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerOwner**](ListVDIPools200ResponseAllOfVdiPoolsInnerOwner.md) |  | [optional] 
**Config** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInnerConfig**](ListVDIPools200ResponseAllOfVdiPoolsInnerConfig.md) |  | [optional] 
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

### NewListVDIPools200ResponseAllOfVdiPoolsInner

`func NewListVDIPools200ResponseAllOfVdiPoolsInner() *ListVDIPools200ResponseAllOfVdiPoolsInner`

NewListVDIPools200ResponseAllOfVdiPoolsInner instantiates a new ListVDIPools200ResponseAllOfVdiPoolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIPools200ResponseAllOfVdiPoolsInnerWithDefaults

`func NewListVDIPools200ResponseAllOfVdiPoolsInnerWithDefaults() *ListVDIPools200ResponseAllOfVdiPoolsInner`

NewListVDIPools200ResponseAllOfVdiPoolsInnerWithDefaults instantiates a new ListVDIPools200ResponseAllOfVdiPoolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMinIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMinIdle() int64`

GetMinIdle returns the MinIdle field if non-nil, zero value otherwise.

### GetMinIdleOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMinIdleOk() (*int64, bool)`

GetMinIdleOk returns a tuple with the MinIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetMinIdle(v int64)`

SetMinIdle sets MinIdle field to given value.

### HasMinIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasMinIdle() bool`

HasMinIdle returns a boolean if a field has been set.

### GetMaxIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMaxIdle() int64`

GetMaxIdle returns the MaxIdle field if non-nil, zero value otherwise.

### GetMaxIdleOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMaxIdleOk() (*int64, bool)`

GetMaxIdleOk returns a tuple with the MaxIdle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetMaxIdle(v int64)`

SetMaxIdle sets MaxIdle field to given value.

### HasMaxIdle

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasMaxIdle() bool`

HasMaxIdle returns a boolean if a field has been set.

### GetInitialPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetInitialPoolSize() int64`

GetInitialPoolSize returns the InitialPoolSize field if non-nil, zero value otherwise.

### GetInitialPoolSizeOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetInitialPoolSizeOk() (*int64, bool)`

GetInitialPoolSizeOk returns a tuple with the InitialPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetInitialPoolSize(v int64)`

SetInitialPoolSize sets InitialPoolSize field to given value.

### HasInitialPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasInitialPoolSize() bool`

HasInitialPoolSize returns a boolean if a field has been set.

### GetMaxPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMaxPoolSize() int64`

GetMaxPoolSize returns the MaxPoolSize field if non-nil, zero value otherwise.

### GetMaxPoolSizeOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetMaxPoolSizeOk() (*int64, bool)`

GetMaxPoolSizeOk returns a tuple with the MaxPoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetMaxPoolSize(v int64)`

SetMaxPoolSize sets MaxPoolSize field to given value.

### HasMaxPoolSize

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasMaxPoolSize() bool`

HasMaxPoolSize returns a boolean if a field has been set.

### GetAllocationTimeoutMinutes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllocationTimeoutMinutes() int64`

GetAllocationTimeoutMinutes returns the AllocationTimeoutMinutes field if non-nil, zero value otherwise.

### GetAllocationTimeoutMinutesOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllocationTimeoutMinutesOk() (*int64, bool)`

GetAllocationTimeoutMinutesOk returns a tuple with the AllocationTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllocationTimeoutMinutes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAllocationTimeoutMinutes(v int64)`

SetAllocationTimeoutMinutes sets AllocationTimeoutMinutes field to given value.

### HasAllocationTimeoutMinutes

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAllocationTimeoutMinutes() bool`

HasAllocationTimeoutMinutes returns a boolean if a field has been set.

### GetPersistentUser

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetPersistentUser() bool`

GetPersistentUser returns the PersistentUser field if non-nil, zero value otherwise.

### GetPersistentUserOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetPersistentUserOk() (*bool, bool)`

GetPersistentUserOk returns a tuple with the PersistentUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistentUser

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetPersistentUser(v bool)`

SetPersistentUser sets PersistentUser field to given value.

### HasPersistentUser

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasPersistentUser() bool`

HasPersistentUser returns a boolean if a field has been set.

### GetRecyclable

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetRecyclable() bool`

GetRecyclable returns the Recyclable field if non-nil, zero value otherwise.

### GetRecyclableOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetRecyclableOk() (*bool, bool)`

GetRecyclableOk returns a tuple with the Recyclable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecyclable

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetRecyclable(v bool)`

SetRecyclable sets Recyclable field to given value.

### HasRecyclable

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasRecyclable() bool`

HasRecyclable returns a boolean if a field has been set.

### GetEnabled

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAutoCreateLocalUserOnReservation

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAutoCreateLocalUserOnReservation() bool`

GetAutoCreateLocalUserOnReservation returns the AutoCreateLocalUserOnReservation field if non-nil, zero value otherwise.

### GetAutoCreateLocalUserOnReservationOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAutoCreateLocalUserOnReservationOk() (*bool, bool)`

GetAutoCreateLocalUserOnReservationOk returns a tuple with the AutoCreateLocalUserOnReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoCreateLocalUserOnReservation

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAutoCreateLocalUserOnReservation(v bool)`

SetAutoCreateLocalUserOnReservation sets AutoCreateLocalUserOnReservation field to given value.

### HasAutoCreateLocalUserOnReservation

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAutoCreateLocalUserOnReservation() bool`

HasAutoCreateLocalUserOnReservation returns a boolean if a field has been set.

### GetAllowHypervisorConsole

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowHypervisorConsole() bool`

GetAllowHypervisorConsole returns the AllowHypervisorConsole field if non-nil, zero value otherwise.

### GetAllowHypervisorConsoleOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowHypervisorConsoleOk() (*bool, bool)`

GetAllowHypervisorConsoleOk returns a tuple with the AllowHypervisorConsole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowHypervisorConsole

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAllowHypervisorConsole(v bool)`

SetAllowHypervisorConsole sets AllowHypervisorConsole field to given value.

### HasAllowHypervisorConsole

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAllowHypervisorConsole() bool`

HasAllowHypervisorConsole returns a boolean if a field has been set.

### GetAllowCopy

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowCopy() bool`

GetAllowCopy returns the AllowCopy field if non-nil, zero value otherwise.

### GetAllowCopyOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowCopyOk() (*bool, bool)`

GetAllowCopyOk returns a tuple with the AllowCopy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCopy

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAllowCopy(v bool)`

SetAllowCopy sets AllowCopy field to given value.

### HasAllowCopy

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAllowCopy() bool`

HasAllowCopy returns a boolean if a field has been set.

### GetAllowPrinter

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowPrinter() bool`

GetAllowPrinter returns the AllowPrinter field if non-nil, zero value otherwise.

### GetAllowPrinterOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowPrinterOk() (*bool, bool)`

GetAllowPrinterOk returns a tuple with the AllowPrinter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPrinter

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAllowPrinter(v bool)`

SetAllowPrinter sets AllowPrinter field to given value.

### HasAllowPrinter

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAllowPrinter() bool`

HasAllowPrinter returns a boolean if a field has been set.

### GetAllowFileshare

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowFileshare() bool`

GetAllowFileshare returns the AllowFileshare field if non-nil, zero value otherwise.

### GetAllowFileshareOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAllowFileshareOk() (*bool, bool)`

GetAllowFileshareOk returns a tuple with the AllowFileshare field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowFileshare

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetAllowFileshare(v bool)`

SetAllowFileshare sets AllowFileshare field to given value.

### HasAllowFileshare

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasAllowFileshare() bool`

HasAllowFileshare returns a boolean if a field has been set.

### GetGuestConsoleJumpHost

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpHost() string`

GetGuestConsoleJumpHost returns the GuestConsoleJumpHost field if non-nil, zero value otherwise.

### GetGuestConsoleJumpHostOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpHostOk() (*string, bool)`

GetGuestConsoleJumpHostOk returns a tuple with the GuestConsoleJumpHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpHost

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGuestConsoleJumpHost(v string)`

SetGuestConsoleJumpHost sets GuestConsoleJumpHost field to given value.

### HasGuestConsoleJumpHost

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGuestConsoleJumpHost() bool`

HasGuestConsoleJumpHost returns a boolean if a field has been set.

### GetGuestConsoleJumpPort

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpPort() string`

GetGuestConsoleJumpPort returns the GuestConsoleJumpPort field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPortOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpPortOk() (*string, bool)`

GetGuestConsoleJumpPortOk returns a tuple with the GuestConsoleJumpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPort

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGuestConsoleJumpPort(v string)`

SetGuestConsoleJumpPort sets GuestConsoleJumpPort field to given value.

### HasGuestConsoleJumpPort

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGuestConsoleJumpPort() bool`

HasGuestConsoleJumpPort returns a boolean if a field has been set.

### GetGuestConsoleJumpUsername

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpUsername() string`

GetGuestConsoleJumpUsername returns the GuestConsoleJumpUsername field if non-nil, zero value otherwise.

### GetGuestConsoleJumpUsernameOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpUsernameOk() (*string, bool)`

GetGuestConsoleJumpUsernameOk returns a tuple with the GuestConsoleJumpUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpUsername

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGuestConsoleJumpUsername(v string)`

SetGuestConsoleJumpUsername sets GuestConsoleJumpUsername field to given value.

### HasGuestConsoleJumpUsername

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGuestConsoleJumpUsername() bool`

HasGuestConsoleJumpUsername returns a boolean if a field has been set.

### GetGuestConsoleJumpPassword

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpPassword() string`

GetGuestConsoleJumpPassword returns the GuestConsoleJumpPassword field if non-nil, zero value otherwise.

### GetGuestConsoleJumpPasswordOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpPasswordOk() (*string, bool)`

GetGuestConsoleJumpPasswordOk returns a tuple with the GuestConsoleJumpPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpPassword

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGuestConsoleJumpPassword(v string)`

SetGuestConsoleJumpPassword sets GuestConsoleJumpPassword field to given value.

### HasGuestConsoleJumpPassword

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGuestConsoleJumpPassword() bool`

HasGuestConsoleJumpPassword returns a boolean if a field has been set.

### GetGuestConsoleJumpKeypair

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpKeypair() string`

GetGuestConsoleJumpKeypair returns the GuestConsoleJumpKeypair field if non-nil, zero value otherwise.

### GetGuestConsoleJumpKeypairOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGuestConsoleJumpKeypairOk() (*string, bool)`

GetGuestConsoleJumpKeypairOk returns a tuple with the GuestConsoleJumpKeypair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestConsoleJumpKeypair

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGuestConsoleJumpKeypair(v string)`

SetGuestConsoleJumpKeypair sets GuestConsoleJumpKeypair field to given value.

### HasGuestConsoleJumpKeypair

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGuestConsoleJumpKeypair() bool`

HasGuestConsoleJumpKeypair returns a boolean if a field has been set.

### GetGateway

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGateway() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGatewayOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGateway(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetIconPath

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetLogo

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetLogo() string`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetLogoOk() (*string, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetLogo(v string)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetApps

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetApps() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetAppsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetApps(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetOwner

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetOwner() ListVDIPools200ResponseAllOfVdiPoolsInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetOwnerOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetOwner(v ListVDIPools200ResponseAllOfVdiPoolsInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetConfig() ListVDIPools200ResponseAllOfVdiPoolsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetConfigOk() (*ListVDIPools200ResponseAllOfVdiPoolsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetConfig(v ListVDIPools200ResponseAllOfVdiPoolsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGroup() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetGroup(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetCloud() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetCloudOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetCloud(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetCloud sets Cloud field to given value.

### HasCloud

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasCloud() bool`

HasCloud returns a boolean if a field has been set.

### GetUsedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetUsedCount() int64`

GetUsedCount returns the UsedCount field if non-nil, zero value otherwise.

### GetUsedCountOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetUsedCountOk() (*int64, bool)`

GetUsedCountOk returns a tuple with the UsedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetUsedCount(v int64)`

SetUsedCount sets UsedCount field to given value.

### HasUsedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasUsedCount() bool`

HasUsedCount returns a boolean if a field has been set.

### GetReservedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetReservedCount() int64`

GetReservedCount returns the ReservedCount field if non-nil, zero value otherwise.

### GetReservedCountOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetReservedCountOk() (*int64, bool)`

GetReservedCountOk returns a tuple with the ReservedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetReservedCount(v int64)`

SetReservedCount sets ReservedCount field to given value.

### HasReservedCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasReservedCount() bool`

HasReservedCount returns a boolean if a field has been set.

### GetPreparingCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetPreparingCount() int64`

GetPreparingCount returns the PreparingCount field if non-nil, zero value otherwise.

### GetPreparingCountOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetPreparingCountOk() (*int64, bool)`

GetPreparingCountOk returns a tuple with the PreparingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreparingCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetPreparingCount(v int64)`

SetPreparingCount sets PreparingCount field to given value.

### HasPreparingCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasPreparingCount() bool`

HasPreparingCount returns a boolean if a field has been set.

### GetIdleCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetIdleCount() int64`

GetIdleCount returns the IdleCount field if non-nil, zero value otherwise.

### GetIdleCountOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetIdleCountOk() (*int64, bool)`

GetIdleCountOk returns a tuple with the IdleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetIdleCount(v int64)`

SetIdleCount sets IdleCount field to given value.

### HasIdleCount

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasIdleCount() bool`

HasIdleCount returns a boolean if a field has been set.

### GetStatus

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListVDIPools200ResponseAllOfVdiPoolsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


