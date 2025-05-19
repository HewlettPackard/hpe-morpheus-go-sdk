# ListGuidances200ResponseAllOfDiscoveriesInner

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
**ActionPlanId** | Pointer to **string** |  | [optional] 
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
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config.md) |  | [optional] 

## Methods

### NewListGuidances200ResponseAllOfDiscoveriesInner

`func NewListGuidances200ResponseAllOfDiscoveriesInner() *ListGuidances200ResponseAllOfDiscoveriesInner`

NewListGuidances200ResponseAllOfDiscoveriesInner instantiates a new ListGuidances200ResponseAllOfDiscoveriesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseAllOfDiscoveriesInnerWithDefaults

`func NewListGuidances200ResponseAllOfDiscoveriesInnerWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInner`

NewListGuidances200ResponseAllOfDiscoveriesInnerWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionPlanId() string`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetActionPlanIdOk() (*string, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetActionPlanId(v string)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### GetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResource() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetResourceOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetResource(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetPlanBeforeAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction`

GetPlanBeforeAction returns the PlanBeforeAction field if non-nil, zero value otherwise.

### GetPlanBeforeActionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetPlanBeforeActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction, bool)`

GetPlanBeforeActionOk returns a tuple with the PlanBeforeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetPlanBeforeAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction)`

SetPlanBeforeAction sets PlanBeforeAction field to given value.

### HasPlanBeforeAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasPlanBeforeAction() bool`

HasPlanBeforeAction returns a boolean if a field has been set.

### GetPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetPlanAfterAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

GetPlanAfterAction returns the PlanAfterAction field if non-nil, zero value otherwise.

### GetPlanAfterActionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetPlanAfterActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction, bool)`

GetPlanAfterActionOk returns a tuple with the PlanAfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetPlanAfterAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction)`

SetPlanAfterAction sets PlanAfterAction field to given value.

### HasPlanAfterAction

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasPlanAfterAction() bool`

HasPlanAfterAction returns a boolean if a field has been set.

### GetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


