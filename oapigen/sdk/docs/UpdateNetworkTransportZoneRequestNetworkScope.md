# UpdateNetworkTransportZoneRequestNetworkScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Network transport zone name | 
**Description** | Pointer to **NullableString** | Network transport zone description | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] 
**Tenants** | Pointer to [**[]UpdateNetworkTransportZoneRequestNetworkScopeTenantsInner**](UpdateNetworkTransportZoneRequestNetworkScopeTenantsInner.md) | Array of tenant account ids that are allowed access | [optional] 

## Methods

### NewUpdateNetworkTransportZoneRequestNetworkScope

`func NewUpdateNetworkTransportZoneRequestNetworkScope(name string, ) *UpdateNetworkTransportZoneRequestNetworkScope`

NewUpdateNetworkTransportZoneRequestNetworkScope instantiates a new UpdateNetworkTransportZoneRequestNetworkScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkTransportZoneRequestNetworkScopeWithDefaults

`func NewUpdateNetworkTransportZoneRequestNetworkScopeWithDefaults() *UpdateNetworkTransportZoneRequestNetworkScope`

NewUpdateNetworkTransportZoneRequestNetworkScopeWithDefaults instantiates a new UpdateNetworkTransportZoneRequestNetworkScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateNetworkTransportZoneRequestNetworkScope) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetVisibility

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetTenants() []UpdateNetworkTransportZoneRequestNetworkScopeTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) GetTenantsOk() (*[]UpdateNetworkTransportZoneRequestNetworkScopeTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) SetTenants(v []UpdateNetworkTransportZoneRequestNetworkScopeTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateNetworkTransportZoneRequestNetworkScope) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


