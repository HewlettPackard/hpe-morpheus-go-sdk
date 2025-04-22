# GuidanceAzureReservations

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
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config.md) |  | [optional] 

## Methods

### NewGuidanceAzureReservations

`func NewGuidanceAzureReservations() *GuidanceAzureReservations`

NewGuidanceAzureReservations instantiates a new GuidanceAzureReservations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceAzureReservationsWithDefaults

`func NewGuidanceAzureReservationsWithDefaults() *GuidanceAzureReservations`

NewGuidanceAzureReservationsWithDefaults instantiates a new GuidanceAzureReservations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuidanceAzureReservations) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuidanceAzureReservations) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuidanceAzureReservations) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GuidanceAzureReservations) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDateCreated

`func (o *GuidanceAzureReservations) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GuidanceAzureReservations) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GuidanceAzureReservations) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GuidanceAzureReservations) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GuidanceAzureReservations) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GuidanceAzureReservations) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GuidanceAzureReservations) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GuidanceAzureReservations) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetActionCategory

`func (o *GuidanceAzureReservations) GetActionCategory() string`

GetActionCategory returns the ActionCategory field if non-nil, zero value otherwise.

### GetActionCategoryOk

`func (o *GuidanceAzureReservations) GetActionCategoryOk() (*string, bool)`

GetActionCategoryOk returns a tuple with the ActionCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionCategory

`func (o *GuidanceAzureReservations) SetActionCategory(v string)`

SetActionCategory sets ActionCategory field to given value.

### HasActionCategory

`func (o *GuidanceAzureReservations) HasActionCategory() bool`

HasActionCategory returns a boolean if a field has been set.

### GetActionMessage

`func (o *GuidanceAzureReservations) GetActionMessage() string`

GetActionMessage returns the ActionMessage field if non-nil, zero value otherwise.

### GetActionMessageOk

`func (o *GuidanceAzureReservations) GetActionMessageOk() (*string, bool)`

GetActionMessageOk returns a tuple with the ActionMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionMessage

`func (o *GuidanceAzureReservations) SetActionMessage(v string)`

SetActionMessage sets ActionMessage field to given value.

### HasActionMessage

`func (o *GuidanceAzureReservations) HasActionMessage() bool`

HasActionMessage returns a boolean if a field has been set.

### GetActionTitle

`func (o *GuidanceAzureReservations) GetActionTitle() string`

GetActionTitle returns the ActionTitle field if non-nil, zero value otherwise.

### GetActionTitleOk

`func (o *GuidanceAzureReservations) GetActionTitleOk() (*string, bool)`

GetActionTitleOk returns a tuple with the ActionTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionTitle

`func (o *GuidanceAzureReservations) SetActionTitle(v string)`

SetActionTitle sets ActionTitle field to given value.

### HasActionTitle

`func (o *GuidanceAzureReservations) HasActionTitle() bool`

HasActionTitle returns a boolean if a field has been set.

### GetActionType

`func (o *GuidanceAzureReservations) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GuidanceAzureReservations) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GuidanceAzureReservations) SetActionType(v string)`

SetActionType sets ActionType field to given value.

### HasActionType

`func (o *GuidanceAzureReservations) HasActionType() bool`

HasActionType returns a boolean if a field has been set.

### GetActionValue

`func (o *GuidanceAzureReservations) GetActionValue() string`

GetActionValue returns the ActionValue field if non-nil, zero value otherwise.

### GetActionValueOk

`func (o *GuidanceAzureReservations) GetActionValueOk() (*string, bool)`

GetActionValueOk returns a tuple with the ActionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValue

`func (o *GuidanceAzureReservations) SetActionValue(v string)`

SetActionValue sets ActionValue field to given value.

### HasActionValue

`func (o *GuidanceAzureReservations) HasActionValue() bool`

HasActionValue returns a boolean if a field has been set.

### GetActionValueType

`func (o *GuidanceAzureReservations) GetActionValueType() string`

GetActionValueType returns the ActionValueType field if non-nil, zero value otherwise.

### GetActionValueTypeOk

`func (o *GuidanceAzureReservations) GetActionValueTypeOk() (*string, bool)`

GetActionValueTypeOk returns a tuple with the ActionValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValueType

`func (o *GuidanceAzureReservations) SetActionValueType(v string)`

SetActionValueType sets ActionValueType field to given value.

### HasActionValueType

`func (o *GuidanceAzureReservations) HasActionValueType() bool`

HasActionValueType returns a boolean if a field has been set.

### GetActionPlanId

`func (o *GuidanceAzureReservations) GetActionPlanId() string`

GetActionPlanId returns the ActionPlanId field if non-nil, zero value otherwise.

### GetActionPlanIdOk

`func (o *GuidanceAzureReservations) GetActionPlanIdOk() (*string, bool)`

GetActionPlanIdOk returns a tuple with the ActionPlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionPlanId

`func (o *GuidanceAzureReservations) SetActionPlanId(v string)`

SetActionPlanId sets ActionPlanId field to given value.

### HasActionPlanId

`func (o *GuidanceAzureReservations) HasActionPlanId() bool`

HasActionPlanId returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GuidanceAzureReservations) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GuidanceAzureReservations) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GuidanceAzureReservations) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GuidanceAzureReservations) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetAccountId

`func (o *GuidanceAzureReservations) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GuidanceAzureReservations) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GuidanceAzureReservations) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GuidanceAzureReservations) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUserId

