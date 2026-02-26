# AddInstanceRequestConfig

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
**PoolProviderType** | Pointer to **string** | The type of pool provider to use for this instance, must be \&quot;mvm\&quot; | [optional] [default to "mvm"]
**KvmHostId** | Pointer to **int64** | The ID of the KVM host to provision the instance on | [optional] 

## Methods

### NewAddInstanceRequestConfig

`func NewAddInstanceRequestConfig() *AddInstanceRequestConfig`

NewAddInstanceRequestConfig instantiates a new AddInstanceRequestConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstanceRequestConfigWithDefaults

`func NewAddInstanceRequestConfigWithDefaults() *AddInstanceRequestConfig`

NewAddInstanceRequestConfigWithDefaults instantiates a new AddInstanceRequestConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *AddInstanceRequestConfig) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddInstanceRequestConfig) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddInstanceRequestConfig) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddInstanceRequestConfig) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *AddInstanceRequestConfig) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *AddInstanceRequestConfig) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *AddInstanceRequestConfig) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *AddInstanceRequestConfig) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *AddInstanceRequestConfig) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *AddInstanceRequestConfig) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *AddInstanceRequestConfig) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *AddInstanceRequestConfig) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AddInstanceRequestConfig) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AddInstanceRequestConfig) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AddInstanceRequestConfig) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AddInstanceRequestConfig) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *AddInstanceRequestConfig) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *AddInstanceRequestConfig) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *AddInstanceRequestConfig) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *AddInstanceRequestConfig) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *AddInstanceRequestConfig) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *AddInstanceRequestConfig) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *AddInstanceRequestConfig) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *AddInstanceRequestConfig) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *AddInstanceRequestConfig) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *AddInstanceRequestConfig) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *AddInstanceRequestConfig) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *AddInstanceRequestConfig) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *AddInstanceRequestConfig) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *AddInstanceRequestConfig) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *AddInstanceRequestConfig) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *AddInstanceRequestConfig) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddInstanceRequestConfig) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddInstanceRequestConfig) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddInstanceRequestConfig) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddInstanceRequestConfig) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNoAgent

`func (o *AddInstanceRequestConfig) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddInstanceRequestConfig) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddInstanceRequestConfig) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddInstanceRequestConfig) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetHostId

`func (o *AddInstanceRequestConfig) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *AddInstanceRequestConfig) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *AddInstanceRequestConfig) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *AddInstanceRequestConfig) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *AddInstanceRequestConfig) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *AddInstanceRequestConfig) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *AddInstanceRequestConfig) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *AddInstanceRequestConfig) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *AddInstanceRequestConfig) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *AddInstanceRequestConfig) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *AddInstanceRequestConfig) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *AddInstanceRequestConfig) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *AddInstanceRequestConfig) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *AddInstanceRequestConfig) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *AddInstanceRequestConfig) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *AddInstanceRequestConfig) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetTemplate

`func (o *AddInstanceRequestConfig) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AddInstanceRequestConfig) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AddInstanceRequestConfig) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AddInstanceRequestConfig) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *AddInstanceRequestConfig) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *AddInstanceRequestConfig) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *AddInstanceRequestConfig) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *AddInstanceRequestConfig) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *AddInstanceRequestConfig) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *AddInstanceRequestConfig) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *AddInstanceRequestConfig) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *AddInstanceRequestConfig) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *AddInstanceRequestConfig) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *AddInstanceRequestConfig) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *AddInstanceRequestConfig) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *AddInstanceRequestConfig) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *AddInstanceRequestConfig) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *AddInstanceRequestConfig) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *AddInstanceRequestConfig) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *AddInstanceRequestConfig) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *AddInstanceRequestConfig) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *AddInstanceRequestConfig) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *AddInstanceRequestConfig) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *AddInstanceRequestConfig) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetIsEC2

`func (o *AddInstanceRequestConfig) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *AddInstanceRequestConfig) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *AddInstanceRequestConfig) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *AddInstanceRequestConfig) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *AddInstanceRequestConfig) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *AddInstanceRequestConfig) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *AddInstanceRequestConfig) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *AddInstanceRequestConfig) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *AddInstanceRequestConfig) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *AddInstanceRequestConfig) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *AddInstanceRequestConfig) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *AddInstanceRequestConfig) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AddInstanceRequestConfig) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AddInstanceRequestConfig) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AddInstanceRequestConfig) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AddInstanceRequestConfig) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *AddInstanceRequestConfig) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *AddInstanceRequestConfig) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *AddInstanceRequestConfig) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *AddInstanceRequestConfig) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *AddInstanceRequestConfig) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *AddInstanceRequestConfig) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *AddInstanceRequestConfig) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *AddInstanceRequestConfig) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetUserData

`func (o *AddInstanceRequestConfig) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *AddInstanceRequestConfig) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *AddInstanceRequestConfig) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *AddInstanceRequestConfig) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### GetPoolProviderType

`func (o *AddInstanceRequestConfig) GetPoolProviderType() string`

GetPoolProviderType returns the PoolProviderType field if non-nil, zero value otherwise.

### GetPoolProviderTypeOk

`func (o *AddInstanceRequestConfig) GetPoolProviderTypeOk() (*string, bool)`

GetPoolProviderTypeOk returns a tuple with the PoolProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolProviderType

`func (o *AddInstanceRequestConfig) SetPoolProviderType(v string)`

SetPoolProviderType sets PoolProviderType field to given value.

### HasPoolProviderType

`func (o *AddInstanceRequestConfig) HasPoolProviderType() bool`

HasPoolProviderType returns a boolean if a field has been set.

### GetKvmHostId

`func (o *AddInstanceRequestConfig) GetKvmHostId() int64`

GetKvmHostId returns the KvmHostId field if non-nil, zero value otherwise.

### GetKvmHostIdOk

`func (o *AddInstanceRequestConfig) GetKvmHostIdOk() (*int64, bool)`

GetKvmHostIdOk returns a tuple with the KvmHostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKvmHostId

`func (o *AddInstanceRequestConfig) SetKvmHostId(v int64)`

SetKvmHostId sets KvmHostId field to given value.

### HasKvmHostId

`func (o *AddInstanceRequestConfig) HasKvmHostId() bool`

HasKvmHostId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


