# ApplyTemplateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceUrl** | Pointer to **string** | Url of desired template to apply to cluster | [optional] 
**SpecTemplate** | Pointer to **string** | Name or ID of desired Spec Template to apply to cluster | [optional] 
**SpecYaml** | Pointer to **string** | Yaml of template to apply to cluster | [optional] 

## Methods

### NewApplyTemplateRequest

`func NewApplyTemplateRequest() *ApplyTemplateRequest`

NewApplyTemplateRequest instantiates a new ApplyTemplateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyTemplateRequestWithDefaults

`func NewApplyTemplateRequestWithDefaults() *ApplyTemplateRequest`

NewApplyTemplateRequestWithDefaults instantiates a new ApplyTemplateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceUrl

`func (o *ApplyTemplateRequest) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ApplyTemplateRequest) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ApplyTemplateRequest) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ApplyTemplateRequest) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetSpecTemplate

`func (o *ApplyTemplateRequest) GetSpecTemplate() string`

GetSpecTemplate returns the SpecTemplate field if non-nil, zero value otherwise.

### GetSpecTemplateOk

`func (o *ApplyTemplateRequest) GetSpecTemplateOk() (*string, bool)`

GetSpecTemplateOk returns a tuple with the SpecTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplate

`func (o *ApplyTemplateRequest) SetSpecTemplate(v string)`

SetSpecTemplate sets SpecTemplate field to given value.

### HasSpecTemplate

`func (o *ApplyTemplateRequest) HasSpecTemplate() bool`

HasSpecTemplate returns a boolean if a field has been set.

### GetSpecYaml

`func (o *ApplyTemplateRequest) GetSpecYaml() string`

GetSpecYaml returns the SpecYaml field if non-nil, zero value otherwise.

### GetSpecYamlOk

`func (o *ApplyTemplateRequest) GetSpecYamlOk() (*string, bool)`

GetSpecYamlOk returns a tuple with the SpecYaml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecYaml

`func (o *ApplyTemplateRequest) SetSpecYaml(v string)`

SetSpecYaml sets SpecYaml field to given value.

### HasSpecYaml

`func (o *ApplyTemplateRequest) HasSpecYaml() bool`

HasSpecYaml returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


