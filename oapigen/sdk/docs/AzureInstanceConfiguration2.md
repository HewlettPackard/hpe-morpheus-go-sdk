# AzureInstanceConfiguration2

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
**CreateUser** | Pointer to **bool** | Create User | [optional] [default to true]

## Methods

### NewAzureInstanceConfiguration2

`func NewAzureInstanceConfiguration2() *AzureInstanceConfiguration2`

NewAzureInstanceConfiguration2 instantiates a new AzureInstanceConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureInstanceConfiguration2WithDefaults

`func NewAzureInstanceConfiguration2WithDefaults() *AzureInstanceConfiguration2`

NewAzureInstanceConfiguration2WithDefaults instantiates a new AzureInstanceConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *AzureInstanceConfiguration2) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AzureInstanceConfiguration2) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AzureInstanceConfiguration2) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AzureInstanceConfiguration2) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *AzureInstanceConfiguration2) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *AzureInstanceConfiguration2) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *AzureInstanceConfiguration2) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *AzureInstanceConfiguration2) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *AzureInstanceConfiguration2) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *AzureInstanceConfiguration2) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *AzureInstanceConfiguration2) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *AzureInstanceConfiguration2) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AzureInstanceConfiguration2) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AzureInstanceConfiguration2) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AzureInstanceConfiguration2) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AzureInstanceConfiguration2) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *AzureInstanceConfiguration2) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *AzureInstanceConfiguration2) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *AzureInstanceConfiguration2) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *AzureInstanceConfiguration2) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *AzureInstanceConfiguration2) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *AzureInstanceConfiguration2) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *AzureInstanceConfiguration2) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *AzureInstanceConfiguration2) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *AzureInstanceConfiguration2) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *AzureInstanceConfiguration2) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *AzureInstanceConfiguration2) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *AzureInstanceConfiguration2) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration2) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *AzureInstanceConfiguration2) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration2) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration2) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AzureInstanceConfiguration2) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AzureInstanceConfiguration2) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AzureInstanceConfiguration2) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AzureInstanceConfiguration2) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


