# NetworkServerCreateNSX

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type Code (nsx-t) | 
**Name** | **string** | Name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network server. | [optional] [default to true]
**ServiceUrl** | **string** | URL | 
**ServiceUsername** | Pointer to **string** | Username | [optional] 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**Config** | Pointer to [**NSXNetworkServerConfig**](NSXNetworkServerConfig.md) |  | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**ZoneId** | **int64** | ID of the Cloud that the NSX network server integration should be attached to. | 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewNetworkServerCreateNSX

`func NewNetworkServerCreateNSX(type_ string, name string, serviceUrl string, zoneId int64, ) *NetworkServerCreateNSX`

NewNetworkServerCreateNSX instantiates a new NetworkServerCreateNSX object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServerCreateNSXWithDefaults

`func NewNetworkServerCreateNSXWithDefaults() *NetworkServerCreateNSX`

NewNetworkServerCreateNSXWithDefaults instantiates a new NetworkServerCreateNSX object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NetworkServerCreateNSX) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkServerCreateNSX) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkServerCreateNSX) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *NetworkServerCreateNSX) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServerCreateNSX) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServerCreateNSX) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *NetworkServerCreateNSX) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkServerCreateNSX) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkServerCreateNSX) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkServerCreateNSX) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkServerCreateNSX) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkServerCreateNSX) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkServerCreateNSX) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *NetworkServerCreateNSX) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkServerCreateNSX) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkServerCreateNSX) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkServerCreateNSX) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *NetworkServerCreateNSX) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkServerCreateNSX) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkServerCreateNSX) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkServerCreateNSX) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkServerCreateNSX) GetConfig() NSXNetworkServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkServerCreateNSX) GetConfigOk() (*NSXNetworkServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkServerCreateNSX) SetConfig(v NSXNetworkServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkServerCreateNSX) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkServerCreateNSX) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkServerCreateNSX) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkServerCreateNSX) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkServerCreateNSX) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetZoneId

`func (o *NetworkServerCreateNSX) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *NetworkServerCreateNSX) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *NetworkServerCreateNSX) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.


### GetCredential

`func (o *NetworkServerCreateNSX) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkServerCreateNSX) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkServerCreateNSX) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkServerCreateNSX) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *NetworkServerCreateNSX) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *NetworkServerCreateNSX) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *NetworkServerCreateNSX) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *NetworkServerCreateNSX) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


