# CatalogItemTypeInstanceScribeConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePoolId** | Pointer to **string** | id of the resource group to be used, can be prefixed with &#x60;pool-&#x60;. A resource pool group can be specified instead by prefixing its ID with &#x60;poolGroup-&#x60;. | [optional] 
**AvailabilityOptions** | Pointer to **string** | Availability Options | [optional] 
**AvailabilitySet** | Pointer to **string** | Availability Set | [optional] 
**AvailabilityZone** | Pointer to **int64** | Availability Zone | [optional] 
**AzurefloatingIp** | Pointer to **string** | Assign Public IP | [optional] 
**BootDiagnostics** | Pointer to **string** | Boot Diagnostics | [optional] 
**OsGuestDiagnostics** | Pointer to **string** | OS Guest Diagnostics | [optional] 
**DiagnosticsStorageAccount** | Pointer to **string** | Diagnostics Storage Account | [optional] 
**CreateUser** | Pointer to **bool** | Create user | [optional] [default to false]
**NoAgent** | Pointer to **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**HostId** | Pointer to **string** | Specific host to deploy to if so desired. | [optional] 
**SmbiosAssetTag** | Pointer to **string** | Sets the asset tag on the SMBIOS for use by the guest operating system. If left blank, the virtual machine name will be used. | [optional] 
**NestedVirtualization** | Pointer to **string** | Enable Nested Virtualization | [optional] [default to "off"]
**VmwareFolderId** | Pointer to **string** | VMWare Folder External ID (as a String) or ID (as an Integer or String) | [optional] 
**Template** | Pointer to **int64** | Image ID. This is the ID of a Virtual Image. | [optional] 
**GoogleZoneId** | Pointer to **int64** | External ID of the Google zone to use for instance. | [optional] 
**ExternalIpType** | Pointer to **int64** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**NetworkTags** | Pointer to **string** | Network Tags | [optional] 
**ServiceAccount** | Pointer to **string** | Service Account | [optional] 
**AccessScope** | Pointer to **string** | Access Scope | [optional] 
**IsEC2** | Pointer to **string** | Amazon Cloud Type | [optional] [default to "false"]
**AvailabilityId** | Pointer to **string** | Amazon Zone | [optional] 
**SecurityId** | Pointer to **string** | Security Group | [optional] 
**PublicIpType** | Pointer to **string** | Public IP | [optional] 
**InstanceProfile** | Pointer to **string** | IAM Profile | [optional] 
**KmsKeyId** | Pointer to **string** | KMS Key ID | [optional] 
**UserData** | Pointer to **string** | User Data. Allows for override of cloud-init based user-data yaml or custom scripts | [optional] 

## Methods

### NewCatalogItemTypeInstanceScribeConfig

`func NewCatalogItemTypeInstanceScribeConfig() *CatalogItemTypeInstanceScribeConfig`

NewCatalogItemTypeInstanceScribeConfig instantiates a new CatalogItemTypeInstanceScribeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeInstanceScribeConfigWithDefaults

`func NewCatalogItemTypeInstanceScribeConfigWithDefaults() *CatalogItemTypeInstanceScribeConfig`

NewCatalogItemTypeInstanceScribeConfigWithDefaults instantiates a new CatalogItemTypeInstanceScribeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *CatalogItemTypeInstanceScribeConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *CatalogItemTypeInstanceScribeConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *CatalogItemTypeInstanceScribeConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *CatalogItemTypeInstanceScribeConfig) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *CatalogItemTypeInstanceScribeConfig) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *CatalogItemTypeInstanceScribeConfig) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *CatalogItemTypeInstanceScribeConfig) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *CatalogItemTypeInstanceScribeConfig) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *CatalogItemTypeInstanceScribeConfig) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *CatalogItemTypeInstanceScribeConfig) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *CatalogItemTypeInstanceScribeConfig) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *CatalogItemTypeInstanceScribeConfig) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *CatalogItemTypeInstanceScribeConfig) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *CatalogItemTypeInstanceScribeConfig) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *CatalogItemTypeInstanceScribeConfig) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *CatalogItemTypeInstanceScribeConfig) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *CatalogItemTypeInstanceScribeConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *CatalogItemTypeInstanceScribeConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *CatalogItemTypeInstanceScribeConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNoAgent

`func (o *CatalogItemTypeInstanceScribeConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *CatalogItemTypeInstanceScribeConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *CatalogItemTypeInstanceScribeConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetHostId

`func (o *CatalogItemTypeInstanceScribeConfig) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *CatalogItemTypeInstanceScribeConfig) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *CatalogItemTypeInstanceScribeConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *CatalogItemTypeInstanceScribeConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *CatalogItemTypeInstanceScribeConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *CatalogItemTypeInstanceScribeConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *CatalogItemTypeInstanceScribeConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *CatalogItemTypeInstanceScribeConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *CatalogItemTypeInstanceScribeConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *CatalogItemTypeInstanceScribeConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *CatalogItemTypeInstanceScribeConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *CatalogItemTypeInstanceScribeConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetTemplate

`func (o *CatalogItemTypeInstanceScribeConfig) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CatalogItemTypeInstanceScribeConfig) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CatalogItemTypeInstanceScribeConfig) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *CatalogItemTypeInstanceScribeConfig) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *CatalogItemTypeInstanceScribeConfig) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *CatalogItemTypeInstanceScribeConfig) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *CatalogItemTypeInstanceScribeConfig) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *CatalogItemTypeInstanceScribeConfig) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *CatalogItemTypeInstanceScribeConfig) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *CatalogItemTypeInstanceScribeConfig) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *CatalogItemTypeInstanceScribeConfig) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *CatalogItemTypeInstanceScribeConfig) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *CatalogItemTypeInstanceScribeConfig) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *CatalogItemTypeInstanceScribeConfig) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *CatalogItemTypeInstanceScribeConfig) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *CatalogItemTypeInstanceScribeConfig) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *CatalogItemTypeInstanceScribeConfig) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *CatalogItemTypeInstanceScribeConfig) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetIsEC2

`func (o *CatalogItemTypeInstanceScribeConfig) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *CatalogItemTypeInstanceScribeConfig) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *CatalogItemTypeInstanceScribeConfig) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *CatalogItemTypeInstanceScribeConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *CatalogItemTypeInstanceScribeConfig) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *CatalogItemTypeInstanceScribeConfig) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *CatalogItemTypeInstanceScribeConfig) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *CatalogItemTypeInstanceScribeConfig) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *CatalogItemTypeInstanceScribeConfig) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *CatalogItemTypeInstanceScribeConfig) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *CatalogItemTypeInstanceScribeConfig) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *CatalogItemTypeInstanceScribeConfig) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *CatalogItemTypeInstanceScribeConfig) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *CatalogItemTypeInstanceScribeConfig) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *CatalogItemTypeInstanceScribeConfig) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *CatalogItemTypeInstanceScribeConfig) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *CatalogItemTypeInstanceScribeConfig) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *CatalogItemTypeInstanceScribeConfig) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetUserData

`func (o *CatalogItemTypeInstanceScribeConfig) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *CatalogItemTypeInstanceScribeConfig) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *CatalogItemTypeInstanceScribeConfig) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *CatalogItemTypeInstanceScribeConfig) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


