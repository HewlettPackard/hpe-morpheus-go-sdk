# UpdatePoliciesCloud200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicyPolicyType**](UpdatePoliciesCloud200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicyZone**](UpdatePoliciesCloud200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicySite**](UpdatePoliciesCloud200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicyUser**](UpdatePoliciesCloud200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicyRole**](UpdatePoliciesCloud200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**UpdatePoliciesCloud200ResponseAllOfPolicyConfig**](UpdatePoliciesCloud200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableUpdatePoliciesCloud200ResponseAllOfPolicyOwner**](UpdatePoliciesCloud200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner**](ListPolicies200ResponseAllOfPoliciesInnerAccountsInner.md) |  | [optional] 

## Methods

### NewUpdatePoliciesCloud200ResponseAllOfPolicy

`func NewUpdatePoliciesCloud200ResponseAllOfPolicy() *UpdatePoliciesCloud200ResponseAllOfPolicy`

NewUpdatePoliciesCloud200ResponseAllOfPolicy instantiates a new UpdatePoliciesCloud200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePoliciesCloud200ResponseAllOfPolicyWithDefaults

`func NewUpdatePoliciesCloud200ResponseAllOfPolicyWithDefaults() *UpdatePoliciesCloud200ResponseAllOfPolicy`

NewUpdatePoliciesCloud200ResponseAllOfPolicyWithDefaults instantiates a new UpdatePoliciesCloud200ResponseAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetPolicyType() UpdatePoliciesCloud200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetPolicyTypeOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetPolicyType(v UpdatePoliciesCloud200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetZone() UpdatePoliciesCloud200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetZoneOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetZone(v UpdatePoliciesCloud200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetSite() UpdatePoliciesCloud200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetSiteOk() (*UpdatePoliciesCloud200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetSite(v UpdatePoliciesCloud200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetUser() UpdatePoliciesCloud200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetUserOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetUser(v UpdatePoliciesCloud200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRole() UpdatePoliciesCloud200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRoleOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetRole(v UpdatePoliciesCloud200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetConfig() UpdatePoliciesCloud200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetConfigOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetConfig(v UpdatePoliciesCloud200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetOwner() UpdatePoliciesCloud200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetOwnerOk() (*UpdatePoliciesCloud200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetOwner(v UpdatePoliciesCloud200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetAccounts() []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) GetAccountsOk() (*[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetAccounts(v []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *UpdatePoliciesCloud200ResponseAllOfPolicy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


