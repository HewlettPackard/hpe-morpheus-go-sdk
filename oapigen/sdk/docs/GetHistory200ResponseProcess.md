# GetHistory200ResponseProcess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**GetHistory200ResponseProcessProcessType**](GetHistory200ResponseProcessProcessType.md) |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**SubType** | Pointer to **NullableString** |  | [optional] 
**SubId** | Pointer to **NullableString** |  | [optional] 
**ZoneId** | Pointer to **int64** |  | [optional] 
**IntegrationId** | Pointer to **NullableString** |  | [optional] 
**AppId** | Pointer to **NullableString** |  | [optional] 
**InstanceId** | Pointer to **int64** |  | [optional] 
**ContainerId** | Pointer to **int64** |  | [optional] 
**ServerId** | Pointer to **int64** |  | [optional] 
**ContainerName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Percent** | Pointer to **float64** |  | [optional] 
**StatusEta** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Output** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**ResultType** | Pointer to **NullableString** |  | [optional] 
**ResultId** | Pointer to **NullableInt64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**GetHistory200ResponseProcessCreatedBy**](GetHistory200ResponseProcessCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**GetHistory200ResponseProcessUpdatedBy**](GetHistory200ResponseProcessUpdatedBy.md) |  | [optional] 
**Events** | Pointer to [**[]GetHistory200ResponseProcessEventsInner**](GetHistory200ResponseProcessEventsInner.md) |  | [optional] 

## Methods

### NewGetHistory200ResponseProcess

`func NewGetHistory200ResponseProcess() *GetHistory200ResponseProcess`

NewGetHistory200ResponseProcess instantiates a new GetHistory200ResponseProcess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHistory200ResponseProcessWithDefaults

`func NewGetHistory200ResponseProcessWithDefaults() *GetHistory200ResponseProcess`

NewGetHistory200ResponseProcessWithDefaults instantiates a new GetHistory200ResponseProcess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetHistory200ResponseProcess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetHistory200ResponseProcess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetHistory200ResponseProcess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetHistory200ResponseProcess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetHistory200ResponseProcess) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetHistory200ResponseProcess) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetHistory200ResponseProcess) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetHistory200ResponseProcess) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetHistory200ResponseProcess) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetHistory200ResponseProcess) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetHistory200ResponseProcess) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetHistory200ResponseProcess) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetHistory200ResponseProcess) GetProcessType() GetHistory200ResponseProcessProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetHistory200ResponseProcess) GetProcessTypeOk() (*GetHistory200ResponseProcessProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetHistory200ResponseProcess) SetProcessType(v GetHistory200ResponseProcessProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetHistory200ResponseProcess) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetHistory200ResponseProcess) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetHistory200ResponseProcess) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetHistory200ResponseProcess) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetHistory200ResponseProcess) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetHistory200ResponseProcess) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetHistory200ResponseProcess) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetHistory200ResponseProcess) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetHistory200ResponseProcess) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetHistory200ResponseProcess) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetHistory200ResponseProcess) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubType

`func (o *GetHistory200ResponseProcess) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetHistory200ResponseProcess) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetHistory200ResponseProcess) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetHistory200ResponseProcess) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *GetHistory200ResponseProcess) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *GetHistory200ResponseProcess) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSubId

`func (o *GetHistory200ResponseProcess) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetHistory200ResponseProcess) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetHistory200ResponseProcess) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetHistory200ResponseProcess) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### SetSubIdNil

`func (o *GetHistory200ResponseProcess) SetSubIdNil(b bool)`

 SetSubIdNil sets the value for SubId to be an explicit nil

### UnsetSubId
`func (o *GetHistory200ResponseProcess) UnsetSubId()`

UnsetSubId ensures that no value is present for SubId, not even an explicit nil
### GetZoneId

`func (o *GetHistory200ResponseProcess) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetHistory200ResponseProcess) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetHistory200ResponseProcess) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetHistory200ResponseProcess) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetHistory200ResponseProcess) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetHistory200ResponseProcess) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetHistory200ResponseProcess) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetHistory200ResponseProcess) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *GetHistory200ResponseProcess) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *GetHistory200ResponseProcess) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetAppId

`func (o *GetHistory200ResponseProcess) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetHistory200ResponseProcess) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetHistory200ResponseProcess) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetHistory200ResponseProcess) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *GetHistory200ResponseProcess) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *GetHistory200ResponseProcess) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetInstanceId

