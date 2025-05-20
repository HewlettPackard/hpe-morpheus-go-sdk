# ListHistory200ResponseAllOfProcessesInnerEventsInner

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
**ZoneId** | Pointer to **int64** |  | [optional] 
**IntegrationId** | Pointer to **string** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
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

### NewListHistory200ResponseAllOfProcessesInnerEventsInner

`func NewListHistory200ResponseAllOfProcessesInnerEventsInner() *ListHistory200ResponseAllOfProcessesInnerEventsInner`

NewListHistory200ResponseAllOfProcessesInnerEventsInner instantiates a new ListHistory200ResponseAllOfProcessesInnerEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults

`func NewListHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults() *ListHistory200ResponseAllOfProcessesInnerEventsInner`

NewListHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults instantiates a new ListHistory200ResponseAllOfProcessesInnerEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetAccountId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRefType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ListHistory200ResponseAllOfProcessesInnerEventsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


