# ListCertificates200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | Pointer to [**[]ListCertificates200ResponseCertificatesInner**](ListCertificates200ResponseCertificatesInner.md) |  | [optional] 
**CertificateCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewListCertificates200Response

`func NewListCertificates200Response() *ListCertificates200Response`

NewListCertificates200Response instantiates a new ListCertificates200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCertificates200ResponseWithDefaults

`func NewListCertificates200ResponseWithDefaults() *ListCertificates200Response`

NewListCertificates200ResponseWithDefaults instantiates a new ListCertificates200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *ListCertificates200Response) GetCertificates() []ListCertificates200ResponseCertificatesInner`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ListCertificates200Response) GetCertificatesOk() (*[]ListCertificates200ResponseCertificatesInner, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ListCertificates200Response) SetCertificates(v []ListCertificates200ResponseCertificatesInner)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ListCertificates200Response) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### GetCertificateCount

`func (o *ListCertificates200Response) GetCertificateCount() int64`

GetCertificateCount returns the CertificateCount field if non-nil, zero value otherwise.

### GetCertificateCountOk

`func (o *ListCertificates200Response) GetCertificateCountOk() (*int64, bool)`

GetCertificateCountOk returns a tuple with the CertificateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCount

`func (o *ListCertificates200Response) SetCertificateCount(v int64)`

SetCertificateCount sets CertificateCount field to given value.

### HasCertificateCount

`func (o *ListCertificates200Response) HasCertificateCount() bool`

HasCertificateCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


