# SaveClusterDatastoreRequestDatastoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | Pointer to **string** | Host name or IP address for target NFS instance. | [optional] 
**SourceDirPath** | Pointer to **string** | Path to the target NFS export directory. | [optional] 
**BlockDevice** | Pointer to **string** | Block device for target GFS2. | [optional] 

## Methods

### NewSaveClusterDatastoreRequestDatastoreConfig

`func NewSaveClusterDatastoreRequestDatastoreConfig() *SaveClusterDatastoreRequestDatastoreConfig`

NewSaveClusterDatastoreRequestDatastoreConfig instantiates a new SaveClusterDatastoreRequestDatastoreConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastoreRequestDatastoreConfigWithDefaults

`func NewSaveClusterDatastoreRequestDatastoreConfigWithDefaults() *SaveClusterDatastoreRequestDatastoreConfig`

NewSaveClusterDatastoreRequestDatastoreConfigWithDefaults instantiates a new SaveClusterDatastoreRequestDatastoreConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHostname

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.

### HasSourceHostname

`func (o *SaveClusterDatastoreRequestDatastoreConfig) HasSourceHostname() bool`

HasSourceHostname returns a boolean if a field has been set.

### GetSourceDirPath

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceDirPath() string`

GetSourceDirPath returns the SourceDirPath field if non-nil, zero value otherwise.

### GetSourceDirPathOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceDirPathOk() (*string, bool)`

GetSourceDirPathOk returns a tuple with the SourceDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirPath

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetSourceDirPath(v string)`

SetSourceDirPath sets SourceDirPath field to given value.

### HasSourceDirPath

`func (o *SaveClusterDatastoreRequestDatastoreConfig) HasSourceDirPath() bool`

HasSourceDirPath returns a boolean if a field has been set.

### GetBlockDevice

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetBlockDevice() string`

GetBlockDevice returns the BlockDevice field if non-nil, zero value otherwise.

### GetBlockDeviceOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetBlockDeviceOk() (*string, bool)`

GetBlockDeviceOk returns a tuple with the BlockDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDevice

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetBlockDevice(v string)`

SetBlockDevice sets BlockDevice field to given value.

### HasBlockDevice

`func (o *SaveClusterDatastoreRequestDatastoreConfig) HasBlockDevice() bool`

HasBlockDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


