# UpdateClusterPackageRequestClusterPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Cluster Package name | [optional] 
**Description** | Pointer to **string** | Cluster Package description | [optional] 
**Code** | Pointer to **string** | Cluster Package code | [optional] 
**PackageVersion** | Pointer to **string** | Version of the cluster package | [optional] 
**PackageType** | Pointer to **string** | Package Type | [optional] 
**Type** | Pointer to **string** | type | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the cluster package. | [optional] [default to true]
**IconPath** | Pointer to **string** | Icon Path, relative location of an icon image, eg. /assets/containers-png/nginx.png. | [optional] 
**SpecTemplates** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of resource spec templates | [optional] 

## Methods

### NewUpdateClusterPackageRequestClusterPackage

`func NewUpdateClusterPackageRequestClusterPackage() *UpdateClusterPackageRequestClusterPackage`

NewUpdateClusterPackageRequestClusterPackage instantiates a new UpdateClusterPackageRequestClusterPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterPackageRequestClusterPackageWithDefaults

`func NewUpdateClusterPackageRequestClusterPackageWithDefaults() *UpdateClusterPackageRequestClusterPackage`

NewUpdateClusterPackageRequestClusterPackageWithDefaults instantiates a new UpdateClusterPackageRequestClusterPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateClusterPackageRequestClusterPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateClusterPackageRequestClusterPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateClusterPackageRequestClusterPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateClusterPackageRequestClusterPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateClusterPackageRequestClusterPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateClusterPackageRequestClusterPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCode

`func (o *UpdateClusterPackageRequestClusterPackage) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateClusterPackageRequestClusterPackage) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateClusterPackageRequestClusterPackage) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPackageVersion

`func (o *UpdateClusterPackageRequestClusterPackage) GetPackageVersion() string`

GetPackageVersion returns the PackageVersion field if non-nil, zero value otherwise.

### GetPackageVersionOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetPackageVersionOk() (*string, bool)`

GetPackageVersionOk returns a tuple with the PackageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageVersion

`func (o *UpdateClusterPackageRequestClusterPackage) SetPackageVersion(v string)`

SetPackageVersion sets PackageVersion field to given value.

### HasPackageVersion

`func (o *UpdateClusterPackageRequestClusterPackage) HasPackageVersion() bool`

HasPackageVersion returns a boolean if a field has been set.

### GetPackageType

`func (o *UpdateClusterPackageRequestClusterPackage) GetPackageType() string`

GetPackageType returns the PackageType field if non-nil, zero value otherwise.

### GetPackageTypeOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetPackageTypeOk() (*string, bool)`

GetPackageTypeOk returns a tuple with the PackageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageType

`func (o *UpdateClusterPackageRequestClusterPackage) SetPackageType(v string)`

SetPackageType sets PackageType field to given value.

### HasPackageType

`func (o *UpdateClusterPackageRequestClusterPackage) HasPackageType() bool`

HasPackageType returns a boolean if a field has been set.

### GetType

`func (o *UpdateClusterPackageRequestClusterPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateClusterPackageRequestClusterPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateClusterPackageRequestClusterPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateClusterPackageRequestClusterPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateClusterPackageRequestClusterPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateClusterPackageRequestClusterPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIconPath

`func (o *UpdateClusterPackageRequestClusterPackage) GetIconPath() string`

GetIconPath returns the IconPath field if non-nil, zero value otherwise.

### GetIconPathOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetIconPathOk() (*string, bool)`

GetIconPathOk returns a tuple with the IconPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconPath

`func (o *UpdateClusterPackageRequestClusterPackage) SetIconPath(v string)`

SetIconPath sets IconPath field to given value.

### HasIconPath

`func (o *UpdateClusterPackageRequestClusterPackage) HasIconPath() bool`

HasIconPath returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *UpdateClusterPackageRequestClusterPackage) GetSpecTemplates() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *UpdateClusterPackageRequestClusterPackage) GetSpecTemplatesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *UpdateClusterPackageRequestClusterPackage) SetSpecTemplates(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *UpdateClusterPackageRequestClusterPackage) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


