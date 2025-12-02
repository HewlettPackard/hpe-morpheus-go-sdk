# ClusterResourceNamePolicyTypeConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerNamingType** | Pointer to **string** | Options: \&quot;user\&quot; (user configurable), \&quot;fixed\&quot; (strict pattern) | [optional] 
**ServerNamingPattern** | Pointer to **string** | Name pattern uses ${variable} string interpolation. Available variables are: groupName, groupCode, cloudName, cloudCode, type, accountId, account, accountType, platform, username, userId, userInitials, provisionType | [optional] 
**ServerNamingConflict** | Pointer to **bool** | Auto-resolve conflicts | [optional] 

## Methods

### NewClusterResourceNamePolicyTypeConfiguration1

`func NewClusterResourceNamePolicyTypeConfiguration1() *ClusterResourceNamePolicyTypeConfiguration1`

NewClusterResourceNamePolicyTypeConfiguration1 instantiates a new ClusterResourceNamePolicyTypeConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterResourceNamePolicyTypeConfiguration1WithDefaults

`func NewClusterResourceNamePolicyTypeConfiguration1WithDefaults() *ClusterResourceNamePolicyTypeConfiguration1`

NewClusterResourceNamePolicyTypeConfiguration1WithDefaults instantiates a new ClusterResourceNamePolicyTypeConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerNamingType

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingType() string`

GetServerNamingType returns the ServerNamingType field if non-nil, zero value otherwise.

### GetServerNamingTypeOk

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingTypeOk() (*string, bool)`

GetServerNamingTypeOk returns a tuple with the ServerNamingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingType

`func (o *ClusterResourceNamePolicyTypeConfiguration1) SetServerNamingType(v string)`

SetServerNamingType sets ServerNamingType field to given value.

### HasServerNamingType

`func (o *ClusterResourceNamePolicyTypeConfiguration1) HasServerNamingType() bool`

HasServerNamingType returns a boolean if a field has been set.

### GetServerNamingPattern

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingPattern() string`

GetServerNamingPattern returns the ServerNamingPattern field if non-nil, zero value otherwise.

### GetServerNamingPatternOk

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingPatternOk() (*string, bool)`

GetServerNamingPatternOk returns a tuple with the ServerNamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingPattern

`func (o *ClusterResourceNamePolicyTypeConfiguration1) SetServerNamingPattern(v string)`

SetServerNamingPattern sets ServerNamingPattern field to given value.

### HasServerNamingPattern

`func (o *ClusterResourceNamePolicyTypeConfiguration1) HasServerNamingPattern() bool`

HasServerNamingPattern returns a boolean if a field has been set.

### GetServerNamingConflict

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingConflict() bool`

GetServerNamingConflict returns the ServerNamingConflict field if non-nil, zero value otherwise.

### GetServerNamingConflictOk

`func (o *ClusterResourceNamePolicyTypeConfiguration1) GetServerNamingConflictOk() (*bool, bool)`

GetServerNamingConflictOk returns a tuple with the ServerNamingConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNamingConflict

`func (o *ClusterResourceNamePolicyTypeConfiguration1) SetServerNamingConflict(v bool)`

SetServerNamingConflict sets ServerNamingConflict field to given value.

### HasServerNamingConflict

`func (o *ClusterResourceNamePolicyTypeConfiguration1) HasServerNamingConflict() bool`

HasServerNamingConflict returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


