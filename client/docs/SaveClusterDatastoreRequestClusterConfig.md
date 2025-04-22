# SaveClusterDatastoreRequestClusterConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | Pointer to **string** | Host name or IP address for target NFS instance. | [optional] 
**SourceDirPath** | Pointer to **string** | Path to the target NFS export directory. | [optional] 
**BlockDevice** | Pointer to **string** | Block device for target GFS2. | [optional] 

## Methods

### NewSaveClusterDatastoreRequestClusterConfig

`func NewSaveClusterDatastoreRequestClusterConfig() *SaveClusterDatastoreRequestClusterConfig`

NewSaveClusterDatastoreRequestClusterConfig instantiates a new SaveClusterDatastoreRequestClusterConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastoreRequestClusterConfigWithDefaults

`func NewSaveClusterDatastoreRequestClusterConfigWithDefaults() *SaveClusterDatastoreRequestClusterConfig`

NewSaveClusterDatastoreRequestClusterConfigWithDefaults instantiates a new SaveClusterDatastoreRequestClusterConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHostname

`func (o *SaveClusterDatastoreRequestClusterConfig) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *SaveClusterDatastoreRequestClusterConfig) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *SaveClusterDatastoreRequestClusterConfig) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.

### HasSourceHostname

`func (o *SaveClusterDatastoreRequestClusterConfig) HasSourceHostname() bool`

HasSourceHostname returns a boolean if a field has been set.

### GetSourceDirPath

`func (o *SaveClusterDatastoreRequestClusterConfig) GetSourceDirPath() string`

GetSourceDirPath returns the SourceDirPath field if non-nil, zero value otherwise.

### GetSourceDirPathOk

`func (o *SaveClusterDatastoreRequestClusterConfig) GetSourceDirPathOk() (*string, bool)`

GetSourceDirPathOk returns a tuple with the SourceDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirPath

`func (o *SaveClusterDatastoreRequestClusterConfig) SetSourceDirPath(v string)`

SetSourceDirPath sets SourceDirPath field to given value.

### HasSourceDirPath

`func (o *SaveClusterDatastoreRequestClusterConfig) HasSourceDirPath() bool`

HasSourceDirPath returns a boolean if a field has been set.

### GetBlockDevice

`func (o *SaveClusterDatastoreRequestClusterConfig) GetBlockDevice() string`

GetBlockDevice returns the BlockDevice field if non-nil, zero value otherwise.

### GetBlockDeviceOk

`func (o *SaveClusterDatastoreRequestClusterConfig) GetBlockDeviceOk() (*string, bool)`

GetBlockDeviceOk returns a tuple with the BlockDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDevice

`func (o *SaveClusterDatastoreRequestClusterConfig) SetBlockDevice(v string)`

SetBlockDevice sets BlockDevice field to given value.

### HasBlockDevice

`func (o *SaveClusterDatastoreRequestClusterConfig) HasBlockDevice() bool`

HasBlockDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


