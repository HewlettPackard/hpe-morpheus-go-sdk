# BlueprintCreateSuccessConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the blueprint | [optional] 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | Pointer to **string** | Blueprint Type | [optional] 
**Arm** | Pointer to [**BlueprintCreateSuccessConfigOneOfArm**](BlueprintCreateSuccessConfigOneOfArm.md) |  | [optional] 
**Visibility** | Pointer to **string** | Private or Public Access | [optional] [default to "private"]
**ResourcePermission** | Pointer to **map[string]interface{}** | Resource Permission Block | [optional] 
**Owner** | Pointer to **map[string]interface{}** | Owner | [optional] 
**Tenant** | Pointer to **map[string]interface{}** | Tenant | [optional] 
**CloudFormation** | Pointer to [**BlueprintCreateSuccessConfigOneOf1CloudFormation**](BlueprintCreateSuccessConfigOneOf1CloudFormation.md) |  | [optional] 
**Helm** | Pointer to [**BlueprintCreateSuccessConfigOneOf2Helm**](BlueprintCreateSuccessConfigOneOf2Helm.md) |  | [optional] 
**Kubernetes** | Pointer to [**BlueprintCreateSuccessConfigOneOf3Kubernetes**](BlueprintCreateSuccessConfigOneOf3Kubernetes.md) |  | [optional] 
**Config** | Pointer to [**BlueprintCreateSuccessConfigOneOf5Config**](BlueprintCreateSuccessConfigOneOf5Config.md) |  | [optional] 
**Terraform** | Pointer to [**BlueprintCreateSuccessConfigOneOf5Terraform**](BlueprintCreateSuccessConfigOneOf5Terraform.md) |  | [optional] 

## Methods

### NewBlueprintCreateSuccessConfig

`func NewBlueprintCreateSuccessConfig() *BlueprintCreateSuccessConfig`

NewBlueprintCreateSuccessConfig instantiates a new BlueprintCreateSuccessConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *BlueprintCreateSuccessConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BlueprintCreateSuccessConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BlueprintCreateSuccessConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BlueprintCreateSuccessConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetImage

`func (o *BlueprintCreateSuccessConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *BlueprintCreateSuccessConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *BlueprintCreateSuccessConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *BlueprintCreateSuccessConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *BlueprintCreateSuccessConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlueprintCreateSuccessConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlueprintCreateSuccessConfig) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BlueprintCreateSuccessConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetArm

`func (o *BlueprintCreateSuccessConfig) GetArm() BlueprintCreateSuccessConfigOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *BlueprintCreateSuccessConfig) GetArmOk() (*BlueprintCreateSuccessConfigOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *BlueprintCreateSuccessConfig) SetArm(v BlueprintCreateSuccessConfigOneOfArm)`

SetArm sets Arm field to given value.

### HasArm

`func (o *BlueprintCreateSuccessConfig) HasArm() bool`

HasArm returns a boolean if a field has been set.

### GetVisibility

`func (o *BlueprintCreateSuccessConfig) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *BlueprintCreateSuccessConfig) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *BlueprintCreateSuccessConfig) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *BlueprintCreateSuccessConfig) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetResourcePermission

`func (o *BlueprintCreateSuccessConfig) GetResourcePermission() map[string]interface{}`

GetResourcePermission returns the ResourcePermission field if non-nil, zero value otherwise.

### GetResourcePermissionOk

`func (o *BlueprintCreateSuccessConfig) GetResourcePermissionOk() (*map[string]interface{}, bool)`

GetResourcePermissionOk returns a tuple with the ResourcePermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePermission

`func (o *BlueprintCreateSuccessConfig) SetResourcePermission(v map[string]interface{})`

SetResourcePermission sets ResourcePermission field to given value.

### HasResourcePermission

`func (o *BlueprintCreateSuccessConfig) HasResourcePermission() bool`

HasResourcePermission returns a boolean if a field has been set.

### GetOwner

`func (o *BlueprintCreateSuccessConfig) GetOwner() map[string]interface{}`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *BlueprintCreateSuccessConfig) GetOwnerOk() (*map[string]interface{}, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *BlueprintCreateSuccessConfig) SetOwner(v map[string]interface{})`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *BlueprintCreateSuccessConfig) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetTenant

`func (o *BlueprintCreateSuccessConfig) GetTenant() map[string]interface{}`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *BlueprintCreateSuccessConfig) GetTenantOk() (*map[string]interface{}, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *BlueprintCreateSuccessConfig) SetTenant(v map[string]interface{})`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *BlueprintCreateSuccessConfig) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetCloudFormation

`func (o *BlueprintCreateSuccessConfig) GetCloudFormation() BlueprintCreateSuccessConfigOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *BlueprintCreateSuccessConfig) GetCloudFormationOk() (*BlueprintCreateSuccessConfigOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *BlueprintCreateSuccessConfig) SetCloudFormation(v BlueprintCreateSuccessConfigOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.

### HasCloudFormation

`func (o *BlueprintCreateSuccessConfig) HasCloudFormation() bool`

HasCloudFormation returns a boolean if a field has been set.

### GetHelm

`func (o *BlueprintCreateSuccessConfig) GetHelm() BlueprintCreateSuccessConfigOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *BlueprintCreateSuccessConfig) GetHelmOk() (*BlueprintCreateSuccessConfigOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *BlueprintCreateSuccessConfig) SetHelm(v BlueprintCreateSuccessConfigOneOf2Helm)`

SetHelm sets Helm field to given value.

### HasHelm

`func (o *BlueprintCreateSuccessConfig) HasHelm() bool`

HasHelm returns a boolean if a field has been set.

### GetKubernetes

`func (o *BlueprintCreateSuccessConfig) GetKubernetes() BlueprintCreateSuccessConfigOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *BlueprintCreateSuccessConfig) GetKubernetesOk() (*BlueprintCreateSuccessConfigOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *BlueprintCreateSuccessConfig) SetKubernetes(v BlueprintCreateSuccessConfigOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.

### HasKubernetes

`func (o *BlueprintCreateSuccessConfig) HasKubernetes() bool`

HasKubernetes returns a boolean if a field has been set.

### GetConfig

`func (o *BlueprintCreateSuccessConfig) GetConfig() BlueprintCreateSuccessConfigOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *BlueprintCreateSuccessConfig) GetConfigOk() (*BlueprintCreateSuccessConfigOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *BlueprintCreateSuccessConfig) SetConfig(v BlueprintCreateSuccessConfigOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *BlueprintCreateSuccessConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTerraform

`func (o *BlueprintCreateSuccessConfig) GetTerraform() BlueprintCreateSuccessConfigOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *BlueprintCreateSuccessConfig) GetTerraformOk() (*BlueprintCreateSuccessConfigOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *BlueprintCreateSuccessConfig) SetTerraform(v BlueprintCreateSuccessConfigOneOf5Terraform)`

SetTerraform sets Terraform field to given value.

### HasTerraform

`func (o *BlueprintCreateSuccessConfig) HasTerraform() bool`

HasTerraform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


