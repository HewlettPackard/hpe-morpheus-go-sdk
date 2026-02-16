# GoogleCloudInstanceConfiguration2

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

### NewGoogleCloudInstanceConfiguration2

`func NewGoogleCloudInstanceConfiguration2() *GoogleCloudInstanceConfiguration2`

NewGoogleCloudInstanceConfiguration2 instantiates a new GoogleCloudInstanceConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleCloudInstanceConfiguration2WithDefaults

`func NewGoogleCloudInstanceConfiguration2WithDefaults() *GoogleCloudInstanceConfiguration2`

NewGoogleCloudInstanceConfiguration2WithDefaults instantiates a new GoogleCloudInstanceConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *GoogleCloudInstanceConfiguration2) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *GoogleCloudInstanceConfiguration2) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *GoogleCloudInstanceConfiguration2) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *GoogleCloudInstanceConfiguration2) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### SetNoAgentNil

`func (o *GoogleCloudInstanceConfiguration2) SetNoAgentNil(b bool)`

 SetNoAgentNil sets the value for NoAgent to be an explicit nil

### UnsetNoAgent
`func (o *GoogleCloudInstanceConfiguration2) UnsetNoAgent()`

UnsetNoAgent ensures that no value is present for NoAgent, not even an explicit nil
### GetGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration2) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *GoogleCloudInstanceConfiguration2) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration2) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *GoogleCloudInstanceConfiguration2) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *GoogleCloudInstanceConfiguration2) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *GoogleCloudInstanceConfiguration2) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *GoogleCloudInstanceConfiguration2) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *GoogleCloudInstanceConfiguration2) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *GoogleCloudInstanceConfiguration2) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *GoogleCloudInstanceConfiguration2) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *GoogleCloudInstanceConfiguration2) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *GoogleCloudInstanceConfiguration2) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *GoogleCloudInstanceConfiguration2) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *GoogleCloudInstanceConfiguration2) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *GoogleCloudInstanceConfiguration2) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *GoogleCloudInstanceConfiguration2) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *GoogleCloudInstanceConfiguration2) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *GoogleCloudInstanceConfiguration2) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *GoogleCloudInstanceConfiguration2) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *GoogleCloudInstanceConfiguration2) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


