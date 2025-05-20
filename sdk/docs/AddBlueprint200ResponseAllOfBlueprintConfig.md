# AddBlueprint200ResponseAllOfBlueprintConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Arm** | Pointer to [**AddBlueprintRequestOneOfArm**](AddBlueprintRequestOneOfArm.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 
**CloudFormation** | Pointer to [**AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation**](AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation.md) |  | [optional] 
**Helm** | Pointer to [**AddBlueprintRequestOneOf2Helm**](AddBlueprintRequestOneOf2Helm.md) |  | [optional] 
**Kubernetes** | Pointer to [**AddBlueprintRequestOneOf3Kubernetes**](AddBlueprintRequestOneOf3Kubernetes.md) |  | [optional] 
**Config** | Pointer to [**AddBlueprintRequestOneOf5Config**](AddBlueprintRequestOneOf5Config.md) |  | [optional] 
**Terraform** | Pointer to [**AddBlueprintRequestOneOf5Terraform**](AddBlueprintRequestOneOf5Terraform.md) |  | [optional] 

## Methods

### NewAddBlueprint200ResponseAllOfBlueprintConfig

`func NewAddBlueprint200ResponseAllOfBlueprintConfig() *AddBlueprint200ResponseAllOfBlueprintConfig`

NewAddBlueprint200ResponseAllOfBlueprintConfig instantiates a new AddBlueprint200ResponseAllOfBlueprintConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprint200ResponseAllOfBlueprintConfigWithDefaults

`func NewAddBlueprint200ResponseAllOfBlueprintConfigWithDefaults() *AddBlueprint200ResponseAllOfBlueprintConfig`

NewAddBlueprint200ResponseAllOfBlueprintConfigWithDefaults instantiates a new AddBlueprint200ResponseAllOfBlueprintConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetArm() AddBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetArmOk() (*AddBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetArm(v AddBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.

### HasArm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasArm() bool`

HasArm returns a boolean if a field has been set.

### GetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCloudFormation

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetCloudFormation() AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetCloudFormationOk() (*AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetCloudFormation(v AddBlueprint200ResponseAllOfBlueprintConfigOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.

### HasCloudFormation

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasCloudFormation() bool`

HasCloudFormation returns a boolean if a field has been set.

### GetHelm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetHelm() AddBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetHelmOk() (*AddBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetHelm(v AddBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.

### HasHelm

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasHelm() bool`

HasHelm returns a boolean if a field has been set.

### GetKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetKubernetes() AddBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetKubernetesOk() (*AddBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetKubernetes(v AddBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetConfig() AddBlueprintRequestOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetConfigOk() (*AddBlueprintRequestOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetConfig(v AddBlueprintRequestOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTerraform

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetTerraform() AddBlueprintRequestOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) GetTerraformOk() (*AddBlueprintRequestOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) SetTerraform(v AddBlueprintRequestOneOf5Terraform)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *AddBlueprint200ResponseAllOfBlueprintConfig) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


