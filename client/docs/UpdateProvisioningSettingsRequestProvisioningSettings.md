# UpdateProvisioningSettingsRequestProvisioningSettings

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

### NewUpdateProvisioningSettingsRequestProvisioningSettings

`func NewUpdateProvisioningSettingsRequestProvisioningSettings() *UpdateProvisioningSettingsRequestProvisioningSettings`

NewUpdateProvisioningSettingsRequestProvisioningSettings instantiates a new UpdateProvisioningSettingsRequestProvisioningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProvisioningSettingsRequestProvisioningSettingsWithDefaults

`func NewUpdateProvisioningSettingsRequestProvisioningSettingsWithDefaults() *UpdateProvisioningSettingsRequestProvisioningSettings`

NewUpdateProvisioningSettingsRequestProvisioningSettingsWithDefaults instantiates a new UpdateProvisioningSettingsRequestProvisioningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowZoneSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetAllowZoneSelection() bool`

GetAllowZoneSelection returns the AllowZoneSelection field if non-nil, zero value otherwise.

### GetAllowZoneSelectionOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetAllowZoneSelectionOk() (*bool, bool)`

GetAllowZoneSelectionOk returns a tuple with the AllowZoneSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowZoneSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetAllowZoneSelection(v bool)`

SetAllowZoneSelection sets AllowZoneSelection field to given value.

### HasAllowZoneSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasAllowZoneSelection() bool`

HasAllowZoneSelection returns a boolean if a field has been set.

### GetAllowServerSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetAllowServerSelection() bool`

GetAllowServerSelection returns the AllowServerSelection field if non-nil, zero value otherwise.

### GetAllowServerSelectionOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetAllowServerSelectionOk() (*bool, bool)`

GetAllowServerSelectionOk returns a tuple with the AllowServerSelection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowServerSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetAllowServerSelection(v bool)`

SetAllowServerSelection sets AllowServerSelection field to given value.

### HasAllowServerSelection

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasAllowServerSelection() bool`

HasAllowServerSelection returns a boolean if a field has been set.

### GetRequireEnvironments

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetRequireEnvironments() bool`

GetRequireEnvironments returns the RequireEnvironments field if non-nil, zero value otherwise.

### GetRequireEnvironmentsOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetRequireEnvironmentsOk() (*bool, bool)`

GetRequireEnvironmentsOk returns a tuple with the RequireEnvironments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireEnvironments

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetRequireEnvironments(v bool)`

SetRequireEnvironments sets RequireEnvironments field to given value.

### HasRequireEnvironments

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasRequireEnvironments() bool`

HasRequireEnvironments returns a boolean if a field has been set.

### GetShowPricing

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetShowPricing() bool`

GetShowPricing returns the ShowPricing field if non-nil, zero value otherwise.

### GetShowPricingOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetShowPricingOk() (*bool, bool)`

GetShowPricingOk returns a tuple with the ShowPricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPricing

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetShowPricing(v bool)`

SetShowPricing sets ShowPricing field to given value.

### HasShowPricing

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasShowPricing() bool`

HasShowPricing returns a boolean if a field has been set.

### GetHideDatastoreStats

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetHideDatastoreStats() bool`

GetHideDatastoreStats returns the HideDatastoreStats field if non-nil, zero value otherwise.

### GetHideDatastoreStatsOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetHideDatastoreStatsOk() (*bool, bool)`

GetHideDatastoreStatsOk returns a tuple with the HideDatastoreStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideDatastoreStats

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetHideDatastoreStats(v bool)`

SetHideDatastoreStats sets HideDatastoreStats field to given value.

### HasHideDatastoreStats

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasHideDatastoreStats() bool`

HasHideDatastoreStats returns a boolean if a field has been set.

### GetCrossTenantNamingPolicies

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCrossTenantNamingPolicies() bool`

GetCrossTenantNamingPolicies returns the CrossTenantNamingPolicies field if non-nil, zero value otherwise.

### GetCrossTenantNamingPoliciesOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCrossTenantNamingPoliciesOk() (*bool, bool)`

GetCrossTenantNamingPoliciesOk returns a tuple with the CrossTenantNamingPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossTenantNamingPolicies

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetCrossTenantNamingPolicies(v bool)`

SetCrossTenantNamingPolicies sets CrossTenantNamingPolicies field to given value.

### HasCrossTenantNamingPolicies

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasCrossTenantNamingPolicies() bool`

