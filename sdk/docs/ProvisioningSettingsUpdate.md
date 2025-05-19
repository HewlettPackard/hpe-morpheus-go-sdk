# ProvisioningSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowZoneSelection** | Pointer to **bool** | Use this to enable / disable allowing cloud selection | [optional] 
**AllowServerSelection** | Pointer to **bool** | Use this to enable / disable allowing host selection | [optional] 
**RequireEnvironments** | Pointer to **bool** | Use this to enable / disable requiring environment selection | [optional] 
**ShowPricing** | Pointer to **bool** | Use this to enable / disable showing pricing | [optional] 
**HideDatastoreStats** | Pointer to **bool** | Use this to enable / disable hiding datastore stats | [optional] 
**CrossTenantNamingPolicies** | Pointer to **bool** | Use this to enable / disable cross-tenant naming policies | [optional] 
**ReuseSequence** | Pointer to **bool** | Use this to enable / disable reusing naming sequence numbers | [optional] 
**CloudInitUsername** | Pointer to **string** | Cloud-init username | [optional] 
**CloudInitPassword** | Pointer to **string** | Cloud-init password | [optional] 
**CloudInitKeyPair** | Pointer to [**UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair**](UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair.md) |  | [optional] 
**DeployStorageProvider** | Pointer to [**UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider**](UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider.md) |  | [optional] 
**WindowsPassword** | Pointer to **string** | Windows administrator password | [optional] 
**PxeRootPassword** | Pointer to **string** | PXE Boot default root password | [optional] 
**DefaultTemplateType** | Pointer to [**UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType**](UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType.md) |  | [optional] 

## Methods

### NewProvisioningSettingsUpdate

`func NewProvisioningSettingsUpdate() *ProvisioningSettingsUpdate`

NewProvisioningSettingsUpdate instantiates a new ProvisioningSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisioningSettingsUpdateWithDefaults

`func NewProvisioningSettingsUpdateWithDefaults() *ProvisioningSettingsUpdate`

NewProvisioningSettingsUpdateWithDefaults instantiates a new ProvisioningSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowZoneSelection

`func (o *ProvisioningSettingsUpdate) GetAllowZoneSelection() bool`

GetAllowZoneSelection returns the AllowZoneSelection field if non-nil, zero value otherwise.

### GetAllowZoneSelectionOk

`func (o *ProvisioningSettingsUpdate) GetAllowZoneSelectionOk() (*bool, bool)`

GetAllowZoneSelectionOk returns a tuple with the AllowZoneSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZoneSelection

`func (o *ProvisioningSettingsUpdate) SetAllowZoneSelection(v bool)`

SetAllowZoneSelection sets AllowZoneSelection field to given value.

### HasAllowZoneSelection

`func (o *ProvisioningSettingsUpdate) HasAllowZoneSelection() bool`

HasAllowZoneSelection returns a boolean if a field has been set.

### GetAllowServerSelection

`func (o *ProvisioningSettingsUpdate) GetAllowServerSelection() bool`

GetAllowServerSelection returns the AllowServerSelection field if non-nil, zero value otherwise.

### GetAllowServerSelectionOk

`func (o *ProvisioningSettingsUpdate) GetAllowServerSelectionOk() (*bool, bool)`

GetAllowServerSelectionOk returns a tuple with the AllowServerSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServerSelection

`func (o *ProvisioningSettingsUpdate) SetAllowServerSelection(v bool)`

SetAllowServerSelection sets AllowServerSelection field to given value.

### HasAllowServerSelection

`func (o *ProvisioningSettingsUpdate) HasAllowServerSelection() bool`

HasAllowServerSelection returns a boolean if a field has been set.

### GetRequireEnvironments

`func (o *ProvisioningSettingsUpdate) GetRequireEnvironments() bool`

GetRequireEnvironments returns the RequireEnvironments field if non-nil, zero value otherwise.

### GetRequireEnvironmentsOk

`func (o *ProvisioningSettingsUpdate) GetRequireEnvironmentsOk() (*bool, bool)`

GetRequireEnvironmentsOk returns a tuple with the RequireEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEnvironments

`func (o *ProvisioningSettingsUpdate) SetRequireEnvironments(v bool)`

