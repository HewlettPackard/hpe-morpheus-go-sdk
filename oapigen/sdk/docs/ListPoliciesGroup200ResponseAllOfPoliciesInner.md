# ListPoliciesGroup200ResponseAllOfPoliciesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerPolicyType**](ListPoliciesGroup200ResponseAllOfPoliciesInnerPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerZone**](ListPoliciesGroup200ResponseAllOfPoliciesInnerZone.md) |  | [optional] 
**Site** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerSite**](ListPoliciesGroup200ResponseAllOfPoliciesInnerSite.md) |  | [optional] 
**User** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerUser**](ListPoliciesGroup200ResponseAllOfPoliciesInnerUser.md) |  | [optional] 
**Role** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerRole**](ListPoliciesGroup200ResponseAllOfPoliciesInnerRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**ListPoliciesGroup200ResponseAllOfPoliciesInnerConfig**](ListPoliciesGroup200ResponseAllOfPoliciesInnerConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**NullableListPoliciesGroup200ResponseAllOfPoliciesInnerOwner**](ListPoliciesGroup200ResponseAllOfPoliciesInnerOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner**](ListPolicies200ResponseAllOfPoliciesInnerAccountsInner.md) |  | [optional] 

## Methods

### NewListPoliciesGroup200ResponseAllOfPoliciesInner

`func NewListPoliciesGroup200ResponseAllOfPoliciesInner() *ListPoliciesGroup200ResponseAllOfPoliciesInner`

NewListPoliciesGroup200ResponseAllOfPoliciesInner instantiates a new ListPoliciesGroup200ResponseAllOfPoliciesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPoliciesGroup200ResponseAllOfPoliciesInnerWithDefaults

`func NewListPoliciesGroup200ResponseAllOfPoliciesInnerWithDefaults() *ListPoliciesGroup200ResponseAllOfPoliciesInner`

NewListPoliciesGroup200ResponseAllOfPoliciesInnerWithDefaults instantiates a new ListPoliciesGroup200ResponseAllOfPoliciesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetPolicyType() ListPoliciesGroup200ResponseAllOfPoliciesInnerPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetPolicyTypeOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetPolicyType(v ListPoliciesGroup200ResponseAllOfPoliciesInnerPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetZone() ListPoliciesGroup200ResponseAllOfPoliciesInnerZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetZoneOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetZone(v ListPoliciesGroup200ResponseAllOfPoliciesInnerZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetSite() ListPoliciesGroup200ResponseAllOfPoliciesInnerSite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetSiteOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerSite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetSite(v ListPoliciesGroup200ResponseAllOfPoliciesInnerSite)`

SetSite sets Site field to given value.

### HasSite

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetUser() ListPoliciesGroup200ResponseAllOfPoliciesInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetUserOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetUser(v ListPoliciesGroup200ResponseAllOfPoliciesInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRole() ListPoliciesGroup200ResponseAllOfPoliciesInnerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRoleOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetRole(v ListPoliciesGroup200ResponseAllOfPoliciesInnerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetConfig() ListPoliciesGroup200ResponseAllOfPoliciesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetConfigOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetConfig(v ListPoliciesGroup200ResponseAllOfPoliciesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetOwner() ListPoliciesGroup200ResponseAllOfPoliciesInnerOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetOwnerOk() (*ListPoliciesGroup200ResponseAllOfPoliciesInnerOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetOwner(v ListPoliciesGroup200ResponseAllOfPoliciesInnerOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### SetOwnerNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetOwnerNil(b bool)`

 SetOwnerNil sets the value for Owner to be an explicit nil

### UnsetOwner
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetOwner()`

UnsetOwner ensures that no value is present for Owner, not even an explicit nil
### GetAccounts

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetAccounts() []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) GetAccountsOk() (*[]ListPolicies200ResponseAllOfPoliciesInnerAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetAccounts(v []ListPolicies200ResponseAllOfPoliciesInnerAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### SetAccountsNil

`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) SetAccountsNil(b bool)`

 SetAccountsNil sets the value for Accounts to be an explicit nil

### UnsetAccounts
`func (o *ListPoliciesGroup200ResponseAllOfPoliciesInner) UnsetAccounts()`

UnsetAccounts ensures that no value is present for Accounts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


