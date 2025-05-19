# ClusterHistoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**SubId** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **string** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Percent** | Pointer to **int64** |  | [optional] 
**StatusEta** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Output** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy**](GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy**](GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy.md) |  | [optional] 
**Events** | Pointer to [**[]GetClusterHistory200ResponseAllOfProcessesInnerEventsInner**](GetClusterHistory200ResponseAllOfProcessesInnerEventsInner.md) |  | [optional] 

## Methods

### NewClusterHistoryItem

`func NewClusterHistoryItem() *ClusterHistoryItem`

NewClusterHistoryItem instantiates a new ClusterHistoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHistoryItemWithDefaults

`func NewClusterHistoryItemWithDefaults() *ClusterHistoryItem`

NewClusterHistoryItemWithDefaults instantiates a new ClusterHistoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterHistoryItem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterHistoryItem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterHistoryItem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterHistoryItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ClusterHistoryItem) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClusterHistoryItem) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClusterHistoryItem) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClusterHistoryItem) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ClusterHistoryItem) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ClusterHistoryItem) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ClusterHistoryItem) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ClusterHistoryItem) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *ClusterHistoryItem) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ClusterHistoryItem) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ClusterHistoryItem) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ClusterHistoryItem) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterHistoryItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterHistoryItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterHistoryItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterHistoryItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubType

`func (o *ClusterHistoryItem) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ClusterHistoryItem) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ClusterHistoryItem) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ClusterHistoryItem) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *ClusterHistoryItem) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ClusterHistoryItem) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ClusterHistoryItem) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *ClusterHistoryItem) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *ClusterHistoryItem) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ClusterHistoryItem) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ClusterHistoryItem) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ClusterHistoryItem) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ClusterHistoryItem) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ClusterHistoryItem) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ClusterHistoryItem) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ClusterHistoryItem) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetAppId

`func (o *ClusterHistoryItem) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ClusterHistoryItem) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ClusterHistoryItem) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ClusterHistoryItem) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ClusterHistoryItem) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClusterHistoryItem) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClusterHistoryItem) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClusterHistoryItem) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *ClusterHistoryItem) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ClusterHistoryItem) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ClusterHistoryItem) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ClusterHistoryItem) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *ClusterHistoryItem) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ClusterHistoryItem) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ClusterHistoryItem) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ClusterHistoryItem) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *ClusterHistoryItem) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ClusterHistoryItem) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ClusterHistoryItem) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ClusterHistoryItem) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ClusterHistoryItem) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterHistoryItem) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterHistoryItem) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterHistoryItem) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterHistoryItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterHistoryItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterHistoryItem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterHistoryItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ClusterHistoryItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterHistoryItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterHistoryItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterHistoryItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *ClusterHistoryItem) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ClusterHistoryItem) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ClusterHistoryItem) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ClusterHistoryItem) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ClusterHistoryItem) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ClusterHistoryItem) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ClusterHistoryItem) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ClusterHistoryItem) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterHistoryItem) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterHistoryItem) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterHistoryItem) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterHistoryItem) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *ClusterHistoryItem) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ClusterHistoryItem) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ClusterHistoryItem) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ClusterHistoryItem) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *ClusterHistoryItem) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterHistoryItem) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterHistoryItem) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ClusterHistoryItem) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *ClusterHistoryItem) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClusterHistoryItem) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClusterHistoryItem) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClusterHistoryItem) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClusterHistoryItem) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClusterHistoryItem) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClusterHistoryItem) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClusterHistoryItem) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ClusterHistoryItem) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClusterHistoryItem) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClusterHistoryItem) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClusterHistoryItem) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterHistoryItem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterHistoryItem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterHistoryItem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterHistoryItem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterHistoryItem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterHistoryItem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterHistoryItem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterHistoryItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClusterHistoryItem) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClusterHistoryItem) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClusterHistoryItem) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClusterHistoryItem) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ClusterHistoryItem) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClusterHistoryItem) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClusterHistoryItem) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ClusterHistoryItem) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *ClusterHistoryItem) GetEvents() []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ClusterHistoryItem) GetEventsOk() (*[]GetClusterHistory200ResponseAllOfProcessesInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ClusterHistoryItem) SetEvents(v []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ClusterHistoryItem) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


