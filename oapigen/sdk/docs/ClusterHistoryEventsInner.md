# ClusterHistoryEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**ProcessId** | Pointer to **int64** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**UniqueId** | Pointer to **string** |  | [optional] 
**ProcessType** | Pointer to [**ClusterHistoryEventsInnerProcessType**](ClusterHistoryEventsInnerProcessType.md) |  | [optional] 
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
**Percent** | Pointer to **int64** |  | [optional] 
**StatusEta** | Pointer to **int64** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Output** | Pointer to **NullableString** |  | [optional] 
**Error** | Pointer to **NullableString** |  | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**EndDate** | Pointer to **time.Time** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**CreatedBy** | Pointer to [**ClusterHistoryEventsInnerCreatedBy**](ClusterHistoryEventsInnerCreatedBy.md) |  | [optional] 
**UpdatedBy** | Pointer to [**ClusterHistoryEventsInnerUpdatedBy**](ClusterHistoryEventsInnerUpdatedBy.md) |  | [optional] 

## Methods

### NewClusterHistoryEventsInner

`func NewClusterHistoryEventsInner() *ClusterHistoryEventsInner`

NewClusterHistoryEventsInner instantiates a new ClusterHistoryEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterHistoryEventsInnerWithDefaults

`func NewClusterHistoryEventsInnerWithDefaults() *ClusterHistoryEventsInner`

NewClusterHistoryEventsInnerWithDefaults instantiates a new ClusterHistoryEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterHistoryEventsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterHistoryEventsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterHistoryEventsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterHistoryEventsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProcessId

`func (o *ClusterHistoryEventsInner) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *ClusterHistoryEventsInner) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *ClusterHistoryEventsInner) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *ClusterHistoryEventsInner) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.

### GetAccountId

`func (o *ClusterHistoryEventsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ClusterHistoryEventsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ClusterHistoryEventsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ClusterHistoryEventsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetUniqueId

`func (o *ClusterHistoryEventsInner) GetUniqueId() string`

GetUniqueId returns the UniqueId field if non-nil, zero value otherwise.

### GetUniqueIdOk

`func (o *ClusterHistoryEventsInner) GetUniqueIdOk() (*string, bool)`

GetUniqueIdOk returns a tuple with the UniqueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniqueId

`func (o *ClusterHistoryEventsInner) SetUniqueId(v string)`

SetUniqueId sets UniqueId field to given value.

### HasUniqueId

`func (o *ClusterHistoryEventsInner) HasUniqueId() bool`

HasUniqueId returns a boolean if a field has been set.

### GetProcessType

`func (o *ClusterHistoryEventsInner) GetProcessType() ClusterHistoryEventsInnerProcessType`

GetProcessType returns the ProcessType field if non-nil, zero value otherwise.

### GetProcessTypeOk

`func (o *ClusterHistoryEventsInner) GetProcessTypeOk() (*ClusterHistoryEventsInnerProcessType, bool)`

GetProcessTypeOk returns a tuple with the ProcessType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessType

`func (o *ClusterHistoryEventsInner) SetProcessType(v ClusterHistoryEventsInnerProcessType)`

SetProcessType sets ProcessType field to given value.

### HasProcessType

`func (o *ClusterHistoryEventsInner) HasProcessType() bool`

HasProcessType returns a boolean if a field has been set.

### GetDescription

`func (o *ClusterHistoryEventsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterHistoryEventsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterHistoryEventsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterHistoryEventsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ClusterHistoryEventsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ClusterHistoryEventsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetRefType

`func (o *ClusterHistoryEventsInner) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *ClusterHistoryEventsInner) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *ClusterHistoryEventsInner) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *ClusterHistoryEventsInner) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *ClusterHistoryEventsInner) GetRefId() int64`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *ClusterHistoryEventsInner) GetRefIdOk() (*int64, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *ClusterHistoryEventsInner) SetRefId(v int64)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *ClusterHistoryEventsInner) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetSubType

`func (o *ClusterHistoryEventsInner) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *ClusterHistoryEventsInner) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *ClusterHistoryEventsInner) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *ClusterHistoryEventsInner) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### SetSubTypeNil

