# AzureInstanceConfiguration1

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

### NewAzureInstanceConfiguration1

`func NewAzureInstanceConfiguration1() *AzureInstanceConfiguration1`

NewAzureInstanceConfiguration1 instantiates a new AzureInstanceConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureInstanceConfiguration1WithDefaults

`func NewAzureInstanceConfiguration1WithDefaults() *AzureInstanceConfiguration1`

NewAzureInstanceConfiguration1WithDefaults instantiates a new AzureInstanceConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *AzureInstanceConfiguration1) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AzureInstanceConfiguration1) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AzureInstanceConfiguration1) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AzureInstanceConfiguration1) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *AzureInstanceConfiguration1) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *AzureInstanceConfiguration1) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *AzureInstanceConfiguration1) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *AzureInstanceConfiguration1) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *AzureInstanceConfiguration1) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *AzureInstanceConfiguration1) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *AzureInstanceConfiguration1) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *AzureInstanceConfiguration1) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AzureInstanceConfiguration1) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AzureInstanceConfiguration1) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AzureInstanceConfiguration1) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AzureInstanceConfiguration1) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *AzureInstanceConfiguration1) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *AzureInstanceConfiguration1) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *AzureInstanceConfiguration1) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *AzureInstanceConfiguration1) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *AzureInstanceConfiguration1) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *AzureInstanceConfiguration1) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *AzureInstanceConfiguration1) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *AzureInstanceConfiguration1) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *AzureInstanceConfiguration1) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *AzureInstanceConfiguration1) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *AzureInstanceConfiguration1) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *AzureInstanceConfiguration1) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration1) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *AzureInstanceConfiguration1) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration1) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *AzureInstanceConfiguration1) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AzureInstanceConfiguration1) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AzureInstanceConfiguration1) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AzureInstanceConfiguration1) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AzureInstanceConfiguration1) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


