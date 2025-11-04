# ClusterDatastoreCreateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | **string** | Host name or IP address for target NFS instance. | 
**SourceDirPath** | **string** | Path to the target NFS export directory. | 
**SourceVersion** | Pointer to **string** | NFS version to use when mounting the export. | [optional] 
**BlockDevice** | **string** | Block device for target GFS2. | 
**AllowReformat** | Pointer to **bool** | Allow reformatting of the device if it is already formatted. | [optional] 
**Enableransomware** | Pointer to **string** | Enable ransomware protection for this datastore | [optional] 
**ProtocolType** | **string** | Storage protocol to use, can be &#39;iSCSI&#39; or &#39;FC&#39; | 

## Methods

### NewClusterDatastoreCreateConfig

`func NewClusterDatastoreCreateConfig(sourceHostname string, sourceDirPath string, blockDevice string, protocolType string, ) *ClusterDatastoreCreateConfig`

NewClusterDatastoreCreateConfig instantiates a new ClusterDatastoreCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoreCreateConfigWithDefaults

`func NewClusterDatastoreCreateConfigWithDefaults() *ClusterDatastoreCreateConfig`

NewClusterDatastoreCreateConfigWithDefaults instantiates a new ClusterDatastoreCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHostname

`func (o *ClusterDatastoreCreateConfig) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *ClusterDatastoreCreateConfig) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *ClusterDatastoreCreateConfig) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.


### GetSourceDirPath

`func (o *ClusterDatastoreCreateConfig) GetSourceDirPath() string`

GetSourceDirPath returns the SourceDirPath field if non-nil, zero value otherwise.

### GetSourceDirPathOk

`func (o *ClusterDatastoreCreateConfig) GetSourceDirPathOk() (*string, bool)`

GetSourceDirPathOk returns a tuple with the SourceDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirPath

`func (o *ClusterDatastoreCreateConfig) SetSourceDirPath(v string)`

SetSourceDirPath sets SourceDirPath field to given value.


### GetSourceVersion

`func (o *ClusterDatastoreCreateConfig) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *ClusterDatastoreCreateConfig) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *ClusterDatastoreCreateConfig) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *ClusterDatastoreCreateConfig) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

### GetBlockDevice

`func (o *ClusterDatastoreCreateConfig) GetBlockDevice() string`

GetBlockDevice returns the BlockDevice field if non-nil, zero value otherwise.

### GetBlockDeviceOk

`func (o *ClusterDatastoreCreateConfig) GetBlockDeviceOk() (*string, bool)`

GetBlockDeviceOk returns a tuple with the BlockDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockDevice

`func (o *ClusterDatastoreCreateConfig) SetBlockDevice(v string)`

SetBlockDevice sets BlockDevice field to given value.


### GetAllowReformat

`func (o *ClusterDatastoreCreateConfig) GetAllowReformat() bool`

GetAllowReformat returns the AllowReformat field if non-nil, zero value otherwise.

### GetAllowReformatOk

`func (o *ClusterDatastoreCreateConfig) GetAllowReformatOk() (*bool, bool)`

GetAllowReformatOk returns a tuple with the AllowReformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReformat

`func (o *ClusterDatastoreCreateConfig) SetAllowReformat(v bool)`

SetAllowReformat sets AllowReformat field to given value.

### HasAllowReformat

`func (o *ClusterDatastoreCreateConfig) HasAllowReformat() bool`

HasAllowReformat returns a boolean if a field has been set.

### GetEnableransomware

`func (o *ClusterDatastoreCreateConfig) GetEnableransomware() string`

GetEnableransomware returns the Enableransomware field if non-nil, zero value otherwise.

### GetEnableransomwareOk

`func (o *ClusterDatastoreCreateConfig) GetEnableransomwareOk() (*string, bool)`

GetEnableransomwareOk returns a tuple with the Enableransomware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableransomware

`func (o *ClusterDatastoreCreateConfig) SetEnableransomware(v string)`

SetEnableransomware sets Enableransomware field to given value.

### HasEnableransomware

`func (o *ClusterDatastoreCreateConfig) HasEnableransomware() bool`

HasEnableransomware returns a boolean if a field has been set.

### GetProtocolType

`func (o *ClusterDatastoreCreateConfig) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *ClusterDatastoreCreateConfig) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *ClusterDatastoreCreateConfig) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


