# WebGetConfig1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebMethod** | **string** | HTTP method to use for testing | 
**WebUrl** | **string** | Web URL you wish to use to run a check on | 
**IgnoreSSL** | Pointer to **bool** | Ignore SSL Errors | [optional] [default to false]
**CheckUser** | Pointer to **string** | If you want to use HTTP Basic Authentication, populate this field with the username | [optional] 
**CheckPassword** | Pointer to **string** | If you want to use HTTP basic Authentication, populate this field with the password | [optional] 
**TextCheckOn** | Pointer to **string** | Set value to &#x60;on&#x60; if you want to turn on text matching | [optional] 
**WebTextMatch** | Pointer to **string** | Set the string you want to look for in the page source | [optional] 

## Methods

### NewWebGetConfig1

`func NewWebGetConfig1(webMethod string, webUrl string, ) *WebGetConfig1`

NewWebGetConfig1 instantiates a new WebGetConfig1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebGetConfig1WithDefaults

`func NewWebGetConfig1WithDefaults() *WebGetConfig1`

NewWebGetConfig1WithDefaults instantiates a new WebGetConfig1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebMethod

`func (o *WebGetConfig1) GetWebMethod() string`

GetWebMethod returns the WebMethod field if non-nil, zero value otherwise.

### GetWebMethodOk

`func (o *WebGetConfig1) GetWebMethodOk() (*string, bool)`

GetWebMethodOk returns a tuple with the WebMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebMethod

`func (o *WebGetConfig1) SetWebMethod(v string)`

SetWebMethod sets WebMethod field to given value.


### GetWebUrl

`func (o *WebGetConfig1) GetWebUrl() string`

GetWebUrl returns the WebUrl field if non-nil, zero value otherwise.

### GetWebUrlOk

`func (o *WebGetConfig1) GetWebUrlOk() (*string, bool)`

GetWebUrlOk returns a tuple with the WebUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebUrl

`func (o *WebGetConfig1) SetWebUrl(v string)`

SetWebUrl sets WebUrl field to given value.


### GetIgnoreSSL

`func (o *WebGetConfig1) GetIgnoreSSL() bool`

GetIgnoreSSL returns the IgnoreSSL field if non-nil, zero value otherwise.

### GetIgnoreSSLOk

`func (o *WebGetConfig1) GetIgnoreSSLOk() (*bool, bool)`

GetIgnoreSSLOk returns a tuple with the IgnoreSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSL

`func (o *WebGetConfig1) SetIgnoreSSL(v bool)`

SetIgnoreSSL sets IgnoreSSL field to given value.

### HasIgnoreSSL

`func (o *WebGetConfig1) HasIgnoreSSL() bool`

HasIgnoreSSL returns a boolean if a field has been set.

### GetCheckUser

`func (o *WebGetConfig1) GetCheckUser() string`

GetCheckUser returns the CheckUser field if non-nil, zero value otherwise.

### GetCheckUserOk

`func (o *WebGetConfig1) GetCheckUserOk() (*string, bool)`

GetCheckUserOk returns a tuple with the CheckUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckUser

`func (o *WebGetConfig1) SetCheckUser(v string)`

SetCheckUser sets CheckUser field to given value.

### HasCheckUser

`func (o *WebGetConfig1) HasCheckUser() bool`

HasCheckUser returns a boolean if a field has been set.

### GetCheckPassword

`func (o *WebGetConfig1) GetCheckPassword() string`

GetCheckPassword returns the CheckPassword field if non-nil, zero value otherwise.

### GetCheckPasswordOk

`func (o *WebGetConfig1) GetCheckPasswordOk() (*string, bool)`

GetCheckPasswordOk returns a tuple with the CheckPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckPassword

`func (o *WebGetConfig1) SetCheckPassword(v string)`

SetCheckPassword sets CheckPassword field to given value.

### HasCheckPassword

`func (o *WebGetConfig1) HasCheckPassword() bool`

HasCheckPassword returns a boolean if a field has been set.

### GetTextCheckOn

`func (o *WebGetConfig1) GetTextCheckOn() string`

GetTextCheckOn returns the TextCheckOn field if non-nil, zero value otherwise.

### GetTextCheckOnOk

`func (o *WebGetConfig1) GetTextCheckOnOk() (*string, bool)`

GetTextCheckOnOk returns a tuple with the TextCheckOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextCheckOn

`func (o *WebGetConfig1) SetTextCheckOn(v string)`

SetTextCheckOn sets TextCheckOn field to given value.

### HasTextCheckOn

`func (o *WebGetConfig1) HasTextCheckOn() bool`

HasTextCheckOn returns a boolean if a field has been set.

### GetWebTextMatch

`func (o *WebGetConfig1) GetWebTextMatch() string`

GetWebTextMatch returns the WebTextMatch field if non-nil, zero value otherwise.

### GetWebTextMatchOk

`func (o *WebGetConfig1) GetWebTextMatchOk() (*string, bool)`

GetWebTextMatchOk returns a tuple with the WebTextMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebTextMatch

`func (o *WebGetConfig1) SetWebTextMatch(v string)`

SetWebTextMatch sets WebTextMatch field to given value.

### HasWebTextMatch

`func (o *WebGetConfig1) HasWebTextMatch() bool`

HasWebTextMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


