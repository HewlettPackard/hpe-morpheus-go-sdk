# GetPolicies200ResponseAllOfPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**PolicyType** | Pointer to [**GetPolicies200ResponseAllOfPolicyPolicyType**](GetPolicies200ResponseAllOfPolicyPolicyType.md) |  | [optional] 
**Zone** | Pointer to [**GetPolicies200ResponseAllOfPolicyZone**](GetPolicies200ResponseAllOfPolicyZone.md) |  | [optional] 
**Site** | Pointer to [**GetPolicies200ResponseAllOfPolicySite**](GetPolicies200ResponseAllOfPolicySite.md) |  | [optional] 
**User** | Pointer to [**GetPolicies200ResponseAllOfPolicyUser**](GetPolicies200ResponseAllOfPolicyUser.md) |  | [optional] 
**Role** | Pointer to [**GetPolicies200ResponseAllOfPolicyRole**](GetPolicies200ResponseAllOfPolicyRole.md) |  | [optional] 
**RefType** | Pointer to **NullableString** |  | [optional] 
**RefId** | Pointer to **NullableInt64** |  | [optional] 
**EachUser** | Pointer to **NullableBool** |  | [optional] 
**Config** | Pointer to [**GetPolicies200ResponseAllOfPolicyConfig**](GetPolicies200ResponseAllOfPolicyConfig.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Owner** | Pointer to [**GetPolicies200ResponseAllOfPolicyOwner**](GetPolicies200ResponseAllOfPolicyOwner.md) |  | [optional] 
**Accounts** | Pointer to [**[]AddPolicies200ResponseAllOfPolicyAccountsInner**](AddPolicies200ResponseAllOfPolicyAccountsInner.md) |  | [optional] 

## Methods

### NewGetPolicies200ResponseAllOfPolicy

`func NewGetPolicies200ResponseAllOfPolicy() *GetPolicies200ResponseAllOfPolicy`

NewGetPolicies200ResponseAllOfPolicy instantiates a new GetPolicies200ResponseAllOfPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPolicies200ResponseAllOfPolicyWithDefaults

`func NewGetPolicies200ResponseAllOfPolicyWithDefaults() *GetPolicies200ResponseAllOfPolicy`

NewGetPolicies200ResponseAllOfPolicyWithDefaults instantiates a new GetPolicies200ResponseAllOfPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPolicies200ResponseAllOfPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPolicies200ResponseAllOfPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPolicies200ResponseAllOfPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPolicies200ResponseAllOfPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPolicies200ResponseAllOfPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPolicies200ResponseAllOfPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetPolicies200ResponseAllOfPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPolicies200ResponseAllOfPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPolicies200ResponseAllOfPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetPolicies200ResponseAllOfPolicy) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetPolicies200ResponseAllOfPolicy) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPolicyType

`func (o *GetPolicies200ResponseAllOfPolicy) GetPolicyType() GetPolicies200ResponseAllOfPolicyPolicyType`

GetPolicyType returns the PolicyType field if non-nil, zero value otherwise.

### GetPolicyTypeOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetPolicyTypeOk() (*GetPolicies200ResponseAllOfPolicyPolicyType, bool)`

GetPolicyTypeOk returns a tuple with the PolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyType

`func (o *GetPolicies200ResponseAllOfPolicy) SetPolicyType(v GetPolicies200ResponseAllOfPolicyPolicyType)`

SetPolicyType sets PolicyType field to given value.

### HasPolicyType

`func (o *GetPolicies200ResponseAllOfPolicy) HasPolicyType() bool`

HasPolicyType returns a boolean if a field has been set.

### GetZone

`func (o *GetPolicies200ResponseAllOfPolicy) GetZone() GetPolicies200ResponseAllOfPolicyZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetZoneOk() (*GetPolicies200ResponseAllOfPolicyZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetPolicies200ResponseAllOfPolicy) SetZone(v GetPolicies200ResponseAllOfPolicyZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetPolicies200ResponseAllOfPolicy) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetSite

`func (o *GetPolicies200ResponseAllOfPolicy) GetSite() GetPolicies200ResponseAllOfPolicySite`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetSiteOk() (*GetPolicies200ResponseAllOfPolicySite, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *GetPolicies200ResponseAllOfPolicy) SetSite(v GetPolicies200ResponseAllOfPolicySite)`

SetSite sets Site field to given value.

### HasSite

