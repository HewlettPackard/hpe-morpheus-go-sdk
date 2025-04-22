# GetClusterApiConfig200Response

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

### NewGetClusterApiConfig200Response

`func NewGetClusterApiConfig200Response() *GetClusterApiConfig200Response`

NewGetClusterApiConfig200Response instantiates a new GetClusterApiConfig200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterApiConfig200ResponseWithDefaults

`func NewGetClusterApiConfig200ResponseWithDefaults() *GetClusterApiConfig200Response`

NewGetClusterApiConfig200ResponseWithDefaults instantiates a new GetClusterApiConfig200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceUrl

`func (o *GetClusterApiConfig200Response) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *GetClusterApiConfig200Response) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *GetClusterApiConfig200Response) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *GetClusterApiConfig200Response) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *GetClusterApiConfig200Response) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetClusterApiConfig200Response) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetClusterApiConfig200Response) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetClusterApiConfig200Response) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePath

`func (o *GetClusterApiConfig200Response) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *GetClusterApiConfig200Response) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *GetClusterApiConfig200Response) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *GetClusterApiConfig200Response) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceHostname

`func (o *GetClusterApiConfig200Response) GetServiceHostname() string`

GetServiceHostname returns the ServiceHostname field if non-nil, zero value otherwise.

### GetServiceHostnameOk

`func (o *GetClusterApiConfig200Response) GetServiceHostnameOk() (*string, bool)`

GetServiceHostnameOk returns a tuple with the ServiceHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHostname

`func (o *GetClusterApiConfig200Response) SetServiceHostname(v string)`

SetServiceHostname sets ServiceHostname field to given value.

### HasServiceHostname

`func (o *GetClusterApiConfig200Response) HasServiceHostname() bool`

HasServiceHostname returns a boolean if a field has been set.

### GetServicePort

`func (o *GetClusterApiConfig200Response) GetServicePort() int64`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GetClusterApiConfig200Response) GetServicePortOk() (*int64, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GetClusterApiConfig200Response) SetServicePort(v int64)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GetClusterApiConfig200Response) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceUsername

`func (o *GetClusterApiConfig200Response) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetClusterApiConfig200Response) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetClusterApiConfig200Response) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetClusterApiConfig200Response) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *GetClusterApiConfig200Response) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetClusterApiConfig200Response) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetClusterApiConfig200Response) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetClusterApiConfig200Response) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *GetClusterApiConfig200Response) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *GetClusterApiConfig200Response) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *GetClusterApiConfig200Response) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *GetClusterApiConfig200Response) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceToken

`func (o *GetClusterApiConfig200Response) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *GetClusterApiConfig200Response) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *GetClusterApiConfig200Response) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *GetClusterApiConfig200Response) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceAccess

`func (o *GetClusterApiConfig200Response) GetServiceAccess() string`

GetServiceAccess returns the ServiceAccess field if non-nil, zero value otherwise.

### GetServiceAccessOk

`func (o *GetClusterApiConfig200Response) GetServiceAccessOk() (*string, bool)`

GetServiceAccessOk returns a tuple with the ServiceAccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccess

`func (o *GetClusterApiConfig200Response) SetServiceAccess(v string)`

SetServiceAccess sets ServiceAccess field to given value.

### HasServiceAccess

`func (o *GetClusterApiConfig200Response) HasServiceAccess() bool`

HasServiceAccess returns a boolean if a field has been set.

### GetServiceCert

`func (o *GetClusterApiConfig200Response) GetServiceCert() string`

GetServiceCert returns the ServiceCert field if non-nil, zero value otherwise.

### GetServiceCertOk

`func (o *GetClusterApiConfig200Response) GetServiceCertOk() (*string, bool)`

GetServiceCertOk returns a tuple with the ServiceCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCert

`func (o *GetClusterApiConfig200Response) SetServiceCert(v string)`

SetServiceCert sets ServiceCert field to given value.

### HasServiceCert

`func (o *GetClusterApiConfig200Response) HasServiceCert() bool`

HasServiceCert returns a boolean if a field has been set.

### GetServiceVersion

`func (o *GetClusterApiConfig200Response) GetServiceVersion() string`

GetServiceVersion returns the ServiceVersion field if non-nil, zero value otherwise.

### GetServiceVersionOk

`func (o *GetClusterApiConfig200Response) GetServiceVersionOk() (*string, bool)`

GetServiceVersionOk returns a tuple with the ServiceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceVersion

`func (o *GetClusterApiConfig200Response) SetServiceVersion(v string)`

SetServiceVersion sets ServiceVersion field to given value.

### HasServiceVersion

`func (o *GetClusterApiConfig200Response) HasServiceVersion() bool`

HasServiceVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


