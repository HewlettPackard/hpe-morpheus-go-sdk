# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf

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

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOfWithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOfWithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOfWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetAvailabilityOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityOptions() string`

GetAvailabilityOptions returns the AvailabilityOptions field if non-nil, zero value otherwise.

### GetAvailabilityOptionsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityOptionsOk() (*string, bool)`

GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilityOptions(v string)`

SetAvailabilityOptions sets AvailabilityOptions field to given value.

### HasAvailabilityOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasAvailabilityOptions() bool`

HasAvailabilityOptions returns a boolean if a field has been set.

### GetAvailabilitySet

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilitySet() string`

GetAvailabilitySet returns the AvailabilitySet field if non-nil, zero value otherwise.

### GetAvailabilitySetOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilitySetOk() (*string, bool)`

GetAvailabilitySetOk returns a tuple with the AvailabilitySet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilitySet

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilitySet(v string)`

SetAvailabilitySet sets AvailabilitySet field to given value.

### HasAvailabilitySet

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasAvailabilitySet() bool`

HasAvailabilitySet returns a boolean if a field has been set.

### GetAvailabilityZone

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityZone() int64`

GetAvailabilityZone returns the AvailabilityZone field if non-nil, zero value otherwise.

### GetAvailabilityZoneOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityZoneOk() (*int64, bool)`

GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityZone

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilityZone(v int64)`

SetAvailabilityZone sets AvailabilityZone field to given value.

### HasAvailabilityZone

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasAvailabilityZone() bool`

HasAvailabilityZone returns a boolean if a field has been set.

### GetAzurefloatingIp

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAzurefloatingIp() string`

GetAzurefloatingIp returns the AzurefloatingIp field if non-nil, zero value otherwise.

### GetAzurefloatingIpOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAzurefloatingIpOk() (*string, bool)`

GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzurefloatingIp

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAzurefloatingIp(v string)`

SetAzurefloatingIp sets AzurefloatingIp field to given value.

### HasAzurefloatingIp

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasAzurefloatingIp() bool`

HasAzurefloatingIp returns a boolean if a field has been set.

### GetBootDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetBootDiagnostics() string`

GetBootDiagnostics returns the BootDiagnostics field if non-nil, zero value otherwise.

### GetBootDiagnosticsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetBootDiagnosticsOk() (*string, bool)`

GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetBootDiagnostics(v string)`

SetBootDiagnostics sets BootDiagnostics field to given value.

### HasBootDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasBootDiagnostics() bool`

HasBootDiagnostics returns a boolean if a field has been set.

### GetOsGuestDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetOsGuestDiagnostics() string`

GetOsGuestDiagnostics returns the OsGuestDiagnostics field if non-nil, zero value otherwise.

### GetOsGuestDiagnosticsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetOsGuestDiagnosticsOk() (*string, bool)`

GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsGuestDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetOsGuestDiagnostics(v string)`

SetOsGuestDiagnostics sets OsGuestDiagnostics field to given value.

### HasOsGuestDiagnostics

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasOsGuestDiagnostics() bool`

HasOsGuestDiagnostics returns a boolean if a field has been set.

### GetDiagnosticsStorageAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetDiagnosticsStorageAccount() string`

GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field if non-nil, zero value otherwise.

### GetDiagnosticsStorageAccountOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetDiagnosticsStorageAccountOk() (*string, bool)`

GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiagnosticsStorageAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetDiagnosticsStorageAccount(v string)`

SetDiagnosticsStorageAccount sets DiagnosticsStorageAccount field to given value.

### HasDiagnosticsStorageAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasDiagnosticsStorageAccount() bool`

HasDiagnosticsStorageAccount returns a boolean if a field has been set.

### GetCreateUser

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


