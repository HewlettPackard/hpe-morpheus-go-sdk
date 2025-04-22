# Roles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | a unique name of the role | [optional] 
**Authority** | Pointer to **string** | Alias for name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LandingUrl** | Pointer to **string** | An optional override for the default landing page after login for a user. | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**RoleType** | Pointer to **string** |  | [optional] 
**Multitenant** | Pointer to **bool** |  | [optional] 
**MultitenantLocked** | Pointer to **bool** |  | [optional] 
**ParentRoleId** | Pointer to **string** |  | [optional] 
**Diverged** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**DefaultPersona** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRoles

`func NewRoles() *Roles`

NewRoles instantiates a new Roles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolesWithDefaults

`func NewRolesWithDefaults() *Roles`

NewRolesWithDefaults instantiates a new Roles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Roles) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Roles) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Roles) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Roles) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Roles) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Roles) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Roles) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Roles) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthority

`func (o *Roles) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *Roles) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *Roles) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *Roles) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *Roles) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Roles) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Roles) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Roles) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLandingUrl

`func (o *Roles) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *Roles) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *Roles) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *Roles) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### GetScope

`func (o *Roles) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Roles) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Roles) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Roles) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRoleType

`func (o *Roles) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *Roles) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *Roles) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *Roles) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetMultitenant

`func (o *Roles) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *Roles) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *Roles) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *Roles) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *Roles) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *Roles) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *Roles) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *Roles) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetParentRoleId

`func (o *Roles) GetParentRoleId() string`

GetParentRoleId returns the ParentRoleId field if non-nil, zero value otherwise.

### GetParentRoleIdOk

`func (o *Roles) GetParentRoleIdOk() (*string, bool)`

GetParentRoleIdOk returns a tuple with the ParentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleId

`func (o *Roles) SetParentRoleId(v string)`

SetParentRoleId sets ParentRoleId field to given value.

### HasParentRoleId

`func (o *Roles) HasParentRoleId() bool`

HasParentRoleId returns a boolean if a field has been set.

### GetDiverged

`func (o *Roles) GetDiverged() bool`

GetDiverged returns the Diverged field if non-nil, zero value otherwise.

### GetDivergedOk

`func (o *Roles) GetDivergedOk() (*bool, bool)`

GetDivergedOk returns a tuple with the Diverged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverged

`func (o *Roles) SetDiverged(v bool)`

SetDiverged sets Diverged field to given value.

### HasDiverged

`func (o *Roles) HasDiverged() bool`

HasDiverged returns a boolean if a field has been set.

### GetOwnerId

`func (o *Roles) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *Roles) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *Roles) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *Roles) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *Roles) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Roles) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Roles) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Roles) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *Roles) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *Roles) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *Roles) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *Roles) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *Roles) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *Roles) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *Roles) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *Roles) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *Roles) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *Roles) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *Roles) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *Roles) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


