# InstanceConfigObjectConfig

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

### NewInstanceConfigObjectConfig

`func NewInstanceConfigObjectConfig() *InstanceConfigObjectConfig`

NewInstanceConfigObjectConfig instantiates a new InstanceConfigObjectConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObjectConfigWithDefaults

`func NewInstanceConfigObjectConfigWithDefaults() *InstanceConfigObjectConfig`

NewInstanceConfigObjectConfigWithDefaults instantiates a new InstanceConfigObjectConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *InstanceConfigObjectConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstanceConfigObjectConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstanceConfigObjectConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstanceConfigObjectConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *InstanceConfigObjectConfig) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *InstanceConfigObjectConfig) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *InstanceConfigObjectConfig) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *InstanceConfigObjectConfig) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *InstanceConfigObjectConfig) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *InstanceConfigObjectConfig) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *InstanceConfigObjectConfig) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *InstanceConfigObjectConfig) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *InstanceConfigObjectConfig) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *InstanceConfigObjectConfig) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *InstanceConfigObjectConfig) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *InstanceConfigObjectConfig) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *InstanceConfigObjectConfig) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *InstanceConfigObjectConfig) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *InstanceConfigObjectConfig) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *InstanceConfigObjectConfig) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *InstanceConfigObjectConfig) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *InstanceConfigObjectConfig) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *InstanceConfigObjectConfig) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *InstanceConfigObjectConfig) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *InstanceConfigObjectConfig) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *InstanceConfigObjectConfig) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *InstanceConfigObjectConfig) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *InstanceConfigObjectConfig) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *InstanceConfigObjectConfig) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *InstanceConfigObjectConfig) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *InstanceConfigObjectConfig) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *InstanceConfigObjectConfig) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *InstanceConfigObjectConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *InstanceConfigObjectConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *InstanceConfigObjectConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *InstanceConfigObjectConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNoAgent

`func (o *InstanceConfigObjectConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *InstanceConfigObjectConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *InstanceConfigObjectConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *InstanceConfigObjectConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetHostId

`func (o *InstanceConfigObjectConfig) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *InstanceConfigObjectConfig) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *InstanceConfigObjectConfig) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *InstanceConfigObjectConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *InstanceConfigObjectConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *InstanceConfigObjectConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *InstanceConfigObjectConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *InstanceConfigObjectConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *InstanceConfigObjectConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *InstanceConfigObjectConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *InstanceConfigObjectConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *InstanceConfigObjectConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *InstanceConfigObjectConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *InstanceConfigObjectConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *InstanceConfigObjectConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *InstanceConfigObjectConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetTemplate

`func (o *InstanceConfigObjectConfig) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *InstanceConfigObjectConfig) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *InstanceConfigObjectConfig) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *InstanceConfigObjectConfig) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *InstanceConfigObjectConfig) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *InstanceConfigObjectConfig) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *InstanceConfigObjectConfig) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *InstanceConfigObjectConfig) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *InstanceConfigObjectConfig) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *InstanceConfigObjectConfig) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *InstanceConfigObjectConfig) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *InstanceConfigObjectConfig) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InstanceConfigObjectConfig) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InstanceConfigObjectConfig) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InstanceConfigObjectConfig) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InstanceConfigObjectConfig) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *InstanceConfigObjectConfig) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *InstanceConfigObjectConfig) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *InstanceConfigObjectConfig) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *InstanceConfigObjectConfig) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *InstanceConfigObjectConfig) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *InstanceConfigObjectConfig) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *InstanceConfigObjectConfig) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *InstanceConfigObjectConfig) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetIsEC2

`func (o *InstanceConfigObjectConfig) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *InstanceConfigObjectConfig) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *InstanceConfigObjectConfig) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *InstanceConfigObjectConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *InstanceConfigObjectConfig) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *InstanceConfigObjectConfig) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *InstanceConfigObjectConfig) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *InstanceConfigObjectConfig) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *InstanceConfigObjectConfig) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *InstanceConfigObjectConfig) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *InstanceConfigObjectConfig) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *InstanceConfigObjectConfig) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *InstanceConfigObjectConfig) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *InstanceConfigObjectConfig) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *InstanceConfigObjectConfig) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *InstanceConfigObjectConfig) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *InstanceConfigObjectConfig) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *InstanceConfigObjectConfig) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *InstanceConfigObjectConfig) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *InstanceConfigObjectConfig) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *InstanceConfigObjectConfig) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *InstanceConfigObjectConfig) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *InstanceConfigObjectConfig) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *InstanceConfigObjectConfig) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetUserData

`func (o *InstanceConfigObjectConfig) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstanceConfigObjectConfig) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstanceConfigObjectConfig) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstanceConfigObjectConfig) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


