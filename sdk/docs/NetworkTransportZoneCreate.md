# NetworkTransportZoneCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Network transport zone name | 
**Description** | Pointer to **string** | Network transport zone description | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewNetworkTransportZoneCreate

`func NewNetworkTransportZoneCreate(name string, ) *NetworkTransportZoneCreate`

NewNetworkTransportZoneCreate instantiates a new NetworkTransportZoneCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkTransportZoneCreateWithDefaults

`func NewNetworkTransportZoneCreateWithDefaults() *NetworkTransportZoneCreate`

NewNetworkTransportZoneCreateWithDefaults instantiates a new NetworkTransportZoneCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkTransportZoneCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkTransportZoneCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkTransportZoneCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *NetworkTransportZoneCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkTransportZoneCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkTransportZoneCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkTransportZoneCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkTransportZoneCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkTransportZoneCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkTransportZoneCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkTransportZoneCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *NetworkTransportZoneCreate) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *NetworkTransportZoneCreate) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *NetworkTransportZoneCreate) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *NetworkTransportZoneCreate) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


