# ClusterDatastoreConfigNFS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | Pointer to **string** | Host name or IP address for target NFS instance. | [optional] 
**SourceDirPath** | Pointer to **string** | Path to the target NFS export directory. | [optional] 

## Methods

### NewClusterDatastoreConfigNFS

`func NewClusterDatastoreConfigNFS() *ClusterDatastoreConfigNFS`

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

### HasSourceHostname

`func (o *ClusterDatastoreConfigNFS) HasSourceHostname() bool`

HasSourceHostname returns a boolean if a field has been set.

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

### HasSourceDirPath

`func (o *ClusterDatastoreConfigNFS) HasSourceDirPath() bool`

HasSourceDirPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


