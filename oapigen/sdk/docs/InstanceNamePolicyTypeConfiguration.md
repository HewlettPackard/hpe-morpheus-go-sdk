# InstanceNamePolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NamingType** | **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | 
**NamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation.  Available variables are:&lt;br&gt;groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 
**NamingConflict** | Pointer to **bool** | Auto-resolve conflicts | [optional] 

## Methods

### NewInstanceNamePolicyTypeConfiguration

`func NewInstanceNamePolicyTypeConfiguration(namingType string, ) *InstanceNamePolicyTypeConfiguration`

NewInstanceNamePolicyTypeConfiguration instantiates a new InstanceNamePolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceNamePolicyTypeConfigurationWithDefaults

`func NewInstanceNamePolicyTypeConfigurationWithDefaults() *InstanceNamePolicyTypeConfiguration`

NewInstanceNamePolicyTypeConfigurationWithDefaults instantiates a new InstanceNamePolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamingType

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingType() string`

GetNamingType returns the NamingType field if non-nil, zero value otherwise.

### GetNamingTypeOk

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingTypeOk() (*string, bool)`

GetNamingTypeOk returns a tuple with the NamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingType

`func (o *InstanceNamePolicyTypeConfiguration) SetNamingType(v string)`

SetNamingType sets NamingType field to given value.


### GetNamingPattern

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *InstanceNamePolicyTypeConfiguration) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *InstanceNamePolicyTypeConfiguration) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.

### GetNamingConflict

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingConflict() bool`

GetNamingConflict returns the NamingConflict field if non-nil, zero value otherwise.

### GetNamingConflictOk

`func (o *InstanceNamePolicyTypeConfiguration) GetNamingConflictOk() (*bool, bool)`

GetNamingConflictOk returns a tuple with the NamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingConflict

`func (o *InstanceNamePolicyTypeConfiguration) SetNamingConflict(v bool)`

SetNamingConflict sets NamingConflict field to given value.

### HasNamingConflict

`func (o *InstanceNamePolicyTypeConfiguration) HasNamingConflict() bool`

HasNamingConflict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


