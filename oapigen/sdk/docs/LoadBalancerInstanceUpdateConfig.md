# LoadBalancerInstanceUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationProfile** | Pointer to **string** | The Load Balancer Application Profile ID. The Options API &#x60;/api/options/nsxt/nsxtLBVirtualServerApplicationProfile?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available. | [optional] 
**Pool** | Pointer to **string** | The backend server pool ID (&#x60;NetworkLoadBalancerPool&#x60;). The Options API &#x60;/api/options/nsxt/nsxtLBPool?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 
**Persistence** | Pointer to **NullableString** | Session persistence mode. The available values depend on the virtual server protocol. For HTTP: &#x60;SOURCE_IP&#x60;, &#x60;COOKIE&#x60;, or empty string (disabled). For TCP/UDP: &#x60;SOURCE_IP&#x60; or empty string (disabled). The Options API &#x60;/api/options/nsxt/nsxtLBPersistence?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available. | [optional] 
**PersistenceProfile** | Pointer to **string** | The ID of the persistence profile to use. Required when &#x60;persistence&#x60; is set to a non-empty value (&#x60;SOURCE_IP&#x60; or &#x60;COOKIE&#x60;). The Options API &#x60;/api/options/nsxt/nsxtLBPersistenceProfile?loadBalancerId&#x3D;42&amp;config.persistence&#x3D;SOURCE_IP&#x60; can be used to see which options are available. | [optional] 
**SslClientProfile** | Pointer to **string** | The SSL client profile ID. Only applicable when &#x60;sslCert&#x60; is set to a non-zero value. The Options API &#x60;/api/options/nsxt/nsxtLBClientSSlProfiles?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 
**SslServerProfile** | Pointer to **string** | The SSL server profile ID. Only applicable when &#x60;sslServerCert&#x60; is set to a non-zero value. The Options API &#x60;/api/options/nsxt/nsxtLBServerSSlProfiles?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 

## Methods

### NewLoadBalancerInstanceUpdateConfig

`func NewLoadBalancerInstanceUpdateConfig() *LoadBalancerInstanceUpdateConfig`

NewLoadBalancerInstanceUpdateConfig instantiates a new LoadBalancerInstanceUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerInstanceUpdateConfigWithDefaults

`func NewLoadBalancerInstanceUpdateConfigWithDefaults() *LoadBalancerInstanceUpdateConfig`

NewLoadBalancerInstanceUpdateConfigWithDefaults instantiates a new LoadBalancerInstanceUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationProfile

`func (o *LoadBalancerInstanceUpdateConfig) GetApplicationProfile() string`

GetApplicationProfile returns the ApplicationProfile field if non-nil, zero value otherwise.

### GetApplicationProfileOk

`func (o *LoadBalancerInstanceUpdateConfig) GetApplicationProfileOk() (*string, bool)`

GetApplicationProfileOk returns a tuple with the ApplicationProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProfile

`func (o *LoadBalancerInstanceUpdateConfig) SetApplicationProfile(v string)`

SetApplicationProfile sets ApplicationProfile field to given value.

### HasApplicationProfile

`func (o *LoadBalancerInstanceUpdateConfig) HasApplicationProfile() bool`

HasApplicationProfile returns a boolean if a field has been set.

### GetPool

`func (o *LoadBalancerInstanceUpdateConfig) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *LoadBalancerInstanceUpdateConfig) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *LoadBalancerInstanceUpdateConfig) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *LoadBalancerInstanceUpdateConfig) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetPersistence

`func (o *LoadBalancerInstanceUpdateConfig) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *LoadBalancerInstanceUpdateConfig) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *LoadBalancerInstanceUpdateConfig) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *LoadBalancerInstanceUpdateConfig) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### SetPersistenceNil

`func (o *LoadBalancerInstanceUpdateConfig) SetPersistenceNil(b bool)`

 SetPersistenceNil sets the value for Persistence to be an explicit nil

### UnsetPersistence
`func (o *LoadBalancerInstanceUpdateConfig) UnsetPersistence()`

UnsetPersistence ensures that no value is present for Persistence, not even an explicit nil
### GetPersistenceProfile

`func (o *LoadBalancerInstanceUpdateConfig) GetPersistenceProfile() string`

GetPersistenceProfile returns the PersistenceProfile field if non-nil, zero value otherwise.

### GetPersistenceProfileOk

`func (o *LoadBalancerInstanceUpdateConfig) GetPersistenceProfileOk() (*string, bool)`

GetPersistenceProfileOk returns a tuple with the PersistenceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceProfile

`func (o *LoadBalancerInstanceUpdateConfig) SetPersistenceProfile(v string)`

SetPersistenceProfile sets PersistenceProfile field to given value.

### HasPersistenceProfile

`func (o *LoadBalancerInstanceUpdateConfig) HasPersistenceProfile() bool`

HasPersistenceProfile returns a boolean if a field has been set.

### GetSslClientProfile

`func (o *LoadBalancerInstanceUpdateConfig) GetSslClientProfile() string`

GetSslClientProfile returns the SslClientProfile field if non-nil, zero value otherwise.

### GetSslClientProfileOk

`func (o *LoadBalancerInstanceUpdateConfig) GetSslClientProfileOk() (*string, bool)`

GetSslClientProfileOk returns a tuple with the SslClientProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientProfile

`func (o *LoadBalancerInstanceUpdateConfig) SetSslClientProfile(v string)`

SetSslClientProfile sets SslClientProfile field to given value.

### HasSslClientProfile

`func (o *LoadBalancerInstanceUpdateConfig) HasSslClientProfile() bool`

HasSslClientProfile returns a boolean if a field has been set.

### GetSslServerProfile

`func (o *LoadBalancerInstanceUpdateConfig) GetSslServerProfile() string`

GetSslServerProfile returns the SslServerProfile field if non-nil, zero value otherwise.

### GetSslServerProfileOk

`func (o *LoadBalancerInstanceUpdateConfig) GetSslServerProfileOk() (*string, bool)`

GetSslServerProfileOk returns a tuple with the SslServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerProfile

`func (o *LoadBalancerInstanceUpdateConfig) SetSslServerProfile(v string)`

SetSslServerProfile sets SslServerProfile field to given value.

### HasSslServerProfile

`func (o *LoadBalancerInstanceUpdateConfig) HasSslServerProfile() bool`

HasSslServerProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


