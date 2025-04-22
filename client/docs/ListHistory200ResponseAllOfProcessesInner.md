# ListHistory200ResponseAllOfProcessesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner**](ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**SubId** | Pointer to **string** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**AppId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
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
**ResultType** | Pointer to **string** |  | [optional] 
**ResultId** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy**](GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy**](GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy.md) |  | [optional] 
**Events** | Pointer to [**[]ListHistory200ResponseAllOfProcessesInnerEventsInner**](ListHistory200ResponseAllOfProcessesInnerEventsInner.md) |  | [optional] 

## Methods

### NewListHistory200ResponseAllOfProcessesInner

`func NewListHistory200ResponseAllOfProcessesInner() *ListHistory200ResponseAllOfProcessesInner`

NewListHistory200ResponseAllOfProcessesInner instantiates a new ListHistory200ResponseAllOfProcessesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHistory200ResponseAllOfProcessesInnerWithDefaults

`func NewListHistory200ResponseAllOfProcessesInnerWithDefaults() *ListHistory200ResponseAllOfProcessesInner`

NewListHistory200ResponseAllOfProcessesInnerWithDefaults instantiates a new ListHistory200ResponseAllOfProcessesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *ListHistory200ResponseAllOfProcessesInner) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ListHistory200ResponseAllOfProcessesInner) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ListHistory200ResponseAllOfProcessesInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListHistory200ResponseAllOfProcessesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListHistory200ResponseAllOfProcessesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubType

`func (o *ListHistory200ResponseAllOfProcessesInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ListHistory200ResponseAllOfProcessesInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ListHistory200ResponseAllOfProcessesInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetAppId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *ListHistory200ResponseAllOfProcessesInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ListHistory200ResponseAllOfProcessesInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ListHistory200ResponseAllOfProcessesInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetStatus

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHistory200ResponseAllOfProcessesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHistory200ResponseAllOfProcessesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ListHistory200ResponseAllOfProcessesInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListHistory200ResponseAllOfProcessesInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListHistory200ResponseAllOfProcessesInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *ListHistory200ResponseAllOfProcessesInner) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ListHistory200ResponseAllOfProcessesInner) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ListHistory200ResponseAllOfProcessesInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *ListHistory200ResponseAllOfProcessesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListHistory200ResponseAllOfProcessesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListHistory200ResponseAllOfProcessesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *ListHistory200ResponseAllOfProcessesInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ListHistory200ResponseAllOfProcessesInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ListHistory200ResponseAllOfProcessesInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *ListHistory200ResponseAllOfProcessesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListHistory200ResponseAllOfProcessesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListHistory200ResponseAllOfProcessesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListHistory200ResponseAllOfProcessesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListHistory200ResponseAllOfProcessesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListHistory200ResponseAllOfProcessesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListHistory200ResponseAllOfProcessesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListHistory200ResponseAllOfProcessesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListHistory200ResponseAllOfProcessesInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ListHistory200ResponseAllOfProcessesInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResultType

`func (o *ListHistory200ResponseAllOfProcessesInner) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *ListHistory200ResponseAllOfProcessesInner) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *ListHistory200ResponseAllOfProcessesInner) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### GetResultId

`func (o *ListHistory200ResponseAllOfProcessesInner) GetResultId() int64`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetResultIdOk() (*int64, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *ListHistory200ResponseAllOfProcessesInner) SetResultId(v int64)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *ListHistory200ResponseAllOfProcessesInner) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *ListHistory200ResponseAllOfProcessesInner) GetEvents() []ListHistory200ResponseAllOfProcessesInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListHistory200ResponseAllOfProcessesInner) GetEventsOk() (*[]ListHistory200ResponseAllOfProcessesInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListHistory200ResponseAllOfProcessesInner) SetEvents(v []ListHistory200ResponseAllOfProcessesInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *ListHistory200ResponseAllOfProcessesInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


