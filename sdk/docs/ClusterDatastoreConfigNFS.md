# ClusterDatastoreConfigNFS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | **string** | Host name or IP address for target NFS instance. | 
**SourceDirPath** | **string** | Path to the target NFS export directory. | 
**SourceVersion** | Pointer to **string** | NFS version to use when mounting the export. | [optional] 

## Methods

### NewClusterDatastoreConfigNFS

`func NewClusterDatastoreConfigNFS(sourceHostname string, sourceDirPath string, ) *ClusterDatastoreConfigNFS`

NewClusterDatastoreConfigNFS instantiates a new ClusterDatastoreConfigNFS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterDatastoreConfigNFSWithDefaults

`func NewClusterDatastoreConfigNFSWithDefaults() *ClusterDatastoreConfigNFS`

NewClusterDatastoreConfigNFSWithDefaults instantiates a new ClusterDatastoreConfigNFS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHostname

`func (o *ClusterDatastoreConfigNFS) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *ClusterDatastoreConfigNFS) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *ClusterDatastoreConfigNFS) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.


### GetSourceDirPath

`func (o *ClusterDatastoreConfigNFS) GetSourceDirPath() string`

GetSourceDirPath returns the SourceDirPath field if non-nil, zero value otherwise.

### GetSourceDirPathOk

`func (o *ClusterDatastoreConfigNFS) GetSourceDirPathOk() (*string, bool)`

GetSourceDirPathOk returns a tuple with the SourceDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirPath

`func (o *ClusterDatastoreConfigNFS) SetSourceDirPath(v string)`

SetSourceDirPath sets SourceDirPath field to given value.


### GetSourceVersion

`func (o *ClusterDatastoreConfigNFS) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *ClusterDatastoreConfigNFS) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *ClusterDatastoreConfigNFS) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *ClusterDatastoreConfigNFS) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


