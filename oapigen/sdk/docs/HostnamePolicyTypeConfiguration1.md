# HostnamePolicyTypeConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostNamingType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**HostNamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation. Available variables are: groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 

## Methods

### NewHostnamePolicyTypeConfiguration1

`func NewHostnamePolicyTypeConfiguration1() *HostnamePolicyTypeConfiguration1`

NewHostnamePolicyTypeConfiguration1 instantiates a new HostnamePolicyTypeConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostnamePolicyTypeConfiguration1WithDefaults

`func NewHostnamePolicyTypeConfiguration1WithDefaults() *HostnamePolicyTypeConfiguration1`

NewHostnamePolicyTypeConfiguration1WithDefaults instantiates a new HostnamePolicyTypeConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostNamingType

`func (o *HostnamePolicyTypeConfiguration1) GetHostNamingType() string`

GetHostNamingType returns the HostNamingType field if non-nil, zero value otherwise.

### GetHostNamingTypeOk

`func (o *HostnamePolicyTypeConfiguration1) GetHostNamingTypeOk() (*string, bool)`

GetHostNamingTypeOk returns a tuple with the HostNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingType

`func (o *HostnamePolicyTypeConfiguration1) SetHostNamingType(v string)`

SetHostNamingType sets HostNamingType field to given value.

### HasHostNamingType

`func (o *HostnamePolicyTypeConfiguration1) HasHostNamingType() bool`

HasHostNamingType returns a boolean if a field has been set.

### GetHostNamingPattern

`func (o *HostnamePolicyTypeConfiguration1) GetHostNamingPattern() string`

GetHostNamingPattern returns the HostNamingPattern field if non-nil, zero value otherwise.

### GetHostNamingPatternOk

`func (o *HostnamePolicyTypeConfiguration1) GetHostNamingPatternOk() (*string, bool)`

GetHostNamingPatternOk returns a tuple with the HostNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostNamingPattern

`func (o *HostnamePolicyTypeConfiguration1) SetHostNamingPattern(v string)`

SetHostNamingPattern sets HostNamingPattern field to given value.

### HasHostNamingPattern

`func (o *HostnamePolicyTypeConfiguration1) HasHostNamingPattern() bool`

HasHostNamingPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


