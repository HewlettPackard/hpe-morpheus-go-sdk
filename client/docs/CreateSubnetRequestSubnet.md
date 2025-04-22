# CreateSubnetRequestSubnet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**CreateSubnetRequestSubnetType**](CreateSubnetRequestSubnetType.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object. Settings vary by type. | [optional] 
**Tenants** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of tenant account ID objects that are allowed access | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 

## Methods

### NewCreateSubnetRequestSubnet

`func NewCreateSubnetRequestSubnet() *CreateSubnetRequestSubnet`

NewCreateSubnetRequestSubnet instantiates a new CreateSubnetRequestSubnet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSubnetRequestSubnetWithDefaults

`func NewCreateSubnetRequestSubnetWithDefaults() *CreateSubnetRequestSubnet`

NewCreateSubnetRequestSubnetWithDefaults instantiates a new CreateSubnetRequestSubnet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateSubnetRequestSubnet) GetType() CreateSubnetRequestSubnetType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateSubnetRequestSubnet) GetTypeOk() (*CreateSubnetRequestSubnetType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateSubnetRequestSubnet) SetType(v CreateSubnetRequestSubnetType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateSubnetRequestSubnet) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *CreateSubnetRequestSubnet) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateSubnetRequestSubnet) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateSubnetRequestSubnet) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateSubnetRequestSubnet) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTenants

`func (o *CreateSubnetRequestSubnet) GetTenants() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *CreateSubnetRequestSubnet) GetTenantsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *CreateSubnetRequestSubnet) SetTenants(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *CreateSubnetRequestSubnet) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetVisibility

`func (o *CreateSubnetRequestSubnet) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *CreateSubnetRequestSubnet) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *CreateSubnetRequestSubnet) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *CreateSubnetRequestSubnet) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLabels

`func (o *CreateSubnetRequestSubnet) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CreateSubnetRequestSubnet) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CreateSubnetRequestSubnet) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CreateSubnetRequestSubnet) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


