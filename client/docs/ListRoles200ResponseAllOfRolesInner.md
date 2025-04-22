# ListRoles200ResponseAllOfRolesInner

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

### NewListRoles200ResponseAllOfRolesInner

`func NewListRoles200ResponseAllOfRolesInner() *ListRoles200ResponseAllOfRolesInner`

NewListRoles200ResponseAllOfRolesInner instantiates a new ListRoles200ResponseAllOfRolesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRoles200ResponseAllOfRolesInnerWithDefaults

`func NewListRoles200ResponseAllOfRolesInnerWithDefaults() *ListRoles200ResponseAllOfRolesInner`

NewListRoles200ResponseAllOfRolesInnerWithDefaults instantiates a new ListRoles200ResponseAllOfRolesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRoles200ResponseAllOfRolesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRoles200ResponseAllOfRolesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListRoles200ResponseAllOfRolesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListRoles200ResponseAllOfRolesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListRoles200ResponseAllOfRolesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListRoles200ResponseAllOfRolesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthority

`func (o *ListRoles200ResponseAllOfRolesInner) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *ListRoles200ResponseAllOfRolesInner) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *ListRoles200ResponseAllOfRolesInner) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *ListRoles200ResponseAllOfRolesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListRoles200ResponseAllOfRolesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListRoles200ResponseAllOfRolesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLandingUrl

`func (o *ListRoles200ResponseAllOfRolesInner) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *ListRoles200ResponseAllOfRolesInner) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *ListRoles200ResponseAllOfRolesInner) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### GetScope

`func (o *ListRoles200ResponseAllOfRolesInner) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListRoles200ResponseAllOfRolesInner) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ListRoles200ResponseAllOfRolesInner) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRoleType

`func (o *ListRoles200ResponseAllOfRolesInner) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *ListRoles200ResponseAllOfRolesInner) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *ListRoles200ResponseAllOfRolesInner) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetMultitenant

`func (o *ListRoles200ResponseAllOfRolesInner) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *ListRoles200ResponseAllOfRolesInner) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *ListRoles200ResponseAllOfRolesInner) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *ListRoles200ResponseAllOfRolesInner) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *ListRoles200ResponseAllOfRolesInner) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *ListRoles200ResponseAllOfRolesInner) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetParentRoleId

`func (o *ListRoles200ResponseAllOfRolesInner) GetParentRoleId() string`

GetParentRoleId returns the ParentRoleId field if non-nil, zero value otherwise.

### GetParentRoleIdOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetParentRoleIdOk() (*string, bool)`

GetParentRoleIdOk returns a tuple with the ParentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleId

`func (o *ListRoles200ResponseAllOfRolesInner) SetParentRoleId(v string)`

SetParentRoleId sets ParentRoleId field to given value.

### HasParentRoleId

`func (o *ListRoles200ResponseAllOfRolesInner) HasParentRoleId() bool`

HasParentRoleId returns a boolean if a field has been set.

### GetDiverged

`func (o *ListRoles200ResponseAllOfRolesInner) GetDiverged() bool`

GetDiverged returns the Diverged field if non-nil, zero value otherwise.

### GetDivergedOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetDivergedOk() (*bool, bool)`

GetDivergedOk returns a tuple with the Diverged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverged

`func (o *ListRoles200ResponseAllOfRolesInner) SetDiverged(v bool)`

SetDiverged sets Diverged field to given value.

### HasDiverged

`func (o *ListRoles200ResponseAllOfRolesInner) HasDiverged() bool`

HasDiverged returns a boolean if a field has been set.

### GetOwnerId

`func (o *ListRoles200ResponseAllOfRolesInner) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *ListRoles200ResponseAllOfRolesInner) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *ListRoles200ResponseAllOfRolesInner) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *ListRoles200ResponseAllOfRolesInner) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListRoles200ResponseAllOfRolesInner) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListRoles200ResponseAllOfRolesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *ListRoles200ResponseAllOfRolesInner) GetDefaultPersona() string`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetDefaultPersonaOk() (*string, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *ListRoles200ResponseAllOfRolesInner) SetDefaultPersona(v string)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *ListRoles200ResponseAllOfRolesInner) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListRoles200ResponseAllOfRolesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListRoles200ResponseAllOfRolesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListRoles200ResponseAllOfRolesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListRoles200ResponseAllOfRolesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListRoles200ResponseAllOfRolesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListRoles200ResponseAllOfRolesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListRoles200ResponseAllOfRolesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


