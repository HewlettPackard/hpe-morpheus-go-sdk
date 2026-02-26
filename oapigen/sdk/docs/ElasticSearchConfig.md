# ElasticSearchConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EsHost** | **string** | Hostname or IP address of the Elasticsearch server | 
**EsPort** | **string** | Port to connect to the HTTP service | 
**CheckUser** | Pointer to **string** |  | [optional] 
**TextCheckOn** | Pointer to **string** |  | [optional] 
**CheckPassword** | Pointer to **string** |  | [optional] 
**WebTextMatch** | Pointer to **string** |  | [optional] 
**CheckPasswordHash** | Pointer to **string** |  | [optional] 

## Methods

### NewElasticSearchConfig

`func NewElasticSearchConfig(esHost string, esPort string, ) *ElasticSearchConfig`

NewElasticSearchConfig instantiates a new ElasticSearchConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewElasticSearchConfigWithDefaults

`func NewElasticSearchConfigWithDefaults() *ElasticSearchConfig`

NewElasticSearchConfigWithDefaults instantiates a new ElasticSearchConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEsHost

`func (o *ElasticSearchConfig) GetEsHost() string`

GetEsHost returns the EsHost field if non-nil, zero value otherwise.

### GetEsHostOk

`func (o *ElasticSearchConfig) GetEsHostOk() (*string, bool)`

GetEsHostOk returns a tuple with the EsHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsHost

`func (o *ElasticSearchConfig) SetEsHost(v string)`

SetEsHost sets EsHost field to given value.


### GetEsPort

`func (o *ElasticSearchConfig) GetEsPort() string`

GetEsPort returns the EsPort field if non-nil, zero value otherwise.

### GetEsPortOk

`func (o *ElasticSearchConfig) GetEsPortOk() (*string, bool)`

GetEsPortOk returns a tuple with the EsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsPort

`func (o *ElasticSearchConfig) SetEsPort(v string)`

SetEsPort sets EsPort field to given value.


### GetCheckUser

`func (o *ElasticSearchConfig) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *ElasticSearchConfig) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *ElasticSearchConfig) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *ElasticSearchConfig) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *ElasticSearchConfig) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *ElasticSearchConfig) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *ElasticSearchConfig) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *ElasticSearchConfig) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetCheckPassword

`func (o *ElasticSearchConfig) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *ElasticSearchConfig) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *ElasticSearchConfig) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *ElasticSearchConfig) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *ElasticSearchConfig) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *ElasticSearchConfig) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *ElasticSearchConfig) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *ElasticSearchConfig) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.

### GetCheckPasswordHash

`func (o *ElasticSearchConfig) GetCheckPasswordHash() string`

GetCheckPasswordHash returns the CheckPasswordHash field if non-nil, zero value otherwise.

### GetCheckPasswordHashOk

`func (o *ElasticSearchConfig) GetCheckPasswordHashOk() (*string, bool)`

GetCheckPasswordHashOk returns a tuple with the CheckPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPasswordHash

`func (o *ElasticSearchConfig) SetCheckPasswordHash(v string)`

SetCheckPasswordHash sets CheckPasswordHash field to given value.

### HasCheckPasswordHash

`func (o *ElasticSearchConfig) HasCheckPasswordHash() bool`

HasCheckPasswordHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


