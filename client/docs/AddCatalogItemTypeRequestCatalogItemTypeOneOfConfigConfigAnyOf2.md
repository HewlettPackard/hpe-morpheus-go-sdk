# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **bool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**GoogleZoneId** | Pointer to **int64** | External ID of the Google zone to use for instance. | [optional] 
**ExternalIpType** | Pointer to **int64** | External IP Type.  &#x60;-1&#x60; for ephemeral IP. | [optional] 
**NetworkTags** | Pointer to **string** | Network Tags | [optional] 
**ServiceAccount** | Pointer to **string** | Service Account | [optional] 
**AccessScope** | Pointer to **string** | Access Scope | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2 instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2WithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2WithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2WithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### GetGoogleZoneId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetGoogleZoneId() int64`

GetGoogleZoneId returns the GoogleZoneId field if non-nil, zero value otherwise.

### GetGoogleZoneIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetGoogleZoneIdOk() (*int64, bool)`

GetGoogleZoneIdOk returns a tuple with the GoogleZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleZoneId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetGoogleZoneId(v int64)`

SetGoogleZoneId sets GoogleZoneId field to given value.

### HasGoogleZoneId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasGoogleZoneId() bool`

HasGoogleZoneId returns a boolean if a field has been set.

### GetExternalIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetExternalIpType() int64`

GetExternalIpType returns the ExternalIpType field if non-nil, zero value otherwise.

### GetExternalIpTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetExternalIpTypeOk() (*int64, bool)`

GetExternalIpTypeOk returns a tuple with the ExternalIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetExternalIpType(v int64)`

SetExternalIpType sets ExternalIpType field to given value.

### HasExternalIpType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasExternalIpType() bool`

HasExternalIpType returns a boolean if a field has been set.

### GetNetworkTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetNetworkTags() string`

GetNetworkTags returns the NetworkTags field if non-nil, zero value otherwise.

### GetNetworkTagsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetNetworkTagsOk() (*string, bool)`

GetNetworkTagsOk returns a tuple with the NetworkTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetNetworkTags(v string)`

SetNetworkTags sets NetworkTags field to given value.

### HasNetworkTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasNetworkTags() bool`

HasNetworkTags returns a boolean if a field has been set.

### GetServiceAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.

### GetAccessScope

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetAccessScope() string`

GetAccessScope returns the AccessScope field if non-nil, zero value otherwise.

### GetAccessScopeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) GetAccessScopeOk() (*string, bool)`

GetAccessScopeOk returns a tuple with the AccessScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessScope

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) SetAccessScope(v string)`

SetAccessScope sets AccessScope field to given value.

### HasAccessScope

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf2) HasAccessScope() bool`

HasAccessScope returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


