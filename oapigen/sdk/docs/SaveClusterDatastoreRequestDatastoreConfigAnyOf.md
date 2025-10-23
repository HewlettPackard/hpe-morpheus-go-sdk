# SaveClusterDatastoreRequestDatastoreConfigAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceHostname** | **string** | Host name or IP address for target NFS instance. | 
**SourceDirPath** | **string** | Path to the target NFS export directory. | 
**SourceVersion** | Pointer to **string** | NFS version to use when mounting the export. | [optional] 

## Methods

### NewSaveClusterDatastoreRequestDatastoreConfigAnyOf

`func NewSaveClusterDatastoreRequestDatastoreConfigAnyOf(sourceHostname string, sourceDirPath string, ) *SaveClusterDatastoreRequestDatastoreConfigAnyOf`

NewSaveClusterDatastoreRequestDatastoreConfigAnyOf instantiates a new SaveClusterDatastoreRequestDatastoreConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveClusterDatastoreRequestDatastoreConfigAnyOfWithDefaults

`func NewSaveClusterDatastoreRequestDatastoreConfigAnyOfWithDefaults() *SaveClusterDatastoreRequestDatastoreConfigAnyOf`

NewSaveClusterDatastoreRequestDatastoreConfigAnyOfWithDefaults instantiates a new SaveClusterDatastoreRequestDatastoreConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceHostname

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceHostname() string`

GetSourceHostname returns the SourceHostname field if non-nil, zero value otherwise.

### GetSourceHostnameOk

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceHostnameOk() (*string, bool)`

GetSourceHostnameOk returns a tuple with the SourceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceHostname

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) SetSourceHostname(v string)`

SetSourceHostname sets SourceHostname field to given value.


### GetSourceDirPath

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceDirPath() string`

GetSourceDirPath returns the SourceDirPath field if non-nil, zero value otherwise.

### GetSourceDirPathOk

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceDirPathOk() (*string, bool)`

GetSourceDirPathOk returns a tuple with the SourceDirPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceDirPath

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) SetSourceDirPath(v string)`

SetSourceDirPath sets SourceDirPath field to given value.


### GetSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceVersion() string`

GetSourceVersion returns the SourceVersion field if non-nil, zero value otherwise.

### GetSourceVersionOk

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) GetSourceVersionOk() (*string, bool)`

GetSourceVersionOk returns a tuple with the SourceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) SetSourceVersion(v string)`

SetSourceVersion sets SourceVersion field to given value.

### HasSourceVersion

`func (o *SaveClusterDatastoreRequestDatastoreConfigAnyOf) HasSourceVersion() bool`

HasSourceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


