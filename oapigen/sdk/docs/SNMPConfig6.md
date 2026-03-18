# SNMPConfig6

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

### NewSNMPConfig6

`func NewSNMPConfig6(host string, port string, ) *SNMPConfig6`

NewSNMPConfig6 instantiates a new SNMPConfig6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSNMPConfig6WithDefaults

`func NewSNMPConfig6WithDefaults() *SNMPConfig6`

NewSNMPConfig6WithDefaults instantiates a new SNMPConfig6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SNMPConfig6) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SNMPConfig6) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SNMPConfig6) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SNMPConfig6) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SNMPConfig6) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SNMPConfig6) SetPort(v string)`

SetPort sets Port field to given value.


### GetOid

`func (o *SNMPConfig6) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *SNMPConfig6) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *SNMPConfig6) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *SNMPConfig6) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetCheckOperator

`func (o *SNMPConfig6) GetCheckOperator() string`

GetCheckOperator returns the CheckOperator field if non-nil, zero value otherwise.

### GetCheckOperatorOk

`func (o *SNMPConfig6) GetCheckOperatorOk() (*string, bool)`

GetCheckOperatorOk returns a tuple with the CheckOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckOperator

`func (o *SNMPConfig6) SetCheckOperator(v string)`

SetCheckOperator sets CheckOperator field to given value.

### HasCheckOperator

`func (o *SNMPConfig6) HasCheckOperator() bool`

HasCheckOperator returns a boolean if a field has been set.

### GetCheckResponse

`func (o *SNMPConfig6) GetCheckResponse() string`

GetCheckResponse returns the CheckResponse field if non-nil, zero value otherwise.

### GetCheckResponseOk

`func (o *SNMPConfig6) GetCheckResponseOk() (*string, bool)`

GetCheckResponseOk returns a tuple with the CheckResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckResponse

`func (o *SNMPConfig6) SetCheckResponse(v string)`

SetCheckResponse sets CheckResponse field to given value.

### HasCheckResponse

`func (o *SNMPConfig6) HasCheckResponse() bool`

HasCheckResponse returns a boolean if a field has been set.

### GetVersion

`func (o *SNMPConfig6) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SNMPConfig6) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SNMPConfig6) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SNMPConfig6) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetCommunity

`func (o *SNMPConfig6) GetCommunity() string`

GetCommunity returns the Community field if non-nil, zero value otherwise.

### GetCommunityOk

`func (o *SNMPConfig6) GetCommunityOk() (*string, bool)`

GetCommunityOk returns a tuple with the Community field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunity

`func (o *SNMPConfig6) SetCommunity(v string)`

SetCommunity sets Community field to given value.

### HasCommunity

`func (o *SNMPConfig6) HasCommunity() bool`

HasCommunity returns a boolean if a field has been set.

### GetUsername

`func (o *SNMPConfig6) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *SNMPConfig6) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *SNMPConfig6) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *SNMPConfig6) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetSecurityLevel

`func (o *SNMPConfig6) GetSecurityLevel() string`

GetSecurityLevel returns the SecurityLevel field if non-nil, zero value otherwise.

### GetSecurityLevelOk

`func (o *SNMPConfig6) GetSecurityLevelOk() (*string, bool)`

GetSecurityLevelOk returns a tuple with the SecurityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevel

`func (o *SNMPConfig6) SetSecurityLevel(v string)`

SetSecurityLevel sets SecurityLevel field to given value.

### HasSecurityLevel

`func (o *SNMPConfig6) HasSecurityLevel() bool`

HasSecurityLevel returns a boolean if a field has been set.

### GetAuth

`func (o *SNMPConfig6) GetAuth() string`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *SNMPConfig6) GetAuthOk() (*string, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *SNMPConfig6) SetAuth(v string)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *SNMPConfig6) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetAuthpassword

`func (o *SNMPConfig6) GetAuthpassword() string`

GetAuthpassword returns the Authpassword field if non-nil, zero value otherwise.

### GetAuthpasswordOk

`func (o *SNMPConfig6) GetAuthpasswordOk() (*string, bool)`

GetAuthpasswordOk returns a tuple with the Authpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthpassword

`func (o *SNMPConfig6) SetAuthpassword(v string)`

SetAuthpassword sets Authpassword field to given value.

### HasAuthpassword

`func (o *SNMPConfig6) HasAuthpassword() bool`

HasAuthpassword returns a boolean if a field has been set.

### GetPriv

`func (o *SNMPConfig6) GetPriv() string`

GetPriv returns the Priv field if non-nil, zero value otherwise.

### GetPrivOk

`func (o *SNMPConfig6) GetPrivOk() (*string, bool)`

GetPrivOk returns a tuple with the Priv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriv

`func (o *SNMPConfig6) SetPriv(v string)`

SetPriv sets Priv field to given value.

### HasPriv

`func (o *SNMPConfig6) HasPriv() bool`

HasPriv returns a boolean if a field has been set.

### GetPrivpassword

`func (o *SNMPConfig6) GetPrivpassword() string`

GetPrivpassword returns the Privpassword field if non-nil, zero value otherwise.

### GetPrivpasswordOk

`func (o *SNMPConfig6) GetPrivpasswordOk() (*string, bool)`

GetPrivpasswordOk returns a tuple with the Privpassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivpassword

`func (o *SNMPConfig6) SetPrivpassword(v string)`

SetPrivpassword sets Privpassword field to given value.

### HasPrivpassword

`func (o *SNMPConfig6) HasPrivpassword() bool`

HasPrivpassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


