# ContainersAttachFloatingIpRequestConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OsExternalNetworkId** | **string** | The Floating IP identifier in the format: \&quot;ip-ID\&quot; or \&quot;pool-ID\&quot;.  The Options API /api/options/openStack/openstackFloatingIpOptions?containerId&#x3D;{{containerId}} can be used to see which options are available.  | 
**FloatingIpBandwidth** | Pointer to **float32** | Bandwidth (Mbit/s) Only the following cloud types support this parameter: Huawei, OpenTelekom  | [optional] 

## Methods

### NewContainersAttachFloatingIpRequestConfig

`func NewContainersAttachFloatingIpRequestConfig(osExternalNetworkId string, ) *ContainersAttachFloatingIpRequestConfig`

NewContainersAttachFloatingIpRequestConfig instantiates a new ContainersAttachFloatingIpRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContainersAttachFloatingIpRequestConfigWithDefaults

`func NewContainersAttachFloatingIpRequestConfigWithDefaults() *ContainersAttachFloatingIpRequestConfig`

NewContainersAttachFloatingIpRequestConfigWithDefaults instantiates a new ContainersAttachFloatingIpRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOsExternalNetworkId

`func (o *ContainersAttachFloatingIpRequestConfig) GetOsExternalNetworkId() string`

GetOsExternalNetworkId returns the OsExternalNetworkId field if non-nil, zero value otherwise.

### GetOsExternalNetworkIdOk

`func (o *ContainersAttachFloatingIpRequestConfig) GetOsExternalNetworkIdOk() (*string, bool)`

GetOsExternalNetworkIdOk returns a tuple with the OsExternalNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsExternalNetworkId

`func (o *ContainersAttachFloatingIpRequestConfig) SetOsExternalNetworkId(v string)`

SetOsExternalNetworkId sets OsExternalNetworkId field to given value.


### GetFloatingIpBandwidth

`func (o *ContainersAttachFloatingIpRequestConfig) GetFloatingIpBandwidth() float32`

GetFloatingIpBandwidth returns the FloatingIpBandwidth field if non-nil, zero value otherwise.

### GetFloatingIpBandwidthOk

`func (o *ContainersAttachFloatingIpRequestConfig) GetFloatingIpBandwidthOk() (*float32, bool)`

GetFloatingIpBandwidthOk returns a tuple with the FloatingIpBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloatingIpBandwidth

`func (o *ContainersAttachFloatingIpRequestConfig) SetFloatingIpBandwidth(v float32)`

SetFloatingIpBandwidth sets FloatingIpBandwidth field to given value.

### HasFloatingIpBandwidth

`func (o *ContainersAttachFloatingIpRequestConfig) HasFloatingIpBandwidth() bool`

HasFloatingIpBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


