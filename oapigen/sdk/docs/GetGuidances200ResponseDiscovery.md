# GetGuidances200ResponseDiscovery

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
**Zone** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone.md) |  | [optional] 
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
**Resource** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOfResource**](GetGuidances200ResponseDiscoveryAnyOfResource.md) |  | [optional] 
**PlanBeforeAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction.md) |  | [optional] 
**PlanAfterAction** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction.md) |  | [optional] 
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config.md) |  | [optional] 

## Methods

### NewGetGuidances200ResponseDiscovery

`func NewGetGuidances200ResponseDiscovery() *GetGuidances200ResponseDiscovery`

NewGetGuidances200ResponseDiscovery instantiates a new GetGuidances200ResponseDiscovery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidances200ResponseDiscoveryWithDefaults

`func NewGetGuidances200ResponseDiscoveryWithDefaults() *GetGuidances200ResponseDiscovery`

NewGetGuidances200ResponseDiscoveryWithDefaults instantiates a new GetGuidances200ResponseDiscovery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuidances200ResponseDiscovery) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuidances200ResponseDiscovery) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuidances200ResponseDiscovery) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGuidances200ResponseDiscovery) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetGuidances200ResponseDiscovery) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGuidances200ResponseDiscovery) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGuidances200ResponseDiscovery) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGuidances200ResponseDiscovery) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGuidances200ResponseDiscovery) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGuidances200ResponseDiscovery) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGuidances200ResponseDiscovery) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGuidances200ResponseDiscovery) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *GetGuidances200ResponseDiscovery) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *GetGuidances200ResponseDiscovery) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *GetGuidances200ResponseDiscovery) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *GetGuidances200ResponseDiscovery) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *GetGuidances200ResponseDiscovery) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *GetGuidances200ResponseDiscovery) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *GetGuidances200ResponseDiscovery) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *GetGuidances200ResponseDiscovery) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *GetGuidances200ResponseDiscovery) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *GetGuidances200ResponseDiscovery) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *GetGuidances200ResponseDiscovery) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *GetGuidances200ResponseDiscovery) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *GetGuidances200ResponseDiscovery) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GetGuidances200ResponseDiscovery) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GetGuidances200ResponseDiscovery) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *GetGuidances200ResponseDiscovery) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *GetGuidances200ResponseDiscovery) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *GetGuidances200ResponseDiscovery) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *GetGuidances200ResponseDiscovery) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *GetGuidances200ResponseDiscovery) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *GetGuidances200ResponseDiscovery) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *GetGuidances200ResponseDiscovery) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *GetGuidances200ResponseDiscovery) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *GetGuidances200ResponseDiscovery) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *GetGuidances200ResponseDiscovery) GetActionPlanId() string`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *GetGuidances200ResponseDiscovery) GetActionPlanIdOk() (*string, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *GetGuidances200ResponseDiscovery) SetActionPlanId(v string)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *GetGuidances200ResponseDiscovery) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetGuidances200ResponseDiscovery) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetGuidances200ResponseDiscovery) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetGuidances200ResponseDiscovery) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetGuidances200ResponseDiscovery) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGuidances200ResponseDiscovery) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGuidances200ResponseDiscovery) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGuidances200ResponseDiscovery) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGuidances200ResponseDiscovery) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *GetGuidances200ResponseDiscovery) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetGuidances200ResponseDiscovery) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetGuidances200ResponseDiscovery) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetGuidances200ResponseDiscovery) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *GetGuidances200ResponseDiscovery) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GetGuidances200ResponseDiscovery) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GetGuidances200ResponseDiscovery) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GetGuidances200ResponseDiscovery) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *GetGuidances200ResponseDiscovery) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GetGuidances200ResponseDiscovery) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GetGuidances200ResponseDiscovery) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Zone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GetGuidances200ResponseDiscovery) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *GetGuidances200ResponseDiscovery) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetGuidances200ResponseDiscovery) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetGuidances200ResponseDiscovery) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GetGuidances200ResponseDiscovery) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *GetGuidances200ResponseDiscovery) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *GetGuidances200ResponseDiscovery) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *GetGuidances200ResponseDiscovery) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *GetGuidances200ResponseDiscovery) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *GetGuidances200ResponseDiscovery) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetGuidances200ResponseDiscovery) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetGuidances200ResponseDiscovery) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetGuidances200ResponseDiscovery) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *GetGuidances200ResponseDiscovery) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GetGuidances200ResponseDiscovery) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GetGuidances200ResponseDiscovery) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GetGuidances200ResponseDiscovery) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *GetGuidances200ResponseDiscovery) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *GetGuidances200ResponseDiscovery) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *GetGuidances200ResponseDiscovery) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *GetGuidances200ResponseDiscovery) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### GetRefType

`func (o *GetGuidances200ResponseDiscovery) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetGuidances200ResponseDiscovery) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetGuidances200ResponseDiscovery) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetGuidances200ResponseDiscovery) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetGuidances200ResponseDiscovery) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetGuidances200ResponseDiscovery) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetGuidances200ResponseDiscovery) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetGuidances200ResponseDiscovery) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GetGuidances200ResponseDiscovery) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GetGuidances200ResponseDiscovery) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GetGuidances200ResponseDiscovery) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GetGuidances200ResponseDiscovery) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *GetGuidances200ResponseDiscovery) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGuidances200ResponseDiscovery) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGuidances200ResponseDiscovery) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *GetGuidances200ResponseDiscovery) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *GetGuidances200ResponseDiscovery) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GetGuidances200ResponseDiscovery) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GetGuidances200ResponseDiscovery) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GetGuidances200ResponseDiscovery) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetResource

