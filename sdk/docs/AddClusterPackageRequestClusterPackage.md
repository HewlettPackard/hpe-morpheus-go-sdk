# AddClusterPackageRequestClusterPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Cluster Package name | 
**Description** | Pointer to **string** | Cluster Package description | [optional] 
**Code** | **string** | Cluster Package code | 
**PackageVersion** | **string** | Version of the cluster package | 
**PackageType** | **string** | Package Type | 
**Type** | **string** | type | 
**Enabled** | **bool** | Can be used to enable / disable the cluster package. | [default to true]
**IconPath** | Pointer to **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**SpecTemplates** | **[]int64** | Array of resource spec templates | 

## Methods

### NewAddClusterPackageRequestClusterPackage

`func NewAddClusterPackageRequestClusterPackage(name string, code string, packageVersion string, packageType string, type_ string, enabled bool, specTemplates []int64, ) *AddClusterPackageRequestClusterPackage`

NewAddClusterPackageRequestClusterPackage instantiates a new AddClusterPackageRequestClusterPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterPackageRequestClusterPackageWithDefaults

`func NewAddClusterPackageRequestClusterPackageWithDefaults() *AddClusterPackageRequestClusterPackage`

NewAddClusterPackageRequestClusterPackageWithDefaults instantiates a new AddClusterPackageRequestClusterPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddClusterPackageRequestClusterPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClusterPackageRequestClusterPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClusterPackageRequestClusterPackage) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddClusterPackageRequestClusterPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddClusterPackageRequestClusterPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddClusterPackageRequestClusterPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddClusterPackageRequestClusterPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *AddClusterPackageRequestClusterPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddClusterPackageRequestClusterPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddClusterPackageRequestClusterPackage) SetCode(v string)`

SetCode sets Code field to given value.


### GetPackageVersion

`func (o *AddClusterPackageRequestClusterPackage) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *AddClusterPackageRequestClusterPackage) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *AddClusterPackageRequestClusterPackage) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.


### GetPackageType

`func (o *AddClusterPackageRequestClusterPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *AddClusterPackageRequestClusterPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *AddClusterPackageRequestClusterPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.


### GetType

`func (o *AddClusterPackageRequestClusterPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddClusterPackageRequestClusterPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddClusterPackageRequestClusterPackage) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *AddClusterPackageRequestClusterPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddClusterPackageRequestClusterPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddClusterPackageRequestClusterPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetIconPath

`func (o *AddClusterPackageRequestClusterPackage) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *AddClusterPackageRequestClusterPackage) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *AddClusterPackageRequestClusterPackage) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *AddClusterPackageRequestClusterPackage) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *AddClusterPackageRequestClusterPackage) GetSpecTemplates() []int64`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *AddClusterPackageRequestClusterPackage) GetSpecTemplatesOk() (*[]int64, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *AddClusterPackageRequestClusterPackage) SetSpecTemplates(v []int64)`

SetSpecTemplates sets SpecTemplates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


