# AddBlueprint200ResponseAllOfBlueprintConfigOneOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Kubernetes** | Pointer to [**AddBlueprintRequestOneOf3Kubernetes**](AddBlueprintRequestOneOf3Kubernetes.md) |  | [optional] 
**Config** | Pointer to [**AddBlueprintRequestOneOf3Config**](AddBlueprintRequestOneOf3Config.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 

## Methods

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3 instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3WithDefaults

`func NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3WithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3`

NewAddBlueprint200ResponseAllOfBlueprintConfigOneOf3WithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfigOneOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetKubernetes() AddBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetKubernetesOk() (*AddBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetKubernetes(v AddBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetConfig() AddBlueprintRequestOneOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetConfigOk() (*AddBlueprintRequestOneOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetConfig(v AddBlueprintRequestOneOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfigOneOf3) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


