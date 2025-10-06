# AddRoles200ResponseAllOfRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | a unique name of the role | [optional] 
**Authority** | Pointer to **string** | Alias for name | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LandingUrl** | Pointer to **NullableString** | An optional override for the default landing page after login for a user. | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**RoleType** | Pointer to **string** |  | [optional] 
**Multitenant** | Pointer to **bool** |  | [optional] 
**MultitenantLocked** | Pointer to **bool** |  | [optional] 
**ParentRoleId** | Pointer to **NullableString** |  | [optional] 
**Diverged** | Pointer to **bool** |  | [optional] 
**OwnerId** | Pointer to **int64** |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DefaultPersona** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddRoles200ResponseAllOfRole

`func NewAddRoles200ResponseAllOfRole() *AddRoles200ResponseAllOfRole`

NewAddRoles200ResponseAllOfRole instantiates a new AddRoles200ResponseAllOfRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddRoles200ResponseAllOfRoleWithDefaults

`func NewAddRoles200ResponseAllOfRoleWithDefaults() *AddRoles200ResponseAllOfRole`

NewAddRoles200ResponseAllOfRoleWithDefaults instantiates a new AddRoles200ResponseAllOfRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddRoles200ResponseAllOfRole) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddRoles200ResponseAllOfRole) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddRoles200ResponseAllOfRole) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddRoles200ResponseAllOfRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddRoles200ResponseAllOfRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddRoles200ResponseAllOfRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddRoles200ResponseAllOfRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddRoles200ResponseAllOfRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthority

`func (o *AddRoles200ResponseAllOfRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *AddRoles200ResponseAllOfRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *AddRoles200ResponseAllOfRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *AddRoles200ResponseAllOfRole) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *AddRoles200ResponseAllOfRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddRoles200ResponseAllOfRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddRoles200ResponseAllOfRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddRoles200ResponseAllOfRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddRoles200ResponseAllOfRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddRoles200ResponseAllOfRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLandingUrl

`func (o *AddRoles200ResponseAllOfRole) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *AddRoles200ResponseAllOfRole) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *AddRoles200ResponseAllOfRole) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *AddRoles200ResponseAllOfRole) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### SetLandingUrlNil

`func (o *AddRoles200ResponseAllOfRole) SetLandingUrlNil(b bool)`

 SetLandingUrlNil sets the value for LandingUrl to be an explicit nil

### UnsetLandingUrl
`func (o *AddRoles200ResponseAllOfRole) UnsetLandingUrl()`

UnsetLandingUrl ensures that no value is present for LandingUrl, not even an explicit nil
### GetScope

`func (o *AddRoles200ResponseAllOfRole) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AddRoles200ResponseAllOfRole) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AddRoles200ResponseAllOfRole) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AddRoles200ResponseAllOfRole) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRoleType

`func (o *AddRoles200ResponseAllOfRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *AddRoles200ResponseAllOfRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *AddRoles200ResponseAllOfRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *AddRoles200ResponseAllOfRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetMultitenant

`func (o *AddRoles200ResponseAllOfRole) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *AddRoles200ResponseAllOfRole) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *AddRoles200ResponseAllOfRole) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *AddRoles200ResponseAllOfRole) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *AddRoles200ResponseAllOfRole) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *AddRoles200ResponseAllOfRole) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *AddRoles200ResponseAllOfRole) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *AddRoles200ResponseAllOfRole) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetParentRoleId

`func (o *AddRoles200ResponseAllOfRole) GetParentRoleId() string`

GetParentRoleId returns the ParentRoleId field if non-nil, zero value otherwise.

### GetParentRoleIdOk

`func (o *AddRoles200ResponseAllOfRole) GetParentRoleIdOk() (*string, bool)`

GetParentRoleIdOk returns a tuple with the ParentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleId

`func (o *AddRoles200ResponseAllOfRole) SetParentRoleId(v string)`

SetParentRoleId sets ParentRoleId field to given value.

### HasParentRoleId

`func (o *AddRoles200ResponseAllOfRole) HasParentRoleId() bool`

HasParentRoleId returns a boolean if a field has been set.

### SetParentRoleIdNil

`func (o *AddRoles200ResponseAllOfRole) SetParentRoleIdNil(b bool)`

 SetParentRoleIdNil sets the value for ParentRoleId to be an explicit nil

### UnsetParentRoleId
`func (o *AddRoles200ResponseAllOfRole) UnsetParentRoleId()`

UnsetParentRoleId ensures that no value is present for ParentRoleId, not even an explicit nil
### GetDiverged

`func (o *AddRoles200ResponseAllOfRole) GetDiverged() bool`

GetDiverged returns the Diverged field if non-nil, zero value otherwise.

### GetDivergedOk

`func (o *AddRoles200ResponseAllOfRole) GetDivergedOk() (*bool, bool)`

GetDivergedOk returns a tuple with the Diverged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverged

`func (o *AddRoles200ResponseAllOfRole) SetDiverged(v bool)`

SetDiverged sets Diverged field to given value.

### HasDiverged

`func (o *AddRoles200ResponseAllOfRole) HasDiverged() bool`

HasDiverged returns a boolean if a field has been set.

### GetOwnerId

`func (o *AddRoles200ResponseAllOfRole) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *AddRoles200ResponseAllOfRole) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *AddRoles200ResponseAllOfRole) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *AddRoles200ResponseAllOfRole) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *AddRoles200ResponseAllOfRole) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddRoles200ResponseAllOfRole) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddRoles200ResponseAllOfRole) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddRoles200ResponseAllOfRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *AddRoles200ResponseAllOfRole) GetDefaultPersona() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *AddRoles200ResponseAllOfRole) GetDefaultPersonaOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *AddRoles200ResponseAllOfRole) SetDefaultPersona(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *AddRoles200ResponseAllOfRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddRoles200ResponseAllOfRole) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddRoles200ResponseAllOfRole) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddRoles200ResponseAllOfRole) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddRoles200ResponseAllOfRole) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddRoles200ResponseAllOfRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddRoles200ResponseAllOfRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddRoles200ResponseAllOfRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddRoles200ResponseAllOfRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


