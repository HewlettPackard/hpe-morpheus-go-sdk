# GetRole200ResponseRole

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
**Owner** | Pointer to [**GetRole200ResponseRoleOwner**](GetRole200ResponseRoleOwner.md) |  | [optional] 
**DefaultPersona** | Pointer to [**GetRole200ResponseRoleDefaultPersona**](GetRole200ResponseRoleDefaultPersona.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetRole200ResponseRole

`func NewGetRole200ResponseRole() *GetRole200ResponseRole`

NewGetRole200ResponseRole instantiates a new GetRole200ResponseRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRole200ResponseRoleWithDefaults

`func NewGetRole200ResponseRoleWithDefaults() *GetRole200ResponseRole`

NewGetRole200ResponseRoleWithDefaults instantiates a new GetRole200ResponseRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetRole200ResponseRole) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRole200ResponseRole) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRole200ResponseRole) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetRole200ResponseRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetRole200ResponseRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetRole200ResponseRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetRole200ResponseRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetRole200ResponseRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAuthority

`func (o *GetRole200ResponseRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *GetRole200ResponseRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *GetRole200ResponseRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *GetRole200ResponseRole) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.

### GetDescription

`func (o *GetRole200ResponseRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetRole200ResponseRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetRole200ResponseRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetRole200ResponseRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetRole200ResponseRole) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetRole200ResponseRole) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLandingUrl

`func (o *GetRole200ResponseRole) GetLandingUrl() string`

GetLandingUrl returns the LandingUrl field if non-nil, zero value otherwise.

### GetLandingUrlOk

`func (o *GetRole200ResponseRole) GetLandingUrlOk() (*string, bool)`

GetLandingUrlOk returns a tuple with the LandingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLandingUrl

`func (o *GetRole200ResponseRole) SetLandingUrl(v string)`

SetLandingUrl sets LandingUrl field to given value.

### HasLandingUrl

`func (o *GetRole200ResponseRole) HasLandingUrl() bool`

HasLandingUrl returns a boolean if a field has been set.

### SetLandingUrlNil

`func (o *GetRole200ResponseRole) SetLandingUrlNil(b bool)`

 SetLandingUrlNil sets the value for LandingUrl to be an explicit nil

### UnsetLandingUrl
`func (o *GetRole200ResponseRole) UnsetLandingUrl()`

UnsetLandingUrl ensures that no value is present for LandingUrl, not even an explicit nil
### GetScope

`func (o *GetRole200ResponseRole) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetRole200ResponseRole) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetRole200ResponseRole) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GetRole200ResponseRole) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRoleType

`func (o *GetRole200ResponseRole) GetRoleType() string`

GetRoleType returns the RoleType field if non-nil, zero value otherwise.

### GetRoleTypeOk

`func (o *GetRole200ResponseRole) GetRoleTypeOk() (*string, bool)`

GetRoleTypeOk returns a tuple with the RoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleType

`func (o *GetRole200ResponseRole) SetRoleType(v string)`

SetRoleType sets RoleType field to given value.

### HasRoleType

`func (o *GetRole200ResponseRole) HasRoleType() bool`

HasRoleType returns a boolean if a field has been set.

### GetMultitenant

`func (o *GetRole200ResponseRole) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *GetRole200ResponseRole) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *GetRole200ResponseRole) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *GetRole200ResponseRole) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMultitenantLocked

`func (o *GetRole200ResponseRole) GetMultitenantLocked() bool`

GetMultitenantLocked returns the MultitenantLocked field if non-nil, zero value otherwise.

### GetMultitenantLockedOk

`func (o *GetRole200ResponseRole) GetMultitenantLockedOk() (*bool, bool)`

GetMultitenantLockedOk returns a tuple with the MultitenantLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenantLocked

`func (o *GetRole200ResponseRole) SetMultitenantLocked(v bool)`

SetMultitenantLocked sets MultitenantLocked field to given value.

### HasMultitenantLocked

`func (o *GetRole200ResponseRole) HasMultitenantLocked() bool`

HasMultitenantLocked returns a boolean if a field has been set.

### GetParentRoleId

`func (o *GetRole200ResponseRole) GetParentRoleId() string`

GetParentRoleId returns the ParentRoleId field if non-nil, zero value otherwise.

### GetParentRoleIdOk

`func (o *GetRole200ResponseRole) GetParentRoleIdOk() (*string, bool)`

GetParentRoleIdOk returns a tuple with the ParentRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentRoleId

`func (o *GetRole200ResponseRole) SetParentRoleId(v string)`

SetParentRoleId sets ParentRoleId field to given value.

### HasParentRoleId

`func (o *GetRole200ResponseRole) HasParentRoleId() bool`

HasParentRoleId returns a boolean if a field has been set.

### SetParentRoleIdNil

`func (o *GetRole200ResponseRole) SetParentRoleIdNil(b bool)`

 SetParentRoleIdNil sets the value for ParentRoleId to be an explicit nil

### UnsetParentRoleId
`func (o *GetRole200ResponseRole) UnsetParentRoleId()`

UnsetParentRoleId ensures that no value is present for ParentRoleId, not even an explicit nil
### GetDiverged

`func (o *GetRole200ResponseRole) GetDiverged() bool`

GetDiverged returns the Diverged field if non-nil, zero value otherwise.

### GetDivergedOk

`func (o *GetRole200ResponseRole) GetDivergedOk() (*bool, bool)`

GetDivergedOk returns a tuple with the Diverged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiverged

`func (o *GetRole200ResponseRole) SetDiverged(v bool)`

SetDiverged sets Diverged field to given value.

### HasDiverged

`func (o *GetRole200ResponseRole) HasDiverged() bool`

HasDiverged returns a boolean if a field has been set.

### GetOwnerId

`func (o *GetRole200ResponseRole) GetOwnerId() int64`

GetOwnerId returns the OwnerId field if non-nil, zero value otherwise.

### GetOwnerIdOk

`func (o *GetRole200ResponseRole) GetOwnerIdOk() (*int64, bool)`

GetOwnerIdOk returns a tuple with the OwnerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerId

`func (o *GetRole200ResponseRole) SetOwnerId(v int64)`

SetOwnerId sets OwnerId field to given value.

### HasOwnerId

`func (o *GetRole200ResponseRole) HasOwnerId() bool`

HasOwnerId returns a boolean if a field has been set.

### GetOwner

`func (o *GetRole200ResponseRole) GetOwner() GetRole200ResponseRoleOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetRole200ResponseRole) GetOwnerOk() (*GetRole200ResponseRoleOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetRole200ResponseRole) SetOwner(v GetRole200ResponseRoleOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetRole200ResponseRole) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *GetRole200ResponseRole) GetDefaultPersona() GetRole200ResponseRoleDefaultPersona`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *GetRole200ResponseRole) GetDefaultPersonaOk() (*GetRole200ResponseRoleDefaultPersona, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *GetRole200ResponseRole) SetDefaultPersona(v GetRole200ResponseRoleDefaultPersona)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *GetRole200ResponseRole) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetRole200ResponseRole) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetRole200ResponseRole) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetRole200ResponseRole) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetRole200ResponseRole) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetRole200ResponseRole) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetRole200ResponseRole) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetRole200ResponseRole) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetRole200ResponseRole) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


