# ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProvisionTypeCode** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**InstanceTypeLayouts** | Pointer to [**[]ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerInstanceTypeLayoutsInner**](ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerInstanceTypeLayoutsInner.md) |  | [optional] 
**ImagePath** | Pointer to **string** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **string** | Dark logo image URL | [optional] 

## Methods

### NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner

`func NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner() *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner`

NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner instantiates a new ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerWithDefaults

`func NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerWithDefaults() *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner`

NewListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerWithDefaults instantiates a new ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetCategory

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetActive

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetVisibility

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVersions

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetInstanceTypeLayouts

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetInstanceTypeLayouts() []ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerInstanceTypeLayoutsInner`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetInstanceTypeLayoutsOk() (*[]ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerInstanceTypeLayoutsInner, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetInstanceTypeLayouts(v []ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInnerInstanceTypeLayoutsInner)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *ListInstanceTypesProvisioning200ResponseAllOfInstanceTypesInner) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


