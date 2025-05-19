# NSXNetworkServerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | Pointer to **string** | NSX Project (NSX 4.1+) | [optional] 

## Methods

### NewNSXNetworkServerConfig

`func NewNSXNetworkServerConfig() *NSXNetworkServerConfig`

NewNSXNetworkServerConfig instantiates a new NSXNetworkServerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXNetworkServerConfigWithDefaults

`func NewNSXNetworkServerConfigWithDefaults() *NSXNetworkServerConfig`

NewNSXNetworkServerConfigWithDefaults instantiates a new NSXNetworkServerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *NSXNetworkServerConfig) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *NSXNetworkServerConfig) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *NSXNetworkServerConfig) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *NSXNetworkServerConfig) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


