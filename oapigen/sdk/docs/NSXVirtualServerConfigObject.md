# NSXVirtualServerConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationProfile** | Pointer to **NullableInt64** | The Load Balancer Application Profile ID. The Options API &#x60;/api/options/nsxt/nsxtLBVirtualServerApplicationProfile?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available. | [optional] 
**Pool** | Pointer to **NullableString** | The backend server pool ID (&#x60;NetworkLoadBalancerPool&#x60;). The Options API &#x60;/api/options/nsxt/nsxtLBPool?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 
**Persistence** | Pointer to **NullableString** | Session persistence mode. The available values depend on the virtual server protocol. For HTTP: &#x60;SOURCE_IP&#x60;, &#x60;COOKIE&#x60;, or empty string (disabled). For TCP/UDP: &#x60;SOURCE_IP&#x60; or empty string (disabled). The Options API &#x60;/api/options/nsxt/nsxtLBPersistence?loadBalancerId&#x3D;42&amp;loadBalancerInstance.vipProtocol&#x3D;tcp&#x60; can be used to see which options are available. | [optional] 
**PersistenceProfile** | Pointer to **NullableInt64** | The ID of the persistence profile to use. Required when &#x60;persistence&#x60; is set to a non-empty value (&#x60;SOURCE_IP&#x60; or &#x60;COOKIE&#x60;). The Options API &#x60;/api/options/nsxt/nsxtLBPersistenceProfile?loadBalancerId&#x3D;42&amp;config.persistence&#x3D;SOURCE_IP&#x60; can be used to see which options are available. | [optional] 
**SslClientProfile** | Pointer to **NullableInt64** | The SSL client profile ID. Only applicable when &#x60;sslCert&#x60; is set to a non-zero value. The Options API &#x60;/api/options/nsxt/nsxtLBClientSSlProfiles?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 
**SslServerProfile** | Pointer to **NullableInt64** | The SSL server profile ID. Only applicable when &#x60;sslServerCert&#x60; is set to a non-zero value. The Options API &#x60;/api/options/nsxt/nsxtLBServerSSlProfiles?loadBalancerId&#x3D;42&#x60; can be used to see which options are available. | [optional] 

## Methods

### NewNSXVirtualServerConfigObject

`func NewNSXVirtualServerConfigObject() *NSXVirtualServerConfigObject`

NewNSXVirtualServerConfigObject instantiates a new NSXVirtualServerConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXVirtualServerConfigObjectWithDefaults

`func NewNSXVirtualServerConfigObjectWithDefaults() *NSXVirtualServerConfigObject`

NewNSXVirtualServerConfigObjectWithDefaults instantiates a new NSXVirtualServerConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationProfile

`func (o *NSXVirtualServerConfigObject) GetApplicationProfile() int64`

GetApplicationProfile returns the ApplicationProfile field if non-nil, zero value otherwise.

### GetApplicationProfileOk

`func (o *NSXVirtualServerConfigObject) GetApplicationProfileOk() (*int64, bool)`

GetApplicationProfileOk returns a tuple with the ApplicationProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationProfile

`func (o *NSXVirtualServerConfigObject) SetApplicationProfile(v int64)`

SetApplicationProfile sets ApplicationProfile field to given value.

### HasApplicationProfile

`func (o *NSXVirtualServerConfigObject) HasApplicationProfile() bool`

HasApplicationProfile returns a boolean if a field has been set.

### SetApplicationProfileNil

`func (o *NSXVirtualServerConfigObject) SetApplicationProfileNil(b bool)`

 SetApplicationProfileNil sets the value for ApplicationProfile to be an explicit nil

### UnsetApplicationProfile
`func (o *NSXVirtualServerConfigObject) UnsetApplicationProfile()`

UnsetApplicationProfile ensures that no value is present for ApplicationProfile, not even an explicit nil
### GetPool

`func (o *NSXVirtualServerConfigObject) GetPool() string`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *NSXVirtualServerConfigObject) GetPoolOk() (*string, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *NSXVirtualServerConfigObject) SetPool(v string)`

SetPool sets Pool field to given value.

### HasPool

`func (o *NSXVirtualServerConfigObject) HasPool() bool`

HasPool returns a boolean if a field has been set.

### SetPoolNil

`func (o *NSXVirtualServerConfigObject) SetPoolNil(b bool)`

 SetPoolNil sets the value for Pool to be an explicit nil

### UnsetPool
`func (o *NSXVirtualServerConfigObject) UnsetPool()`

UnsetPool ensures that no value is present for Pool, not even an explicit nil
### GetPersistence

`func (o *NSXVirtualServerConfigObject) GetPersistence() string`

GetPersistence returns the Persistence field if non-nil, zero value otherwise.

### GetPersistenceOk

`func (o *NSXVirtualServerConfigObject) GetPersistenceOk() (*string, bool)`

GetPersistenceOk returns a tuple with the Persistence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistence

`func (o *NSXVirtualServerConfigObject) SetPersistence(v string)`

SetPersistence sets Persistence field to given value.

### HasPersistence

`func (o *NSXVirtualServerConfigObject) HasPersistence() bool`

HasPersistence returns a boolean if a field has been set.

### SetPersistenceNil

`func (o *NSXVirtualServerConfigObject) SetPersistenceNil(b bool)`

 SetPersistenceNil sets the value for Persistence to be an explicit nil

### UnsetPersistence
`func (o *NSXVirtualServerConfigObject) UnsetPersistence()`

UnsetPersistence ensures that no value is present for Persistence, not even an explicit nil
### GetPersistenceProfile

`func (o *NSXVirtualServerConfigObject) GetPersistenceProfile() int64`

GetPersistenceProfile returns the PersistenceProfile field if non-nil, zero value otherwise.

### GetPersistenceProfileOk

`func (o *NSXVirtualServerConfigObject) GetPersistenceProfileOk() (*int64, bool)`

GetPersistenceProfileOk returns a tuple with the PersistenceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPersistenceProfile

`func (o *NSXVirtualServerConfigObject) SetPersistenceProfile(v int64)`

SetPersistenceProfile sets PersistenceProfile field to given value.

### HasPersistenceProfile

`func (o *NSXVirtualServerConfigObject) HasPersistenceProfile() bool`

HasPersistenceProfile returns a boolean if a field has been set.

### SetPersistenceProfileNil

`func (o *NSXVirtualServerConfigObject) SetPersistenceProfileNil(b bool)`

 SetPersistenceProfileNil sets the value for PersistenceProfile to be an explicit nil

### UnsetPersistenceProfile
`func (o *NSXVirtualServerConfigObject) UnsetPersistenceProfile()`

UnsetPersistenceProfile ensures that no value is present for PersistenceProfile, not even an explicit nil
### GetSslClientProfile

`func (o *NSXVirtualServerConfigObject) GetSslClientProfile() int64`

GetSslClientProfile returns the SslClientProfile field if non-nil, zero value otherwise.

### GetSslClientProfileOk

`func (o *NSXVirtualServerConfigObject) GetSslClientProfileOk() (*int64, bool)`

GetSslClientProfileOk returns a tuple with the SslClientProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslClientProfile

`func (o *NSXVirtualServerConfigObject) SetSslClientProfile(v int64)`

SetSslClientProfile sets SslClientProfile field to given value.

### HasSslClientProfile

`func (o *NSXVirtualServerConfigObject) HasSslClientProfile() bool`

HasSslClientProfile returns a boolean if a field has been set.

### SetSslClientProfileNil

`func (o *NSXVirtualServerConfigObject) SetSslClientProfileNil(b bool)`

 SetSslClientProfileNil sets the value for SslClientProfile to be an explicit nil

### UnsetSslClientProfile
`func (o *NSXVirtualServerConfigObject) UnsetSslClientProfile()`

UnsetSslClientProfile ensures that no value is present for SslClientProfile, not even an explicit nil
### GetSslServerProfile

`func (o *NSXVirtualServerConfigObject) GetSslServerProfile() int64`

GetSslServerProfile returns the SslServerProfile field if non-nil, zero value otherwise.

### GetSslServerProfileOk

`func (o *NSXVirtualServerConfigObject) GetSslServerProfileOk() (*int64, bool)`

GetSslServerProfileOk returns a tuple with the SslServerProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerProfile

`func (o *NSXVirtualServerConfigObject) SetSslServerProfile(v int64)`

SetSslServerProfile sets SslServerProfile field to given value.

### HasSslServerProfile

`func (o *NSXVirtualServerConfigObject) HasSslServerProfile() bool`

HasSslServerProfile returns a boolean if a field has been set.

### SetSslServerProfileNil

`func (o *NSXVirtualServerConfigObject) SetSslServerProfileNil(b bool)`

 SetSslServerProfileNil sets the value for SslServerProfile to be an explicit nil

### UnsetSslServerProfile
`func (o *NSXVirtualServerConfigObject) UnsetSslServerProfile()`

UnsetSslServerProfile ensures that no value is present for SslServerProfile, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