`func (o *ClusterHistoryEventsInner) SetSubTypeNil(b bool)`

 SetSubTypeNil sets the value for SubType to be an explicit nil

### UnsetSubType
`func (o *ClusterHistoryEventsInner) UnsetSubType()`

UnsetSubType ensures that no value is present for SubType, not even an explicit nil
### GetSubId

`func (o *ClusterHistoryEventsInner) GetSubId() int64`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *ClusterHistoryEventsInner) GetSubIdOk() (*int64, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *ClusterHistoryEventsInner) SetSubId(v int64)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *ClusterHistoryEventsInner) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### SetSubIdNil

`func (o *ClusterHistoryEventsInner) SetSubIdNil(b bool)`

 SetSubIdNil sets the value for SubId to be an explicit nil

### UnsetSubId
`func (o *ClusterHistoryEventsInner) UnsetSubId()`

UnsetSubId ensures that no value is present for SubId, not even an explicit nil
### GetZoneId

`func (o *ClusterHistoryEventsInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ClusterHistoryEventsInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ClusterHistoryEventsInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ClusterHistoryEventsInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### SetZoneIdNil

`func (o *ClusterHistoryEventsInner) SetZoneIdNil(b bool)`

 SetZoneIdNil sets the value for ZoneId to be an explicit nil

### UnsetZoneId
`func (o *ClusterHistoryEventsInner) UnsetZoneId()`

UnsetZoneId ensures that no value is present for ZoneId, not even an explicit nil
### GetIntegrationId

`func (o *ClusterHistoryEventsInner) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *ClusterHistoryEventsInner) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *ClusterHistoryEventsInner) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *ClusterHistoryEventsInner) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *ClusterHistoryEventsInner) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *ClusterHistoryEventsInner) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetInstanceId

`func (o *ClusterHistoryEventsInner) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *ClusterHistoryEventsInner) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *ClusterHistoryEventsInner) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.

### HasInstanceId

`func (o *ClusterHistoryEventsInner) HasInstanceId() bool`

HasInstanceId returns a boolean if a field has been set.

### SetInstanceIdNil

`func (o *ClusterHistoryEventsInner) SetInstanceIdNil(b bool)`

 SetInstanceIdNil sets the value for InstanceId to be an explicit nil

### UnsetInstanceId
`func (o *ClusterHistoryEventsInner) UnsetInstanceId()`

UnsetInstanceId ensures that no value is present for InstanceId, not even an explicit nil
### GetContainerId

`func (o *ClusterHistoryEventsInner) GetContainerId() int64`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ClusterHistoryEventsInner) GetContainerIdOk() (*int64, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ClusterHistoryEventsInner) SetContainerId(v int64)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ClusterHistoryEventsInner) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ClusterHistoryEventsInner) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ClusterHistoryEventsInner) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetServerId

`func (o *ClusterHistoryEventsInner) GetServerId() int64`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ClusterHistoryEventsInner) GetServerIdOk() (*int64, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ClusterHistoryEventsInner) SetServerId(v int64)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ClusterHistoryEventsInner) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### SetServerIdNil

`func (o *ClusterHistoryEventsInner) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *ClusterHistoryEventsInner) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetContainerName

