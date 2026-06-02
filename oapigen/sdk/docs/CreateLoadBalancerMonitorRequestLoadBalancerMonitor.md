# CreateLoadBalancerMonitorRequestLoadBalancerMonitor

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
**Config** | Pointer to [**CreateLoadBalancerMonitorRequestLoadBalancerMonitorConfig**](CreateLoadBalancerMonitorRequestLoadBalancerMonitorConfig.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor

`func NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor() *CreateLoadBalancerMonitorRequestLoadBalancerMonitor`

NewCreateLoadBalancerMonitorRequestLoadBalancerMonitor instantiates a new CreateLoadBalancerMonitorRequestLoadBalancerMonitor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorType() string`

GetMonitorType returns the MonitorType field if non-nil, zero value otherwise.

### GetMonitorTypeOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTypeOk() (*string, bool)`

GetMonitorTypeOk returns a tuple with the MonitorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorType(v string)`

SetMonitorType sets MonitorType field to given value.

### HasMonitorType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorType() bool`

HasMonitorType returns a boolean if a field has been set.

### GetMonitorInterval

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorInterval() int64`

GetMonitorInterval returns the MonitorInterval field if non-nil, zero value otherwise.

### GetMonitorIntervalOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorIntervalOk() (*int64, bool)`

GetMonitorIntervalOk returns a tuple with the MonitorInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorInterval

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorInterval(v int64)`

SetMonitorInterval sets MonitorInterval field to given value.

### HasMonitorInterval

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorInterval() bool`

HasMonitorInterval returns a boolean if a field has been set.

### GetMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeout() int64`

GetMonitorTimeout returns the MonitorTimeout field if non-nil, zero value otherwise.

### GetMonitorTimeoutOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorTimeoutOk() (*int64, bool)`

GetMonitorTimeoutOk returns a tuple with the MonitorTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorTimeout(v int64)`

SetMonitorTimeout sets MonitorTimeout field to given value.

### HasMonitorTimeout

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorTimeout() bool`

HasMonitorTimeout returns a boolean if a field has been set.

### GetSendData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendData() string`

GetSendData returns the SendData field if non-nil, zero value otherwise.

### GetSendDataOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendDataOk() (*string, bool)`

GetSendDataOk returns a tuple with the SendData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendData(v string)`

SetSendData sets SendData field to given value.

### HasSendData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendData() bool`

HasSendData returns a boolean if a field has been set.

### SetSendDataNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendDataNil(b bool)`

 SetSendDataNil sets the value for SendData to be an explicit nil

### UnsetSendData
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendData()`

UnsetSendData ensures that no value is present for SendData, not even an explicit nil
### GetSendVersion

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendVersion() string`

GetSendVersion returns the SendVersion field if non-nil, zero value otherwise.

### GetSendVersionOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendVersionOk() (*string, bool)`

GetSendVersionOk returns a tuple with the SendVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendVersion

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendVersion(v string)`

SetSendVersion sets SendVersion field to given value.

### HasSendVersion

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendVersion() bool`

HasSendVersion returns a boolean if a field has been set.

### SetSendVersionNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendVersionNil(b bool)`

 SetSendVersionNil sets the value for SendVersion to be an explicit nil

### UnsetSendVersion
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendVersion()`

UnsetSendVersion ensures that no value is present for SendVersion, not even an explicit nil
### GetSendType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendType() string`

GetSendType returns the SendType field if non-nil, zero value otherwise.

### GetSendTypeOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetSendTypeOk() (*string, bool)`

GetSendTypeOk returns a tuple with the SendType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendType(v string)`

SetSendType sets SendType field to given value.

### HasSendType

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasSendType() bool`

HasSendType returns a boolean if a field has been set.

### SetSendTypeNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetSendTypeNil(b bool)`

 SetSendTypeNil sets the value for SendType to be an explicit nil

### UnsetSendType
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetSendType()`

UnsetSendType ensures that no value is present for SendType, not even an explicit nil
### GetReceiveData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveData() string`

GetReceiveData returns the ReceiveData field if non-nil, zero value otherwise.

### GetReceiveDataOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveDataOk() (*string, bool)`

GetReceiveDataOk returns a tuple with the ReceiveData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveData(v string)`

SetReceiveData sets ReceiveData field to given value.

### HasReceiveData

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasReceiveData() bool`

HasReceiveData returns a boolean if a field has been set.

### SetReceiveDataNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveDataNil(b bool)`

 SetReceiveDataNil sets the value for ReceiveData to be an explicit nil

### UnsetReceiveData
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetReceiveData()`

UnsetReceiveData ensures that no value is present for ReceiveData, not even an explicit nil
### GetReceiveCode

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveCode() string`

GetReceiveCode returns the ReceiveCode field if non-nil, zero value otherwise.

### GetReceiveCodeOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetReceiveCodeOk() (*string, bool)`

GetReceiveCodeOk returns a tuple with the ReceiveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveCode

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveCode(v string)`

SetReceiveCode sets ReceiveCode field to given value.

### HasReceiveCode

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasReceiveCode() bool`

HasReceiveCode returns a boolean if a field has been set.

### SetReceiveCodeNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetReceiveCodeNil(b bool)`

 SetReceiveCodeNil sets the value for ReceiveCode to be an explicit nil

### UnsetReceiveCode
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetReceiveCode()`

UnsetReceiveCode ensures that no value is present for ReceiveCode, not even an explicit nil
### GetMonitorUsername

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorUsername() string`

GetMonitorUsername returns the MonitorUsername field if non-nil, zero value otherwise.

### GetMonitorUsernameOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorUsernameOk() (*string, bool)`

GetMonitorUsernameOk returns a tuple with the MonitorUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorUsername

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorUsername(v string)`

SetMonitorUsername sets MonitorUsername field to given value.

### HasMonitorUsername

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorUsername() bool`

HasMonitorUsername returns a boolean if a field has been set.

### SetMonitorUsernameNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorUsernameNil(b bool)`

 SetMonitorUsernameNil sets the value for MonitorUsername to be an explicit nil

### UnsetMonitorUsername
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorUsername()`

UnsetMonitorUsername ensures that no value is present for MonitorUsername, not even an explicit nil
### GetMonitorPassword

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorPassword() string`

GetMonitorPassword returns the MonitorPassword field if non-nil, zero value otherwise.

### GetMonitorPasswordOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorPasswordOk() (*string, bool)`

GetMonitorPasswordOk returns a tuple with the MonitorPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorPassword

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorPassword(v string)`

SetMonitorPassword sets MonitorPassword field to given value.

### HasMonitorPassword

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorPassword() bool`

HasMonitorPassword returns a boolean if a field has been set.

### SetMonitorPasswordNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorPasswordNil(b bool)`

 SetMonitorPasswordNil sets the value for MonitorPassword to be an explicit nil

### UnsetMonitorPassword
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorPassword()`

UnsetMonitorPassword ensures that no value is present for MonitorPassword, not even an explicit nil
### GetMonitorDestination

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorDestination() string`

GetMonitorDestination returns the MonitorDestination field if non-nil, zero value otherwise.

### GetMonitorDestinationOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMonitorDestinationOk() (*string, bool)`

GetMonitorDestinationOk returns a tuple with the MonitorDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonitorDestination

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorDestination(v string)`

SetMonitorDestination sets MonitorDestination field to given value.

### HasMonitorDestination

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMonitorDestination() bool`

HasMonitorDestination returns a boolean if a field has been set.

### SetMonitorDestinationNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMonitorDestinationNil(b bool)`

 SetMonitorDestinationNil sets the value for MonitorDestination to be an explicit nil

### UnsetMonitorDestination
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetMonitorDestination()`

UnsetMonitorDestination ensures that no value is present for MonitorDestination, not even an explicit nil
### GetFallCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetFallCount() int64`

GetFallCount returns the FallCount field if non-nil, zero value otherwise.

### GetFallCountOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetFallCountOk() (*int64, bool)`

GetFallCountOk returns a tuple with the FallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetFallCount(v int64)`

SetFallCount sets FallCount field to given value.

### HasFallCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasFallCount() bool`

HasFallCount returns a boolean if a field has been set.

### GetRiseCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetRiseCount() int64`

GetRiseCount returns the RiseCount field if non-nil, zero value otherwise.

### GetRiseCountOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetRiseCountOk() (*int64, bool)`

GetRiseCountOk returns a tuple with the RiseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiseCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetRiseCount(v int64)`

SetRiseCount sets RiseCount field to given value.

### HasRiseCount

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasRiseCount() bool`

HasRiseCount returns a boolean if a field has been set.

### GetAliasPort

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetAliasPort() int64`

GetAliasPort returns the AliasPort field if non-nil, zero value otherwise.

### GetAliasPortOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetAliasPortOk() (*int64, bool)`

GetAliasPortOk returns a tuple with the AliasPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasPort

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetAliasPort(v int64)`

SetAliasPort sets AliasPort field to given value.

### HasAliasPort

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasAliasPort() bool`

HasAliasPort returns a boolean if a field has been set.

### GetDataLength

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDataLength() int64`

GetDataLength returns the DataLength field if non-nil, zero value otherwise.

### GetDataLengthOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetDataLengthOk() (*int64, bool)`

GetDataLengthOk returns a tuple with the DataLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataLength

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetDataLength(v int64)`

SetDataLength sets DataLength field to given value.

### HasDataLength

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasDataLength() bool`

HasDataLength returns a boolean if a field has been set.

### GetMaxRetry

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMaxRetry() int64`

GetMaxRetry returns the MaxRetry field if non-nil, zero value otherwise.

### GetMaxRetryOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetMaxRetryOk() (*int64, bool)`

GetMaxRetryOk returns a tuple with the MaxRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRetry

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetMaxRetry(v int64)`

SetMaxRetry sets MaxRetry field to given value.

### HasMaxRetry

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasMaxRetry() bool`

HasMaxRetry returns a boolean if a field has been set.

### GetExtraConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetExtraConfig() string`

GetExtraConfig returns the ExtraConfig field if non-nil, zero value otherwise.

### GetExtraConfigOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetExtraConfigOk() (*string, bool)`

GetExtraConfigOk returns a tuple with the ExtraConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetExtraConfig(v string)`

SetExtraConfig sets ExtraConfig field to given value.

### HasExtraConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasExtraConfig() bool`

HasExtraConfig returns a boolean if a field has been set.

### SetExtraConfigNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetExtraConfigNil(b bool)`

 SetExtraConfigNil sets the value for ExtraConfig to be an explicit nil

### UnsetExtraConfig
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetExtraConfig()`

UnsetExtraConfig ensures that no value is present for ExtraConfig, not even an explicit nil
### GetPartition

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetPartition() string`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetPartitionOk() (*string, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetPartition(v string)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### SetPartitionNil

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetPartitionNil(b bool)`

 SetPartitionNil sets the value for Partition to be an explicit nil

### UnsetPartition
`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) UnsetPartition()`

UnsetPartition ensures that no value is present for Partition, not even an explicit nil
### GetConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfig() CreateLoadBalancerMonitorRequestLoadBalancerMonitorConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) GetConfigOk() (*CreateLoadBalancerMonitorRequestLoadBalancerMonitorConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) SetConfig(v CreateLoadBalancerMonitorRequestLoadBalancerMonitorConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerMonitorRequestLoadBalancerMonitor) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


