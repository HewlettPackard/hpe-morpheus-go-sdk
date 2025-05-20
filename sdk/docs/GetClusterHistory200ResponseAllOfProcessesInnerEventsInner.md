# GetClusterHistory200ResponseAllOfProcessesInnerEventsInner

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
**SubId** | Pointer to **int64** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**IntegrationId** | Pointer to **int64** |  | [optional] 
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

### NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInner

`func NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInner() *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner`

NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInner instantiates a new GetClusterHistory200ResponseAllOfProcessesInnerEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults

`func NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults() *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner`

NewGetClusterHistory200ResponseAllOfProcessesInnerEventsInnerWithDefaults instantiates a new GetClusterHistory200ResponseAllOfProcessesInnerEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRefType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetSubId() int64`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetSubIdOk() (*int64, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetSubId(v int64)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInnerEventsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


