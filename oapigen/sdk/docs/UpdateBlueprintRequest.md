# UpdateBlueprintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**UpdateBlueprintRequestOneOfArm**](UpdateBlueprintRequestOneOfArm.md) |  | 
**CloudFormation** | [**UpdateBlueprintRequestOneOf1CloudFormation**](UpdateBlueprintRequestOneOf1CloudFormation.md) |  | 
**Helm** | [**UpdateBlueprintRequestOneOf2Helm**](UpdateBlueprintRequestOneOf2Helm.md) |  | 
**Kubernetes** | [**UpdateBlueprintRequestOneOf3Kubernetes**](UpdateBlueprintRequestOneOf3Kubernetes.md) |  | 
**Config** | Pointer to [**UpdateBlueprintRequestOneOf5Config**](UpdateBlueprintRequestOneOf5Config.md) |  | [optional] 
**Tiers** | **map[string]interface{}** | Tier definitions - Create in UI to view a baseline for object | 
**Terraform** | [**UpdateBlueprintRequestOneOf5Terraform**](UpdateBlueprintRequestOneOf5Terraform.md) |  | 

## Methods

### NewUpdateBlueprintRequest

`func NewUpdateBlueprintRequest(name string, type_ string, arm UpdateBlueprintRequestOneOfArm, cloudFormation UpdateBlueprintRequestOneOf1CloudFormation, helm UpdateBlueprintRequestOneOf2Helm, kubernetes UpdateBlueprintRequestOneOf3Kubernetes, tiers map[string]interface{}, terraform UpdateBlueprintRequestOneOf5Terraform, ) *UpdateBlueprintRequest`

NewUpdateBlueprintRequest instantiates a new UpdateBlueprintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBlueprintRequestWithDefaults

`func NewUpdateBlueprintRequestWithDefaults() *UpdateBlueprintRequest`

NewUpdateBlueprintRequestWithDefaults instantiates a new UpdateBlueprintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateBlueprintRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBlueprintRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBlueprintRequest) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *UpdateBlueprintRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *UpdateBlueprintRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *UpdateBlueprintRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *UpdateBlueprintRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *UpdateBlueprintRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateBlueprintRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateBlueprintRequest) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *UpdateBlueprintRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateBlueprintRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateBlueprintRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateBlueprintRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetArm

`func (o *UpdateBlueprintRequest) GetArm() UpdateBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *UpdateBlueprintRequest) GetArmOk() (*UpdateBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *UpdateBlueprintRequest) SetArm(v UpdateBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.


### GetCloudFormation

`func (o *UpdateBlueprintRequest) GetCloudFormation() UpdateBlueprintRequestOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *UpdateBlueprintRequest) GetCloudFormationOk() (*UpdateBlueprintRequestOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *UpdateBlueprintRequest) SetCloudFormation(v UpdateBlueprintRequestOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.


### GetHelm

`func (o *UpdateBlueprintRequest) GetHelm() UpdateBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *UpdateBlueprintRequest) GetHelmOk() (*UpdateBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *UpdateBlueprintRequest) SetHelm(v UpdateBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.


### GetKubernetes

`func (o *UpdateBlueprintRequest) GetKubernetes() UpdateBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *UpdateBlueprintRequest) GetKubernetesOk() (*UpdateBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *UpdateBlueprintRequest) SetKubernetes(v UpdateBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.


### GetConfig

`func (o *UpdateBlueprintRequest) GetConfig() UpdateBlueprintRequestOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateBlueprintRequest) GetConfigOk() (*UpdateBlueprintRequestOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateBlueprintRequest) SetConfig(v UpdateBlueprintRequestOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateBlueprintRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTiers

`func (o *UpdateBlueprintRequest) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *UpdateBlueprintRequest) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *UpdateBlueprintRequest) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.


### GetTerraform

`func (o *UpdateBlueprintRequest) GetTerraform() UpdateBlueprintRequestOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *UpdateBlueprintRequest) GetTerraformOk() (*UpdateBlueprintRequestOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *UpdateBlueprintRequest) SetTerraform(v UpdateBlueprintRequestOneOf5Terraform)`

SetTerraform sets Terraform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


