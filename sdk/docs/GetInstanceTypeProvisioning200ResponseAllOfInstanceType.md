# GetInstanceTypeProvisioning200ResponseAllOfInstanceType

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
**InstanceTypeLayouts** | Pointer to [**[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner.md) |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ImagePath** | Pointer to **string** | Logo image URL | [optional] 
**DarkImagePath** | Pointer to **string** | Dark logo image URL | [optional] 

## Methods

### NewGetInstanceTypeProvisioning200ResponseAllOfInstanceType

`func NewGetInstanceTypeProvisioning200ResponseAllOfInstanceType() *GetInstanceTypeProvisioning200ResponseAllOfInstanceType`

NewGetInstanceTypeProvisioning200ResponseAllOfInstanceType instantiates a new GetInstanceTypeProvisioning200ResponseAllOfInstanceType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeWithDefaults

`func NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeWithDefaults() *GetInstanceTypeProvisioning200ResponseAllOfInstanceType`

NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeWithDefaults instantiates a new GetInstanceTypeProvisioning200ResponseAllOfInstanceType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetCategory

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetActive

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetHasProvisioningStep

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasProvisioningStep() bool`

GetHasProvisioningStep returns the HasProvisioningStep field if non-nil, zero value otherwise.

### GetHasProvisioningStepOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasProvisioningStepOk() (*bool, bool)`

GetHasProvisioningStepOk returns a tuple with the HasProvisioningStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProvisioningStep

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetHasProvisioningStep(v bool)`

SetHasProvisioningStep sets HasProvisioningStep field to given value.

### HasHasProvisioningStep

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasHasProvisioningStep() bool`

HasHasProvisioningStep returns a boolean if a field has been set.

### GetHasDeployment

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasDeployment() bool`

GetHasDeployment returns the HasDeployment field if non-nil, zero value otherwise.

### GetHasDeploymentOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasDeploymentOk() (*bool, bool)`

GetHasDeploymentOk returns a tuple with the HasDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDeployment

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetHasDeployment(v bool)`

SetHasDeployment sets HasDeployment field to given value.

### HasHasDeployment

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasHasDeployment() bool`

HasHasDeployment returns a boolean if a field has been set.

### GetHasConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasConfig() bool`

GetHasConfig returns the HasConfig field if non-nil, zero value otherwise.

### GetHasConfigOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasConfigOk() (*bool, bool)`

GetHasConfigOk returns a tuple with the HasConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetHasConfig(v bool)`

SetHasConfig sets HasConfig field to given value.

### HasHasConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasHasConfig() bool`

HasHasConfig returns a boolean if a field has been set.

### GetHasSettings

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasSettings() bool`

GetHasSettings returns the HasSettings field if non-nil, zero value otherwise.

### GetHasSettingsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasSettingsOk() (*bool, bool)`

GetHasSettingsOk returns a tuple with the HasSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSettings

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetHasSettings(v bool)`

SetHasSettings sets HasSettings field to given value.

### HasHasSettings

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasHasSettings() bool`

HasHasSettings returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetProxyType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyType() string`

GetProxyType returns the ProxyType field if non-nil, zero value otherwise.

### GetProxyTypeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyTypeOk() (*string, bool)`

GetProxyTypeOk returns a tuple with the ProxyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetProxyType(v string)`

SetProxyType sets ProxyType field to given value.

### HasProxyType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasProxyType() bool`

HasProxyType returns a boolean if a field has been set.

### GetProxyPort

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyPort() string`

GetProxyPort returns the ProxyPort field if non-nil, zero value otherwise.

### GetProxyPortOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyPortOk() (*string, bool)`

GetProxyPortOk returns a tuple with the ProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyPort

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetProxyPort(v string)`

SetProxyPort sets ProxyPort field to given value.

### HasProxyPort

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasProxyPort() bool`

HasProxyPort returns a boolean if a field has been set.

### GetProxyProtocol

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyProtocol() string`

GetProxyProtocol returns the ProxyProtocol field if non-nil, zero value otherwise.

### GetProxyProtocolOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetProxyProtocolOk() (*string, bool)`

GetProxyProtocolOk returns a tuple with the ProxyProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyProtocol

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetProxyProtocol(v string)`

SetProxyProtocol sets ProxyProtocol field to given value.

### HasProxyProtocol

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasProxyProtocol() bool`

HasProxyProtocol returns a boolean if a field has been set.

### GetEnvironmentPrefix

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetEnvironmentPrefix() string`

GetEnvironmentPrefix returns the EnvironmentPrefix field if non-nil, zero value otherwise.

### GetEnvironmentPrefixOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetEnvironmentPrefixOk() (*string, bool)`

GetEnvironmentPrefixOk returns a tuple with the EnvironmentPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentPrefix

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetEnvironmentPrefix(v string)`

SetEnvironmentPrefix sets EnvironmentPrefix field to given value.

### HasEnvironmentPrefix

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasEnvironmentPrefix() bool`

HasEnvironmentPrefix returns a boolean if a field has been set.

### GetBackupType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetBackupType() string`

GetBackupType returns the BackupType field if non-nil, zero value otherwise.

### GetBackupTypeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetBackupTypeOk() (*string, bool)`

GetBackupTypeOk returns a tuple with the BackupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetBackupType(v string)`

SetBackupType sets BackupType field to given value.

### HasBackupType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasBackupType() bool`

HasBackupType returns a boolean if a field has been set.

### GetConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetVisibility

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetFeatured

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetVersions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetVersions() []string`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetVersionsOk() (*[]string, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetVersions(v []string)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetInstanceTypeLayouts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetInstanceTypeLayouts() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner`

GetInstanceTypeLayouts returns the InstanceTypeLayouts field if non-nil, zero value otherwise.

### GetInstanceTypeLayoutsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetInstanceTypeLayoutsOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner, bool)`

GetInstanceTypeLayoutsOk returns a tuple with the InstanceTypeLayouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeLayouts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetInstanceTypeLayouts(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner)`

SetInstanceTypeLayouts sets InstanceTypeLayouts field to given value.

### HasInstanceTypeLayouts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasInstanceTypeLayouts() bool`

HasInstanceTypeLayouts returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetImagePath() string`

GetImagePath returns the ImagePath field if non-nil, zero value otherwise.

### GetImagePathOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetImagePathOk() (*string, bool)`

GetImagePathOk returns a tuple with the ImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetImagePath(v string)`

SetImagePath sets ImagePath field to given value.

### HasImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasImagePath() bool`

HasImagePath returns a boolean if a field has been set.

### GetDarkImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetDarkImagePath() string`

GetDarkImagePath returns the DarkImagePath field if non-nil, zero value otherwise.

### GetDarkImagePathOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) GetDarkImagePathOk() (*string, bool)`

GetDarkImagePathOk returns a tuple with the DarkImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDarkImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) SetDarkImagePath(v string)`

SetDarkImagePath sets DarkImagePath field to given value.

### HasDarkImagePath

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceType) HasDarkImagePath() bool`

HasDarkImagePath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


