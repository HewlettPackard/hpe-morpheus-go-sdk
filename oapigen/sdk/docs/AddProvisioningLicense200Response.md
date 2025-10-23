# AddProvisioningLicense200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**License** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewAddProvisioningLicense200Response

`func NewAddProvisioningLicense200Response() *AddProvisioningLicense200Response`

NewAddProvisioningLicense200Response instantiates a new AddProvisioningLicense200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddProvisioningLicense200ResponseWithDefaults

`func NewAddProvisioningLicense200ResponseWithDefaults() *AddProvisioningLicense200Response`

NewAddProvisioningLicense200ResponseWithDefaults instantiates a new AddProvisioningLicense200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddProvisioningLicense200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddProvisioningLicense200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddProvisioningLicense200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddProvisioningLicense200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetLicense

`func (o *AddProvisioningLicense200Response) GetLicense() GetAlerts200ResponseAllOfChecksInnerAccount`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *AddProvisioningLicense200Response) GetLicenseOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *AddProvisioningLicense200Response) SetLicense(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetLicense sets License field to given value.

### HasLicense

`func (o *AddProvisioningLicense200Response) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