SetRequireEnvironments sets RequireEnvironments field to given value.

### HasRequireEnvironments

`func (o *ProvisioningSettingsUpdate) HasRequireEnvironments() bool`

HasRequireEnvironments returns a boolean if a field has been set.

### GetShowPricing

`func (o *ProvisioningSettingsUpdate) GetShowPricing() bool`

GetShowPricing returns the ShowPricing field if non-nil, zero value otherwise.

### GetShowPricingOk

`func (o *ProvisioningSettingsUpdate) GetShowPricingOk() (*bool, bool)`

GetShowPricingOk returns a tuple with the ShowPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPricing

`func (o *ProvisioningSettingsUpdate) SetShowPricing(v bool)`

SetShowPricing sets ShowPricing field to given value.

### HasShowPricing

`func (o *ProvisioningSettingsUpdate) HasShowPricing() bool`

HasShowPricing returns a boolean if a field has been set.

### GetHideDatastoreStats

`func (o *ProvisioningSettingsUpdate) GetHideDatastoreStats() bool`

GetHideDatastoreStats returns the HideDatastoreStats field if non-nil, zero value otherwise.

### GetHideDatastoreStatsOk

`func (o *ProvisioningSettingsUpdate) GetHideDatastoreStatsOk() (*bool, bool)`

GetHideDatastoreStatsOk returns a tuple with the HideDatastoreStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDatastoreStats

`func (o *ProvisioningSettingsUpdate) SetHideDatastoreStats(v bool)`

SetHideDatastoreStats sets HideDatastoreStats field to given value.

### HasHideDatastoreStats

`func (o *ProvisioningSettingsUpdate) HasHideDatastoreStats() bool`

HasHideDatastoreStats returns a boolean if a field has been set.

### GetCrossTenantNamingPolicies

`func (o *ProvisioningSettingsUpdate) GetCrossTenantNamingPolicies() bool`

GetCrossTenantNamingPolicies returns the CrossTenantNamingPolicies field if non-nil, zero value otherwise.

### GetCrossTenantNamingPoliciesOk

`func (o *ProvisioningSettingsUpdate) GetCrossTenantNamingPoliciesOk() (*bool, bool)`

GetCrossTenantNamingPoliciesOk returns a tuple with the CrossTenantNamingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossTenantNamingPolicies

`func (o *ProvisioningSettingsUpdate) SetCrossTenantNamingPolicies(v bool)`

SetCrossTenantNamingPolicies sets CrossTenantNamingPolicies field to given value.

### HasCrossTenantNamingPolicies

`func (o *ProvisioningSettingsUpdate) HasCrossTenantNamingPolicies() bool`

HasCrossTenantNamingPolicies returns a boolean if a field has been set.

### GetReuseSequence

`func (o *ProvisioningSettingsUpdate) GetReuseSequence() bool`

GetReuseSequence returns the ReuseSequence field if non-nil, zero value otherwise.

### GetReuseSequenceOk

`func (o *ProvisioningSettingsUpdate) GetReuseSequenceOk() (*bool, bool)`

GetReuseSequenceOk returns a tuple with the ReuseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseSequence

`func (o *ProvisioningSettingsUpdate) SetReuseSequence(v bool)`

SetReuseSequence sets ReuseSequence field to given value.

### HasReuseSequence

`func (o *ProvisioningSettingsUpdate) HasReuseSequence() bool`

HasReuseSequence returns a boolean if a field has been set.

### GetCloudInitUsername

`func (o *ProvisioningSettingsUpdate) GetCloudInitUsername() string`

GetCloudInitUsername returns the CloudInitUsername field if non-nil, zero value otherwise.

### GetCloudInitUsernameOk

`func (o *ProvisioningSettingsUpdate) GetCloudInitUsernameOk() (*string, bool)`

GetCloudInitUsernameOk returns a tuple with the CloudInitUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitUsername

`func (o *ProvisioningSettingsUpdate) SetCloudInitUsername(v string)`

SetCloudInitUsername sets CloudInitUsername field to given value.

### HasCloudInitUsername

`func (o *ProvisioningSettingsUpdate) HasCloudInitUsername() bool`

