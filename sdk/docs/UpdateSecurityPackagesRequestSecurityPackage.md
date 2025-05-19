# UpdateSecurityPackagesRequestSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the security package | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Security Package Type Code | [optional] [default to "scap-package"]
**Description** | Pointer to **string** | A description for the security package | [optional] 
**Url** | Pointer to **string** | URL to download the security package zip file from | [optional] 
**Enabled** | Pointer to **bool** | Can be used to disable the security package | [optional] 

## Methods

### NewUpdateSecurityPackagesRequestSecurityPackage

`func NewUpdateSecurityPackagesRequestSecurityPackage() *UpdateSecurityPackagesRequestSecurityPackage`

NewUpdateSecurityPackagesRequestSecurityPackage instantiates a new UpdateSecurityPackagesRequestSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityPackagesRequestSecurityPackageWithDefaults

`func NewUpdateSecurityPackagesRequestSecurityPackageWithDefaults() *UpdateSecurityPackagesRequestSecurityPackage`

NewUpdateSecurityPackagesRequestSecurityPackageWithDefaults instantiates a new UpdateSecurityPackagesRequestSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSecurityPackagesRequestSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSecurityPackagesRequestSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSecurityPackagesRequestSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


