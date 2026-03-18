# HealthLogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LogSignature** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SignatureVerified** | Pointer to **bool** |  | [optional] 

## Methods

### NewHealthLogs

`func NewHealthLogs() *HealthLogs`

NewHealthLogs instantiates a new HealthLogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthLogsWithDefaults

`func NewHealthLogsWithDefaults() *HealthLogs`

NewHealthLogsWithDefaults instantiates a new HealthLogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeCode

`func (o *HealthLogs) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *HealthLogs) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *HealthLogs) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *HealthLogs) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetTs

`func (o *HealthLogs) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *HealthLogs) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *HealthLogs) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *HealthLogs) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetLevel

`func (o *HealthLogs) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *HealthLogs) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *HealthLogs) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *HealthLogs) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSourceType

`func (o *HealthLogs) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *HealthLogs) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *HealthLogs) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *HealthLogs) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetMessage

`func (o *HealthLogs) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *HealthLogs) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *HealthLogs) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *HealthLogs) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHostname

`func (o *HealthLogs) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *HealthLogs) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *HealthLogs) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *HealthLogs) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTitle

`func (o *HealthLogs) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HealthLogs) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HealthLogs) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HealthLogs) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLogSignature

`func (o *HealthLogs) GetLogSignature() string`

GetLogSignature returns the LogSignature field if non-nil, zero value otherwise.

### GetLogSignatureOk

`func (o *HealthLogs) GetLogSignatureOk() (*string, bool)`

GetLogSignatureOk returns a tuple with the LogSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSignature

`func (o *HealthLogs) SetLogSignature(v string)`

SetLogSignature sets LogSignature field to given value.

### HasLogSignature

`func (o *HealthLogs) HasLogSignature() bool`

HasLogSignature returns a boolean if a field has been set.

### GetObjectId

`func (o *HealthLogs) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *HealthLogs) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *HealthLogs) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *HealthLogs) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSeq

`func (o *HealthLogs) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *HealthLogs) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *HealthLogs) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *HealthLogs) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetId

`func (o *HealthLogs) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HealthLogs) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HealthLogs) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HealthLogs) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignatureVerified

`func (o *HealthLogs) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *HealthLogs) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *HealthLogs) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *HealthLogs) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


