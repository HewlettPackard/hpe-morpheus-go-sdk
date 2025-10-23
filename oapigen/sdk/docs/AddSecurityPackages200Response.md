# AddSecurityPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackage** | Pointer to [**ListSecurityPackages200ResponseAllOfSecurityPackagesInner**](ListSecurityPackages200ResponseAllOfSecurityPackagesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddSecurityPackages200Response

`func NewAddSecurityPackages200Response() *AddSecurityPackages200Response`

NewAddSecurityPackages200Response instantiates a new AddSecurityPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityPackages200ResponseWithDefaults

`func NewAddSecurityPackages200ResponseWithDefaults() *AddSecurityPackages200Response`

NewAddSecurityPackages200ResponseWithDefaults instantiates a new AddSecurityPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityPackage

`func (o *AddSecurityPackages200Response) GetSecurityPackage() ListSecurityPackages200ResponseAllOfSecurityPackagesInner`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *AddSecurityPackages200Response) GetSecurityPackageOk() (*ListSecurityPackages200ResponseAllOfSecurityPackagesInner, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *AddSecurityPackages200Response) SetSecurityPackage(v ListSecurityPackages200ResponseAllOfSecurityPackagesInner)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *AddSecurityPackages200Response) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetSuccess

`func (o *AddSecurityPackages200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddSecurityPackages200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddSecurityPackages200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddSecurityPackages200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


