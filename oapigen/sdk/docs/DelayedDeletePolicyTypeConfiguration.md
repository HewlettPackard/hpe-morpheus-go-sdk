# DelayedDeletePolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountIntegrationId** | Pointer to **string** |  | [optional] 
**RemovalAge** | **string** |  | 

## Methods

### NewDelayedDeletePolicyTypeConfiguration

`func NewDelayedDeletePolicyTypeConfiguration(removalAge string, ) *DelayedDeletePolicyTypeConfiguration`

NewDelayedDeletePolicyTypeConfiguration instantiates a new DelayedDeletePolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDelayedDeletePolicyTypeConfigurationWithDefaults

`func NewDelayedDeletePolicyTypeConfigurationWithDefaults() *DelayedDeletePolicyTypeConfiguration`

NewDelayedDeletePolicyTypeConfigurationWithDefaults instantiates a new DelayedDeletePolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountIntegrationId

`func (o *DelayedDeletePolicyTypeConfiguration) GetAccountIntegrationId() string`

GetAccountIntegrationId returns the AccountIntegrationId field if non-nil, zero value otherwise.

### GetAccountIntegrationIdOk

`func (o *DelayedDeletePolicyTypeConfiguration) GetAccountIntegrationIdOk() (*string, bool)`

GetAccountIntegrationIdOk returns a tuple with the AccountIntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountIntegrationId

`func (o *DelayedDeletePolicyTypeConfiguration) SetAccountIntegrationId(v string)`

SetAccountIntegrationId sets AccountIntegrationId field to given value.

### HasAccountIntegrationId

`func (o *DelayedDeletePolicyTypeConfiguration) HasAccountIntegrationId() bool`

HasAccountIntegrationId returns a boolean if a field has been set.

### GetRemovalAge

`func (o *DelayedDeletePolicyTypeConfiguration) GetRemovalAge() string`

GetRemovalAge returns the RemovalAge field if non-nil, zero value otherwise.

### GetRemovalAgeOk

`func (o *DelayedDeletePolicyTypeConfiguration) GetRemovalAgeOk() (*string, bool)`

GetRemovalAgeOk returns a tuple with the RemovalAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovalAge

`func (o *DelayedDeletePolicyTypeConfiguration) SetRemovalAge(v string)`

SetRemovalAge sets RemovalAge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


