# ClusterCreateConfigTemplateParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsPrefix** | Pointer to **string** | Optional DNS prefix to use with hosted Kubernetes API server FQDN. | [optional] 
**OnDiskSizeGB** | Pointer to **int64** | Disk size (in GB) to provision for each of the agent pool nodes. This value ranges from 0 to 1023. Specifying 0 will apply the default disk size for that agentVMSize. | [optional] 
**NodeVMSize** | Pointer to **int64** | The size of the Virtual Machine. | [optional] 
**VnetSubnetID** | Pointer to **string** | Resource ID of virtual network subnet used for nodes and/or pods IP assignment. | [optional] 
**ServiceCidr** | Pointer to **string** | A CIDR notation IP range from which to assign service cluster IPs. | [optional] 
**DnsServiceIP** | Pointer to **string** | Containers DNS server IP address. | [optional] 
**DockerBridgeCidr** | Pointer to **string** | A CIDR notation IP for Docker bridge. | [optional] 

## Methods

### NewClusterCreateConfigTemplateParameter

`func NewClusterCreateConfigTemplateParameter() *ClusterCreateConfigTemplateParameter`

NewClusterCreateConfigTemplateParameter instantiates a new ClusterCreateConfigTemplateParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateConfigTemplateParameterWithDefaults

`func NewClusterCreateConfigTemplateParameterWithDefaults() *ClusterCreateConfigTemplateParameter`

NewClusterCreateConfigTemplateParameterWithDefaults instantiates a new ClusterCreateConfigTemplateParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsPrefix

`func (o *ClusterCreateConfigTemplateParameter) GetDnsPrefix() string`

GetDnsPrefix returns the DnsPrefix field if non-nil, zero value otherwise.

### GetDnsPrefixOk

`func (o *ClusterCreateConfigTemplateParameter) GetDnsPrefixOk() (*string, bool)`

GetDnsPrefixOk returns a tuple with the DnsPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsPrefix

`func (o *ClusterCreateConfigTemplateParameter) SetDnsPrefix(v string)`

SetDnsPrefix sets DnsPrefix field to given value.

### HasDnsPrefix

`func (o *ClusterCreateConfigTemplateParameter) HasDnsPrefix() bool`

HasDnsPrefix returns a boolean if a field has been set.

### GetOnDiskSizeGB

`func (o *ClusterCreateConfigTemplateParameter) GetOnDiskSizeGB() int64`

GetOnDiskSizeGB returns the OnDiskSizeGB field if non-nil, zero value otherwise.

### GetOnDiskSizeGBOk

`func (o *ClusterCreateConfigTemplateParameter) GetOnDiskSizeGBOk() (*int64, bool)`

GetOnDiskSizeGBOk returns a tuple with the OnDiskSizeGB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnDiskSizeGB

`func (o *ClusterCreateConfigTemplateParameter) SetOnDiskSizeGB(v int64)`

SetOnDiskSizeGB sets OnDiskSizeGB field to given value.

### HasOnDiskSizeGB

`func (o *ClusterCreateConfigTemplateParameter) HasOnDiskSizeGB() bool`

HasOnDiskSizeGB returns a boolean if a field has been set.

### GetNodeVMSize

`func (o *ClusterCreateConfigTemplateParameter) GetNodeVMSize() int64`

GetNodeVMSize returns the NodeVMSize field if non-nil, zero value otherwise.

### GetNodeVMSizeOk

`func (o *ClusterCreateConfigTemplateParameter) GetNodeVMSizeOk() (*int64, bool)`

GetNodeVMSizeOk returns a tuple with the NodeVMSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeVMSize

`func (o *ClusterCreateConfigTemplateParameter) SetNodeVMSize(v int64)`

SetNodeVMSize sets NodeVMSize field to given value.

### HasNodeVMSize

`func (o *ClusterCreateConfigTemplateParameter) HasNodeVMSize() bool`

HasNodeVMSize returns a boolean if a field has been set.

### GetVnetSubnetID

`func (o *ClusterCreateConfigTemplateParameter) GetVnetSubnetID() string`

GetVnetSubnetID returns the VnetSubnetID field if non-nil, zero value otherwise.

### GetVnetSubnetIDOk

`func (o *ClusterCreateConfigTemplateParameter) GetVnetSubnetIDOk() (*string, bool)`

GetVnetSubnetIDOk returns a tuple with the VnetSubnetID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVnetSubnetID

`func (o *ClusterCreateConfigTemplateParameter) SetVnetSubnetID(v string)`

SetVnetSubnetID sets VnetSubnetID field to given value.

### HasVnetSubnetID

`func (o *ClusterCreateConfigTemplateParameter) HasVnetSubnetID() bool`

HasVnetSubnetID returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterCreateConfigTemplateParameter) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterCreateConfigTemplateParameter) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterCreateConfigTemplateParameter) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterCreateConfigTemplateParameter) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetDnsServiceIP

`func (o *ClusterCreateConfigTemplateParameter) GetDnsServiceIP() string`

GetDnsServiceIP returns the DnsServiceIP field if non-nil, zero value otherwise.

### GetDnsServiceIPOk

`func (o *ClusterCreateConfigTemplateParameter) GetDnsServiceIPOk() (*string, bool)`

GetDnsServiceIPOk returns a tuple with the DnsServiceIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsServiceIP

`func (o *ClusterCreateConfigTemplateParameter) SetDnsServiceIP(v string)`

SetDnsServiceIP sets DnsServiceIP field to given value.

### HasDnsServiceIP

`func (o *ClusterCreateConfigTemplateParameter) HasDnsServiceIP() bool`

HasDnsServiceIP returns a boolean if a field has been set.

### GetDockerBridgeCidr

`func (o *ClusterCreateConfigTemplateParameter) GetDockerBridgeCidr() string`

GetDockerBridgeCidr returns the DockerBridgeCidr field if non-nil, zero value otherwise.

### GetDockerBridgeCidrOk

`func (o *ClusterCreateConfigTemplateParameter) GetDockerBridgeCidrOk() (*string, bool)`

GetDockerBridgeCidrOk returns a tuple with the DockerBridgeCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDockerBridgeCidr

`func (o *ClusterCreateConfigTemplateParameter) SetDockerBridgeCidr(v string)`

SetDockerBridgeCidr sets DockerBridgeCidr field to given value.

### HasDockerBridgeCidr

`func (o *ClusterCreateConfigTemplateParameter) HasDockerBridgeCidr() bool`

HasDockerBridgeCidr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