`func (o *GetHistory200ResponseProcess) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetHistory200ResponseProcess) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetHistory200ResponseProcess) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetHistory200ResponseProcess) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetHistory200ResponseProcess) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetHistory200ResponseProcess) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetHistory200ResponseProcess) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetHistory200ResponseProcess) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetHistory200ResponseProcess) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetHistory200ResponseProcess) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetHistory200ResponseProcess) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetHistory200ResponseProcess) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetHistory200ResponseProcess) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetHistory200ResponseProcess) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetHistory200ResponseProcess) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetHistory200ResponseProcess) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetStatus

`func (o *GetHistory200ResponseProcess) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetHistory200ResponseProcess) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetHistory200ResponseProcess) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetHistory200ResponseProcess) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetHistory200ResponseProcess) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetHistory200ResponseProcess) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetHistory200ResponseProcess) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetHistory200ResponseProcess) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GetHistory200ResponseProcess) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetHistory200ResponseProcess) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetPercent

`func (o *GetHistory200ResponseProcess) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetHistory200ResponseProcess) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetHistory200ResponseProcess) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetHistory200ResponseProcess) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetHistory200ResponseProcess) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetHistory200ResponseProcess) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetHistory200ResponseProcess) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetHistory200ResponseProcess) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetHistory200ResponseProcess) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetHistory200ResponseProcess) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetHistory200ResponseProcess) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetHistory200ResponseProcess) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetHistory200ResponseProcess) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetHistory200ResponseProcess) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOutput

`func (o *GetHistory200ResponseProcess) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetHistory200ResponseProcess) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetHistory200ResponseProcess) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetHistory200ResponseProcess) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *GetHistory200ResponseProcess) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *GetHistory200ResponseProcess) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetError

`func (o *GetHistory200ResponseProcess) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetHistory200ResponseProcess) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetHistory200ResponseProcess) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetHistory200ResponseProcess) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetHistory200ResponseProcess) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetHistory200ResponseProcess) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStartDate

`func (o *GetHistory200ResponseProcess) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetHistory200ResponseProcess) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetHistory200ResponseProcess) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetHistory200ResponseProcess) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetHistory200ResponseProcess) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetHistory200ResponseProcess) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetHistory200ResponseProcess) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetHistory200ResponseProcess) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetHistory200ResponseProcess) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetHistory200ResponseProcess) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetHistory200ResponseProcess) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetHistory200ResponseProcess) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResultType

`func (o *GetHistory200ResponseProcess) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *GetHistory200ResponseProcess) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *GetHistory200ResponseProcess) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *GetHistory200ResponseProcess) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *GetHistory200ResponseProcess) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *GetHistory200ResponseProcess) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetResultId

`func (o *GetHistory200ResponseProcess) GetResultId() int64`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *GetHistory200ResponseProcess) GetResultIdOk() (*int64, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *GetHistory200ResponseProcess) SetResultId(v int64)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *GetHistory200ResponseProcess) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### SetResultIdNil

`func (o *GetHistory200ResponseProcess) SetResultIdNil(b bool)`

 SetResultIdNil sets the value for ResultId to be an explicit nil

### UnsetResultId
`func (o *GetHistory200ResponseProcess) UnsetResultId()`

UnsetResultId ensures that no value is present for ResultId, not even an explicit nil
### GetDateCreated

`func (o *GetHistory200ResponseProcess) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetHistory200ResponseProcess) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetHistory200ResponseProcess) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetHistory200ResponseProcess) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetHistory200ResponseProcess) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetHistory200ResponseProcess) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetHistory200ResponseProcess) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetHistory200ResponseProcess) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetHistory200ResponseProcess) GetCreatedBy() GetHistory200ResponseProcessCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetHistory200ResponseProcess) GetCreatedByOk() (*GetHistory200ResponseProcessCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetHistory200ResponseProcess) SetCreatedBy(v GetHistory200ResponseProcessCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetHistory200ResponseProcess) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetHistory200ResponseProcess) GetUpdatedBy() GetHistory200ResponseProcessUpdatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetHistory200ResponseProcess) GetUpdatedByOk() (*GetHistory200ResponseProcessUpdatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetHistory200ResponseProcess) SetUpdatedBy(v GetHistory200ResponseProcessUpdatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetHistory200ResponseProcess) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *GetHistory200ResponseProcess) GetEvents() []GetHistory200ResponseProcessEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetHistory200ResponseProcess) GetEventsOk() (*[]GetHistory200ResponseProcessEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetHistory200ResponseProcess) SetEvents(v []GetHistory200ResponseProcessEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetHistory200ResponseProcess) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


