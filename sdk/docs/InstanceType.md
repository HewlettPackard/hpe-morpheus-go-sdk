# InstanceType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**ProvisionTypeCode** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**HasProvisioningStep** | Pointer to **bool** |  | [optional] 
**HasDeployment** | Pointer to **bool** |  | [optional] 
**HasConfig** | Pointer to **bool** |  | [optional] 
**HasSettings** | Pointer to **bool** |  | [optional] 
**HasAutoScale** | Pointer to **bool** |  | [optional] 
**ProxyType** | Pointer to **string** |  | [optional] 
**ProxyPort** | Pointer to **string** |  | [optional] 
**ProxyProtocol** | Pointer to **string** |  | [optional] 
**EnvironmentPrefix** | Pointer to **string** |  | [optional] 
**BackupType** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**Versions** | Pointer to **[]string** |  | [optional] 
**InstanceTypeLayouts** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ImagePath** | Pointer to **string** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **string** | Dark logo image URL | [optional] 

## Methods

### NewInstanceType

`func NewInstanceType() *InstanceType`

NewInstanceType instantiates a new InstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeWithDefaults

`func NewInstanceTypeWithDefaults() *InstanceType`

NewInstanceTypeWithDefaults instantiates a new InstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceType) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceType) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceType) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *InstanceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *InstanceType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *InstanceType) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *InstanceType) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *InstanceType) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *InstanceType) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetCategory

`func (o *InstanceType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetActive

`func (o *InstanceType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *InstanceType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *InstanceType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *InstanceType) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetHasProvisioningStep

`func (o *InstanceType) GetHasProvisioningStep() bool`

GetHasProvisioningStep returns the HasProvisioningStep field if non-nil, zero value otherwise.

### GetHasProvisioningStepOk

`func (o *InstanceType) GetHasProvisioningStepOk() (*bool, bool)`

GetHasProvisioningStepOk returns a tuple with the HasProvisioningStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProvisioningStep

`func (o *InstanceType) SetHasProvisioningStep(v bool)`

SetHasProvisioningStep sets HasProvisioningStep field to given value.

### HasHasProvisioningStep

`func (o *InstanceType) HasHasProvisioningStep() bool`

HasHasProvisioningStep returns a boolean if a field has been set.

### GetHasDeployment

`func (o *InstanceType) GetHasDeployment() bool`

GetHasDeployment returns the HasDeployment field if non-nil, zero value otherwise.

### GetHasDeploymentOk

`func (o *InstanceType) GetHasDeploymentOk() (*bool, bool)`

GetHasDeploymentOk returns a tuple with the HasDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeployment

`func (o *InstanceType) SetHasDeployment(v bool)`

SetHasDeployment sets HasDeployment field to given value.

### HasHasDeployment

`func (o *InstanceType) HasHasDeployment() bool`

HasHasDeployment returns a boolean if a field has been set.

### GetHasConfig

`func (o *InstanceType) GetHasConfig() bool`

GetHasConfig returns the HasConfig field if non-nil, zero value otherwise.

### GetHasConfigOk

`func (o *InstanceType) GetHasConfigOk() (*bool, bool)`

GetHasConfigOk returns a tuple with the HasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfig

`func (o *InstanceType) SetHasConfig(v bool)`

SetHasConfig sets HasConfig field to given value.

### HasHasConfig

`func (o *InstanceType) HasHasConfig() bool`

HasHasConfig returns a boolean if a field has been set.

### GetHasSettings

`func (o *InstanceType) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *InstanceType) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *InstanceType) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *InstanceType) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *InstanceType) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *InstanceType) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *InstanceType) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *InstanceType) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetProxyType

`func (o *InstanceType) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *InstanceType) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *InstanceType) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *InstanceType) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetProxyPort

`func (o *InstanceType) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *InstanceType) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *InstanceType) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *InstanceType) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *InstanceType) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *InstanceType) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *InstanceType) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *InstanceType) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *InstanceType) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *InstanceType) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *InstanceType) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *InstanceType) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetBackupType

`func (o *InstanceType) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *InstanceType) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *InstanceType) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *InstanceType) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetConfig

`func (o *InstanceType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InstanceType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *InstanceType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *InstanceType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *InstanceType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *InstanceType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *InstanceType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *InstanceType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *InstanceType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *InstanceType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVersions

`func (o *InstanceType) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *InstanceType) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *InstanceType) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *InstanceType) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetInstanceTypeLayouts

`func (o *InstanceType) GetInstanceTypeLayouts() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *InstanceType) GetInstanceTypeLayoutsOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *InstanceType) SetInstanceTypeLayouts(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *InstanceType) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *InstanceType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *InstanceType) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *InstanceType) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *InstanceType) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *InstanceType) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetImagePath

`func (o *InstanceType) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *InstanceType) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *InstanceType) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *InstanceType) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *InstanceType) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *InstanceType) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *InstanceType) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *InstanceType) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


