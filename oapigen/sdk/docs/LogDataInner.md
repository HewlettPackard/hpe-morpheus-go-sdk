# LogDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TypeCode** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**Ts** | Pointer to **time.Time** |  | [optional] 
**SourceType** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**LogSignature** | Pointer to **string** |  | [optional] 
**ObjectId** | Pointer to **string** |  | [optional] 
**Seq** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**SignatureVerified** | Pointer to **bool** |  | [optional] 

## Methods

### NewLogDataInner

`func NewLogDataInner() *LogDataInner`

NewLogDataInner instantiates a new LogDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLogDataInnerWithDefaults

`func NewLogDataInnerWithDefaults() *LogDataInner`

NewLogDataInnerWithDefaults instantiates a new LogDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeCode

`func (o *LogDataInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *LogDataInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *LogDataInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *LogDataInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetMessage

`func (o *LogDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *LogDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *LogDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *LogDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLevel

`func (o *LogDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *LogDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *LogDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *LogDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTs

`func (o *LogDataInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *LogDataInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *LogDataInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *LogDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSourceType

`func (o *LogDataInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *LogDataInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *LogDataInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *LogDataInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *LogDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *LogDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *LogDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *LogDataInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLogSignature

`func (o *LogDataInner) GetLogSignature() string`

GetLogSignature returns the LogSignature field if non-nil, zero value otherwise.

### GetLogSignatureOk

`func (o *LogDataInner) GetLogSignatureOk() (*string, bool)`

GetLogSignatureOk returns a tuple with the LogSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSignature

`func (o *LogDataInner) SetLogSignature(v string)`

SetLogSignature sets LogSignature field to given value.

### HasLogSignature

`func (o *LogDataInner) HasLogSignature() bool`

HasLogSignature returns a boolean if a field has been set.

### GetObjectId

`func (o *LogDataInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *LogDataInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *LogDataInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *LogDataInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSeq

`func (o *LogDataInner) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *LogDataInner) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *LogDataInner) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *LogDataInner) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetId

`func (o *LogDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LogDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LogDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LogDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignatureVerified

`func (o *LogDataInner) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *LogDataInner) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *LogDataInner) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *LogDataInner) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


