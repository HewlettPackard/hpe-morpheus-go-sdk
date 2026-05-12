# UpdateLoadBalancerMonitorRequestLoadBalancerMonitor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**MonitorType** | Pointer to **string** |  | [optional] 
**MonitorInterval** | Pointer to **int64** |  | [optional] 
**MonitorTimeout** | Pointer to **int64** |  | [optional] 
**SendData** | Pointer to **NullableString** |  | [optional] 
**SendVersion** | Pointer to **NullableString** |  | [optional] 
**SendType** | Pointer to **NullableString** |  | [optional] 
**ReceiveData** | Pointer to **NullableString** |  | [optional] 
**ReceiveCode** | Pointer to **NullableString** |  | [optional] 
**MonitorUsername** | Pointer to **NullableString** |  | [optional] 
**MonitorPassword** | Pointer to **NullableString** |  | [optional] 
**MonitorDestination** | Pointer to **NullableString** |  | [optional] 
**FallCount** | Pointer to **int64** |  | [optional] 
**RiseCount** | Pointer to **int64** |  | [optional] 
**AliasPort** | Pointer to **int64** |  | [optional] 
**DataLength** | Pointer to **int64** |  | [optional] 
**MaxRetry** | Pointer to **int64** |  | [optional] 
**ExtraConfig** | Pointer to **NullableString** |  | [optional] 
**Partition** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to [**UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig**](UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig.md) |  | [optional] 

## Methods

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitor instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults

`func NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults() *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewUpdateLoadBalancerMonitorRequestLoadBalancerMonitorWithDefaults instantiates a new UpdateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### SetSendDataNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendDataNil(b bool)`

 SetSendDataNil sets the value for SendData to be an explicit nil

### UnsetSendData
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendData()`

UnsetSendData ensures that no value is present for SendData, not even an explicit nil
### GetSendVersion

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### SetSendVersionNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendVersionNil(b bool)`

 SetSendVersionNil sets the value for SendVersion to be an explicit nil

### UnsetSendVersion
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendVersion()`

UnsetSendVersion ensures that no value is present for SendVersion, not even an explicit nil
### GetSendType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil
### GetReceiveData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### SetReceiveDataNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveDataNil(b bool)`

 SetReceiveDataNil sets the value for ReceiveData to be an explicit nil

### UnsetReceiveData
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetReceiveData()`

UnsetReceiveData ensures that no value is present for ReceiveData, not even an explicit nil
### GetReceiveCode

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### SetReceiveCodeNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveCodeNil(b bool)`

 SetReceiveCodeNil sets the value for ReceiveCode to be an explicit nil

### UnsetReceiveCode
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetReceiveCode()`

UnsetReceiveCode ensures that no value is present for ReceiveCode, not even an explicit nil
### GetMonitorUsername

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### SetMonitorUsernameNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorUsernameNil(b bool)`

 SetMonitorUsernameNil sets the value for MonitorUsername to be an explicit nil

### UnsetMonitorUsername
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorUsername()`

UnsetMonitorUsername ensures that no value is present for MonitorUsername, not even an explicit nil
### GetMonitorPassword

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### SetMonitorPasswordNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorPasswordNil(b bool)`

 SetMonitorPasswordNil sets the value for MonitorPassword to be an explicit nil

### UnsetMonitorPassword
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorPassword()`

UnsetMonitorPassword ensures that no value is present for MonitorPassword, not even an explicit nil
### GetMonitorDestination

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### SetMonitorDestinationNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorDestinationNil(b bool)`

 SetMonitorDestinationNil sets the value for MonitorDestination to be an explicit nil

### UnsetMonitorDestination
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorDestination()`

UnsetMonitorDestination ensures that no value is present for MonitorDestination, not even an explicit nil
### GetFallCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetAliasPort

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetDataLength

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDataLength() int64`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDataLengthOk() (*int64, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDataLength(v int64)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### GetMaxRetry

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetExtraConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetPartition

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfig() UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfigOk() (*UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) SetConfig(v UpdateLoadBalancerMonitorRequestLoadBalancerMonitorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerMonitorRequestLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


