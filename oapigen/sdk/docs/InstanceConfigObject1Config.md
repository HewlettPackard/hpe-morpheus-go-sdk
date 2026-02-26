# InstanceConfigObject1Config

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

### NewInstanceConfigObject1Config

`func NewInstanceConfigObject1Config() *InstanceConfigObject1Config`

NewInstanceConfigObject1Config instantiates a new InstanceConfigObject1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObject1ConfigWithDefaults

`func NewInstanceConfigObject1ConfigWithDefaults() *InstanceConfigObject1Config`

NewInstanceConfigObject1ConfigWithDefaults instantiates a new InstanceConfigObject1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *InstanceConfigObject1Config) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *InstanceConfigObject1Config) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *InstanceConfigObject1Config) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *InstanceConfigObject1Config) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *InstanceConfigObject1Config) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *InstanceConfigObject1Config) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *InstanceConfigObject1Config) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *InstanceConfigObject1Config) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *InstanceConfigObject1Config) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *InstanceConfigObject1Config) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *InstanceConfigObject1Config) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *InstanceConfigObject1Config) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *InstanceConfigObject1Config) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *InstanceConfigObject1Config) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *InstanceConfigObject1Config) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *InstanceConfigObject1Config) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *InstanceConfigObject1Config) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *InstanceConfigObject1Config) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *InstanceConfigObject1Config) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *InstanceConfigObject1Config) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *InstanceConfigObject1Config) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *InstanceConfigObject1Config) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *InstanceConfigObject1Config) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *InstanceConfigObject1Config) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *InstanceConfigObject1Config) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *InstanceConfigObject1Config) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *InstanceConfigObject1Config) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *InstanceConfigObject1Config) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *InstanceConfigObject1Config) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *InstanceConfigObject1Config) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *InstanceConfigObject1Config) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *InstanceConfigObject1Config) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *InstanceConfigObject1Config) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *InstanceConfigObject1Config) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *InstanceConfigObject1Config) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *InstanceConfigObject1Config) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.

### GetNoAgent

`func (o *InstanceConfigObject1Config) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *InstanceConfigObject1Config) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *InstanceConfigObject1Config) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *InstanceConfigObject1Config) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetHostId

`func (o *InstanceConfigObject1Config) GetHostId() string`

GetHostId returns the HostId field if non-nil, zero value otherwise.

### GetHostIdOk

`func (o *InstanceConfigObject1Config) GetHostIdOk() (*string, bool)`

GetHostIdOk returns a tuple with the HostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostId

`func (o *InstanceConfigObject1Config) SetHostId(v string)`

SetHostId sets HostId field to given value.

### HasHostId

`func (o *InstanceConfigObject1Config) HasHostId() bool`

HasHostId returns a boolean if a field has been set.

### GetSmbiosAssetTag

`func (o *InstanceConfigObject1Config) GetSmbiosAssetTag() string`

GetSmbiosAssetTag returns the SmbiosAssetTag field if non-nil, zero value otherwise.

### GetSmbiosAssetTagOk

`func (o *InstanceConfigObject1Config) GetSmbiosAssetTagOk() (*string, bool)`

GetSmbiosAssetTagOk returns a tuple with the SmbiosAssetTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmbiosAssetTag

`func (o *InstanceConfigObject1Config) SetSmbiosAssetTag(v string)`

SetSmbiosAssetTag sets SmbiosAssetTag field to given value.

### HasSmbiosAssetTag

`func (o *InstanceConfigObject1Config) HasSmbiosAssetTag() bool`

HasSmbiosAssetTag returns a boolean if a field has been set.

### GetNestedVirtualization

`func (o *InstanceConfigObject1Config) GetNestedVirtualization() string`

GetNestedVirtualization returns the NestedVirtualization field if non-nil, zero value otherwise.

### GetNestedVirtualizationOk

`func (o *InstanceConfigObject1Config) GetNestedVirtualizationOk() (*string, bool)`

GetNestedVirtualizationOk returns a tuple with the NestedVirtualization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNestedVirtualization

`func (o *InstanceConfigObject1Config) SetNestedVirtualization(v string)`

SetNestedVirtualization sets NestedVirtualization field to given value.

### HasNestedVirtualization

`func (o *InstanceConfigObject1Config) HasNestedVirtualization() bool`

HasNestedVirtualization returns a boolean if a field has been set.

### GetVmwareFolderId

`func (o *InstanceConfigObject1Config) GetVmwareFolderId() string`

GetVmwareFolderId returns the VmwareFolderId field if non-nil, zero value otherwise.

### GetVmwareFolderIdOk

`func (o *InstanceConfigObject1Config) GetVmwareFolderIdOk() (*string, bool)`

GetVmwareFolderIdOk returns a tuple with the VmwareFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmwareFolderId

`func (o *InstanceConfigObject1Config) SetVmwareFolderId(v string)`

SetVmwareFolderId sets VmwareFolderId field to given value.

### HasVmwareFolderId

`func (o *InstanceConfigObject1Config) HasVmwareFolderId() bool`

HasVmwareFolderId returns a boolean if a field has been set.

### GetTemplate

`func (o *InstanceConfigObject1Config) GetTemplate() int64`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *InstanceConfigObject1Config) GetTemplateOk() (*int64, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *InstanceConfigObject1Config) SetTemplate(v int64)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *InstanceConfigObject1Config) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *InstanceConfigObject1Config) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *InstanceConfigObject1Config) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *InstanceConfigObject1Config) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *InstanceConfigObject1Config) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *InstanceConfigObject1Config) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *InstanceConfigObject1Config) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *InstanceConfigObject1Config) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *InstanceConfigObject1Config) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *InstanceConfigObject1Config) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *InstanceConfigObject1Config) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *InstanceConfigObject1Config) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *InstanceConfigObject1Config) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *InstanceConfigObject1Config) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *InstanceConfigObject1Config) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *InstanceConfigObject1Config) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *InstanceConfigObject1Config) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *InstanceConfigObject1Config) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *InstanceConfigObject1Config) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *InstanceConfigObject1Config) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *InstanceConfigObject1Config) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.

### GetIsEC2

`func (o *InstanceConfigObject1Config) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *InstanceConfigObject1Config) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *InstanceConfigObject1Config) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *InstanceConfigObject1Config) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *InstanceConfigObject1Config) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *InstanceConfigObject1Config) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *InstanceConfigObject1Config) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *InstanceConfigObject1Config) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *InstanceConfigObject1Config) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *InstanceConfigObject1Config) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *InstanceConfigObject1Config) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *InstanceConfigObject1Config) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *InstanceConfigObject1Config) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *InstanceConfigObject1Config) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *InstanceConfigObject1Config) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *InstanceConfigObject1Config) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *InstanceConfigObject1Config) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *InstanceConfigObject1Config) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *InstanceConfigObject1Config) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *InstanceConfigObject1Config) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *InstanceConfigObject1Config) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *InstanceConfigObject1Config) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *InstanceConfigObject1Config) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *InstanceConfigObject1Config) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.

### GetUserData

`func (o *InstanceConfigObject1Config) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstanceConfigObject1Config) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstanceConfigObject1Config) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstanceConfigObject1Config) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


