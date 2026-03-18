# UpdateSecurityPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackage** | Pointer to [**UpdateSecurityPackages200ResponseAllOfSecurityPackage**](UpdateSecurityPackages200ResponseAllOfSecurityPackage.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateSecurityPackages200Response

`func NewUpdateSecurityPackages200Response() *UpdateSecurityPackages200Response`

NewUpdateSecurityPackages200Response instantiates a new UpdateSecurityPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityPackages200ResponseWithDefaults

`func NewUpdateSecurityPackages200ResponseWithDefaults() *UpdateSecurityPackages200Response`

NewUpdateSecurityPackages200ResponseWithDefaults instantiates a new UpdateSecurityPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityPackage

`func (o *UpdateSecurityPackages200Response) GetSecurityPackage() UpdateSecurityPackages200ResponseAllOfSecurityPackage`

GetSecurityPackage returns the SecurityPackage field if non-nil, zero value otherwise.

### GetSecurityPackageOk

`func (o *UpdateSecurityPackages200Response) GetSecurityPackageOk() (*UpdateSecurityPackages200ResponseAllOfSecurityPackage, bool)`

GetSecurityPackageOk returns a tuple with the SecurityPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackage

`func (o *UpdateSecurityPackages200Response) SetSecurityPackage(v UpdateSecurityPackages200ResponseAllOfSecurityPackage)`

SetSecurityPackage sets SecurityPackage field to given value.

### HasSecurityPackage

`func (o *UpdateSecurityPackages200Response) HasSecurityPackage() bool`

HasSecurityPackage returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateSecurityPackages200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateSecurityPackages200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateSecurityPackages200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateSecurityPackages200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


