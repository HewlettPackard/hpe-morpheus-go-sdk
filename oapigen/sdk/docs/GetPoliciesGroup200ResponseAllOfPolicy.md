# GetPoliciesGroup200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicyPolicyType**](GetPoliciesGroup200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicyZone**](GetPoliciesGroup200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicySite**](GetPoliciesGroup200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicyUser**](GetPoliciesGroup200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicyRole**](GetPoliciesGroup200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**GetPoliciesGroup200ResponseAllOfPolicyConfig**](GetPoliciesGroup200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableGetPoliciesGroup200ResponseAllOfPolicyOwner**](GetPoliciesGroup200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner**](ListPolicies200ResponseAllOfPoliciesInnerAccountsInner.md) |  | [optional] 

## Methods

### NewGetPoliciesGroup200ResponseAllOfPolicy

`func NewGetPoliciesGroup200ResponseAllOfPolicy() *GetPoliciesGroup200ResponseAllOfPolicy`

NewGetPoliciesGroup200ResponseAllOfPolicy instantiates a new GetPoliciesGroup200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoliciesGroup200ResponseAllOfPolicyWithDefaults

`func NewGetPoliciesGroup200ResponseAllOfPolicyWithDefaults() *GetPoliciesGroup200ResponseAllOfPolicy`

NewGetPoliciesGroup200ResponseAllOfPolicyWithDefaults instantiates a new GetPoliciesGroup200ResponseAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetPolicyType() GetPoliciesGroup200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetPolicyTypeOk() (*GetPoliciesGroup200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetPolicyType(v GetPoliciesGroup200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetZone() GetPoliciesGroup200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetZoneOk() (*GetPoliciesGroup200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetZone(v GetPoliciesGroup200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetSite() GetPoliciesGroup200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetSiteOk() (*GetPoliciesGroup200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetSite(v GetPoliciesGroup200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetUser() GetPoliciesGroup200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetUserOk() (*GetPoliciesGroup200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetUser(v GetPoliciesGroup200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRole() GetPoliciesGroup200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRoleOk() (*GetPoliciesGroup200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetRole(v GetPoliciesGroup200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetConfig() GetPoliciesGroup200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetConfigOk() (*GetPoliciesGroup200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetConfig(v GetPoliciesGroup200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetOwner() GetPoliciesGroup200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetOwnerOk() (*GetPoliciesGroup200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetOwner(v GetPoliciesGroup200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetAccounts() []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) GetAccountsOk() (*[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetAccounts(v []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *GetPoliciesGroup200ResponseAllOfPolicy) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *GetPoliciesGroup200ResponseAllOfPolicy) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