`func (o *GetGuidances200ResponseDiscovery) GetResource() GetGuidances200ResponseDiscoveryAnyOfResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *GetGuidances200ResponseDiscovery) GetResourceOk() (*GetGuidances200ResponseDiscoveryAnyOfResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *GetGuidances200ResponseDiscovery) SetResource(v GetGuidances200ResponseDiscoveryAnyOfResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *GetGuidances200ResponseDiscovery) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetPlanBeforeAction

`func (o *GetGuidances200ResponseDiscovery) GetPlanBeforeAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction`

GetPlanBeforeAction returns the PlanBeforeAction field if non-nil, zero value otherwise.

### GetPlanBeforeActionOk

`func (o *GetGuidances200ResponseDiscovery) GetPlanBeforeActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction, bool)`

GetPlanBeforeActionOk returns a tuple with the PlanBeforeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanBeforeAction

`func (o *GetGuidances200ResponseDiscovery) SetPlanBeforeAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeAction)`

SetPlanBeforeAction sets PlanBeforeAction field to given value.

### HasPlanBeforeAction

`func (o *GetGuidances200ResponseDiscovery) HasPlanBeforeAction() bool`

HasPlanBeforeAction returns a boolean if a field has been set.

### GetPlanAfterAction

`func (o *GetGuidances200ResponseDiscovery) GetPlanAfterAction() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

GetPlanAfterAction returns the PlanAfterAction field if non-nil, zero value otherwise.

### GetPlanAfterActionOk

`func (o *GetGuidances200ResponseDiscovery) GetPlanAfterActionOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction, bool)`

GetPlanAfterActionOk returns a tuple with the PlanAfterAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanAfterAction

`func (o *GetGuidances200ResponseDiscovery) SetPlanAfterAction(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction)`

SetPlanAfterAction sets PlanAfterAction field to given value.

### HasPlanAfterAction

`func (o *GetGuidances200ResponseDiscovery) HasPlanAfterAction() bool`

HasPlanAfterAction returns a boolean if a field has been set.

### GetConfig

`func (o *GetGuidances200ResponseDiscovery) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetGuidances200ResponseDiscovery) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetGuidances200ResponseDiscovery) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetGuidances200ResponseDiscovery) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


