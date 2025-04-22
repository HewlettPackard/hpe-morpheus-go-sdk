# GuidanceVmwareSizing

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

### NewGuidanceVmwareSizing

`func NewGuidanceVmwareSizing() *GuidanceVmwareSizing`

NewGuidanceVmwareSizing instantiates a new GuidanceVmwareSizing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceVmwareSizingWithDefaults

`func NewGuidanceVmwareSizingWithDefaults() *GuidanceVmwareSizing`

NewGuidanceVmwareSizingWithDefaults instantiates a new GuidanceVmwareSizing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuidanceVmwareSizing) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuidanceVmwareSizing) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuidanceVmwareSizing) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GuidanceVmwareSizing) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GuidanceVmwareSizing) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GuidanceVmwareSizing) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GuidanceVmwareSizing) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GuidanceVmwareSizing) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GuidanceVmwareSizing) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GuidanceVmwareSizing) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GuidanceVmwareSizing) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GuidanceVmwareSizing) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *GuidanceVmwareSizing) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *GuidanceVmwareSizing) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *GuidanceVmwareSizing) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *GuidanceVmwareSizing) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *GuidanceVmwareSizing) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *GuidanceVmwareSizing) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *GuidanceVmwareSizing) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *GuidanceVmwareSizing) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *GuidanceVmwareSizing) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *GuidanceVmwareSizing) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *GuidanceVmwareSizing) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *GuidanceVmwareSizing) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *GuidanceVmwareSizing) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GuidanceVmwareSizing) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GuidanceVmwareSizing) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *GuidanceVmwareSizing) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *GuidanceVmwareSizing) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *GuidanceVmwareSizing) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *GuidanceVmwareSizing) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *GuidanceVmwareSizing) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *GuidanceVmwareSizing) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *GuidanceVmwareSizing) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *GuidanceVmwareSizing) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *GuidanceVmwareSizing) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *GuidanceVmwareSizing) GetActionPlanId() int64`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *GuidanceVmwareSizing) GetActionPlanIdOk() (*int64, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *GuidanceVmwareSizing) SetActionPlanId(v int64)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *GuidanceVmwareSizing) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GuidanceVmwareSizing) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GuidanceVmwareSizing) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GuidanceVmwareSizing) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GuidanceVmwareSizing) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *GuidanceVmwareSizing) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GuidanceVmwareSizing) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GuidanceVmwareSizing) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GuidanceVmwareSizing) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *GuidanceVmwareSizing) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GuidanceVmwareSizing) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GuidanceVmwareSizing) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GuidanceVmwareSizing) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *GuidanceVmwareSizing) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GuidanceVmwareSizing) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GuidanceVmwareSizing) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GuidanceVmwareSizing) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *GuidanceVmwareSizing) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GuidanceVmwareSizing) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GuidanceVmwareSizing) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GuidanceVmwareSizing) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *GuidanceVmwareSizing) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GuidanceVmwareSizing) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GuidanceVmwareSizing) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GuidanceVmwareSizing) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *GuidanceVmwareSizing) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *GuidanceVmwareSizing) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *GuidanceVmwareSizing) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *GuidanceVmwareSizing) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *GuidanceVmwareSizing) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GuidanceVmwareSizing) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GuidanceVmwareSizing) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GuidanceVmwareSizing) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *GuidanceVmwareSizing) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GuidanceVmwareSizing) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GuidanceVmwareSizing) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GuidanceVmwareSizing) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *GuidanceVmwareSizing) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *GuidanceVmwareSizing) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *GuidanceVmwareSizing) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *GuidanceVmwareSizing) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### GetRefType

`func (o *GuidanceVmwareSizing) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GuidanceVmwareSizing) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GuidanceVmwareSizing) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GuidanceVmwareSizing) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GuidanceVmwareSizing) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GuidanceVmwareSizing) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GuidanceVmwareSizing) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GuidanceVmwareSizing) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GuidanceVmwareSizing) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GuidanceVmwareSizing) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GuidanceVmwareSizing) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GuidanceVmwareSizing) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *GuidanceVmwareSizing) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuidanceVmwareSizing) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuidanceVmwareSizing) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *GuidanceVmwareSizing) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *GuidanceVmwareSizing) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GuidanceVmwareSizing) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GuidanceVmwareSizing) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GuidanceVmwareSizing) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetResource

`func (o *GuidanceVmwareSizing) GetResource() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GuidanceVmwareSizing) GetResourceOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GuidanceVmwareSizing) SetResource(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GuidanceVmwareSizing) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPlanBeforeAction

`func (o *GuidanceVmwareSizing) GetPlanBeforeAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction`

GetPlanBeforeAction returns the PlanBeforeAction field if non-nil, zero value otherwise.

### GetPlanBeforeActionOk

`func (o *GuidanceVmwareSizing) GetPlanBeforeActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction, bool)`

GetPlanBeforeActionOk returns a tuple with the PlanBeforeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBeforeAction

`func (o *GuidanceVmwareSizing) SetPlanBeforeAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction)`

SetPlanBeforeAction sets PlanBeforeAction field to given value.

### HasPlanBeforeAction

`func (o *GuidanceVmwareSizing) HasPlanBeforeAction() bool`

HasPlanBeforeAction returns a boolean if a field has been set.

### GetPlanAfterAction

`func (o *GuidanceVmwareSizing) GetPlanAfterAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

GetPlanAfterAction returns the PlanAfterAction field if non-nil, zero value otherwise.

### GetPlanAfterActionOk

`func (o *GuidanceVmwareSizing) GetPlanAfterActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction, bool)`

GetPlanAfterActionOk returns a tuple with the PlanAfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAfterAction

`func (o *GuidanceVmwareSizing) SetPlanAfterAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction)`

SetPlanAfterAction sets PlanAfterAction field to given value.

### HasPlanAfterAction

`func (o *GuidanceVmwareSizing) HasPlanAfterAction() bool`

HasPlanAfterAction returns a boolean if a field has been set.

### GetConfig

`func (o *GuidanceVmwareSizing) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GuidanceVmwareSizing) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GuidanceVmwareSizing) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GuidanceVmwareSizing) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