`func (o *GetPolicies200ResponseAllOfPolicy) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetUser

`func (o *GetPolicies200ResponseAllOfPolicy) GetUser() GetPolicies200ResponseAllOfPolicyUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetUserOk() (*GetPolicies200ResponseAllOfPolicyUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GetPolicies200ResponseAllOfPolicy) SetUser(v GetPolicies200ResponseAllOfPolicyUser)`

SetUser sets User field to given value.

### HasUser

`func (o *GetPolicies200ResponseAllOfPolicy) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetRole

`func (o *GetPolicies200ResponseAllOfPolicy) GetRole() GetPolicies200ResponseAllOfPolicyRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetRoleOk() (*GetPolicies200ResponseAllOfPolicyRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetPolicies200ResponseAllOfPolicy) SetRole(v GetPolicies200ResponseAllOfPolicyRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *GetPolicies200ResponseAllOfPolicy) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetRefType

`func (o *GetPolicies200ResponseAllOfPolicy) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetPolicies200ResponseAllOfPolicy) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetPolicies200ResponseAllOfPolicy) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### SetRefTypeNil

`func (o *GetPolicies200ResponseAllOfPolicy) SetRefTypeNil(b bool)`

 SetRefTypeNil sets the value for RefType to be an explicit nil

### UnsetRefType
`func (o *GetPolicies200ResponseAllOfPolicy) UnsetRefType()`

UnsetRefType ensures that no value is present for RefType, not even an explicit nil
### GetRefId

`func (o *GetPolicies200ResponseAllOfPolicy) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetPolicies200ResponseAllOfPolicy) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetPolicies200ResponseAllOfPolicy) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### SetRefIdNil

`func (o *GetPolicies200ResponseAllOfPolicy) SetRefIdNil(b bool)`

 SetRefIdNil sets the value for RefId to be an explicit nil

### UnsetRefId
`func (o *GetPolicies200ResponseAllOfPolicy) UnsetRefId()`

UnsetRefId ensures that no value is present for RefId, not even an explicit nil
### GetEachUser

`func (o *GetPolicies200ResponseAllOfPolicy) GetEachUser() bool`

GetEachUser returns the EachUser field if non-nil, zero value otherwise.

### GetEachUserOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetEachUserOk() (*bool, bool)`

GetEachUserOk returns a tuple with the EachUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEachUser

`func (o *GetPolicies200ResponseAllOfPolicy) SetEachUser(v bool)`

SetEachUser sets EachUser field to given value.

### HasEachUser

`func (o *GetPolicies200ResponseAllOfPolicy) HasEachUser() bool`

HasEachUser returns a boolean if a field has been set.

### SetEachUserNil

`func (o *GetPolicies200ResponseAllOfPolicy) SetEachUserNil(b bool)`

 SetEachUserNil sets the value for EachUser to be an explicit nil

### UnsetEachUser
`func (o *GetPolicies200ResponseAllOfPolicy) UnsetEachUser()`

UnsetEachUser ensures that no value is present for EachUser, not even an explicit nil
### GetConfig

`func (o *GetPolicies200ResponseAllOfPolicy) GetConfig() GetPolicies200ResponseAllOfPolicyConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetConfigOk() (*GetPolicies200ResponseAllOfPolicyConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetPolicies200ResponseAllOfPolicy) SetConfig(v GetPolicies200ResponseAllOfPolicyConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetPolicies200ResponseAllOfPolicy) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPolicies200ResponseAllOfPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPolicies200ResponseAllOfPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPolicies200ResponseAllOfPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOwner

`func (o *GetPolicies200ResponseAllOfPolicy) GetOwner() GetPolicies200ResponseAllOfPolicyOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetOwnerOk() (*GetPolicies200ResponseAllOfPolicyOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetPolicies200ResponseAllOfPolicy) SetOwner(v GetPolicies200ResponseAllOfPolicyOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetPolicies200ResponseAllOfPolicy) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetAccounts

`func (o *GetPolicies200ResponseAllOfPolicy) GetAccounts() []AddPolicies200ResponseAllOfPolicyAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetPolicies200ResponseAllOfPolicy) GetAccountsOk() (*[]AddPolicies200ResponseAllOfPolicyAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetPolicies200ResponseAllOfPolicy) SetAccounts(v []AddPolicies200ResponseAllOfPolicyAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetPolicies200ResponseAllOfPolicy) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


