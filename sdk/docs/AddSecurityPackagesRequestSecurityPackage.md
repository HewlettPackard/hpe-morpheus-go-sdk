# AddSecurityPackagesRequestSecurityPackage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A name for the security package | 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Type** | Pointer to **string** | Security Package Type Code | [optional] [default to "scap-package"]
**Description** | Pointer to **string** | A description for the security package | [optional] 
**Url** | **string** | URL to download the security package zip file from | 
**Enabled** | Pointer to **bool** | Can be used to disable the security package | [optional] [default to true]

## Methods

### NewAddSecurityPackagesRequestSecurityPackage

`func NewAddSecurityPackagesRequestSecurityPackage(name string, url string, ) *AddSecurityPackagesRequestSecurityPackage`

NewAddSecurityPackagesRequestSecurityPackage instantiates a new AddSecurityPackagesRequestSecurityPackage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityPackagesRequestSecurityPackageWithDefaults

`func NewAddSecurityPackagesRequestSecurityPackageWithDefaults() *AddSecurityPackagesRequestSecurityPackage`

NewAddSecurityPackagesRequestSecurityPackageWithDefaults instantiates a new AddSecurityPackagesRequestSecurityPackage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddSecurityPackagesRequestSecurityPackage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSecurityPackagesRequestSecurityPackage) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddSecurityPackagesRequestSecurityPackage) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddSecurityPackagesRequestSecurityPackage) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddSecurityPackagesRequestSecurityPackage) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *AddSecurityPackagesRequestSecurityPackage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddSecurityPackagesRequestSecurityPackage) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddSecurityPackagesRequestSecurityPackage) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *AddSecurityPackagesRequestSecurityPackage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSecurityPackagesRequestSecurityPackage) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSecurityPackagesRequestSecurityPackage) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUrl

`func (o *AddSecurityPackagesRequestSecurityPackage) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddSecurityPackagesRequestSecurityPackage) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEnabled

`func (o *AddSecurityPackagesRequestSecurityPackage) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddSecurityPackagesRequestSecurityPackage) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddSecurityPackagesRequestSecurityPackage) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddSecurityPackagesRequestSecurityPackage) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


