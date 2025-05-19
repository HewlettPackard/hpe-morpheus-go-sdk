# ClusterApiConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceUrl** | Pointer to **string** |  | [optional] 
**ServiceHost** | Pointer to **string** |  | [optional] 
**ServicePath** | Pointer to **string** |  | [optional] 
**ServiceHostname** | Pointer to **string** |  | [optional] 
**ServicePort** | Pointer to **int64** |  | [optional] 
**ServiceUsername** | Pointer to **string** |  | [optional] 
**ServicePassword** | Pointer to **string** |  | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** | API Token | [optional] 
**ServiceAccess** | Pointer to **string** | Kube Config | [optional] 
**ServiceCert** | Pointer to **string** |  | [optional] 
**ServiceVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterApiConfig

`func NewClusterApiConfig() *ClusterApiConfig`

NewClusterApiConfig instantiates a new ClusterApiConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterApiConfigWithDefaults

`func NewClusterApiConfigWithDefaults() *ClusterApiConfig`

NewClusterApiConfigWithDefaults instantiates a new ClusterApiConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceUrl

`func (o *ClusterApiConfig) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ClusterApiConfig) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ClusterApiConfig) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ClusterApiConfig) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *ClusterApiConfig) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *ClusterApiConfig) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *ClusterApiConfig) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *ClusterApiConfig) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePath

`func (o *ClusterApiConfig) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *ClusterApiConfig) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *ClusterApiConfig) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *ClusterApiConfig) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceHostname

`func (o *ClusterApiConfig) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *ClusterApiConfig) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *ClusterApiConfig) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *ClusterApiConfig) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### GetServicePort

`func (o *ClusterApiConfig) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ClusterApiConfig) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ClusterApiConfig) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ClusterApiConfig) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ClusterApiConfig) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ClusterApiConfig) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ClusterApiConfig) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ClusterApiConfig) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *ClusterApiConfig) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ClusterApiConfig) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ClusterApiConfig) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ClusterApiConfig) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *ClusterApiConfig) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *ClusterApiConfig) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *ClusterApiConfig) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *ClusterApiConfig) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceToken

`func (o *ClusterApiConfig) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *ClusterApiConfig) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *ClusterApiConfig) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *ClusterApiConfig) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceAccess

`func (o *ClusterApiConfig) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *ClusterApiConfig) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *ClusterApiConfig) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *ClusterApiConfig) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetServiceCert

`func (o *ClusterApiConfig) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *ClusterApiConfig) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *ClusterApiConfig) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *ClusterApiConfig) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### GetServiceVersion

`func (o *ClusterApiConfig) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *ClusterApiConfig) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *ClusterApiConfig) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *ClusterApiConfig) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


