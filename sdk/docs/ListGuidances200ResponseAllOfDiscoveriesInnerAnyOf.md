# ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf

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
**UserId** | Pointer to **string** |  | [optional] 
**SiteId** | Pointer to **int64** |  | [optional] 
**Zone** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone.md) |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StateMessage** | Pointer to **string** |  | [optional] 
**Severity** | Pointer to **string** |  | [optional] 
**Resolved** | Pointer to **bool** |  | [optional] 
**ResolvedMessage** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**RefName** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType.md) |  | [optional] 
**Savings** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings.md) |  | [optional] 
**Resource** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource.md) |  | [optional] 
**PlanBeforeAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction.md) |  | [optional] 
**PlanAfterAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction.md) |  | [optional] 
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig.md) |  | [optional] 

## Methods

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOf

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOf() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOf instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfWithDefaults

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionPlanId() int64`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetActionPlanIdOk() (*int64, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetActionPlanId(v int64)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### GetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResource() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetResourceOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetResource(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetPlanBeforeAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction`

GetPlanBeforeAction returns the PlanBeforeAction field if non-nil, zero value otherwise.

### GetPlanBeforeActionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetPlanBeforeActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction, bool)`

GetPlanBeforeActionOk returns a tuple with the PlanBeforeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetPlanBeforeAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction)`

SetPlanBeforeAction sets PlanBeforeAction field to given value.

### HasPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasPlanBeforeAction() bool`

HasPlanBeforeAction returns a boolean if a field has been set.

### GetPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetPlanAfterAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

GetPlanAfterAction returns the PlanAfterAction field if non-nil, zero value otherwise.

### GetPlanAfterActionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetPlanAfterActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction, bool)`

GetPlanAfterActionOk returns a tuple with the PlanAfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetPlanAfterAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction)`

SetPlanAfterAction sets PlanAfterAction field to given value.

### HasPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasPlanAfterAction() bool`

HasPlanAfterAction returns a boolean if a field has been set.

### GetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


