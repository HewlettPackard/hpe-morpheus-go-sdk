# UpdatePoliciesGroup200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicyPolicyType**](UpdatePoliciesGroup200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicyZone**](UpdatePoliciesGroup200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicySite**](UpdatePoliciesGroup200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicyUser**](UpdatePoliciesGroup200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicyRole**](UpdatePoliciesGroup200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**UpdatePoliciesGroup200ResponseAllOfPolicyConfig**](UpdatePoliciesGroup200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableUpdatePoliciesGroup200ResponseAllOfPolicyOwner**](UpdatePoliciesGroup200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner**](ListPolicies200ResponseAllOfPoliciesInnerAccountsInner.md) |  | [optional] 

## Methods

### NewUpdatePoliciesGroup200ResponseAllOfPolicy

`func NewUpdatePoliciesGroup200ResponseAllOfPolicy() *UpdatePoliciesGroup200ResponseAllOfPolicy`

NewUpdatePoliciesGroup200ResponseAllOfPolicy instantiates a new UpdatePoliciesGroup200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetPolicyType() UpdatePoliciesGroup200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetPolicyTypeOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetPolicyType(v UpdatePoliciesGroup200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetZone() UpdatePoliciesGroup200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetZoneOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetZone(v UpdatePoliciesGroup200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetSite() UpdatePoliciesGroup200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetSiteOk() (*UpdatePoliciesGroup200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetSite(v UpdatePoliciesGroup200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetUser() UpdatePoliciesGroup200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetUserOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetUser(v UpdatePoliciesGroup200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRole() UpdatePoliciesGroup200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRoleOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetRole(v UpdatePoliciesGroup200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetConfig() UpdatePoliciesGroup200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetConfigOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetConfig(v UpdatePoliciesGroup200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetOwner() UpdatePoliciesGroup200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetOwnerOk() (*UpdatePoliciesGroup200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetOwner(v UpdatePoliciesGroup200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetAccounts() []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) GetAccountsOk() (*[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetAccounts(v []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *UpdatePoliciesGroup200ResponseAllOfPolicy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


