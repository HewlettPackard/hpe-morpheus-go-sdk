# GoogleCloudInstanceConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **NullableBool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**GoogleZoneId** | Pointer to **int64** | External ID of the Google zone to use for instance. | [optional] 
**ExternalIpType** | Pointer to **int64** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**NetworkTags** | Pointer to **string** | Network Tags | [optional] 
**ServiceAccount** | Pointer to **string** | Service Account | [optional] 
**AccessScope** | Pointer to **string** | Access Scope | [optional] 

## Methods

### NewGoogleCloudInstanceConfiguration1

`func NewGoogleCloudInstanceConfiguration1() *GoogleCloudInstanceConfiguration1`

NewGoogleCloudInstanceConfiguration1 instantiates a new GoogleCloudInstanceConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudInstanceConfiguration1WithDefaults

`func NewGoogleCloudInstanceConfiguration1WithDefaults() *GoogleCloudInstanceConfiguration1`

NewGoogleCloudInstanceConfiguration1WithDefaults instantiates a new GoogleCloudInstanceConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *GoogleCloudInstanceConfiguration1) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *GoogleCloudInstanceConfiguration1) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *GoogleCloudInstanceConfiguration1) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *GoogleCloudInstanceConfiguration1) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### SetNoAgentNil

`func (o *GoogleCloudInstanceConfiguration1) SetNoAgentNil(b bool)`

 SetNoAgentNil sets the value for NoAgent to be an explicit nil

### UnsetNoAgent
`func (o *GoogleCloudInstanceConfiguration1) UnsetNoAgent()`

UnsetNoAgent ensures that no value is present for NoAgent, not even an explicit nil
### GetGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration1) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *GoogleCloudInstanceConfiguration1) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration1) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration1) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *GoogleCloudInstanceConfiguration1) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *GoogleCloudInstanceConfiguration1) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *GoogleCloudInstanceConfiguration1) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *GoogleCloudInstanceConfiguration1) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *GoogleCloudInstanceConfiguration1) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *GoogleCloudInstanceConfiguration1) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *GoogleCloudInstanceConfiguration1) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *GoogleCloudInstanceConfiguration1) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *GoogleCloudInstanceConfiguration1) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GoogleCloudInstanceConfiguration1) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GoogleCloudInstanceConfiguration1) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *GoogleCloudInstanceConfiguration1) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *GoogleCloudInstanceConfiguration1) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *GoogleCloudInstanceConfiguration1) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *GoogleCloudInstanceConfiguration1) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *GoogleCloudInstanceConfiguration1) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


