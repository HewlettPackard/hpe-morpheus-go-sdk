# GetClusterHistoryEventDetail200ResponseProcessEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ProcessId** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**SubId** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
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

## Methods

### NewGetClusterHistoryEventDetail200ResponseProcessEvent

`func NewGetClusterHistoryEventDetail200ResponseProcessEvent() *GetClusterHistoryEventDetail200ResponseProcessEvent`

NewGetClusterHistoryEventDetail200ResponseProcessEvent instantiates a new GetClusterHistoryEventDetail200ResponseProcessEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterHistoryEventDetail200ResponseProcessEventWithDefaults

`func NewGetClusterHistoryEventDetail200ResponseProcessEventWithDefaults() *GetClusterHistoryEventDetail200ResponseProcessEvent`

NewGetClusterHistoryEventDetail200ResponseProcessEventWithDefaults instantiates a new GetClusterHistoryEventDetail200ResponseProcessEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRefType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetInstanceId() string`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetInstanceIdOk() (*string, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetInstanceId(v string)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetContainerId() string`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetContainerIdOk() (*string, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetContainerId(v string)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetClusterHistoryEventDetail200ResponseProcessEvent) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