`func (o *ClusterHistoryEventsInner) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ClusterHistoryEventsInner) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ClusterHistoryEventsInner) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ClusterHistoryEventsInner) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *ClusterHistoryEventsInner) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *ClusterHistoryEventsInner) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetDisplayName

`func (o *ClusterHistoryEventsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ClusterHistoryEventsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ClusterHistoryEventsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ClusterHistoryEventsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterHistoryEventsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterHistoryEventsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterHistoryEventsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ClusterHistoryEventsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetReason

`func (o *ClusterHistoryEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClusterHistoryEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClusterHistoryEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClusterHistoryEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ClusterHistoryEventsInner) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ClusterHistoryEventsInner) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetPercent

`func (o *ClusterHistoryEventsInner) GetPercent() int64`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *ClusterHistoryEventsInner) GetPercentOk() (*int64, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *ClusterHistoryEventsInner) SetPercent(v int64)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *ClusterHistoryEventsInner) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetStatusEta

`func (o *ClusterHistoryEventsInner) GetStatusEta() int64`

GetStatusEta returns the StatusEta field if non-nil, zero value otherwise.

### GetStatusEtaOk

`func (o *ClusterHistoryEventsInner) GetStatusEtaOk() (*int64, bool)`

GetStatusEtaOk returns a tuple with the StatusEta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusEta

`func (o *ClusterHistoryEventsInner) SetStatusEta(v int64)`

SetStatusEta sets StatusEta field to given value.

### HasStatusEta

`func (o *ClusterHistoryEventsInner) HasStatusEta() bool`

HasStatusEta returns a boolean if a field has been set.

### GetMessage

`func (o *ClusterHistoryEventsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ClusterHistoryEventsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ClusterHistoryEventsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ClusterHistoryEventsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *ClusterHistoryEventsInner) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *ClusterHistoryEventsInner) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetOutput

`func (o *ClusterHistoryEventsInner) GetOutput() string`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *ClusterHistoryEventsInner) GetOutputOk() (*string, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *ClusterHistoryEventsInner) SetOutput(v string)`

SetOutput sets Output field to given value.

### HasOutput

`func (o *ClusterHistoryEventsInner) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *ClusterHistoryEventsInner) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *ClusterHistoryEventsInner) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetError

`func (o *ClusterHistoryEventsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ClusterHistoryEventsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ClusterHistoryEventsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ClusterHistoryEventsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *ClusterHistoryEventsInner) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *ClusterHistoryEventsInner) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetStartDate

`func (o *ClusterHistoryEventsInner) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *ClusterHistoryEventsInner) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *ClusterHistoryEventsInner) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *ClusterHistoryEventsInner) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetEndDate

`func (o *ClusterHistoryEventsInner) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *ClusterHistoryEventsInner) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *ClusterHistoryEventsInner) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *ClusterHistoryEventsInner) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetDuration

`func (o *ClusterHistoryEventsInner) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClusterHistoryEventsInner) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClusterHistoryEventsInner) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClusterHistoryEventsInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetDateCreated

`func (o *ClusterHistoryEventsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ClusterHistoryEventsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ClusterHistoryEventsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ClusterHistoryEventsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ClusterHistoryEventsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ClusterHistoryEventsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ClusterHistoryEventsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ClusterHistoryEventsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ClusterHistoryEventsInner) GetCreatedBy() ClusterHistoryEventsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ClusterHistoryEventsInner) GetCreatedByOk() (*ClusterHistoryEventsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ClusterHistoryEventsInner) SetCreatedBy(v ClusterHistoryEventsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ClusterHistoryEventsInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetUpdatedBy

`func (o *ClusterHistoryEventsInner) GetUpdatedBy() ClusterHistoryEventsInnerUpdatedBy`

GetUpdatedBy returns the UpdatedBy field if non-nil, zero value otherwise.

### GetUpdatedByOk

`func (o *ClusterHistoryEventsInner) GetUpdatedByOk() (*ClusterHistoryEventsInnerUpdatedBy, bool)`

GetUpdatedByOk returns a tuple with the UpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedBy

`func (o *ClusterHistoryEventsInner) SetUpdatedBy(v ClusterHistoryEventsInnerUpdatedBy)`

SetUpdatedBy sets UpdatedBy field to given value.

### HasUpdatedBy

`func (o *ClusterHistoryEventsInner) HasUpdatedBy() bool`

HasUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