HasCloudInitUsername returns a boolean if a field has been set.

### GetCloudInitPassword

`func (o *ProvisioningSettingsUpdate) GetCloudInitPassword() string`

GetCloudInitPassword returns the CloudInitPassword field if non-nil, zero value otherwise.

### GetCloudInitPasswordOk

`func (o *ProvisioningSettingsUpdate) GetCloudInitPasswordOk() (*string, bool)`

GetCloudInitPasswordOk returns a tuple with the CloudInitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitPassword

`func (o *ProvisioningSettingsUpdate) SetCloudInitPassword(v string)`

SetCloudInitPassword sets CloudInitPassword field to given value.

### HasCloudInitPassword

`func (o *ProvisioningSettingsUpdate) HasCloudInitPassword() bool`

HasCloudInitPassword returns a boolean if a field has been set.

### GetCloudInitKeyPair

`func (o *ProvisioningSettingsUpdate) GetCloudInitKeyPair() UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair`

GetCloudInitKeyPair returns the CloudInitKeyPair field if non-nil, zero value otherwise.

### GetCloudInitKeyPairOk

`func (o *ProvisioningSettingsUpdate) GetCloudInitKeyPairOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair, bool)`

GetCloudInitKeyPairOk returns a tuple with the CloudInitKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitKeyPair

`func (o *ProvisioningSettingsUpdate) SetCloudInitKeyPair(v UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair)`

SetCloudInitKeyPair sets CloudInitKeyPair field to given value.

### HasCloudInitKeyPair

`func (o *ProvisioningSettingsUpdate) HasCloudInitKeyPair() bool`

HasCloudInitKeyPair returns a boolean if a field has been set.

### GetDeployStorageProvider

`func (o *ProvisioningSettingsUpdate) GetDeployStorageProvider() UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider`

GetDeployStorageProvider returns the DeployStorageProvider field if non-nil, zero value otherwise.

### GetDeployStorageProviderOk

`func (o *ProvisioningSettingsUpdate) GetDeployStorageProviderOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider, bool)`

GetDeployStorageProviderOk returns a tuple with the DeployStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStorageProvider

`func (o *ProvisioningSettingsUpdate) SetDeployStorageProvider(v UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider)`

SetDeployStorageProvider sets DeployStorageProvider field to given value.

### HasDeployStorageProvider

`func (o *ProvisioningSettingsUpdate) HasDeployStorageProvider() bool`

HasDeployStorageProvider returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *ProvisioningSettingsUpdate) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *ProvisioningSettingsUpdate) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *ProvisioningSettingsUpdate) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *ProvisioningSettingsUpdate) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetPxeRootPassword

`func (o *ProvisioningSettingsUpdate) GetPxeRootPassword() string`

GetPxeRootPassword returns the PxeRootPassword field if non-nil, zero value otherwise.

### GetPxeRootPasswordOk

`func (o *ProvisioningSettingsUpdate) GetPxeRootPasswordOk() (*string, bool)`

GetPxeRootPasswordOk returns a tuple with the PxeRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeRootPassword

`func (o *ProvisioningSettingsUpdate) SetPxeRootPassword(v string)`

SetPxeRootPassword sets PxeRootPassword field to given value.

### HasPxeRootPassword

`func (o *ProvisioningSettingsUpdate) HasPxeRootPassword() bool`

HasPxeRootPassword returns a boolean if a field has been set.

### GetDefaultTemplateType

`func (o *ProvisioningSettingsUpdate) GetDefaultTemplateType() UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType`

GetDefaultTemplateType returns the DefaultTemplateType field if non-nil, zero value otherwise.

### GetDefaultTemplateTypeOk

`func (o *ProvisioningSettingsUpdate) GetDefaultTemplateTypeOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType, bool)`

GetDefaultTemplateTypeOk returns a tuple with the DefaultTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplateType

`func (o *ProvisioningSettingsUpdate) SetDefaultTemplateType(v UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType)`

SetDefaultTemplateType sets DefaultTemplateType field to given value.

### HasDefaultTemplateType

`func (o *ProvisioningSettingsUpdate) HasDefaultTemplateType() bool`

HasDefaultTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


