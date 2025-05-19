# GetClusterHistoryDetail200ResponseProcess

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

### NewGetClusterHistoryDetail200ResponseProcess

`func NewGetClusterHistoryDetail200ResponseProcess() *GetClusterHistoryDetail200ResponseProcess`

NewGetClusterHistoryDetail200ResponseProcess instantiates a new GetClusterHistoryDetail200ResponseProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterHistoryDetail200ResponseProcessWithDefaults

`func NewGetClusterHistoryDetail200ResponseProcessWithDefaults() *GetClusterHistoryDetail200ResponseProcess`

NewGetClusterHistoryDetail200ResponseProcessWithDefaults instantiates a new GetClusterHistoryDetail200ResponseProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetClusterHistoryDetail200ResponseProcess) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetClusterHistoryDetail200ResponseProcess) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetClusterHistoryDetail200ResponseProcess) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterHistoryDetail200ResponseProcess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterHistoryDetail200ResponseProcess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubType

`func (o *GetClusterHistoryDetail200ResponseProcess) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetClusterHistoryDetail200ResponseProcess) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetClusterHistoryDetail200ResponseProcess) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetAppId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetClusterHistoryDetail200ResponseProcess) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetClusterHistoryDetail200ResponseProcess) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetClusterHistoryDetail200ResponseProcess) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetClusterHistoryDetail200ResponseProcess) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetClusterHistoryDetail200ResponseProcess) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetClusterHistoryDetail200ResponseProcess) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetClusterHistoryDetail200ResponseProcess) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetClusterHistoryDetail200ResponseProcess) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterHistoryDetail200ResponseProcess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterHistoryDetail200ResponseProcess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterHistoryDetail200ResponseProcess) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterHistoryDetail200ResponseProcess) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterHistoryDetail200ResponseProcess) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *GetClusterHistoryDetail200ResponseProcess) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetClusterHistoryDetail200ResponseProcess) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetClusterHistoryDetail200ResponseProcess) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetClusterHistoryDetail200ResponseProcess) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetClusterHistoryDetail200ResponseProcess) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetClusterHistoryDetail200ResponseProcess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetClusterHistoryDetail200ResponseProcess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetClusterHistoryDetail200ResponseProcess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *GetClusterHistoryDetail200ResponseProcess) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetClusterHistoryDetail200ResponseProcess) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetClusterHistoryDetail200ResponseProcess) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *GetClusterHistoryDetail200ResponseProcess) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetClusterHistoryDetail200ResponseProcess) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetClusterHistoryDetail200ResponseProcess) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetClusterHistoryDetail200ResponseProcess) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetClusterHistoryDetail200ResponseProcess) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetClusterHistoryDetail200ResponseProcess) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetClusterHistoryDetail200ResponseProcess) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetClusterHistoryDetail200ResponseProcess) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetClusterHistoryDetail200ResponseProcess) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetClusterHistoryDetail200ResponseProcess) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterHistoryDetail200ResponseProcess) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterHistoryDetail200ResponseProcess) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterHistoryDetail200ResponseProcess) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterHistoryDetail200ResponseProcess) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterHistoryDetail200ResponseProcess) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetClusterHistoryDetail200ResponseProcess) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *GetClusterHistoryDetail200ResponseProcess) GetEvents() []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetClusterHistoryDetail200ResponseProcess) GetEventsOk() (*[]GetClusterHistory200ResponseAllOfProcessesInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetClusterHistoryDetail200ResponseProcess) SetEvents(v []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetClusterHistoryDetail200ResponseProcess) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


