# UpdateSubnetRequestSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**UpdateSubnetRequestSubnetType**](UpdateSubnetRequestSubnetType.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**Tenants** | Pointer to [**[]UpdateSubnetRequestSubnetTenantsInner**](UpdateSubnetRequestSubnetTenantsInner.md) | Array of tenant account ID objects that are allowed access | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 

## Methods

### NewUpdateSubnetRequestSubnet

`func NewUpdateSubnetRequestSubnet() *UpdateSubnetRequestSubnet`

NewUpdateSubnetRequestSubnet instantiates a new UpdateSubnetRequestSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSubnetRequestSubnetWithDefaults

`func NewUpdateSubnetRequestSubnetWithDefaults() *UpdateSubnetRequestSubnet`

NewUpdateSubnetRequestSubnetWithDefaults instantiates a new UpdateSubnetRequestSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UpdateSubnetRequestSubnet) GetType() UpdateSubnetRequestSubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSubnetRequestSubnet) GetTypeOk() (*UpdateSubnetRequestSubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSubnetRequestSubnet) SetType(v UpdateSubnetRequestSubnetType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSubnetRequestSubnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateSubnetRequestSubnet) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSubnetRequestSubnet) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSubnetRequestSubnet) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSubnetRequestSubnet) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateSubnetRequestSubnet) GetTenants() []UpdateSubnetRequestSubnetTenantsInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateSubnetRequestSubnet) GetTenantsOk() (*[]UpdateSubnetRequestSubnetTenantsInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateSubnetRequestSubnet) SetTenants(v []UpdateSubnetRequestSubnetTenantsInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateSubnetRequestSubnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateSubnetRequestSubnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateSubnetRequestSubnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateSubnetRequestSubnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateSubnetRequestSubnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateSubnetRequestSubnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateSubnetRequestSubnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateSubnetRequestSubnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateSubnetRequestSubnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateSubnetRequestSubnet) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateSubnetRequestSubnet) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