HasCrossTenantNamingPolicies returns a boolean if a field has been set.

### GetReuseSequence

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetReuseSequence() bool`

GetReuseSequence returns the ReuseSequence field if non-nil, zero value otherwise.

### GetReuseSequenceOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetReuseSequenceOk() (*bool, bool)`

GetReuseSequenceOk returns a tuple with the ReuseSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReuseSequence

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetReuseSequence(v bool)`

SetReuseSequence sets ReuseSequence field to given value.

### HasReuseSequence

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasReuseSequence() bool`

HasReuseSequence returns a boolean if a field has been set.

### GetCloudInitUsername

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitUsername() string`

GetCloudInitUsername returns the CloudInitUsername field if non-nil, zero value otherwise.

### GetCloudInitUsernameOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitUsernameOk() (*string, bool)`

GetCloudInitUsernameOk returns a tuple with the CloudInitUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitUsername

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetCloudInitUsername(v string)`

SetCloudInitUsername sets CloudInitUsername field to given value.

### HasCloudInitUsername

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasCloudInitUsername() bool`

HasCloudInitUsername returns a boolean if a field has been set.

### GetCloudInitPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitPassword() string`

GetCloudInitPassword returns the CloudInitPassword field if non-nil, zero value otherwise.

### GetCloudInitPasswordOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitPasswordOk() (*string, bool)`

GetCloudInitPasswordOk returns a tuple with the CloudInitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetCloudInitPassword(v string)`

SetCloudInitPassword sets CloudInitPassword field to given value.

### HasCloudInitPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasCloudInitPassword() bool`

HasCloudInitPassword returns a boolean if a field has been set.

### GetCloudInitKeyPair

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitKeyPair() UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair`

GetCloudInitKeyPair returns the CloudInitKeyPair field if non-nil, zero value otherwise.

### GetCloudInitKeyPairOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetCloudInitKeyPairOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair, bool)`

GetCloudInitKeyPairOk returns a tuple with the CloudInitKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudInitKeyPair

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetCloudInitKeyPair(v UpdateProvisioningSettingsRequestProvisioningSettingsCloudInitKeyPair)`

SetCloudInitKeyPair sets CloudInitKeyPair field to given value.

### HasCloudInitKeyPair

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasCloudInitKeyPair() bool`

HasCloudInitKeyPair returns a boolean if a field has been set.

### GetDeployStorageProvider

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetDeployStorageProvider() UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider`

GetDeployStorageProvider returns the DeployStorageProvider field if non-nil, zero value otherwise.

### GetDeployStorageProviderOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetDeployStorageProviderOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider, bool)`

GetDeployStorageProviderOk returns a tuple with the DeployStorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployStorageProvider

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetDeployStorageProvider(v UpdateProvisioningSettingsRequestProvisioningSettingsDeployStorageProvider)`

SetDeployStorageProvider sets DeployStorageProvider field to given value.

### HasDeployStorageProvider

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasDeployStorageProvider() bool`

HasDeployStorageProvider returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetPxeRootPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetPxeRootPassword() string`

GetPxeRootPassword returns the PxeRootPassword field if non-nil, zero value otherwise.

### GetPxeRootPasswordOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetPxeRootPasswordOk() (*string, bool)`

GetPxeRootPasswordOk returns a tuple with the PxeRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPxeRootPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetPxeRootPassword(v string)`

SetPxeRootPassword sets PxeRootPassword field to given value.

### HasPxeRootPassword

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasPxeRootPassword() bool`

HasPxeRootPassword returns a boolean if a field has been set.

### GetDefaultTemplateType

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetDefaultTemplateType() UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType`

GetDefaultTemplateType returns the DefaultTemplateType field if non-nil, zero value otherwise.

### GetDefaultTemplateTypeOk

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) GetDefaultTemplateTypeOk() (*UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType, bool)`

GetDefaultTemplateTypeOk returns a tuple with the DefaultTemplateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTemplateType

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) SetDefaultTemplateType(v UpdateProvisioningSettingsRequestProvisioningSettingsDefaultTemplateType)`

SetDefaultTemplateType sets DefaultTemplateType field to given value.

### HasDefaultTemplateType

`func (o *UpdateProvisioningSettingsRequestProvisioningSettings) HasDefaultTemplateType() bool`

HasDefaultTemplateType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