`func (o *GuidanceAzureReservations) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GuidanceAzureReservations) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GuidanceAzureReservations) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GuidanceAzureReservations) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSiteId

`func (o *GuidanceAzureReservations) GetSiteId() int64`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GuidanceAzureReservations) GetSiteIdOk() (*int64, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GuidanceAzureReservations) SetSiteId(v int64)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GuidanceAzureReservations) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetZone

`func (o *GuidanceAzureReservations) GetZone() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *GuidanceAzureReservations) GetZoneOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *GuidanceAzureReservations) SetZone(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfZone)`

SetZone sets Zone field to given value.

### HasZone

`func (o *GuidanceAzureReservations) HasZone() bool`

HasZone returns a boolean if a field has been set.

### GetState

`func (o *GuidanceAzureReservations) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GuidanceAzureReservations) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GuidanceAzureReservations) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *GuidanceAzureReservations) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateMessage

`func (o *GuidanceAzureReservations) GetStateMessage() string`

GetStateMessage returns the StateMessage field if non-nil, zero value otherwise.

### GetStateMessageOk

`func (o *GuidanceAzureReservations) GetStateMessageOk() (*string, bool)`

GetStateMessageOk returns a tuple with the StateMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateMessage

`func (o *GuidanceAzureReservations) SetStateMessage(v string)`

SetStateMessage sets StateMessage field to given value.

### HasStateMessage

`func (o *GuidanceAzureReservations) HasStateMessage() bool`

HasStateMessage returns a boolean if a field has been set.

### GetSeverity

`func (o *GuidanceAzureReservations) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GuidanceAzureReservations) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GuidanceAzureReservations) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GuidanceAzureReservations) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetResolved

`func (o *GuidanceAzureReservations) GetResolved() bool`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *GuidanceAzureReservations) GetResolvedOk() (*bool, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *GuidanceAzureReservations) SetResolved(v bool)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *GuidanceAzureReservations) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### GetResolvedMessage

`func (o *GuidanceAzureReservations) GetResolvedMessage() string`

GetResolvedMessage returns the ResolvedMessage field if non-nil, zero value otherwise.

### GetResolvedMessageOk

`func (o *GuidanceAzureReservations) GetResolvedMessageOk() (*string, bool)`

GetResolvedMessageOk returns a tuple with the ResolvedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedMessage

`func (o *GuidanceAzureReservations) SetResolvedMessage(v string)`

SetResolvedMessage sets ResolvedMessage field to given value.

### HasResolvedMessage

`func (o *GuidanceAzureReservations) HasResolvedMessage() bool`

HasResolvedMessage returns a boolean if a field has been set.

### GetRefType

`func (o *GuidanceAzureReservations) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GuidanceAzureReservations) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GuidanceAzureReservations) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GuidanceAzureReservations) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GuidanceAzureReservations) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GuidanceAzureReservations) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GuidanceAzureReservations) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GuidanceAzureReservations) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetRefName

`func (o *GuidanceAzureReservations) GetRefName() string`

GetRefName returns the RefName field if non-nil, zero value otherwise.

### GetRefNameOk

`func (o *GuidanceAzureReservations) GetRefNameOk() (*string, bool)`

GetRefNameOk returns a tuple with the RefName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefName

`func (o *GuidanceAzureReservations) SetRefName(v string)`

SetRefName sets RefName field to given value.

### HasRefName

`func (o *GuidanceAzureReservations) HasRefName() bool`

HasRefName returns a boolean if a field has been set.

### GetType

`func (o *GuidanceAzureReservations) GetType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GuidanceAzureReservations) GetTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GuidanceAzureReservations) SetType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfType)`

SetType sets Type field to given value.

### HasType

`func (o *GuidanceAzureReservations) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSavings

`func (o *GuidanceAzureReservations) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GuidanceAzureReservations) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GuidanceAzureReservations) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GuidanceAzureReservations) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetConfig

`func (o *GuidanceAzureReservations) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GuidanceAzureReservations) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GuidanceAzureReservations) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOf1Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GuidanceAzureReservations) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


