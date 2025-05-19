# AddCertificate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**ListCertificates200ResponseCertificatesInner**](ListCertificates200ResponseCertificatesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCertificate200Response

`func NewAddCertificate200Response() *AddCertificate200Response`

NewAddCertificate200Response instantiates a new AddCertificate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCertificate200ResponseWithDefaults

`func NewAddCertificate200ResponseWithDefaults() *AddCertificate200Response`

NewAddCertificate200ResponseWithDefaults instantiates a new AddCertificate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *AddCertificate200Response) GetCertificate() ListCertificates200ResponseCertificatesInner`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *AddCertificate200Response) GetCertificateOk() (*ListCertificates200ResponseCertificatesInner, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *AddCertificate200Response) SetCertificate(v ListCertificates200ResponseCertificatesInner)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *AddCertificate200Response) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCertificate200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCertificate200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCertificate200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCertificate200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


