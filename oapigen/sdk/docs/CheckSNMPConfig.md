# CheckSNMPConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | Hostname or IP address of the SNMP network entity | 
**Port** | **string** | Port to connect to the SNMP entity | 
**Oid** | Pointer to **string** | Object ID to get from the network entity | [optional] 
**CheckOperator** | Pointer to **string** | Operator to use when comparing returned value with the expected check response value. One of &#39;lt&#39;, &#39;equal&#39;, or &#39;gt&#39;. | [optional] 
**CheckResponse** | Pointer to **string** | Value to use with the check operator | [optional] 
**Version** | Pointer to **string** | SNMP Version 1/2c/3 of snmp get command | [optional] 
**Community** | Pointer to **string** | Community string acts as simple user or ID password (only valid for v1/2c) | [optional] 
**Username** | Pointer to **string** | Username used with SNMPv3 auth/privacy protocol passwords | [optional] 
**SecurityLevel** | Pointer to **string** | Level of security for authentication and privacty | [optional] 
**Auth** | Pointer to **string** | Authentication protocol | [optional] 
**Authpassword** | Pointer to **string** |  | [optional] 
**Priv** | Pointer to **string** | Privacy protocol | [optional] 
**Privpassword** | Pointer to **string** |  | [optional] 

## Methods

### NewCheckSNMPConfig

`func NewCheckSNMPConfig(host string, port string, ) *CheckSNMPConfig`

NewCheckSNMPConfig instantiates a new CheckSNMPConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckSNMPConfigWithDefaults

`func NewCheckSNMPConfigWithDefaults() *CheckSNMPConfig`

NewCheckSNMPConfigWithDefaults instantiates a new CheckSNMPConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *CheckSNMPConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CheckSNMPConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CheckSNMPConfig) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *CheckSNMPConfig) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CheckSNMPConfig) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CheckSNMPConfig) SetPort(v string)`

SetPort sets Port field to given value.


### GetOid

`func (o *CheckSNMPConfig) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *CheckSNMPConfig) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *CheckSNMPConfig) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *CheckSNMPConfig) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetCheckOperator

`func (o *CheckSNMPConfig) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *CheckSNMPConfig) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *CheckSNMPConfig) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *CheckSNMPConfig) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResponse

`func (o *CheckSNMPConfig) GetCheckResponse() string`

GetCheckResponse returns the CheckResponse field if non-nil, zero value otherwise.

### GetCheckResponseOk

`func (o *CheckSNMPConfig) GetCheckResponseOk() (*string, bool)`

GetCheckResponseOk returns a tuple with the CheckResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponse

`func (o *CheckSNMPConfig) SetCheckResponse(v string)`

SetCheckResponse sets CheckResponse field to given value.

### HasCheckResponse

`func (o *CheckSNMPConfig) HasCheckResponse() bool`

HasCheckResponse returns a boolean if a field has been set.

### GetVersion

`func (o *CheckSNMPConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CheckSNMPConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CheckSNMPConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CheckSNMPConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommunity

`func (o *CheckSNMPConfig) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *CheckSNMPConfig) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *CheckSNMPConfig) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *CheckSNMPConfig) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetUsername

`func (o *CheckSNMPConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CheckSNMPConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CheckSNMPConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CheckSNMPConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *CheckSNMPConfig) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *CheckSNMPConfig) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *CheckSNMPConfig) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *CheckSNMPConfig) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetAuth

`func (o *CheckSNMPConfig) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *CheckSNMPConfig) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *CheckSNMPConfig) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *CheckSNMPConfig) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthpassword

`func (o *CheckSNMPConfig) GetAuthpassword() string`

GetAuthpassword returns the Authpassword field if non-nil, zero value otherwise.

### GetAuthpasswordOk

`func (o *CheckSNMPConfig) GetAuthpasswordOk() (*string, bool)`

GetAuthpasswordOk returns a tuple with the Authpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthpassword

`func (o *CheckSNMPConfig) SetAuthpassword(v string)`

SetAuthpassword sets Authpassword field to given value.

### HasAuthpassword

`func (o *CheckSNMPConfig) HasAuthpassword() bool`

HasAuthpassword returns a boolean if a field has been set.

### GetPriv

`func (o *CheckSNMPConfig) GetPriv() string`

GetPriv returns the Priv field if non-nil, zero value otherwise.

### GetPrivOk

`func (o *CheckSNMPConfig) GetPrivOk() (*string, bool)`

GetPrivOk returns a tuple with the Priv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriv

`func (o *CheckSNMPConfig) SetPriv(v string)`

SetPriv sets Priv field to given value.

### HasPriv

`func (o *CheckSNMPConfig) HasPriv() bool`

HasPriv returns a boolean if a field has been set.

### GetPrivpassword

`func (o *CheckSNMPConfig) GetPrivpassword() string`

GetPrivpassword returns the Privpassword field if non-nil, zero value otherwise.

### GetPrivpasswordOk

`func (o *CheckSNMPConfig) GetPrivpasswordOk() (*string, bool)`

GetPrivpasswordOk returns a tuple with the Privpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivpassword

`func (o *CheckSNMPConfig) SetPrivpassword(v string)`

SetPrivpassword sets Privpassword field to given value.

### HasPrivpassword

`func (o *CheckSNMPConfig) HasPrivpassword() bool`

HasPrivpassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


