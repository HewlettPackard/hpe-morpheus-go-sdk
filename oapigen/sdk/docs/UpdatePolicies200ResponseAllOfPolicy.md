# UpdatePolicies200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyPolicyType**](UpdatePolicies200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyZone**](UpdatePolicies200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**UpdatePolicies200ResponseAllOfPolicySite**](UpdatePolicies200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyUser**](UpdatePolicies200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyRole**](UpdatePolicies200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyConfig**](UpdatePolicies200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**UpdatePolicies200ResponseAllOfPolicyOwner**](UpdatePolicies200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]AddPolicies200ResponseAllOfPolicyAccountsInner**](AddPolicies200ResponseAllOfPolicyAccountsInner.md) |  | [optional] 

## Methods

### NewUpdatePolicies200ResponseAllOfPolicy

`func NewUpdatePolicies200ResponseAllOfPolicy() *UpdatePolicies200ResponseAllOfPolicy`

NewUpdatePolicies200ResponseAllOfPolicy instantiates a new UpdatePolicies200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePolicies200ResponseAllOfPolicyWithDefaults

`func NewUpdatePolicies200ResponseAllOfPolicyWithDefaults() *UpdatePolicies200ResponseAllOfPolicy`

NewUpdatePolicies200ResponseAllOfPolicyWithDefaults instantiates a new UpdatePolicies200ResponseAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePolicies200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetPolicyType() UpdatePolicies200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetPolicyTypeOk() (*UpdatePolicies200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetPolicyType(v UpdatePolicies200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetZone() UpdatePolicies200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetZoneOk() (*UpdatePolicies200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetZone(v UpdatePolicies200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetSite() UpdatePolicies200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetSiteOk() (*UpdatePolicies200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetSite(v UpdatePolicies200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetUser() UpdatePolicies200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetUserOk() (*UpdatePolicies200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetUser(v UpdatePolicies200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRole() UpdatePolicies200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRoleOk() (*UpdatePolicies200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetRole(v UpdatePolicies200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdatePolicies200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdatePolicies200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *UpdatePolicies200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetConfig() UpdatePolicies200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetConfigOk() (*UpdatePolicies200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetConfig(v UpdatePolicies200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetOwner() UpdatePolicies200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetOwnerOk() (*UpdatePolicies200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetOwner(v UpdatePolicies200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetAccounts() []AddPolicies200ResponseAllOfPolicyAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdatePolicies200ResponseAllOfPolicy) GetAccountsOk() (*[]AddPolicies200ResponseAllOfPolicyAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdatePolicies200ResponseAllOfPolicy) SetAccounts(v []AddPolicies200ResponseAllOfPolicyAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdatePolicies200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


