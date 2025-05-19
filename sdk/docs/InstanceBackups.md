# InstanceBackups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**BackupsInstance200ResponseInstance**](BackupsInstance200ResponseInstance.md) |  | [optional] 
**Backups** | Pointer to **[]map[string]interface{}** | List of backup objects | [optional] 

## Methods

### NewInstanceBackups

`func NewInstanceBackups() *InstanceBackups`

NewInstanceBackups instantiates a new InstanceBackups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceBackupsWithDefaults

`func NewInstanceBackupsWithDefaults() *InstanceBackups`

NewInstanceBackupsWithDefaults instantiates a new InstanceBackups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *InstanceBackups) GetInstance() BackupsInstance200ResponseInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceBackups) GetInstanceOk() (*BackupsInstance200ResponseInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceBackups) SetInstance(v BackupsInstance200ResponseInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *InstanceBackups) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetBackups

`func (o *InstanceBackups) GetBackups() []map[string]interface{}`

GetBackups returns the Backups field if non-nil, zero value otherwise.

### GetBackupsOk

`func (o *InstanceBackups) GetBackupsOk() (*[]map[string]interface{}, bool)`

GetBackupsOk returns a tuple with the Backups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackups

`func (o *InstanceBackups) SetBackups(v []map[string]interface{})`

SetBackups sets Backups field to given value.

### HasBackups

`func (o *InstanceBackups) HasBackups() bool`

HasBackups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


