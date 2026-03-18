# GetGuidances200ResponseDiscoveryAnyOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**ActionCategory** | Pointer to **string** |  | [optional] 
**ActionMessage** | Pointer to **string** |  | [optional] 
**ActionTitle** | Pointer to **string** |  | [optional] 
**ActionType** | Pointer to **string** |  | [optional] 
**ActionValue** | Pointer to **string** |  | [optional] 
**ActionValueType** | Pointer to **string** |  | [optional] 
**ActionPlanId** | Pointer to **NullableString** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Zone** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOf1Zone**](GetGuidances200ResponseDiscoveryAnyOf1Zone.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateMessage** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**ResolvedMessage** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOf1Type**](GetGuidances200ResponseDiscoveryAnyOf1Type.md) |  | [optional] 
**Savings** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOf1Savings**](GetGuidances200ResponseDiscoveryAnyOf1Savings.md) |  | [optional] 
**Config** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOf1Config**](GetGuidances200ResponseDiscoveryAnyOf1Config.md) |  | [optional] 

## Methods

### NewGetGuidances200ResponseDiscoveryAnyOf1

`func NewGetGuidances200ResponseDiscoveryAnyOf1() *GetGuidances200ResponseDiscoveryAnyOf1`

NewGetGuidances200ResponseDiscoveryAnyOf1 instantiates a new GetGuidances200ResponseDiscoveryAnyOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidances200ResponseDiscoveryAnyOf1WithDefaults

`func NewGetGuidances200ResponseDiscoveryAnyOf1WithDefaults() *GetGuidances200ResponseDiscoveryAnyOf1`

NewGetGuidances200ResponseDiscoveryAnyOf1WithDefaults instantiates a new GetGuidances200ResponseDiscoveryAnyOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionPlanId() string`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetActionPlanIdOk() (*string, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionPlanId(v string)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### SetActionPlanIdNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetActionPlanIdNil(b bool)`

 SetActionPlanIdNil sets the value for ActionPlanId to be an explicit nil

### UnsetActionPlanId
`func (o *GetGuidances200ResponseDiscoveryAnyOf1) UnsetActionPlanId()`

UnsetActionPlanId ensures that no value is present for ActionPlanId, not even an explicit nil
### GetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetGuidances200ResponseDiscoveryAnyOf1) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetZone() GetGuidances200ResponseDiscoveryAnyOf1Zone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetZoneOk() (*GetGuidances200ResponseDiscoveryAnyOf1Zone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetZone(v GetGuidances200ResponseDiscoveryAnyOf1Zone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### SetStateMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetStateMessageNil(b bool)`

 SetStateMessageNil sets the value for StateMessage to be an explicit nil

### UnsetStateMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOf1) UnsetStateMessage()`

UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
### GetSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### SetResolvedMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetResolvedMessageNil(b bool)`

 SetResolvedMessageNil sets the value for ResolvedMessage to be an explicit nil

### UnsetResolvedMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOf1) UnsetResolvedMessage()`

UnsetResolvedMessage ensures that no value is present for ResolvedMessage, not even an explicit nil
### GetRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetType() GetGuidances200ResponseDiscoveryAnyOf1Type`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetTypeOk() (*GetGuidances200ResponseDiscoveryAnyOf1Type, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetType(v GetGuidances200ResponseDiscoveryAnyOf1Type)`

SetType sets Type field to given value.

### HasType

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSavings() GetGuidances200ResponseDiscoveryAnyOf1Savings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetSavingsOk() (*GetGuidances200ResponseDiscoveryAnyOf1Savings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetSavings(v GetGuidances200ResponseDiscoveryAnyOf1Savings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetConfig() GetGuidances200ResponseDiscoveryAnyOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) GetConfigOk() (*GetGuidances200ResponseDiscoveryAnyOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) SetConfig(v GetGuidances200ResponseDiscoveryAnyOf1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf1) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


