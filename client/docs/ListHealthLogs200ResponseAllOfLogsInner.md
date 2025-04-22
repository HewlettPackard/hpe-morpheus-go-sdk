# ListHealthLogs200ResponseAllOfLogsInner

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

### NewListHealthLogs200ResponseAllOfLogsInner

`func NewListHealthLogs200ResponseAllOfLogsInner() *ListHealthLogs200ResponseAllOfLogsInner`

NewListHealthLogs200ResponseAllOfLogsInner instantiates a new ListHealthLogs200ResponseAllOfLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealthLogs200ResponseAllOfLogsInnerWithDefaults

`func NewListHealthLogs200ResponseAllOfLogsInnerWithDefaults() *ListHealthLogs200ResponseAllOfLogsInner`

NewListHealthLogs200ResponseAllOfLogsInnerWithDefaults instantiates a new ListHealthLogs200ResponseAllOfLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeCode

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetTs

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetLevel

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetSourceType

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetMessage

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetHostname

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetTitle

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLogSignature

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetLogSignature() string`

GetLogSignature returns the LogSignature field if non-nil, zero value otherwise.

### GetLogSignatureOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetLogSignatureOk() (*string, bool)`

GetLogSignatureOk returns a tuple with the LogSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSignature

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetLogSignature(v string)`

SetLogSignature sets LogSignature field to given value.

### HasLogSignature

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasLogSignature() bool`

HasLogSignature returns a boolean if a field has been set.

### GetObjectId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSeq

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignatureVerified

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *ListHealthLogs200ResponseAllOfLogsInner) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *ListHealthLogs200ResponseAllOfLogsInner) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *ListHealthLogs200ResponseAllOfLogsInner) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


