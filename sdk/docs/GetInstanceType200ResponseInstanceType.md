# GetInstanceType200ResponseInstanceType

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

### NewGetInstanceType200ResponseInstanceType

`func NewGetInstanceType200ResponseInstanceType() *GetInstanceType200ResponseInstanceType`

NewGetInstanceType200ResponseInstanceType instantiates a new GetInstanceType200ResponseInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeWithDefaults() *GetInstanceType200ResponseInstanceType`

NewGetInstanceType200ResponseInstanceTypeWithDefaults instantiates a new GetInstanceType200ResponseInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetInstanceType200ResponseInstanceType) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInstanceType200ResponseInstanceType) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInstanceType200ResponseInstanceType) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetInstanceType200ResponseInstanceType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceType200ResponseInstanceType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceType200ResponseInstanceType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceType200ResponseInstanceType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceType200ResponseInstanceType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceType200ResponseInstanceType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceType200ResponseInstanceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceType200ResponseInstanceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceType200ResponseInstanceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceType200ResponseInstanceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *GetInstanceType200ResponseInstanceType) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *GetInstanceType200ResponseInstanceType) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *GetInstanceType200ResponseInstanceType) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *GetInstanceType200ResponseInstanceType) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetCategory

`func (o *GetInstanceType200ResponseInstanceType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetInstanceType200ResponseInstanceType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetInstanceType200ResponseInstanceType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetInstanceType200ResponseInstanceType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetActive

`func (o *GetInstanceType200ResponseInstanceType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetInstanceType200ResponseInstanceType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetInstanceType200ResponseInstanceType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetInstanceType200ResponseInstanceType) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetHasProvisioningStep

`func (o *GetInstanceType200ResponseInstanceType) GetHasProvisioningStep() bool`

GetHasProvisioningStep returns the HasProvisioningStep field if non-nil, zero value otherwise.

### GetHasProvisioningStepOk

`func (o *GetInstanceType200ResponseInstanceType) GetHasProvisioningStepOk() (*bool, bool)`

GetHasProvisioningStepOk returns a tuple with the HasProvisioningStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProvisioningStep

`func (o *GetInstanceType200ResponseInstanceType) SetHasProvisioningStep(v bool)`

SetHasProvisioningStep sets HasProvisioningStep field to given value.

### HasHasProvisioningStep

`func (o *GetInstanceType200ResponseInstanceType) HasHasProvisioningStep() bool`

HasHasProvisioningStep returns a boolean if a field has been set.

### GetHasDeployment

`func (o *GetInstanceType200ResponseInstanceType) GetHasDeployment() bool`

GetHasDeployment returns the HasDeployment field if non-nil, zero value otherwise.

### GetHasDeploymentOk

`func (o *GetInstanceType200ResponseInstanceType) GetHasDeploymentOk() (*bool, bool)`

GetHasDeploymentOk returns a tuple with the HasDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeployment

`func (o *GetInstanceType200ResponseInstanceType) SetHasDeployment(v bool)`

SetHasDeployment sets HasDeployment field to given value.

### HasHasDeployment

`func (o *GetInstanceType200ResponseInstanceType) HasHasDeployment() bool`

HasHasDeployment returns a boolean if a field has been set.

### GetHasConfig

`func (o *GetInstanceType200ResponseInstanceType) GetHasConfig() bool`

GetHasConfig returns the HasConfig field if non-nil, zero value otherwise.

### GetHasConfigOk

`func (o *GetInstanceType200ResponseInstanceType) GetHasConfigOk() (*bool, bool)`

GetHasConfigOk returns a tuple with the HasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfig

`func (o *GetInstanceType200ResponseInstanceType) SetHasConfig(v bool)`

SetHasConfig sets HasConfig field to given value.

### HasHasConfig

`func (o *GetInstanceType200ResponseInstanceType) HasHasConfig() bool`

HasHasConfig returns a boolean if a field has been set.

### GetHasSettings

`func (o *GetInstanceType200ResponseInstanceType) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *GetInstanceType200ResponseInstanceType) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *GetInstanceType200ResponseInstanceType) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *GetInstanceType200ResponseInstanceType) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *GetInstanceType200ResponseInstanceType) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *GetInstanceType200ResponseInstanceType) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *GetInstanceType200ResponseInstanceType) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *GetInstanceType200ResponseInstanceType) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetProxyType

`func (o *GetInstanceType200ResponseInstanceType) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *GetInstanceType200ResponseInstanceType) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *GetInstanceType200ResponseInstanceType) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *GetInstanceType200ResponseInstanceType) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetProxyPort

`func (o *GetInstanceType200ResponseInstanceType) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *GetInstanceType200ResponseInstanceType) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *GetInstanceType200ResponseInstanceType) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *GetInstanceType200ResponseInstanceType) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *GetInstanceType200ResponseInstanceType) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *GetInstanceType200ResponseInstanceType) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *GetInstanceType200ResponseInstanceType) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *GetInstanceType200ResponseInstanceType) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *GetInstanceType200ResponseInstanceType) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetInstanceType200ResponseInstanceType) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetInstanceType200ResponseInstanceType) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetInstanceType200ResponseInstanceType) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetBackupType

`func (o *GetInstanceType200ResponseInstanceType) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *GetInstanceType200ResponseInstanceType) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *GetInstanceType200ResponseInstanceType) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *GetInstanceType200ResponseInstanceType) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetConfig

`func (o *GetInstanceType200ResponseInstanceType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetInstanceType200ResponseInstanceType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetInstanceType200ResponseInstanceType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetInstanceType200ResponseInstanceType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *GetInstanceType200ResponseInstanceType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetInstanceType200ResponseInstanceType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetInstanceType200ResponseInstanceType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetInstanceType200ResponseInstanceType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *GetInstanceType200ResponseInstanceType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *GetInstanceType200ResponseInstanceType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *GetInstanceType200ResponseInstanceType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *GetInstanceType200ResponseInstanceType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVersions

`func (o *GetInstanceType200ResponseInstanceType) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetInstanceType200ResponseInstanceType) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetInstanceType200ResponseInstanceType) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetInstanceType200ResponseInstanceType) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetInstanceTypeLayouts

`func (o *GetInstanceType200ResponseInstanceType) GetInstanceTypeLayouts() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *GetInstanceType200ResponseInstanceType) GetInstanceTypeLayoutsOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *GetInstanceType200ResponseInstanceType) SetInstanceTypeLayouts(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *GetInstanceType200ResponseInstanceType) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetInstanceType200ResponseInstanceType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceType200ResponseInstanceType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceType200ResponseInstanceType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetInstanceType200ResponseInstanceType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *GetInstanceType200ResponseInstanceType) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetInstanceType200ResponseInstanceType) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetInstanceType200ResponseInstanceType) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetInstanceType200ResponseInstanceType) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetImagePath

`func (o *GetInstanceType200ResponseInstanceType) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetInstanceType200ResponseInstanceType) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetInstanceType200ResponseInstanceType) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetInstanceType200ResponseInstanceType) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *GetInstanceType200ResponseInstanceType) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetInstanceType200ResponseInstanceType) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetInstanceType200ResponseInstanceType) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetInstanceType200ResponseInstanceType) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


