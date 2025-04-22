# ListLogs200ResponseAllOfDataInner

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

### NewListLogs200ResponseAllOfDataInner

`func NewListLogs200ResponseAllOfDataInner() *ListLogs200ResponseAllOfDataInner`

NewListLogs200ResponseAllOfDataInner instantiates a new ListLogs200ResponseAllOfDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseAllOfDataInnerWithDefaults

`func NewListLogs200ResponseAllOfDataInnerWithDefaults() *ListLogs200ResponseAllOfDataInner`

NewListLogs200ResponseAllOfDataInnerWithDefaults instantiates a new ListLogs200ResponseAllOfDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTypeCode

`func (o *ListLogs200ResponseAllOfDataInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *ListLogs200ResponseAllOfDataInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *ListLogs200ResponseAllOfDataInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *ListLogs200ResponseAllOfDataInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetMessage

`func (o *ListLogs200ResponseAllOfDataInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ListLogs200ResponseAllOfDataInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ListLogs200ResponseAllOfDataInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ListLogs200ResponseAllOfDataInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetLevel

`func (o *ListLogs200ResponseAllOfDataInner) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListLogs200ResponseAllOfDataInner) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListLogs200ResponseAllOfDataInner) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *ListLogs200ResponseAllOfDataInner) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetTs

`func (o *ListLogs200ResponseAllOfDataInner) GetTs() time.Time`

GetTs returns the Ts field if non-nil, zero value otherwise.

### GetTsOk

`func (o *ListLogs200ResponseAllOfDataInner) GetTsOk() (*time.Time, bool)`

GetTsOk returns a tuple with the Ts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTs

`func (o *ListLogs200ResponseAllOfDataInner) SetTs(v time.Time)`

SetTs sets Ts field to given value.

### HasTs

`func (o *ListLogs200ResponseAllOfDataInner) HasTs() bool`

HasTs returns a boolean if a field has been set.

### GetSourceType

`func (o *ListLogs200ResponseAllOfDataInner) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *ListLogs200ResponseAllOfDataInner) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *ListLogs200ResponseAllOfDataInner) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *ListLogs200ResponseAllOfDataInner) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### GetTitle

`func (o *ListLogs200ResponseAllOfDataInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListLogs200ResponseAllOfDataInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListLogs200ResponseAllOfDataInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListLogs200ResponseAllOfDataInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetLogSignature

`func (o *ListLogs200ResponseAllOfDataInner) GetLogSignature() string`

GetLogSignature returns the LogSignature field if non-nil, zero value otherwise.

### GetLogSignatureOk

`func (o *ListLogs200ResponseAllOfDataInner) GetLogSignatureOk() (*string, bool)`

GetLogSignatureOk returns a tuple with the LogSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogSignature

`func (o *ListLogs200ResponseAllOfDataInner) SetLogSignature(v string)`

SetLogSignature sets LogSignature field to given value.

### HasLogSignature

`func (o *ListLogs200ResponseAllOfDataInner) HasLogSignature() bool`

HasLogSignature returns a boolean if a field has been set.

### GetObjectId

`func (o *ListLogs200ResponseAllOfDataInner) GetObjectId() string`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *ListLogs200ResponseAllOfDataInner) GetObjectIdOk() (*string, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *ListLogs200ResponseAllOfDataInner) SetObjectId(v string)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *ListLogs200ResponseAllOfDataInner) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetSeq

`func (o *ListLogs200ResponseAllOfDataInner) GetSeq() int64`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *ListLogs200ResponseAllOfDataInner) GetSeqOk() (*int64, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *ListLogs200ResponseAllOfDataInner) SetSeq(v int64)`

SetSeq sets Seq field to given value.

### HasSeq

`func (o *ListLogs200ResponseAllOfDataInner) HasSeq() bool`

HasSeq returns a boolean if a field has been set.

### GetId

`func (o *ListLogs200ResponseAllOfDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLogs200ResponseAllOfDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLogs200ResponseAllOfDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListLogs200ResponseAllOfDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSignatureVerified

`func (o *ListLogs200ResponseAllOfDataInner) GetSignatureVerified() bool`

GetSignatureVerified returns the SignatureVerified field if non-nil, zero value otherwise.

### GetSignatureVerifiedOk

`func (o *ListLogs200ResponseAllOfDataInner) GetSignatureVerifiedOk() (*bool, bool)`

GetSignatureVerifiedOk returns a tuple with the SignatureVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureVerified

`func (o *ListLogs200ResponseAllOfDataInner) SetSignatureVerified(v bool)`

SetSignatureVerified sets SignatureVerified field to given value.

### HasSignatureVerified

`func (o *ListLogs200ResponseAllOfDataInner) HasSignatureVerified() bool`

HasSignatureVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


