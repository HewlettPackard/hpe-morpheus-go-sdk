# GetGuidances200ResponseDiscoveryAnyOf

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
**ActionPlanId** | Pointer to **int64** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**SiteId** | Pointer to **NullableInt64** |  | [optional] 
**Zone** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateMessage** | Pointer to **NullableString** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**ResolvedMessage** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType.md) |  | [optional] 
**Savings** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings.md) |  | [optional] 
**Resource** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOfResource**](GetGuidances200ResponseDiscoveryAnyOfResource.md) |  | [optional] 
**PlanBeforeAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction.md) |  | [optional] 
**PlanAfterAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction.md) |  | [optional] 
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig.md) |  | [optional] 

## Methods

### NewGetGuidances200ResponseDiscoveryAnyOf

`func NewGetGuidances200ResponseDiscoveryAnyOf() *GetGuidances200ResponseDiscoveryAnyOf`

NewGetGuidances200ResponseDiscoveryAnyOf instantiates a new GetGuidances200ResponseDiscoveryAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidances200ResponseDiscoveryAnyOfWithDefaults

`func NewGetGuidances200ResponseDiscoveryAnyOfWithDefaults() *GetGuidances200ResponseDiscoveryAnyOf`

NewGetGuidances200ResponseDiscoveryAnyOfWithDefaults instantiates a new GetGuidances200ResponseDiscoveryAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionPlanId() int64`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetActionPlanIdOk() (*int64, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetActionPlanId(v int64)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetGuidances200ResponseDiscoveryAnyOf) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### SetSiteIdNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetSiteIdNil(b bool)`

 SetSiteIdNil sets the value for SiteId to be an explicit nil

### UnsetSiteId
`func (o *GetGuidances200ResponseDiscoveryAnyOf) UnsetSiteId()`

UnsetSiteId ensures that no value is present for SiteId, not even an explicit nil
### GetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### SetStateMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetStateMessageNil(b bool)`

 SetStateMessageNil sets the value for StateMessage to be an explicit nil

### UnsetStateMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOf) UnsetStateMessage()`

UnsetStateMessage ensures that no value is present for StateMessage, not even an explicit nil
### GetSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### SetResolvedMessageNil

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetResolvedMessageNil(b bool)`

 SetResolvedMessageNil sets the value for ResolvedMessage to be an explicit nil

### UnsetResolvedMessage
`func (o *GetGuidances200ResponseDiscoveryAnyOf) UnsetResolvedMessage()`

UnsetResolvedMessage ensures that no value is present for ResolvedMessage, not even an explicit nil
### GetRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetResource

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResource() GetGuidances200ResponseDiscoveryAnyOfResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetResourceOk() (*GetGuidances200ResponseDiscoveryAnyOfResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetResource(v GetGuidances200ResponseDiscoveryAnyOfResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPlanBeforeAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetPlanBeforeAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction`

GetPlanBeforeAction returns the PlanBeforeAction field if non-nil, zero value otherwise.

### GetPlanBeforeActionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetPlanBeforeActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction, bool)`

GetPlanBeforeActionOk returns a tuple with the PlanBeforeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBeforeAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetPlanBeforeAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction)`

SetPlanBeforeAction sets PlanBeforeAction field to given value.

### HasPlanBeforeAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasPlanBeforeAction() bool`

HasPlanBeforeAction returns a boolean if a field has been set.

### GetPlanAfterAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetPlanAfterAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

GetPlanAfterAction returns the PlanAfterAction field if non-nil, zero value otherwise.

### GetPlanAfterActionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetPlanAfterActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction, bool)`

GetPlanAfterActionOk returns a tuple with the PlanAfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAfterAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetPlanAfterAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction)`

SetPlanAfterAction sets PlanAfterAction field to given value.

### HasPlanAfterAction

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasPlanAfterAction() bool`

HasPlanAfterAction returns a boolean if a field has been set.

### GetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetGuidances200ResponseDiscoveryAnyOf) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


