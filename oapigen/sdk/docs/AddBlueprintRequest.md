# AddBlueprintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the blueprint | 
**Image** | Pointer to **string** | Path to display image. Defaults to an internal Morpheus image. | [optional] 
**Type** | **string** | Blueprint Type | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Arm** | [**AddBlueprintRequestOneOfArm**](AddBlueprintRequestOneOfArm.md) |  | 
**CloudFormation** | [**AddBlueprintRequestOneOf1CloudFormation**](AddBlueprintRequestOneOf1CloudFormation.md) |  | 
**Helm** | [**AddBlueprintRequestOneOf2Helm**](AddBlueprintRequestOneOf2Helm.md) |  | 
**Kubernetes** | [**AddBlueprintRequestOneOf3Kubernetes**](AddBlueprintRequestOneOf3Kubernetes.md) |  | 
**Config** | Pointer to [**AddBlueprintRequestOneOf5Config**](AddBlueprintRequestOneOf5Config.md) |  | [optional] 
**Tiers** | **map[string]interface{}** | Tier definitions - Create in UI to view a baseline for object | 
**Terraform** | [**AddBlueprintRequestOneOf5Terraform**](AddBlueprintRequestOneOf5Terraform.md) |  | 

## Methods

### NewAddBlueprintRequest

`func NewAddBlueprintRequest(name string, type_ string, arm AddBlueprintRequestOneOfArm, cloudFormation AddBlueprintRequestOneOf1CloudFormation, helm AddBlueprintRequestOneOf2Helm, kubernetes AddBlueprintRequestOneOf3Kubernetes, tiers map[string]interface{}, terraform AddBlueprintRequestOneOf5Terraform, ) *AddBlueprintRequest`

NewAddBlueprintRequest instantiates a new AddBlueprintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBlueprintRequestWithDefaults

`func NewAddBlueprintRequestWithDefaults() *AddBlueprintRequest`

NewAddBlueprintRequestWithDefaults instantiates a new AddBlueprintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddBlueprintRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddBlueprintRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddBlueprintRequest) SetName(v string)`

SetName sets Name field to given value.


### GetImage

`func (o *AddBlueprintRequest) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *AddBlueprintRequest) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *AddBlueprintRequest) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *AddBlueprintRequest) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetType

`func (o *AddBlueprintRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddBlueprintRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddBlueprintRequest) SetType(v string)`

SetType sets Type field to given value.


### GetLabels

`func (o *AddBlueprintRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddBlueprintRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddBlueprintRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddBlueprintRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetArm

`func (o *AddBlueprintRequest) GetArm() AddBlueprintRequestOneOfArm`

GetArm returns the Arm field if non-nil, zero value otherwise.

### GetArmOk

`func (o *AddBlueprintRequest) GetArmOk() (*AddBlueprintRequestOneOfArm, bool)`

GetArmOk returns a tuple with the Arm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArm

`func (o *AddBlueprintRequest) SetArm(v AddBlueprintRequestOneOfArm)`

SetArm sets Arm field to given value.


### GetCloudFormation

`func (o *AddBlueprintRequest) GetCloudFormation() AddBlueprintRequestOneOf1CloudFormation`

GetCloudFormation returns the CloudFormation field if non-nil, zero value otherwise.

### GetCloudFormationOk

`func (o *AddBlueprintRequest) GetCloudFormationOk() (*AddBlueprintRequestOneOf1CloudFormation, bool)`

GetCloudFormationOk returns a tuple with the CloudFormation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudFormation

`func (o *AddBlueprintRequest) SetCloudFormation(v AddBlueprintRequestOneOf1CloudFormation)`

SetCloudFormation sets CloudFormation field to given value.


### GetHelm

`func (o *AddBlueprintRequest) GetHelm() AddBlueprintRequestOneOf2Helm`

GetHelm returns the Helm field if non-nil, zero value otherwise.

### GetHelmOk

`func (o *AddBlueprintRequest) GetHelmOk() (*AddBlueprintRequestOneOf2Helm, bool)`

GetHelmOk returns a tuple with the Helm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelm

`func (o *AddBlueprintRequest) SetHelm(v AddBlueprintRequestOneOf2Helm)`

SetHelm sets Helm field to given value.


### GetKubernetes

`func (o *AddBlueprintRequest) GetKubernetes() AddBlueprintRequestOneOf3Kubernetes`

GetKubernetes returns the Kubernetes field if non-nil, zero value otherwise.

### GetKubernetesOk

`func (o *AddBlueprintRequest) GetKubernetesOk() (*AddBlueprintRequestOneOf3Kubernetes, bool)`

GetKubernetesOk returns a tuple with the Kubernetes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetes

`func (o *AddBlueprintRequest) SetKubernetes(v AddBlueprintRequestOneOf3Kubernetes)`

SetKubernetes sets Kubernetes field to given value.


### GetConfig

`func (o *AddBlueprintRequest) GetConfig() AddBlueprintRequestOneOf5Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddBlueprintRequest) GetConfigOk() (*AddBlueprintRequestOneOf5Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddBlueprintRequest) SetConfig(v AddBlueprintRequestOneOf5Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddBlueprintRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetTiers

`func (o *AddBlueprintRequest) GetTiers() map[string]interface{}`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *AddBlueprintRequest) GetTiersOk() (*map[string]interface{}, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *AddBlueprintRequest) SetTiers(v map[string]interface{})`

SetTiers sets Tiers field to given value.


### GetTerraform

`func (o *AddBlueprintRequest) GetTerraform() AddBlueprintRequestOneOf5Terraform`

GetTerraform returns the Terraform field if non-nil, zero value otherwise.

### GetTerraformOk

`func (o *AddBlueprintRequest) GetTerraformOk() (*AddBlueprintRequestOneOf5Terraform, bool)`

GetTerraformOk returns a tuple with the Terraform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerraform

`func (o *AddBlueprintRequest) SetTerraform(v AddBlueprintRequestOneOf5Terraform)`

SetTerraform sets Terraform field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


