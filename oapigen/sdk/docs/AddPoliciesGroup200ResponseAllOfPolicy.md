# AddPoliciesGroup200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicyPolicyType**](AddPoliciesGroup200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicyZone**](AddPoliciesGroup200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicySite**](AddPoliciesGroup200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicyUser**](AddPoliciesGroup200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicyRole**](AddPoliciesGroup200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**AddPoliciesGroup200ResponseAllOfPolicyConfig**](AddPoliciesGroup200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableAddPoliciesGroup200ResponseAllOfPolicyOwner**](AddPoliciesGroup200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner**](ListPolicies200ResponseAllOfPoliciesInnerAccountsInner.md) |  | [optional] 

## Methods

### NewAddPoliciesGroup200ResponseAllOfPolicy

`func NewAddPoliciesGroup200ResponseAllOfPolicy() *AddPoliciesGroup200ResponseAllOfPolicy`

NewAddPoliciesGroup200ResponseAllOfPolicy instantiates a new AddPoliciesGroup200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetPolicyType() AddPoliciesGroup200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetPolicyTypeOk() (*AddPoliciesGroup200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetPolicyType(v AddPoliciesGroup200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetZone() AddPoliciesGroup200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetZoneOk() (*AddPoliciesGroup200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetZone(v AddPoliciesGroup200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetSite() AddPoliciesGroup200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetSiteOk() (*AddPoliciesGroup200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetSite(v AddPoliciesGroup200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetUser() AddPoliciesGroup200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetUserOk() (*AddPoliciesGroup200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetUser(v AddPoliciesGroup200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRole() AddPoliciesGroup200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRoleOk() (*AddPoliciesGroup200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetRole(v AddPoliciesGroup200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetConfig() AddPoliciesGroup200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetConfigOk() (*AddPoliciesGroup200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetConfig(v AddPoliciesGroup200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetOwner() AddPoliciesGroup200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetOwnerOk() (*AddPoliciesGroup200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetOwner(v AddPoliciesGroup200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetAccounts() []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) GetAccountsOk() (*[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetAccounts(v []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *AddPoliciesGroup200ResponseAllOfPolicy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *AddPoliciesGroup200ResponseAllOfPolicy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


