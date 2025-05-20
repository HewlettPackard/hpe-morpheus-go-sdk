# GetClusterHistory200ResponseAllOfProcessesInner

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
**IntegrationId** | Pointer to **int64** |  | [optional] 
**AppId** | Pointer to **int64** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Percent** | Pointer to **float32** |  | [optional] 
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

### NewGetClusterHistory200ResponseAllOfProcessesInner

`func NewGetClusterHistory200ResponseAllOfProcessesInner() *GetClusterHistory200ResponseAllOfProcessesInner`

NewGetClusterHistory200ResponseAllOfProcessesInner instantiates a new GetClusterHistory200ResponseAllOfProcessesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterHistory200ResponseAllOfProcessesInnerWithDefaults

`func NewGetClusterHistory200ResponseAllOfProcessesInnerWithDefaults() *GetClusterHistory200ResponseAllOfProcessesInner`

NewGetClusterHistory200ResponseAllOfProcessesInnerWithDefaults instantiates a new GetClusterHistory200ResponseAllOfProcessesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetProcessType() ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetProcessTypeOk() (*ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetProcessType(v ListClusterContainers200ResponseAllOfContainersInnerAvailableActionsInner)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### GetAppId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetAppId() int64`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetAppIdOk() (*int64, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetAppId(v int64)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### GetInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetPercent() float32`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetPercentOk() (*float32, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetPercent(v float32)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### GetError

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetCreatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetCreatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetCreatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetUpdatedBy() GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetUpdatedByOk() (*GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetUpdatedBy(v GetClusterHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetEvents() []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) GetEventsOk() (*[]GetClusterHistory200ResponseAllOfProcessesInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) SetEvents(v []GetClusterHistory200ResponseAllOfProcessesInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetClusterHistory200ResponseAllOfProcessesInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


