# GetInstanceHistory200ResponseAllOfProcessesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**GetInstanceHistory200ResponseAllOfProcessesInnerProcessType**](GetInstanceHistory200ResponseAllOfProcessesInnerProcessType.md) |  | [optional] 
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
**CreatedBy** | Pointer to [**GetInstanceHistory200ResponseAllOfProcessesInnerCreatedBy**](GetInstanceHistory200ResponseAllOfProcessesInnerCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**GetInstanceHistory200ResponseAllOfProcessesInnerUpdatedBy**](GetInstanceHistory200ResponseAllOfProcessesInnerUpdatedBy.md) |  | [optional] 
**Events** | Pointer to [**[]GetInstanceHistory200ResponseAllOfProcessesInnerEventsInner**](GetInstanceHistory200ResponseAllOfProcessesInnerEventsInner.md) |  | [optional] 

## Methods

### NewGetInstanceHistory200ResponseAllOfProcessesInner

`func NewGetInstanceHistory200ResponseAllOfProcessesInner() *GetInstanceHistory200ResponseAllOfProcessesInner`

NewGetInstanceHistory200ResponseAllOfProcessesInner instantiates a new GetInstanceHistory200ResponseAllOfProcessesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceHistory200ResponseAllOfProcessesInnerWithDefaults

`func NewGetInstanceHistory200ResponseAllOfProcessesInnerWithDefaults() *GetInstanceHistory200ResponseAllOfProcessesInner`

NewGetInstanceHistory200ResponseAllOfProcessesInnerWithDefaults instantiates a new GetInstanceHistory200ResponseAllOfProcessesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetProcessType() GetInstanceHistory200ResponseAllOfProcessesInnerProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetProcessTypeOk() (*GetInstanceHistory200ResponseAllOfProcessesInnerProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetProcessType(v GetInstanceHistory200ResponseAllOfProcessesInnerProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSubType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSubId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### SetSubIdNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetSubIdNil(b bool)`

 SetSubIdNil sets the value for SubId to be an explicit nil

### UnsetSubId
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetSubId()`

UnsetSubId ensures that no value is present for SubId, not even an explicit nil
### GetZoneId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetIntegrationId() string`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetIntegrationIdOk() (*string, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetIntegrationId(v string)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetAppId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetAppId(v string)`

SetAppId sets AppId field to given value.

### HasAppId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasAppId() bool`

HasAppId returns a boolean if a field has been set.

### SetAppIdNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetAppIdNil(b bool)`

 SetAppIdNil sets the value for AppId to be an explicit nil

### UnsetAppId
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetAppId()`

UnsetAppId ensures that no value is present for AppId, not even an explicit nil
### GetInstanceId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### GetContainerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetServerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetContainerName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### GetStatus

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetPercent

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOutput

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetError

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStartDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetResultType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetResultType() string`

GetResultType returns the ResultType field if non-nil, zero value otherwise.

### GetResultTypeOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetResultTypeOk() (*string, bool)`

GetResultTypeOk returns a tuple with the ResultType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetResultType(v string)`

SetResultType sets ResultType field to given value.

### HasResultType

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasResultType() bool`

HasResultType returns a boolean if a field has been set.

### SetResultTypeNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetResultTypeNil(b bool)`

 SetResultTypeNil sets the value for ResultType to be an explicit nil

### UnsetResultType
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetResultType()`

UnsetResultType ensures that no value is present for ResultType, not even an explicit nil
### GetResultId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetResultId() int64`

GetResultId returns the ResultId field if non-nil, zero value otherwise.

### GetResultIdOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetResultIdOk() (*int64, bool)`

GetResultIdOk returns a tuple with the ResultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetResultId(v int64)`

SetResultId sets ResultId field to given value.

### HasResultId

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasResultId() bool`

HasResultId returns a boolean if a field has been set.

### SetResultIdNil

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetResultIdNil(b bool)`

 SetResultIdNil sets the value for ResultId to be an explicit nil

### UnsetResultId
`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) UnsetResultId()`

UnsetResultId ensures that no value is present for ResultId, not even an explicit nil
### GetDateCreated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetCreatedBy() GetInstanceHistory200ResponseAllOfProcessesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetCreatedByOk() (*GetInstanceHistory200ResponseAllOfProcessesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetCreatedBy(v GetInstanceHistory200ResponseAllOfProcessesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetUpdatedBy() GetInstanceHistory200ResponseAllOfProcessesInnerUpdatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetUpdatedByOk() (*GetInstanceHistory200ResponseAllOfProcessesInnerUpdatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetUpdatedBy(v GetInstanceHistory200ResponseAllOfProcessesInnerUpdatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.

### GetEvents

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetEvents() []GetInstanceHistory200ResponseAllOfProcessesInnerEventsInner`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) GetEventsOk() (*[]GetInstanceHistory200ResponseAllOfProcessesInnerEventsInner, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) SetEvents(v []GetInstanceHistory200ResponseAllOfProcessesInnerEventsInner)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *GetInstanceHistory200ResponseAllOfProcessesInner) HasEvents() bool`

HasEvents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


