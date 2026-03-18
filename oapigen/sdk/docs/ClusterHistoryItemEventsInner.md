# ClusterHistoryItemEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ProcessId** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**ClusterHistoryItemEventsInnerProcessType**](ClusterHistoryItemEventsInnerProcessType.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **int64** |  | [optional] 
**SubType** | Pointer to **NullableString** |  | [optional] 
**SubId** | Pointer to **NullableInt64** |  | [optional] 
**ZoneId** | Pointer to **NullableInt64** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt64** |  | [optional] 
**InstanceId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerId** | Pointer to **NullableInt64** |  | [optional] 
**ServerId** | Pointer to **NullableInt64** |  | [optional] 
**ContainerName** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
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
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**ClusterHistoryItemEventsInnerCreatedBy**](ClusterHistoryItemEventsInnerCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**ClusterHistoryItemEventsInnerUpdatedBy**](ClusterHistoryItemEventsInnerUpdatedBy.md) |  | [optional] 

## Methods

### NewClusterHistoryItemEventsInner

`func NewClusterHistoryItemEventsInner() *ClusterHistoryItemEventsInner`

NewClusterHistoryItemEventsInner instantiates a new ClusterHistoryItemEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHistoryItemEventsInnerWithDefaults

`func NewClusterHistoryItemEventsInnerWithDefaults() *ClusterHistoryItemEventsInner`

NewClusterHistoryItemEventsInnerWithDefaults instantiates a new ClusterHistoryItemEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterHistoryItemEventsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterHistoryItemEventsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterHistoryItemEventsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterHistoryItemEventsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessId

`func (o *ClusterHistoryItemEventsInner) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ClusterHistoryItemEventsInner) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ClusterHistoryItemEventsInner) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ClusterHistoryItemEventsInner) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetAccountId

`func (o *ClusterHistoryItemEventsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClusterHistoryItemEventsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClusterHistoryItemEventsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClusterHistoryItemEventsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ClusterHistoryItemEventsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ClusterHistoryItemEventsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ClusterHistoryItemEventsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ClusterHistoryItemEventsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *ClusterHistoryItemEventsInner) GetProcessType() ClusterHistoryItemEventsInnerProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ClusterHistoryItemEventsInner) GetProcessTypeOk() (*ClusterHistoryItemEventsInnerProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ClusterHistoryItemEventsInner) SetProcessType(v ClusterHistoryItemEventsInnerProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ClusterHistoryItemEventsInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterHistoryItemEventsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterHistoryItemEventsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterHistoryItemEventsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterHistoryItemEventsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterHistoryItemEventsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterHistoryItemEventsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRefType

`func (o *ClusterHistoryItemEventsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ClusterHistoryItemEventsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ClusterHistoryItemEventsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ClusterHistoryItemEventsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ClusterHistoryItemEventsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ClusterHistoryItemEventsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ClusterHistoryItemEventsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ClusterHistoryItemEventsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubType

`func (o *ClusterHistoryItemEventsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ClusterHistoryItemEventsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ClusterHistoryItemEventsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ClusterHistoryItemEventsInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *ClusterHistoryItemEventsInner) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *ClusterHistoryItemEventsInner) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSubId

`func (o *ClusterHistoryItemEventsInner) GetSubId() int64`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ClusterHistoryItemEventsInner) GetSubIdOk() (*int64, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ClusterHistoryItemEventsInner) SetSubId(v int64)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *ClusterHistoryItemEventsInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### SetSubIdNil

`func (o *ClusterHistoryItemEventsInner) SetSubIdNil(b bool)`

 SetSubIdNil sets the value for SubId to be an explicit nil

### UnsetSubId
`func (o *ClusterHistoryItemEventsInner) UnsetSubId()`

UnsetSubId ensures that no value is present for SubId, not even an explicit nil
### GetZoneId

`func (o *ClusterHistoryItemEventsInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ClusterHistoryItemEventsInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ClusterHistoryItemEventsInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ClusterHistoryItemEventsInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ClusterHistoryItemEventsInner) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ClusterHistoryItemEventsInner) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetIntegrationId

`func (o *ClusterHistoryItemEventsInner) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ClusterHistoryItemEventsInner) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ClusterHistoryItemEventsInner) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ClusterHistoryItemEventsInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *ClusterHistoryItemEventsInner) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *ClusterHistoryItemEventsInner) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetInstanceId

`func (o *ClusterHistoryItemEventsInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClusterHistoryItemEventsInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClusterHistoryItemEventsInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClusterHistoryItemEventsInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ClusterHistoryItemEventsInner) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ClusterHistoryItemEventsInner) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *ClusterHistoryItemEventsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ClusterHistoryItemEventsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ClusterHistoryItemEventsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ClusterHistoryItemEventsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ClusterHistoryItemEventsInner) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ClusterHistoryItemEventsInner) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *ClusterHistoryItemEventsInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ClusterHistoryItemEventsInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ClusterHistoryItemEventsInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ClusterHistoryItemEventsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ClusterHistoryItemEventsInner) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ClusterHistoryItemEventsInner) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetContainerName

`func (o *ClusterHistoryItemEventsInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ClusterHistoryItemEventsInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ClusterHistoryItemEventsInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ClusterHistoryItemEventsInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *ClusterHistoryItemEventsInner) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *ClusterHistoryItemEventsInner) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetDisplayName

`func (o *ClusterHistoryItemEventsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterHistoryItemEventsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterHistoryItemEventsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterHistoryItemEventsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterHistoryItemEventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterHistoryItemEventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterHistoryItemEventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterHistoryItemEventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ClusterHistoryItemEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterHistoryItemEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterHistoryItemEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterHistoryItemEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ClusterHistoryItemEventsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ClusterHistoryItemEventsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetPercent

`func (o *ClusterHistoryItemEventsInner) GetPercent() float64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ClusterHistoryItemEventsInner) GetPercentOk() (*float64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ClusterHistoryItemEventsInner) SetPercent(v float64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ClusterHistoryItemEventsInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ClusterHistoryItemEventsInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ClusterHistoryItemEventsInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ClusterHistoryItemEventsInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ClusterHistoryItemEventsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterHistoryItemEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterHistoryItemEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterHistoryItemEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterHistoryItemEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterHistoryItemEventsInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterHistoryItemEventsInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOutput

`func (o *ClusterHistoryItemEventsInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ClusterHistoryItemEventsInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ClusterHistoryItemEventsInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ClusterHistoryItemEventsInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ClusterHistoryItemEventsInner) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ClusterHistoryItemEventsInner) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetError

`func (o *ClusterHistoryItemEventsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterHistoryItemEventsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterHistoryItemEventsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ClusterHistoryItemEventsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ClusterHistoryItemEventsInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ClusterHistoryItemEventsInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStartDate

`func (o *ClusterHistoryItemEventsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClusterHistoryItemEventsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClusterHistoryItemEventsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClusterHistoryItemEventsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClusterHistoryItemEventsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClusterHistoryItemEventsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClusterHistoryItemEventsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClusterHistoryItemEventsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ClusterHistoryItemEventsInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClusterHistoryItemEventsInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClusterHistoryItemEventsInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClusterHistoryItemEventsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterHistoryItemEventsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterHistoryItemEventsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterHistoryItemEventsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterHistoryItemEventsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterHistoryItemEventsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterHistoryItemEventsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterHistoryItemEventsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterHistoryItemEventsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClusterHistoryItemEventsInner) GetCreatedBy() ClusterHistoryItemEventsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClusterHistoryItemEventsInner) GetCreatedByOk() (*ClusterHistoryItemEventsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClusterHistoryItemEventsInner) SetCreatedBy(v ClusterHistoryItemEventsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClusterHistoryItemEventsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ClusterHistoryItemEventsInner) GetUpdatedBy() ClusterHistoryItemEventsInnerUpdatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClusterHistoryItemEventsInner) GetUpdatedByOk() (*ClusterHistoryItemEventsInnerUpdatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClusterHistoryItemEventsInner) SetUpdatedBy(v ClusterHistoryItemEventsInnerUpdatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ClusterHistoryItemEventsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


