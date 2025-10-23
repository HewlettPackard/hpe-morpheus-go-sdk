# SaveClusterDatastoreRequestDatastoreConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | **string** | Host name or IP address for target NFS instance. | 
**SourceDirPath** | **string** | Path to the target NFS export directory. | 
**SourceVersion** | Pointer to **string** | NFS version to use when mounting the export. | [optional] 
**BlockDevice** | **string** | Block device for target GFS2. | 
**AllowReformat** | Pointer to **bool** | Allow reformatting of the device if it is already formatted. | [optional] 
**EnableRansomware** | **bool** | Enable ransomware protection for this datastore | 
**ProtocolType** | **string** | Storage protocol to use, can be &#39;iSCSI&#39; or &#39;FC&#39; | 

## Methods

### NewSaveClusterDatastoreRequestDatastoreConfig

`func NewSaveClusterDatastoreRequestDatastoreConfig(sourceHostname string, sourceDirPath string, blockDevice string, enableRansomware bool, protocolType string, ) *SaveClusterDatastoreRequestDatastoreConfig`

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


### GetSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfig) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.

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


### GetAllowReformat

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetAllowReformat() bool`

GetAllowReformat returns the AllowReformat field if non-nil, zero value otherwise.

### GetAllowReformatOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetAllowReformatOk() (*bool, bool)`

GetAllowReformatOk returns a tuple with the AllowReformat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowReformat

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetAllowReformat(v bool)`

SetAllowReformat sets AllowReformat field to given value.

### HasAllowReformat

`func (o *SaveClusterDatastoreRequestDatastoreConfig) HasAllowReformat() bool`

HasAllowReformat returns a boolean if a field has been set.

### GetEnableRansomware

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetEnableRansomware() bool`

GetEnableRansomware returns the EnableRansomware field if non-nil, zero value otherwise.

### GetEnableRansomwareOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetEnableRansomwareOk() (*bool, bool)`

GetEnableRansomwareOk returns a tuple with the EnableRansomware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableRansomware

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetEnableRansomware(v bool)`

SetEnableRansomware sets EnableRansomware field to given value.


### GetProtocolType

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetProtocolType() string`

GetProtocolType returns the ProtocolType field if non-nil, zero value otherwise.

### GetProtocolTypeOk

`func (o *SaveClusterDatastoreRequestDatastoreConfig) GetProtocolTypeOk() (*string, bool)`

GetProtocolTypeOk returns a tuple with the ProtocolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocolType

`func (o *SaveClusterDatastoreRequestDatastoreConfig) SetProtocolType(v string)`

SetProtocolType sets ProtocolType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


